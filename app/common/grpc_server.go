package common

// gRPC server implementation for the Common module (Documents & Notifications).

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"swiftschool/domain"
	pb "swiftschool/proto"
)

// GRPCServer implements pb.CommonModuleServiceServer.
type GRPCServer struct {
	pb.UnimplementedCommonModuleServiceServer
	service ServiceInterface
}

// NewGRPCServer creates a new gRPC server backed by the existing common service.
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

func optionalUUIDStringValue(v *wrapperspb.StringValue) *uuid.UUID {
	if v == nil {
		return nil
	}
	id, err := uuid.Parse(v.GetValue())
	if err != nil {
		return nil
	}
	return &id
}

func toUUIDStringValue(u *uuid.UUID) *wrapperspb.StringValue {
	if u == nil {
		return nil
	}
	return wrapperspb.String(u.String())
}

// ========================= DOCUMENT RPCs =========================

func (g *GRPCServer) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.DocumentResponse, error) {
	doc := domain.Document{
		OwnerID:   parseUUID(req.GetOwnerId()),
		OwnerType: domain.OwnerType(req.GetOwnerType()),
		DocType:   domain.DocumentType(req.GetDocType()),
		FileName:  optionalStringValue(req.GetFileName()),
		FileURL:   req.GetFileUrl(),
	}
	doc.InstituteID = parseUUID(req.GetInstituteId())

	data, err := g.service.CreateDocument(ctx, doc)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "document creation failed: %v", err)
	}
	return documentToProto(data), nil
}

func (g *GRPCServer) ListDocuments(ctx context.Context, req *pb.TenantListRequest) (*pb.ListDocumentsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	// ListDocuments in the service requires ownerID; for now pass uuid.Nil to list all
	data, err := g.service.ListDocuments(ctx, instID, uuid.Nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list documents: %v", err)
	}
	resp := &pb.ListDocumentsResponse{}
	for _, d := range data {
		resp.Documents = append(resp.Documents, documentToProto(d))
	}
	return resp, nil
}

// ========================= NOTIFICATION RPCs =========================

func (g *GRPCServer) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error) {
	n := domain.Notification{
		UserID:  optionalUUIDStringValue(req.GetUserId()),
		Title:   optionalStringValue(req.GetTitle()),
		Message: optionalStringValue(req.GetMessage()),
		IsRead:  req.GetIsRead(),
	}
	n.InstituteID = parseUUID(req.GetInstituteId())

	data, err := g.service.CreateNotification(ctx, n)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "notification creation failed: %v", err)
	}
	return notificationToProto(data), nil
}

// ========================= CONVERTERS =========================

func documentToProto(d *domain.Document) *pb.DocumentResponse {
	if d == nil {
		return nil
	}
	return &pb.DocumentResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		OwnerId:     d.OwnerID.String(),
		OwnerType:   string(d.OwnerType),
		DocType:     string(d.DocType),
		FileName:    toStringValue(d.FileName),
		FileUrl:     d.FileURL,
		CreatedAt:   timestamppb.New(d.CreatedAt),
		UpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}

func notificationToProto(d *domain.Notification) *pb.NotificationResponse {
	if d == nil {
		return nil
	}
	return &pb.NotificationResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		UserId:      toUUIDStringValue(d.UserID),
		Title:       toStringValue(d.Title),
		Message:     toStringValue(d.Message),
		IsRead:      d.IsRead,
		CreatedAt:   timestamppb.New(d.CreatedAt),
		UpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}
