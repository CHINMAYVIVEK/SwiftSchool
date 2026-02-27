package auth

// gRPC server implementation for the Auth module.

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

// GRPCServer implements pb.AuthServiceServer.
type GRPCServer struct {
	pb.UnimplementedAuthServiceServer
	service ServiceInterface
}

// NewGRPCServer creates a new gRPC server backed by the existing auth service.
func NewGRPCServer(service ServiceInterface) *GRPCServer {
	return &GRPCServer{service: service}
}

// ========================= HELPERS =========================

func parseUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
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

// ========================= AUTH RPCs =========================

func (g *GRPCServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if err := g.service.Login(ctx, req.GetIdentifier(), req.GetUserType()); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "login failed: %v", err)
	}
	return &pb.LoginResponse{
		Success: true,
		Message: "OTP sent successfully. Valid for 5 minutes.",
	}, nil
}

func (g *GRPCServer) VerifyOTP(ctx context.Context, req *pb.VerifyOTPRequest) (*pb.VerifyOTPResponse, error) {
	user, err := g.service.VerifyOTP(ctx, req.GetIdentifier(), req.GetUserType(), req.GetOtp())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "OTP verification failed: %v", err)
	}
	return &pb.VerifyOTPResponse{
		Success: true,
		Message: "OTP verified successfully",
		User:    userToProto(user),
	}, nil
}

// ========================= USER MANAGEMENT RPCs =========================

func (g *GRPCServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	user := domain.User{
		Username:       req.GetUsername(),
		PasswordHash:   req.GetPassword(),
		RoleType:       domain.UserRole(req.GetRoleType()),
		LinkedEntityID: parseUUID(req.GetLinkedEntityId()),
		InstituteID:    optionalUUIDStringValue(req.GetInstituteId()),
		IsActive:       req.GetIsActive(),
	}

	data, err := g.service.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}
	return userToProto(data), nil
}

func (g *GRPCServer) GetUserByUsername(ctx context.Context, req *pb.GetUserByUsernameRequest) (*pb.UserResponse, error) {
	data, err := g.service.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user: %v", err)
	}
	return userToProto(data), nil
}

func (g *GRPCServer) GetUserById(ctx context.Context, req *pb.IdRequest) (*pb.UserResponse, error) {
	id := parseUUID(req.GetId())
	data, err := g.service.GetUserById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch user: %v", err)
	}
	return userToProto(data), nil
}

func (g *GRPCServer) UpdateUserPassword(ctx context.Context, req *pb.UpdateUserPasswordRequest) (*pb.Empty, error) {
	id := parseUUID(req.GetId())
	if err := g.service.UpdateUserPassword(ctx, id, req.GetPassword()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update password: %v", err)
	}
	return &pb.Empty{}, nil
}

func (g *GRPCServer) UpdateUserStatus(ctx context.Context, req *pb.UpdateUserStatusRequest) (*pb.Empty, error) {
	id := parseUUID(req.GetId())
	if err := g.service.UpdateUserStatus(ctx, id, req.GetIsActive()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update user status: %v", err)
	}
	return &pb.Empty{}, nil
}

func (g *GRPCServer) ListUsersByRole(ctx context.Context, req *pb.ListUsersByRoleRequest) (*pb.ListUsersResponse, error) {
	instID := parseUUID(req.GetInstituteId().GetValue())
	role := domain.UserRole(req.GetRoleType())
	data, err := g.service.ListUsersByRole(ctx, instID, role)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users: %v", err)
	}
	resp := &pb.ListUsersResponse{}
	for _, d := range data {
		resp.Users = append(resp.Users, userToProto(d))
	}
	return resp, nil
}

// ========================= CONVERTERS =========================

func userToProto(d *domain.User) *pb.UserResponse {
	if d == nil {
		return nil
	}
	return &pb.UserResponse{
		Id:             d.ID.String(),
		Username:       d.Username,
		RoleType:       string(d.RoleType),
		LinkedEntityId: d.LinkedEntityID.String(),
		InstituteId:    toUUIDStringValue(d.InstituteID),
		IsActive:       d.IsActive,
		CreatedAt:      timestamppb.New(d.CreatedAt),
		UpdatedAt:      timestamppb.New(d.UpdatedAt),
	}
}
