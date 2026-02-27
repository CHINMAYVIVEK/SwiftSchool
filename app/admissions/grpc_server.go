package admissions

// gRPC server implementation for the Admissions module.

import (
	"context"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"swiftschool/domain"
	pb "swiftschool/proto"
)

// GRPCServer implements pb.AdmissionsServiceServer.
type GRPCServer struct {
	pb.UnimplementedAdmissionsServiceServer
	service ServiceInterface
}

// NewGRPCServer creates a new gRPC server backed by the existing admissions service.
func NewGRPCServer(service ServiceInterface) *GRPCServer {
	return &GRPCServer{service: service}
}

// ========================= HELPERS =========================

func parseUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
}

func optionalStringValue(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	s := v.GetValue()
	return &s
}

func toStringValue(s *string) *wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	return wrapperspb.String(*s)
}

func fromOptionalTimestamp(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

func toOptionalTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// ========================= ENQUIRY RPCs =========================

func (g *GRPCServer) CreateEnquiry(ctx context.Context, req *pb.CreateEnquiryRequest) (*pb.EnquiryResponse, error) {
	enq := domain.AdmissionEnquiry{
		StudentName:      optionalStringValue(req.GetStudentName()),
		GuardianName:     optionalStringValue(req.GetGuardianName()),
		Phone:            optionalStringValue(req.GetPhone()),
		Email:            optionalStringValue(req.GetEmail()),
		ClassApplyingFor: optionalStringValue(req.GetClassApplyingFor()),
		PreviousSchool:   optionalStringValue(req.GetPreviousSchool()),
		Status:           domain.AdmissionStatus(req.GetStatus()),
		EnquiryDate:      fromOptionalTimestamp(req.GetEnquiryDate()),
		FollowUpDate:     fromOptionalTimestamp(req.GetFollowUpDate()),
	}
	enq.InstituteID = parseUUID(req.GetInstituteId())

	data, err := g.service.CreateEnquiry(ctx, enq)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "enquiry creation failed: %v", err)
	}
	return enquiryToProto(data), nil
}

func (g *GRPCServer) ListEnquiries(ctx context.Context, req *pb.TenantListRequest) (*pb.ListEnquiriesResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListEnquiries(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list enquiries: %v", err)
	}
	resp := &pb.ListEnquiriesResponse{}
	for _, d := range data {
		resp.Enquiries = append(resp.Enquiries, enquiryToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) UpdateEnquiryStatus(ctx context.Context, req *pb.UpdateEnquiryStatusRequest) (*pb.Empty, error) {
	id := parseUUID(req.GetId())
	instID := parseUUID(req.GetInstituteId())
	if err := g.service.UpdateEnquiryStatus(ctx, id, instID, domain.AdmissionStatus(req.GetStatus())); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update enquiry status: %v", err)
	}
	return &pb.Empty{}, nil
}

// ========================= CONVERTERS =========================

func enquiryToProto(d *domain.AdmissionEnquiry) *pb.EnquiryResponse {
	if d == nil {
		return nil
	}
	return &pb.EnquiryResponse{
		Id:               d.ID.String(),
		InstituteId:      d.InstituteID.String(),
		StudentName:      toStringValue(d.StudentName),
		GuardianName:     toStringValue(d.GuardianName),
		Phone:            toStringValue(d.Phone),
		Email:            toStringValue(d.Email),
		ClassApplyingFor: toStringValue(d.ClassApplyingFor),
		PreviousSchool:   toStringValue(d.PreviousSchool),
		Status:           string(d.Status),
		EnquiryDate:      toOptionalTimestamp(d.EnquiryDate),
		FollowUpDate:     toOptionalTimestamp(d.FollowUpDate),
		CreatedAt:        timestamppb.New(d.CreatedAt),
		UpdatedAt:        timestamppb.New(d.UpdatedAt),
	}
}
