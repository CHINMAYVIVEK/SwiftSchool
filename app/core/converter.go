package core

// Proto-Domain converters for the Core module.
// Bidirectional conversion between protobuf messages and domain structs.

import (
	"time"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"swiftschool/domain"
	pb "swiftschool/proto"
)

// ========================= HELPERS =========================

// parseUUID is a safe wrapper that returns uuid.Nil on invalid input.
func parseUUID(s string) uuid.UUID {
	id, _ := uuid.Parse(s)
	return id
}

// optionalStringValue extracts Go *string from proto StringValue.
func optionalStringValue(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	s := v.GetValue()
	return &s
}

// toStringValue converts Go *string to proto StringValue.
func toStringValue(s *string) *wrapperspb.StringValue {
	if s == nil {
		return nil
	}
	return wrapperspb.String(*s)
}

// optionalUUIDStringValue extracts Go *uuid.UUID from proto StringValue.
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

// toUUIDStringValue converts Go *uuid.UUID to proto StringValue.
func toUUIDStringValue(u *uuid.UUID) *wrapperspb.StringValue {
	if u == nil {
		return nil
	}
	return wrapperspb.String(u.String())
}

// toTimestamp converts Go time.Time to proto Timestamp.
func toTimestamp(t time.Time) *timestamppb.Timestamp {
	if t.IsZero() {
		return nil
	}
	return timestamppb.New(t)
}

// toOptionalTimestamp converts Go *time.Time to proto Timestamp.
func toOptionalTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

// fromTimestamp converts proto Timestamp to Go time.Time.
func fromTimestamp(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}

// fromOptionalTimestamp converts proto Timestamp to Go *time.Time.
func fromOptionalTimestamp(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}
	t := ts.AsTime()
	return &t
}

// optionalDoubleValue extracts Go *float64 from proto DoubleValue.
func optionalDoubleValue(v *wrapperspb.DoubleValue) *float64 {
	if v == nil {
		return nil
	}
	f := v.GetValue()
	return &f
}

// toDoubleValue converts Go *float64 to proto DoubleValue.
func toDoubleValue(f *float64) *wrapperspb.DoubleValue {
	if f == nil {
		return nil
	}
	return wrapperspb.Double(*f)
}

// ========================= INSTITUTE =========================

func instituteFromProto(req *pb.CreateInstituteRequest) domain.Institute {
	return domain.Institute{
		Name:               req.GetName(),
		Code:               req.GetCode(),
		CurrencyCode:       optionalStringValue(req.GetCurrencyCode()),
		LogoURL:            optionalStringValue(req.GetLogoUrl()),
		Website:            optionalStringValue(req.GetWebsite()),
		Timezone:           req.GetTimezone(),
		FiscalYearStartMon: int(req.GetFiscalYearStartMonth()),
		IsActive:           req.GetIsActive(),
	}
}

func instituteUpdateFromProto(req *pb.UpdateInstituteRequest) domain.Institute {
	inst := domain.Institute{
		Name:               req.GetName(),
		Code:               req.GetCode(),
		CurrencyCode:       optionalStringValue(req.GetCurrencyCode()),
		LogoURL:            optionalStringValue(req.GetLogoUrl()),
		Website:            optionalStringValue(req.GetWebsite()),
		Timezone:           req.GetTimezone(),
		FiscalYearStartMon: int(req.GetFiscalYearStartMonth()),
		IsActive:           req.GetIsActive(),
	}
	inst.ID = parseUUID(req.GetId())
	return inst
}

func instituteToProto(d *domain.Institute) *pb.InstituteResponse {
	if d == nil {
		return nil
	}
	return &pb.InstituteResponse{
		Id:                   d.ID.String(),
		Name:                 d.Name,
		Code:                 d.Code,
		CurrencyCode:         toStringValue(d.CurrencyCode),
		LogoUrl:              toStringValue(d.LogoURL),
		Website:              toStringValue(d.Website),
		Timezone:             d.Timezone,
		FiscalYearStartMonth: int32(d.FiscalYearStartMon),
		IsActive:             d.IsActive,
		CreatedAt:            toTimestamp(d.CreatedAt),
		UpdatedAt:            toTimestamp(d.UpdatedAt),
		DeletedAt:            toOptionalTimestamp(d.DeletedAt),
	}
}

// ========================= CLASS =========================

func classFromProto(req *pb.CreateClassRequest) domain.Class {
	c := domain.Class{
		Name:           req.GetName(),
		Section:        req.GetSection(),
		ClassTeacherID: optionalUUIDStringValue(req.GetClassTeacherId()),
	}
	c.InstituteID = parseUUID(req.GetInstituteId())
	c.AcademicSessionID = parseUUID(req.GetAcademicSessionId())
	return c
}

func classUpdateFromProto(req *pb.UpdateClassRequest) domain.Class {
	c := domain.Class{
		Name:           req.GetName(),
		Section:        req.GetSection(),
		ClassTeacherID: optionalUUIDStringValue(req.GetClassTeacherId()),
	}
	c.ID = parseUUID(req.GetId())
	c.InstituteID = parseUUID(req.GetInstituteId())
	c.AcademicSessionID = parseUUID(req.GetAcademicSessionId())
	return c
}

func classToProto(d *domain.Class) *pb.ClassResponse {
	if d == nil {
		return nil
	}
	return &pb.ClassResponse{
		Id:                d.ID.String(),
		InstituteId:       d.InstituteID.String(),
		AcademicSessionId: d.AcademicSessionID.String(),
		Name:              d.Name,
		Section:           d.Section,
		ClassTeacherId:    toUUIDStringValue(d.ClassTeacherID),
		CreatedAt:         toTimestamp(d.CreatedAt),
		UpdatedAt:         toTimestamp(d.UpdatedAt),
	}
}

// ========================= DEPARTMENT =========================

func departmentFromProto(req *pb.CreateDepartmentRequest) domain.Department {
	d := domain.Department{
		Name: req.GetName(),
	}
	d.InstituteID = parseUUID(req.GetInstituteId())
	return d
}

func departmentUpdateFromProto(req *pb.UpdateDepartmentRequest) domain.Department {
	d := domain.Department{
		Name: req.GetName(),
	}
	d.ID = parseUUID(req.GetId())
	d.InstituteID = parseUUID(req.GetInstituteId())
	return d
}

func departmentToProto(d *domain.Department) *pb.DepartmentResponse {
	if d == nil {
		return nil
	}
	return &pb.DepartmentResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		Name:        d.Name,
		CreatedAt:   toTimestamp(d.CreatedAt),
		UpdatedAt:   toTimestamp(d.UpdatedAt),
	}
}

// ========================= STUDENT =========================

func languageSkillsFromProto(skills []*pb.LanguageSkill) []domain.LanguageSkill {
	if len(skills) == 0 {
		return nil
	}
	out := make([]domain.LanguageSkill, len(skills))
	for i, s := range skills {
		out[i] = domain.LanguageSkill{
			Language:    s.GetLanguage(),
			Proficiency: domain.LanguageProficiency(s.GetProficiency()),
			CanRead:     s.GetCanRead(),
			CanWrite:    s.GetCanWrite(),
			CanSpeak:    s.GetCanSpeak(),
		}
	}
	return out
}

func languageSkillsToProto(skills []domain.LanguageSkill) []*pb.LanguageSkill {
	if len(skills) == 0 {
		return nil
	}
	out := make([]*pb.LanguageSkill, len(skills))
	for i, s := range skills {
		out[i] = &pb.LanguageSkill{
			Language:    s.Language,
			Proficiency: string(s.Proficiency),
			CanRead:     s.CanRead,
			CanWrite:    s.CanWrite,
			CanSpeak:    s.CanSpeak,
		}
	}
	return out
}

func socialMediaFromProto(smh *pb.SocialMediaHandles) domain.SocialMediaHandles {
	if smh == nil || len(smh.GetHandles()) == 0 {
		return nil
	}
	return domain.SocialMediaHandles(smh.GetHandles())
}

func socialMediaToProto(smh domain.SocialMediaHandles) *pb.SocialMediaHandles {
	if len(smh) == 0 {
		return nil
	}
	return &pb.SocialMediaHandles{Handles: map[string]string(smh)}
}

func studentFromProto(req *pb.CreateStudentRequest) domain.Student {
	s := domain.Student{
		AdmissionNo:        req.GetAdmissionNo(),
		FirstName:          req.GetFirstName(),
		LastName:           optionalStringValue(req.GetLastName()),
		DOB:                fromOptionalTimestamp(req.GetDob()),
		Gender:             domain.Gender(req.GetGender()),
		BloodGroup:         domain.BloodGroup(req.GetBloodGroup()),
		SocialCategory:     domain.SocialCategory(req.GetSocialCategory()),
		CurrentClassID:     optionalUUIDStringValue(req.GetCurrentClassId()),
		Nationality:        optionalStringValue(req.GetNationality()),
		PreferredLanguage:  optionalStringValue(req.GetPreferredLanguage()),
		SocialMediaHandles: socialMediaFromProto(req.GetSocialMediaHandles()),
		LanguageSkills:     languageSkillsFromProto(req.GetLanguageSkills()),
	}
	s.InstituteID = parseUUID(req.GetInstituteId())
	return s
}

func studentUpdateFromProto(req *pb.UpdateStudentRequest) domain.Student {
	s := domain.Student{
		AdmissionNo:        req.GetAdmissionNo(),
		FirstName:          req.GetFirstName(),
		LastName:           optionalStringValue(req.GetLastName()),
		DOB:                fromOptionalTimestamp(req.GetDob()),
		Gender:             domain.Gender(req.GetGender()),
		BloodGroup:         domain.BloodGroup(req.GetBloodGroup()),
		SocialCategory:     domain.SocialCategory(req.GetSocialCategory()),
		CurrentClassID:     optionalUUIDStringValue(req.GetCurrentClassId()),
		Nationality:        optionalStringValue(req.GetNationality()),
		PreferredLanguage:  optionalStringValue(req.GetPreferredLanguage()),
		SocialMediaHandles: socialMediaFromProto(req.GetSocialMediaHandles()),
		LanguageSkills:     languageSkillsFromProto(req.GetLanguageSkills()),
	}
	s.ID = parseUUID(req.GetId())
	s.InstituteID = parseUUID(req.GetInstituteId())
	return s
}

func studentToProto(d *domain.Student) *pb.StudentResponse {
	if d == nil {
		return nil
	}
	return &pb.StudentResponse{
		Id:                 d.ID.String(),
		InstituteId:        d.InstituteID.String(),
		AdmissionNo:        d.AdmissionNo,
		FirstName:          d.FirstName,
		LastName:           toStringValue(d.LastName),
		Dob:                toOptionalTimestamp(d.DOB),
		Gender:             string(d.Gender),
		BloodGroup:         string(d.BloodGroup),
		SocialCategory:     string(d.SocialCategory),
		CurrentClassId:     toUUIDStringValue(d.CurrentClassID),
		Nationality:        toStringValue(d.Nationality),
		PreferredLanguage:  toStringValue(d.PreferredLanguage),
		SocialMediaHandles: socialMediaToProto(d.SocialMediaHandles),
		LanguageSkills:     languageSkillsToProto(d.LanguageSkills),
		CreatedAt:          toTimestamp(d.CreatedAt),
		UpdatedAt:          toTimestamp(d.UpdatedAt),
	}
}

// ========================= GUARDIAN =========================

func guardianFromProto(req *pb.CreateGuardianRequest) domain.Guardian {
	return domain.Guardian{
		FirstName:    req.GetFirstName(),
		LastName:     optionalStringValue(req.GetLastName()),
		Email:        optionalStringValue(req.GetEmail()),
		Phone:        optionalStringValue(req.GetPhone()),
		Profession:   optionalStringValue(req.GetProfession()),
		AnnualIncome: optionalDoubleValue(req.GetAnnualIncome()),
	}
}

func guardianToProto(d *domain.Guardian) *pb.GuardianResponse {
	if d == nil {
		return nil
	}
	return &pb.GuardianResponse{
		Id:           d.ID.String(),
		FirstName:    d.FirstName,
		LastName:     toStringValue(d.LastName),
		Email:        toStringValue(d.Email),
		Phone:        toStringValue(d.Phone),
		Profession:   toStringValue(d.Profession),
		AnnualIncome: toDoubleValue(d.AnnualIncome),
		CreatedAt:    toTimestamp(d.CreatedAt),
		UpdatedAt:    toTimestamp(d.UpdatedAt),
	}
}

// ========================= ACADEMIC SESSION =========================

func academicSessionFromProto(req *pb.CreateAcademicSessionRequest) domain.AcademicSession {
	s := domain.AcademicSession{
		Name:      req.GetName(),
		StartDate: fromTimestamp(req.GetStartDate()),
		EndDate:   fromTimestamp(req.GetEndDate()),
		IsActive:  req.GetIsActive(),
	}
	s.InstituteID = parseUUID(req.GetInstituteId())
	return s
}

func academicSessionUpdateFromProto(req *pb.UpdateAcademicSessionRequest) domain.AcademicSession {
	s := domain.AcademicSession{
		Name:      req.GetName(),
		StartDate: fromTimestamp(req.GetStartDate()),
		EndDate:   fromTimestamp(req.GetEndDate()),
		IsActive:  req.GetIsActive(),
	}
	s.ID = parseUUID(req.GetId())
	s.InstituteID = parseUUID(req.GetInstituteId())
	return s
}

func academicSessionToProto(d *domain.AcademicSession) *pb.AcademicSessionResponse {
	if d == nil {
		return nil
	}
	return &pb.AcademicSessionResponse{
		Id:          d.ID.String(),
		InstituteId: d.InstituteID.String(),
		Name:        d.Name,
		StartDate:   toTimestamp(d.StartDate),
		EndDate:     toTimestamp(d.EndDate),
		IsActive:    d.IsActive,
		CreatedAt:   toTimestamp(d.CreatedAt),
		UpdatedAt:   toTimestamp(d.UpdatedAt),
	}
}

// ========================= ADDRESS =========================

func addressFromProto(req *pb.CreateAddressRequest) domain.Address {
	return domain.Address{
		OwnerID:      parseUUID(req.GetOwnerId()),
		OwnerType:    domain.OwnerType(req.GetOwnerType()),
		AddressType:  domain.AddressType(req.GetAddressType()),
		AddressLine1: req.GetAddressLine_1(),
		AddressLine2: optionalStringValue(req.GetAddressLine_2()),
		CountryID:    optionalUUIDStringValue(req.GetCountryId()),
		StateID:      optionalUUIDStringValue(req.GetStateId()),
		DistrictID:   optionalUUIDStringValue(req.GetDistrictId()),
		PostalCode:   optionalStringValue(req.GetPostalCode()),
	}
}

func addressToProto(d *domain.Address) *pb.AddressResponse {
	if d == nil {
		return nil
	}
	return &pb.AddressResponse{
		Id:            d.ID.String(),
		OwnerId:       d.OwnerID.String(),
		OwnerType:     string(d.OwnerType),
		AddressType:   string(d.AddressType),
		AddressLine_1: d.AddressLine1,
		AddressLine_2: toStringValue(d.AddressLine2),
		CountryId:     toUUIDStringValue(d.CountryID),
		StateId:       toUUIDStringValue(d.StateID),
		DistrictId:    toUUIDStringValue(d.DistrictID),
		PostalCode:    toStringValue(d.PostalCode),
		CreatedAt:     toTimestamp(d.CreatedAt),
		UpdatedAt:     toTimestamp(d.UpdatedAt),
	}
}
