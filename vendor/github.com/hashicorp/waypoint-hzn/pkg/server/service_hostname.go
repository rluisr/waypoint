package server

import (
	"context"
	"fmt"
	"strings"

	petname "github.com/dustinkirkland/golang-petname"
	empty "github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/horizon/pkg/dbx"
	hznpb "github.com/hashicorp/horizon/pkg/pb"
	"github.com/jinzhu/gorm"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hashicorp/waypoint-hzn/pkg/models"
	"github.com/hashicorp/waypoint-hzn/pkg/pb"
)

func (s *service) RegisterHostname(
	ctx context.Context,
	req *pb.RegisterHostnameRequest,
) (*pb.RegisterHostnameResponse, error) {
	L := hclog.FromContext(ctx)

	// Auth required.
	token, err := s.checkAuth(ctx)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Parse our labels
	var labels hznpb.LabelSet
	if err := mapstructure.Decode(req.Labels, &labels); err != nil {
		return nil, err
	}

	// Get our account registration
	var reg models.Registration
	if err = dbx.Check(
		s.DB.Where("account_id = ?", token.Account().AccountId.Bytes()).First(&reg),
	); err != nil {
		return nil, status.Errorf(codes.PermissionDenied,
			"unregistered account")
	}

	// Determine the full hostname
	var hostname, fqdn string

trying:
	for {
		switch v := req.Hostname.(type) {
		case *pb.RegisterHostnameRequest_Generate:
			hostname = petname.Generate(3, "-")

			if strings.Contains(hostname, "--") {
				// extremely odd, but go ahead and just retry
				continue trying
			}
		case *pb.RegisterHostnameRequest_Exact:
			hostname = v.Exact

			if strings.Contains(hostname, "--") {
				return nil, fmt.Errorf("hostname must not contain a double hyphen")
			}
		}

		var host models.Hostname
		host.RegistrationId = reg.Id
		host.Hostname = hostname
		host.Labels = labels.AsStringArray()

		if err := dbx.Check(s.DB.Create(&host)); err != nil {
			// For now, assume the failure is because of failing the unique
			// constraint. If we autogenerated the name, retry, otherwise return
			// an error.
			if _, ok := req.Hostname.(*pb.RegisterHostnameRequest_Generate); ok {
				continue
			}

			L.Error("error creating hostname", "error", err)
			return nil, fmt.Errorf("requested hostname is not available")
		}

		// Add the domain
		fqdn = hostname + "." + s.Domain

		break
	}

	L.Debug("adding label link", "hostname", fqdn, "target", req.Labels)
	_, err = s.HznControl.AddLabelLink(ctx, &hznpb.AddLabelLinkRequest{
		Labels:  hznpb.MakeLabels(":hostname", fqdn),
		Account: token.Account(),
		Target:  &labels,
	})
	if err != nil {
		return nil, err
	}
	L.Info("added label link", "hostname", fqdn, "target", req.Labels)

	return &pb.RegisterHostnameResponse{
		Hostname: hostname,
		Fqdn:     fqdn,
	}, nil
}

func (s *service) ListHostnames(
	ctx context.Context,
	req *pb.ListHostnamesRequest,
) (*pb.ListHostnamesResponse, error) {
	L := hclog.FromContext(ctx)

	// Auth required.
	token, err := s.checkAuth(ctx)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Get our account registration
	var reg models.Registration
	if err = dbx.Check(
		s.DB.Where("account_id = ?", token.Account().AccountId.Bytes()).First(&reg),
	); err != nil {
		return nil, status.Errorf(codes.PermissionDenied,
			"unregistered account")
	}

	var hostnames []*models.Hostname
	err = dbx.Check(s.DB.Find(&hostnames, "registration_id = ?", reg.Id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Errorf(codes.NotFound, "unregistered account")
		}

		L.Error("error looking up hostnames for account", "registration-id", reg.Id)
		return nil, status.Errorf(codes.Internal, "error querying hostnames")
	}

	var resp pb.ListHostnamesResponse
	for _, h := range hostnames {
		var hznlabels hznpb.LabelSet
		if err := hznlabels.Scan(h.Labels); err != nil {
			return nil, status.Errorf(codes.Internal, "error querying hostnames")
		}

		// Parse our labels
		var labels pb.LabelSet
		if err := mapstructure.Decode(hznlabels, &labels); err != nil {
			return nil, err
		}

		resp.Hostnames = append(resp.Hostnames, &pb.ListHostnamesResponse_Hostname{
			Hostname: h.Hostname,
			Fqdn:     h.Hostname + "." + s.Domain,
			Labels:   &labels,
		})
	}

	return &resp, nil
}

func (s *service) DeleteHostname(
	ctx context.Context,
	req *pb.DeleteHostnameRequest,
) (*empty.Empty, error) {
	L := hclog.FromContext(ctx)

	// Auth required.
	token, err := s.checkAuth(ctx)
	if err != nil {
		return nil, err
	}

	// Validate
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// Get our account registration
	var reg models.Registration
	if err = dbx.Check(
		s.DB.Where("account_id = ?", token.Account().AccountId.Bytes()).First(&reg),
	); err != nil {
		return nil, status.Errorf(codes.PermissionDenied,
			"unregistered account")
	}

	// Delete from our DB
	err = dbx.Check(s.DB.Delete(
		models.Hostname{},
		"registration_id = ? and hostname = ?", reg.Id, req.Hostname),
	)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &empty.Empty{}, nil
		}

		L.Error("error looking up hostnames for account", "registration-id", reg.Id)
		return nil, status.Errorf(codes.Internal, "error deleting hostname")
	}

	_, err = s.HznControl.RemoveLabelLink(ctx, &hznpb.RemoveLabelLinkRequest{
		Labels:  hznpb.MakeLabels(":hostname", req.Hostname),
		Account: token.Account(),
	})
	if err != nil {
		L.Error("error removing label link", "error", err)
		return nil, status.Errorf(codes.Internal, "error deleting hostname")
	}

	return &empty.Empty{}, nil
}