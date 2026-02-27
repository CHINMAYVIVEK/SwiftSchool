package academics

// gRPC server implementation for the Academics module.

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	pb "swiftschool/app/academics/proto"
	"swiftschool/domain"
)

// GRPCServer implements pb.AcademicsServiceServer.
type GRPCServer struct {
	pb.UnimplementedAcademicsServiceServer
	service ServiceInterface
}

// NewGRPCServer creates a new gRPC server backed by the existing academics service.
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

// ========================= SUBJECTS =========================

func (g *GRPCServer) CreateSubject(ctx context.Context, req *pb.CreateSubjectRequest) (*pb.SubjectResponse, error) {
	s := domain.Subject{
		Name:    req.GetName(),
		Code:    optionalStringValue(req.GetCode()),
		Type:    req.GetType(),
		Credits: req.GetCredits(),
	}
	s.InstituteID = parseUUID(req.GetInstituteId())

	data, err := g.service.CreateSubject(ctx, s)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "subject creation failed: %v", err)
	}
	return subjectToProto(data), nil
}

func (g *GRPCServer) ListSubjects(ctx context.Context, req *pb.TenantListRequest) (*pb.ListSubjectsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListSubjects(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list subjects: %v", err)
	}
	resp := &pb.ListSubjectsResponse{}
	for _, d := range data {
		resp.Subjects = append(resp.Subjects, subjectToProto(d))
	}
	return resp, nil
}

// ========================= CLASS PERIODS =========================

func (g *GRPCServer) CreateClassPeriod(ctx context.Context, req *pb.CreateClassPeriodRequest) (*pb.ClassPeriodResponse, error) {
	cp := domain.ClassPeriod{
		Name:      optionalStringValue(req.GetName()),
		StartTime: req.GetStartTime(),
		EndTime:   req.GetEndTime(),
		IsBreak:   req.GetIsBreak(),
	}
	cp.InstituteID = parseUUID(req.GetInstituteId())

	data, err := g.service.CreateClassPeriod(ctx, cp)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "class period creation failed: %v", err)
	}
	return classPeriodToProto(data), nil
}

func (g *GRPCServer) ListClassPeriods(ctx context.Context, req *pb.TenantListRequest) (*pb.ListClassPeriodsResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	data, err := g.service.ListClassPeriods(ctx, instID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list class periods: %v", err)
	}
	resp := &pb.ListClassPeriodsResponse{}
	for _, d := range data {
		resp.Periods = append(resp.Periods, classPeriodToProto(d))
	}
	return resp, nil
}

// ========================= TIMETABLE =========================

func (g *GRPCServer) CreateTimetableEntry(ctx context.Context, req *pb.CreateTimetableEntryRequest) (*pb.TimetableEntryResponse, error) {
	te := domain.TimetableEntry{
		DayOfWeek: domain.DayOfWeek(req.GetDayOfWeek()),
		ClassID:   optionalUUIDStringValue(req.GetClassId()),
		PeriodID:  optionalUUIDStringValue(req.GetPeriodId()),
		SubjectID: optionalUUIDStringValue(req.GetSubjectId()),
		TeacherID: optionalUUIDStringValue(req.GetTeacherId()),
	}
	te.InstituteID = parseUUID(req.GetInstituteId())
	te.AcademicSessionID = parseUUID(req.GetAcademicSessionId())

	data, err := g.service.CreateTimetableEntry(ctx, te)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "timetable entry creation failed: %v", err)
	}
	return timetableEntryToProto(data), nil
}

func (g *GRPCServer) GetClassTimetable(ctx context.Context, req *pb.GetClassTimetableRequest) (*pb.ListTimetableEntriesResponse, error) {
	instID := parseUUID(req.GetInstituteId())
	// Parse optional class_id — use uuid.Nil if absent
	var classID uuid.UUID
	if req.GetClassId() != nil {
		classID = parseUUID(req.GetClassId().GetValue())
	}
	// The existing service requires a day parameter but the proto doesn't include it;
	// we pass an empty day to get all entries (the repo should handle this).
	data, err := g.service.GetClassTimetable(ctx, instID, classID, "")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get class timetable: %v", err)
	}
	resp := &pb.ListTimetableEntriesResponse{}
	for _, d := range data {
		resp.Entries = append(resp.Entries, timetableEntryToProto(d))
	}
	return resp, nil
}

// ========================= CONVERTERS =========================

func subjectToProto(d *domain.Subject) *pb.SubjectResponse {
	if d == nil {
		return nil
	}
	return &pb.SubjectResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		Name:        d.Name,
		Code:        toStringValue(d.Code),
		Type:        d.Type,
		Credits:     d.Credits,
		CreatedAt:   timestamppb.New(d.CreatedAt),
		UpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}

func classPeriodToProto(d *domain.ClassPeriod) *pb.ClassPeriodResponse {
	if d == nil {
		return nil
	}
	return &pb.ClassPeriodResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		Name:        toStringValue(d.Name),
		StartTime:   d.StartTime,
		EndTime:     d.EndTime,
		IsBreak:     d.IsBreak,
		CreatedAt:   timestamppb.New(d.CreatedAt),
		UpdatedAt:   timestamppb.New(d.UpdatedAt),
	}
}

func timetableEntryToProto(d *domain.TimetableEntry) *pb.TimetableEntryResponse {
	if d == nil {
		return nil
	}
	return &pb.TimetableEntryResponse{
		Id:                d.ID.String(),
		InstituteId:       d.InstituteID.String(),
		AcademicSessionId: d.AcademicSessionID.String(),
		ClassId:           toUUIDStringValue(d.ClassID),
		DayOfWeek:         string(d.DayOfWeek),
		PeriodId:          toUUIDStringValue(d.PeriodID),
		SubjectId:         toUUIDStringValue(d.SubjectID),
		TeacherId:         toUUIDStringValue(d.TeacherID),
		CreatedAt:         timestamppb.New(d.CreatedAt),
		UpdatedAt:         timestamppb.New(d.UpdatedAt),
	}
}
