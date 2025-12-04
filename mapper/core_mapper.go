package mapper

import (
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
			CreatedAt: helper.NullTimeToValue(i.CreatedAt),
			UpdatedAt: helper.NullTimeToValue(i.UpdatedAt),
			CreatedBy: helper.NullUUIDToPtr(i.CreatedBy),
		},
		Name:         i.Name,
		Code:         i.Code,
		CurrencyCode: helper.NullStringToPtr(i.CurrencyCode),
		LogoURL:      helper.NullStringToPtr(i.LogoUrl),
		Website:      helper.NullStringToPtr(i.Website),
		IsActive:     helper.NullBoolToValue(i.IsActive),
	}
}

func MapDomainInstituteToDBParams(i domain.Institute) db.CreateInstituteParams {
	return db.CreateInstituteParams{
		Name:         i.Name,
		Code:         i.Code,
		CurrencyCode: helper.ToNullString(helper.StrOrEmpty(i.CurrencyCode)),
		LogoUrl:      helper.ToNullString(helper.StrOrEmpty(i.LogoURL)),
		Website:      helper.ToNullString(helper.StrOrEmpty(i.Website)),
		IsActive:     helper.ToNullBool(i.IsActive),
		CreatedBy:    helper.ToNullUUID(helper.DerefUUID(i.CreatedBy)),
	}
}

// ------------------ CLASS ------------------

func MapDBClassToDomain(c db.CoreClass) domain.Class {
	return domain.Class{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        c.ID,
				CreatedAt: helper.NullTimeToValue(c.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(c.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(c.CreatedBy),
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
		Section:           c.Section,
		AcademicSessionID: c.AcademicSessionID,
		ClassTeacherID:    helper.ToNullUUID(helper.DerefUUID(c.ClassTeacherID)),
		CreatedBy:         helper.ToNullUUID(helper.DerefUUID(c.CreatedBy)),
	}
}

// ------------------ ACADEMIC SESSION ------------------

func MapDBAcademicSessionToDomain(s db.CoreAcademicSession) domain.AcademicSession {
	return domain.AcademicSession{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        s.ID,
				CreatedAt: helper.NullTimeToValue(s.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(s.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(s.CreatedBy),
			},
			InstituteID: s.InstituteID,
		},
		Name:      s.Name,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		IsActive:  helper.NullBoolToValue(s.IsActive),
	}
}

func MapDomainAcademicSessionToDBParams(s domain.AcademicSession) db.CreateAcademicSessionParams {
	return db.CreateAcademicSessionParams{
		InstituteID: s.InstituteID,
		Name:        s.Name,
		StartDate:   s.StartDate,
		EndDate:     s.EndDate,
		IsActive:    helper.ToNullBool(s.IsActive),
		CreatedBy:   helper.ToNullUUID(helper.DerefUUID(s.CreatedBy)),
	}
}

// ------------------ DEPARTMENT ------------------

func MapDBDepartmentToDomain(d db.CoreDepartment) domain.Department {
	return domain.Department{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        d.ID,
				CreatedAt: helper.NullTimeToValue(d.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(d.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(d.CreatedBy),
			},
			InstituteID: helper.NullUUIDToValue(d.InstituteID),
		},
		Name: d.Name,
	}
}

func MapDomainDepartmentToDBParams(d domain.Department) db.CreateDepartmentParams {
	return db.CreateDepartmentParams{
		InstituteID: helper.ToNullUUID(d.InstituteID),
		Name:        d.Name,
		CreatedBy:   helper.ToNullUUID(helper.DerefUUID(d.CreatedBy)),
	}
}

// ------------------ EMPLOYEE ------------------

func MapDBEmployeeToDomain(e db.CoreEmployee) (*domain.Employee, error) {
	langSkills, _ := helper.DecodeJSONB[[]domain.LanguageSkill](e.LanguageSkills)
	socialHandles, _ := helper.DecodeJSONB[domain.SocialMediaHandles](e.SocialMediaHandles)

	return &domain.Employee{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        e.ID,
				CreatedAt: helper.NullTimeToValue(e.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(e.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(e.CreatedBy),
			},
			InstituteID: e.InstituteID,
		},
		EmployeeCode:       e.EmployeeCode,
		FirstName:          e.FirstName,
		LastName:           helper.NullStringToPtr(e.LastName),
		DepartmentID:       helper.NullUUIDToPtr(e.DepartmentID),
		Gender:             domain.Gender(helper.NullStringToValue(e.Gender)),
		MaritalStatus:      domain.MaritalStatus(helper.NullStringToValue(e.MaritalStatus)),
		DateOfJoining:      helper.NullTimeToPtr(e.DateOfJoining),
		Nationality:        helper.NullStringToPtr(e.Nationality),
		PreferredLanguage:  helper.NullStringToPtr(e.PreferredLanguage),
		SocialMediaHandles: socialHandles,
		LanguageSkills:     langSkills,
		IsActive:           helper.NullBoolToValue(e.IsActive),
	}, nil
}

func MapDomainEmployeeToDBParams(e domain.Employee) db.CreateEmployeeParams {
	return db.CreateEmployeeParams{
		InstituteID:        e.InstituteID,
		EmployeeCode:       e.EmployeeCode,
		FirstName:          e.FirstName,
		LastName:           helper.ToNullString(helper.StrOrEmpty(e.LastName)),
		DepartmentID:       helper.ToNullUUID(helper.DerefUUID(e.DepartmentID)),
		Gender:             helper.ToNullString(string(e.Gender)),
		MaritalStatus:      helper.ToNullString(string(e.MaritalStatus)),
		DateOfJoining:      helper.ToNullTime(helper.DerefTime(e.DateOfJoining)),
		Nationality:        helper.ToNullString(helper.StrOrEmpty(e.Nationality)),
		PreferredLanguage:  helper.ToNullString(helper.StrOrEmpty(e.PreferredLanguage)),
		SocialMediaHandles: helper.EncodeJSONB(e.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(e.LanguageSkills),
		// IsActive:           helper.ToNullBool(e.IsActive),
		CreatedBy: helper.ToNullUUID(helper.DerefUUID(e.CreatedBy)),
	}
}

// ------------------ STUDENT ------------------

func MapDBStudentToDomain(s db.CoreStudent) (*domain.Student, error) {
	langSkills, _ := helper.DecodeJSONB[[]domain.LanguageSkill](s.LanguageSkills)
	socialHandles, _ := helper.DecodeJSONB[domain.SocialMediaHandles](s.SocialMediaHandles)

	return &domain.Student{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        s.ID,
				CreatedAt: helper.NullTimeToValue(s.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(s.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(s.CreatedBy),
			},
			InstituteID: s.InstituteID,
		},
		AdmissionNo:        s.AdmissionNo,
		FirstName:          s.FirstName,
		LastName:           helper.NullStringToPtr(s.LastName),
		DOB:                helper.NullTimeToPtr(s.Dob),
		Gender:             domain.Gender(helper.NullStringToValue(s.Gender)),
		BloodGroup:         domain.BloodGroup(helper.NullStringToValue(s.BloodGroup)),
		SocialCategory:     domain.SocialCategory(helper.NullStringToValue(s.SocialCategory)),
		CurrentClassID:     helper.NullUUIDToPtr(s.CurrentClassID),
		Nationality:        helper.NullStringToPtr(s.Nationality),
		PreferredLanguage:  helper.NullStringToPtr(s.PreferredLanguage),
		SocialMediaHandles: socialHandles,
		LanguageSkills:     langSkills,
	}, nil
}

func MapDomainStudentToDBParams(s domain.Student) db.CreateStudentParams {
	return db.CreateStudentParams{
		InstituteID:        s.InstituteID,
		AdmissionNo:        s.AdmissionNo,
		FirstName:          s.FirstName,
		LastName:           helper.ToNullString(helper.StrOrEmpty(s.LastName)),
		Dob:                helper.ToNullTime(helper.DerefTime(s.DOB)),
		Gender:             helper.ToNullString(string(s.Gender)),
		BloodGroup:         helper.ToNullString(string(s.BloodGroup)),
		SocialCategory:     helper.ToNullString(string(s.SocialCategory)),
		CurrentClassID:     helper.ToNullUUID(helper.DerefUUID(s.CurrentClassID)),
		Nationality:        helper.ToNullString(helper.StrOrEmpty(s.Nationality)),
		PreferredLanguage:  helper.ToNullString(helper.StrOrEmpty(s.PreferredLanguage)),
		SocialMediaHandles: helper.EncodeJSONB(s.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(s.LanguageSkills),
		CreatedBy:          helper.ToNullUUID(helper.DerefUUID(s.CreatedBy)),
	}
}

// ------------------ GUARDIAN ------------------

func MapDBGuardianToDomain(g db.CoreGuardian) domain.Guardian {
	return domain.Guardian{
		BaseUUIDModel: domain.BaseUUIDModel{
			ID:        g.ID,
			CreatedAt: helper.NullTimeToValue(g.CreatedAt),
			UpdatedAt: helper.NullTimeToValue(g.UpdatedAt),
			CreatedBy: helper.NullUUIDToPtr(g.CreatedBy),
		},
		FirstName: g.FirstName,
		LastName:  helper.NullStringToPtr(g.LastName),
		Phone:     helper.NullStringToPtr(g.Phone),
		Email:     helper.NullStringToPtr(g.Email),
		// IsActive:  helper.NullBoolToValue(g.IsActive),
	}
}

func MapDomainGuardianToDBParams(g domain.Guardian) db.CreateGuardianParams {
	return db.CreateGuardianParams{
		FirstName: g.FirstName,
		LastName:  helper.ToNullString(helper.StrOrEmpty(g.LastName)),
		Phone:     helper.ToNullString(helper.StrOrEmpty(g.Phone)),
		Email:     helper.ToNullString(helper.StrOrEmpty(g.Email)),
		// IsActive:    helper.ToNullBool(g.IsActive),
		CreatedBy: helper.ToNullUUID(helper.DerefUUID(g.CreatedBy)),
	}
}

// ------------------ ADDRESS ------------------

func MapDBAddressToDomain(a db.CoreAddress) domain.Address {
	return domain.Address{
		BaseUUIDModel: domain.BaseUUIDModel{
			ID:        a.ID,
			CreatedAt: helper.NullTimeToValue(a.CreatedAt),
			UpdatedAt: helper.NullTimeToValue(a.UpdatedAt),
			CreatedBy: helper.NullUUIDToPtr(a.CreatedBy),
		},
		OwnerID:      a.OwnerID,
		OwnerType:    domain.OwnerType(helper.NullStringToValue(a.OwnerType)),
		AddressType:  domain.AddressType(helper.NullStringToValue(a.AddressType)),
		AddressLine1: a.AddressLine1,
		AddressLine2: helper.NullStringToPtr(a.AddressLine2),
		CountryID:    helper.NullUUIDToPtr(a.CountryID),
		StateID:      helper.NullUUIDToPtr(a.StateID),
		DistrictID:   helper.NullUUIDToPtr(a.DistrictID),
		PostalCode:   helper.NullStringToPtr(a.PostalCode),
	}
}

func MapDomainAddressToDBParams(a domain.Address) db.CreateAddressParams {
	return db.CreateAddressParams{
		OwnerID:      a.OwnerID,
		OwnerType:    helper.ToNullString(string(a.OwnerType)),
		AddressType:  helper.ToNullString(string(a.AddressType)),
		AddressLine1: a.AddressLine1,
		// AddressLine2: helper.ToNullString(helper.StrOrEmpty(a.AddressLine2)),
		CountryID:  helper.ToNullUUID(helper.DerefUUID(a.CountryID)),
		StateID:    helper.ToNullUUID(helper.DerefUUID(a.StateID)),
		DistrictID: helper.ToNullUUID(helper.DerefUUID(a.DistrictID)),
		PostalCode: helper.ToNullString(helper.StrOrEmpty(a.PostalCode)),
	}
}

// ------------------ STUDENT-GUARDIAN LINK ------------------

func MapDomainLinkStudentGuardianToDBParams(studentID, guardianID uuid.UUID) db.LinkStudentGuardianParams {
	return db.LinkStudentGuardianParams{
		StudentID:  studentID,
		GuardianID: guardianID,
	}
}
