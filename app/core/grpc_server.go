package core

// gRPC server implementation for the Core module.
// Each RPC method converts proto → domain, calls the existing ServiceInterface,
// then converts domain → proto response.

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "swiftschool/proto"
)

// GRPCServer implements pb.CoreServiceServer.
type GRPCServer struct {
	pb.UnimplementedCoreServiceServer
	service ServiceInterface
}

// NewGRPCServer creates a new gRPC server backed by the existing service layer.
func NewGRPCServer(service ServiceInterface) *GRPCServer {
	return &GRPCServer{service: service}
}

// ========================= INSTITUTE =========================

func (g *GRPCServer) CreateInstitute(ctx context.Context, req *pb.CreateInstituteRequest) (*pb.InstituteResponse, error) {
	inst := instituteFromProto(req)
	data, err := g.service.CreateInstitute(ctx, inst)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "institute creation failed: %v", err)
	}
	return instituteToProto(data), nil
}

func (g *GRPCServer) GetInstitute(ctx context.Context, req *pb.IdRequest) (*pb.InstituteResponse, error) {
	id := parseUUID(req.GetId())
	data, err := g.service.GetInstituteById(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get institute: %v", err)
	}
	return instituteToProto(data), nil
}

func (g *GRPCServer) ListInstitutes(ctx context.Context, _ *pb.Empty) (*pb.ListInstitutesResponse, error) {
	data, err := g.service.ListInstitutes(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list institutes: %v", err)
	}
	resp := &pb.ListInstitutesResponse{}
	for _, d := range data {
		resp.Institutes = append(resp.Institutes, instituteToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) UpdateInstitute(ctx context.Context, req *pb.UpdateInstituteRequest) (*pb.InstituteResponse, error) {
	inst := instituteUpdateFromProto(req)
	data, err := g.service.UpdateInstitute(ctx, inst)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "institute update failed: %v", err)
	}
	return instituteToProto(data), nil
}

func (g *GRPCServer) DeleteInstitute(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	id := parseUUID(req.GetId())
	if err := g.service.DeleteInstitute(ctx, id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete institute: %v", err)
	}
	return &pb.Empty{}, nil
}

// ========================= CLASS =========================

func (g *GRPCServer) CreateClass(ctx context.Context, req *pb.CreateClassRequest) (*pb.ClassResponse, error) {
	c := classFromProto(req)
	data, err := g.service.CreateClass(ctx, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "class creation failed: %v", err)
	}
	return classToProto(data), nil
}

func (g *GRPCServer) ListClasses(ctx context.Context, req *pb.TenantListRequest) (*pb.ListClassesResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListClasses(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list classes: %v", err)
	}
	resp := &pb.ListClassesResponse{}
	for _, d := range data {
		resp.Classes = append(resp.Classes, classToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) UpdateClass(ctx context.Context, req *pb.UpdateClassRequest) (*pb.ClassResponse, error) {
	c := classUpdateFromProto(req)
	data, err := g.service.UpdateClass(ctx, c)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "class update failed: %v", err)
	}
	return classToProto(data), nil
}

func (g *GRPCServer) DeleteClass(ctx context.Context, req *pb.DeleteRequest) (*pb.Empty, error) {
	id := parseUUID(req.GetId())
	if err := g.service.DeleteClass(ctx, id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete class: %v", err)
	}
	return &pb.Empty{}, nil
}

// ========================= DEPARTMENT =========================

func (g *GRPCServer) CreateDepartment(ctx context.Context, req *pb.CreateDepartmentRequest) (*pb.DepartmentResponse, error) {
	d := departmentFromProto(req)
	data, err := g.service.CreateDepartment(ctx, d)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "department creation failed: %v", err)
	}
	return departmentToProto(data), nil
}

func (g *GRPCServer) ListDepartments(ctx context.Context, req *pb.TenantListRequest) (*pb.ListDepartmentsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListDepartments(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list departments: %v", err)
	}
	resp := &pb.ListDepartmentsResponse{}
	for _, d := range data {
		resp.Departments = append(resp.Departments, departmentToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) UpdateDepartment(ctx context.Context, req *pb.UpdateDepartmentRequest) (*pb.DepartmentResponse, error) {
	d := departmentUpdateFromProto(req)
	data, err := g.service.UpdateDepartment(ctx, d)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "department update failed: %v", err)
	}
	return departmentToProto(data), nil
}

func (g *GRPCServer) DeleteDepartment(ctx context.Context, req *pb.TenantDeleteRequest) (*pb.Empty, error) {
	instID := parseUUID(req.GetInstituteId())
	id := parseUUID(req.GetId())
	if err := g.service.DeleteDepartment(ctx, instID, id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete department: %v", err)
	}
	return &pb.Empty{}, nil
}

// ========================= STUDENT =========================

func (g *GRPCServer) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.StudentResponse, error) {
	s := studentFromProto(req)
	data, err := g.service.CreateStudent(ctx, s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "student creation failed: %v", err)
	}
	return studentToProto(data), nil
}

func (g *GRPCServer) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.StudentResponse, error) {
	s := studentUpdateFromProto(req)
	data, err := g.service.UpdateStudent(ctx, s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "student update failed: %v", err)
	}
	return studentToProto(data), nil
}

func (g *GRPCServer) DeleteStudent(ctx context.Context, req *pb.TenantDeleteRequest) (*pb.Empty, error) {
	instID := parseUUID(req.GetInstituteId())
	id := parseUUID(req.GetId())
	if err := g.service.DeleteStudent(ctx, instID, id); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete student: %v", err)
	}
	return &pb.Empty{}, nil
}

func (g *GRPCServer) GetStudentFullProfile(ctx context.Context, req *pb.TenantIdRequest) (*pb.StudentResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	id := parseUUID(req.GetId())
	data, err := g.service.GetStudentFullProfile(ctx, instID, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get student profile: %v", err)
	}
	return studentToProto(data), nil
}

func (g *GRPCServer) ListStudentsByClass(ctx context.Context, req *pb.ListStudentsByClassRequest) (*pb.ListStudentsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	classID := parseUUID(req.GetClassId())
	data, err := g.service.ListStudentsByClass(ctx, instID, classID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list students: %v", err)
	}
	resp := &pb.ListStudentsResponse{}
	for _, d := range data {
		resp.Students = append(resp.Students, studentToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) SearchStudents(ctx context.Context, req *pb.SearchStudentsRequest) (*pb.ListStudentsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.SearchStudents(ctx, instID, req.GetQuery())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to search students: %v", err)
	}
	resp := &pb.ListStudentsResponse{}
	for _, d := range data {
		resp.Students = append(resp.Students, studentToProto(d))
	}
	return resp, nil
}

// ========================= GUARDIAN =========================

func (g *GRPCServer) CreateGuardian(ctx context.Context, req *pb.CreateGuardianRequest) (*pb.GuardianResponse, error) {
	guard := guardianFromProto(req)
	data, err := g.service.CreateGuardian(ctx, guard)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "guardian creation failed: %v", err)
	}
	return guardianToProto(data), nil
}

func (g *GRPCServer) LinkStudentGuardian(ctx context.Context, req *pb.LinkStudentGuardianRequest) (*pb.Empty, error) {
	studentID := parseUUID(req.GetStudentId())
	guardianID := parseUUID(req.GetGuardianId())
	if err := g.service.LinkStudentGuardian(ctx, studentID, guardianID, req.GetRelationship(), req.GetIsPrimaryContact()); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to link student-guardian: %v", err)
	}
	return &pb.Empty{}, nil
}

// ========================= ACADEMIC SESSION =========================

func (g *GRPCServer) CreateAcademicSession(ctx context.Context, req *pb.CreateAcademicSessionRequest) (*pb.AcademicSessionResponse, error) {
	s := academicSessionFromProto(req)
	data, err := g.service.CreateAcademicSession(ctx, s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "academic session creation failed: %v", err)
	}
	return academicSessionToProto(data), nil
}

func (g *GRPCServer) ListAcademicSessions(ctx context.Context, req *pb.TenantListRequest) (*pb.ListAcademicSessionsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListAcademicSessions(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list academic sessions: %v", err)
	}
	resp := &pb.ListAcademicSessionsResponse{}
	for _, d := range data {
		resp.Sessions = append(resp.Sessions, academicSessionToProto(d))
	}
	return resp, nil
}

func (g *GRPCServer) GetActiveSession(ctx context.Context, req *pb.TenantListRequest) (*pb.AcademicSessionResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.GetActiveSession(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get active session: %v", err)
	}
	return academicSessionToProto(data), nil
}

func (g *GRPCServer) UpdateAcademicSession(ctx context.Context, req *pb.UpdateAcademicSessionRequest) (*pb.AcademicSessionResponse, error) {
	s := academicSessionUpdateFromProto(req)
	data, err := g.service.UpdateAcademicSession(ctx, s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "academic session update failed: %v", err)
	}
	return academicSessionToProto(data), nil
}

// ========================= ADDRESS =========================

func (g *GRPCServer) CreateAddress(ctx context.Context, req *pb.CreateAddressRequest) (*pb.AddressResponse, error) {
	addr := addressFromProto(req)
	data, err := g.service.CreateAddress(ctx, addr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "address creation failed: %v", err)
	}
	return addressToProto(data), nil
}
