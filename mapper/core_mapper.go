package mapper

import (
	"database/sql"
	"encoding/json"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"

	"github.com/google/uuid"
)

// ------------------ INSTITUTE ------------------

func MapDBInstituteToDomain(i db.CoreInstitute) *domain.Institute {
	return &domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{
			ID:        i.ID,
			CreatedAt: i.CreatedAt.Time,
			UpdatedAt: i.UpdatedAt.Time,
		},
		Name:         i.Name,
		Code:         i.Code,
		CurrencyCode: helper.NullStringToPtr(i.CurrencyCode),
		LogoURL:      helper.NullStringToPtr(i.LogoUrl),
		Website:      helper.NullStringToPtr(i.Website),
		IsActive:     i.IsActive.Bool,
	}
}

func MapDomainInstituteToDBParams(i domain.Institute) db.CreateInstituteParams {
	return db.CreateInstituteParams{
		Name:         i.Name,
		Code:         i.Code,
		CurrencyCode: helper.PtrToNullString(i.CurrencyCode),
		LogoUrl:      helper.PtrToNullString(i.LogoURL),
		Website:      helper.PtrToNullString(i.Website),
		IsActive:     helper.BoolToNullBool(i.IsActive),
	}
}

// ------------------ CLASS ------------------

func MapDBClassToDomain(c db.CoreClass) domain.Class {
	return domain.Class{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        c.ID, // map ID from NullUUID
				CreatedAt: c.CreatedAt.Time,
				UpdatedAt: c.UpdatedAt.Time,
				CreatedBy: helper.UUIDToPtr(c.CreatedBy.UUID),
			},
			InstituteID: c.InstituteID,
		},
		Name:              c.Name,
		Section:           c.Section,
		AcademicSessionID: c.AcademicSessionID,
		ClassTeacherID:    helper.NullUUIDToPtr(c.ClassTeacherID),
	}
}

func MapDomainClassToDBParams(c domain.Class) db.CreateClassParams {
	return db.CreateClassParams{
		InstituteID:       c.InstituteID,
		Name:              c.Name,
		AcademicSessionID: c.AcademicSessionID,
		Section:           c.Section,
		ClassTeacherID:    helper.PtrToNullUUID(c.ClassTeacherID),
		CreatedBy:         helper.PtrToNullUUID(c.CreatedBy),
	}
}

// ------------------ ACADEMIC SESSION ------------------

func MapDBAcademicSessionToDomain(s db.CoreAcademicSession) domain.AcademicSession {
	return domain.AcademicSession{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        s.ID, // map ID from NullUUID
				CreatedAt: s.CreatedAt.Time,
				UpdatedAt: s.UpdatedAt.Time,
				CreatedBy: helper.UUIDToPtr(s.CreatedBy.UUID),
			},
			InstituteID: s.InstituteID,
		},
		Name:      s.Name,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		IsActive:  s.IsActive.Bool,
	}
}

func MapDomainAcademicSessionToDBParams(s domain.AcademicSession) db.CreateAcademicSessionParams {
	return db.CreateAcademicSessionParams{
		InstituteID: s.InstituteID,
		Name:        s.Name,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
		IsActive:    helper.BoolToNullBool(s.IsActive),
	}
}

// ------------------ DEPARTMENT ------------------

func MapDBDepartmentToDomain(d db.CoreDepartment) domain.Department {
	return domain.Department{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        d.ID, // map ID from NullUUID
				CreatedAt: d.CreatedAt.Time,
				UpdatedAt: d.UpdatedAt.Time,
				CreatedBy: helper.UUIDToPtr(d.CreatedBy.UUID),
			},
			InstituteID: helper.NullUUIDToUUID(d.InstituteID),
		},
		Name: d.Name,
	}
}

func MapDomainDepartmentToDBParams(d domain.Department) db.CreateDepartmentParams {
	return db.CreateDepartmentParams{
		InstituteID: d.InstituteID,
		Name:        d.Name,
	}
}

// ------------------ EMPLOYEE ------------------

func MapDBEmployeeToDomain(e db.CoreEmployee) (*domain.Employee, error) {
	langSkills, _ := helper.DecodeJSONB[domain.LanguageSkill](e.LanguageSkills)
	var socialHandles domain.SocialMediaHandles
	if e.SocialMediaHandles.Valid {
		_ = json.Unmarshal(e.SocialMediaHandles.RawMessage, &socialHandles)
	}
	return &domain.Employee{
		TenantUUIDModel: domain.TenantUUIDModel{
			ID:          e.ID,
			InstituteID: e.InstituteID,
			CreatedAt:   e.CreatedAt.Time,
			UpdatedAt:   e.UpdatedAt.Time,
		},
		EmployeeCode:       e.EmployeeCode,
		FirstName:          e.FirstName,
		LastName:           helper.ToStr(&e.LastName.String),
		DepartmentID:       helper.UUIDToString(&e.DepartmentID.UUID),
		Gender:             domain.Gender(e.Gender.String),
		MaritalStatus:      domain.MaritalStatus(e.MaritalStatus.String),
		DateOfJoining:      e.DateOfJoining.Time,
		Nationality:        helper.ToStr(&e.Nationality.String),
		PreferredLanguage:  helper.ToStr(&e.PreferredLanguage.String),
		SocialMediaHandles: socialHandles,
		LanguageSkills:     langSkills,
		IsActive:           e.IsActive.Bool,
	}, nil
}

func MapDomainEmployeeToDBParams(e domain.Employee) db.CreateEmployeeParams {
	return db.CreateEmployeeParams{
		InstituteID:        e.InstituteID,
		EmployeeCode:       e.EmployeeCode,
		FirstName:          e.FirstName,
		LastName:           helper.ToNullString(e.LastName),
		DepartmentID:       helper.ToNullUUID(e.DepartmentID),
		Gender:             sql.NullString{String: string(e.Gender), Valid: true},
		MaritalStatus:      sql.NullString{String: string(e.MaritalStatus), Valid: true},
		DateOfJoining:      helper.ToNullTime(e.DateOfJoining),
		Nationality:        helper.ToNullString(e.Nationality),
		PreferredLanguage:  helper.ToNullString(e.PreferredLanguage),
		SocialMediaHandles: helper.EncodeJSONB(e.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(e.LanguageSkills),
		IsActive:           helper.ToNullBool(e.IsActive),
	}
}

// ------------------ STUDENT ------------------

func MapDBStudentToDomain(s db.CoreStudent) (*domain.Student, error) {
	langSkills, _ := helper.DecodeJSONB[domain.LanguageSkill](s.LanguageSkills)
	var socialHandles domain.SocialMediaHandles
	if s.SocialMediaHandles.Valid {
		_ = json.Unmarshal(s.SocialMediaHandles.RawMessage, &socialHandles)
	}
	return &domain.Student{
		TenantUUIDModel: domain.TenantUUIDModel{
			ID:          s.ID,
			InstituteID: s.InstituteID,
			CreatedAt:   s.CreatedAt.Time,
			UpdatedAt:   s.UpdatedAt.Time,
		},
		AdmissionNo:        s.AdmissionNo,
		FirstName:          s.FirstName,
		LastName:           helper.ToStr(&s.LastName.String),
		DOB:                s.Dob.Time,
		Gender:             domain.Gender(s.Gender.String),
		BloodGroup:         domain.BloodGroup(s.BloodGroup.String),
		SocialCategory:     domain.SocialCategory(s.SocialCategory.String),
		CurrentClassID:     helper.UUIDToString(&s.CurrentClassID.UUID),
		Nationality:        helper.ToStr(&s.Nationality.String),
		PreferredLanguage:  helper.ToStr(&s.PreferredLanguage.String),
		SocialMediaHandles: socialHandles,
		LanguageSkills:     langSkills,
	}, nil
}

func MapDomainStudentToDBParams(s domain.Student) db.CreateStudentParams {
	return db.CreateStudentParams{
		InstituteID:        s.InstituteID,
		AdmissionNo:        s.AdmissionNo,
		FirstName:          s.FirstName,
		LastName:           helper.ToNullString(s.LastName),
		Dob:                helper.ToNullTime(s.DOB),
		Gender:             sql.NullString{String: string(s.Gender), Valid: true},
		BloodGroup:         sql.NullString{String: string(s.BloodGroup), Valid: true},
		SocialCategory:     sql.NullString{String: string(s.SocialCategory), Valid: true},
		CurrentClassID:     helper.ToNullUUID(s.CurrentClassID),
		Nationality:        helper.ToNullString(s.Nationality),
		PreferredLanguage:  helper.ToNullString(s.PreferredLanguage),
		SocialMediaHandles: helper.EncodeJSONB(s.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(s.LanguageSkills),
	}
}

// ------------------ GUARDIAN ------------------

func MapDBGuardianToDomain(g db.CoreGuardian) domain.Guardian {
	return domain.Guardian{
		TenantUUIDModel: domain.TenantUUIDModel{
			ID:          g.ID,
			InstituteID: g.InstituteID,
			CreatedAt:   g.CreatedAt.Time,
			UpdatedAt:   g.UpdatedAt.Time,
		},
		FirstName: g.FirstName,
		LastName:  helper.ToStr(&g.LastName.String),
		Phone:     helper.ToStr(&g.Phone.String),
		Email:     helper.ToStr(&g.Email.String),
		IsActive:  g.IsActive.Bool,
	}
}

func MapDomainGuardianToDBParams(g domain.Guardian) db.CreateGuardianParams {
	return db.CreateGuardianParams{
		InstituteID: g.InstituteID,
		FirstName:   g.FirstName,
		LastName:    helper.ToNullString(g.LastName),
		Phone:       helper.ToNullString(g.Phone),
		Email:       helper.ToNullString(g.Email),
		IsActive:    helper.ToNullBool(g.IsActive),
	}
}

// ------------------ ADDRESS ------------------

func MapDBAddressToDomain(a db.CoreAddress) domain.Address {
	return domain.Address{
		BaseUUIDModel: domain.BaseUUIDModel{
			ID:        a.ID,
			CreatedAt: a.CreatedAt.Time,
			UpdatedAt: a.UpdatedAt.Time,
		},
		OwnerID:      a.OwnerID,
		OwnerType:    domain.OwnerType(a.OwnerType.String),
		AddressType:  domain.AddressType(a.AddressType.String),
		AddressLine1: a.AddressLine1,
		AddressLine2: helper.ToStr(&a.AddressLine2.String),
		CountryID:    helper.UUIDToString(&a.CountryID.UUID),
		StateID:      helper.UUIDToString(&a.StateID.UUID),
		DistrictID:   helper.UUIDToString(&a.DistrictID.UUID),
		PostalCode:   helper.ToStr(&a.PostalCode.String),
	}
}

func MapDomainAddressToDBParams(a domain.Address) db.CreateAddressParams {
	return db.CreateAddressParams{
		OwnerID:      a.OwnerID,
		OwnerType:    sql.NullString{String: string(a.OwnerType), Valid: true},
		AddressType:  sql.NullString{String: string(a.AddressType), Valid: true},
		AddressLine1: a.AddressLine1,
		AddressLine2: helper.ToNullString(a.AddressLine2),
		CountryID:    helper.ToNullUUID(a.CountryID),
		StateID:      helper.ToNullUUID(a.StateID),
		DistrictID:   helper.ToNullUUID(a.DistrictID),
		PostalCode:   helper.ToNullString(a.PostalCode),
	}
}

// ------------------ STUDENT-GUARDIAN LINK ------------------

func MapDomainLinkStudentGuardianToDBParams(studentID, guardianID uuid.UUID) db.LinkStudentGuardianParams {
	return db.LinkStudentGuardianParams{
		StudentID:  studentID,
		GuardianID: guardianID,
	}
}
