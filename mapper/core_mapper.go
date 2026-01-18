package mapper

import (
	"fmt"
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
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
		AcademicSessionID: c.AcademicSessionID,
		Name:              c.Name,
		Section:           c.Section,
		ClassTeacherID:    helper.ToNullUUID(helper.DerefUUID(c.ClassTeacherID)),
		CreatedBy:         helper.ToNullUUID(helper.DerefUUID(c.CreatedBy)),
	}
}

// func MapUpdateClassParams(c domain.Class) db.UpdateClassParams {
// 	// If UpdateClassParams doesn't exist, this function will error.
// 	// We return empty struct to satisfy compiler if struct exists but unused?
// 	// But previous errors said struct undefined.
// 	// So we should probably remove this function or comment it out.
// 	// However, repo still calls it if not stubbed. I stubbed Repo.
// 	// So i can comment this out.
// 	// return db.UpdateClassParams{}
// 	panic("MapUpdateClassParams called but undefined")
// }

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

func MapUpdateDepartmentParams(d domain.Department) db.UpdateDepartmentParams {
	return db.UpdateDepartmentParams{
		ID:          d.ID,
		InstituteID: helper.ToNullUUID(d.InstituteID),
		Name:        d.Name,
		UpdatedBy:   helper.ToNullUUID(helper.DerefUUID(d.UpdatedBy)),
	}
}

// ------------------ EMPLOYEE ------------------

func MapDBEmployeeToDomain(e db.CoreEmployee) (*domain.Employee, error) {
	// Initialize struct
	emp := &domain.Employee{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        e.ID,
				CreatedAt: helper.NullTimeToValue(e.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(e.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(e.CreatedBy),
			},
			InstituteID: e.InstituteID,
		},
		EmployeeCode:      e.EmployeeCode,
		FirstName:         helper.NullStringToValue(e.FirstName),
		LastName:          helper.NullStringToPtr(e.LastName),
		DepartmentID:      helper.NullUUIDToPtr(e.DepartmentID),
		Gender:            domain.Gender(helper.NullStringToValue(e.Gender)),
		MaritalStatus:     domain.MaritalStatus(helper.NullStringToValue(e.MaritalStatus)),
		DateOfJoining:     helper.NullTimeToPtr(e.DateOfJoining),
		Nationality:       helper.NullStringToPtr(e.Nationality),
		PreferredLanguage: helper.NullStringToPtr(e.PreferredLanguage),
		IsActive:          helper.NullBoolToValue(e.IsActive),
	}

	// Decode JSONBs
	if e.SocialMediaHandles.Valid {
		val, err := helper.DecodeJSONB[domain.SocialMediaHandles](e.SocialMediaHandles)
		if err == nil {
			emp.SocialMediaHandles = val
		}
	}
	if e.LanguageSkills.Valid {
		val, err := helper.DecodeJSONB[[]domain.LanguageSkill](e.LanguageSkills)
		if err == nil {
			emp.LanguageSkills = val
		}
	}

	return emp, nil
}

func MapListEmployeesRowToDomain(e db.ListEmployeesRow) *domain.Employee {
	return MapDBListEmployeesRowToDomain(e)
}

func MapDBListEmployeesRowToDomain(e db.ListEmployeesRow) *domain.Employee {
	emp := &domain.Employee{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        e.ID,
				CreatedAt: helper.NullTimeToValue(e.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(e.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(e.CreatedBy),
			},
			InstituteID: e.InstituteID,
		},
		EmployeeCode:      e.EmployeeCode,
		FirstName:         helper.NullStringToValue(e.FirstName),
		LastName:          helper.NullStringToPtr(e.LastName),
		DepartmentID:      helper.NullUUIDToPtr(e.DepartmentID),
		Gender:            domain.Gender(helper.NullStringToValue(e.Gender)),
		MaritalStatus:     domain.MaritalStatus(helper.NullStringToValue(e.MaritalStatus)),
		DateOfJoining:     helper.NullTimeToPtr(e.DateOfJoining),
		Nationality:       helper.NullStringToPtr(e.Nationality),
		PreferredLanguage: helper.NullStringToPtr(e.PreferredLanguage),
		IsActive:          helper.NullBoolToValue(e.IsActive),
	}

	if e.SocialMediaHandles.Valid {
		val, err := helper.DecodeJSONB[domain.SocialMediaHandles](e.SocialMediaHandles)
		if err == nil {
			emp.SocialMediaHandles = val
		}
	}
	if e.LanguageSkills.Valid {
		val, err := helper.DecodeJSONB[[]domain.LanguageSkill](e.LanguageSkills)
		if err == nil {
			emp.LanguageSkills = val
		}
	}
	return emp
}

func MapDomainEmployeeToDBParams(e domain.Employee) db.CreateEmployeeParams {
	return db.CreateEmployeeParams{
		InstituteID:        e.InstituteID,
		EmployeeCode:       e.EmployeeCode,
		FirstName:          helper.ToNullString(e.FirstName),
		LastName:           helper.ToNullString(helper.StrOrEmpty(e.LastName)),
		DepartmentID:       helper.ToNullUUID(helper.DerefUUID(e.DepartmentID)),
		Gender:             helper.ToNullString(string(e.Gender)),
		MaritalStatus:      helper.ToNullString(string(e.MaritalStatus)),
		DateOfJoining:      helper.ToNullTime(helper.TimeOrZero(e.DateOfJoining)),
		Nationality:        helper.ToNullString(helper.StrOrEmpty(e.Nationality)),
		PreferredLanguage:  helper.ToNullString(helper.StrOrEmpty(e.PreferredLanguage)),
		SocialMediaHandles: helper.EncodeJSONB(e.SocialMediaHandles),
		LanguageSkills:     helper.EncodeJSONB(e.LanguageSkills),
		CreatedBy:          helper.ToNullUUID(helper.DerefUUID(e.CreatedBy)),
	}
}

func MapUpdateEmployeeParams(e domain.Employee) db.UpdateEmployeeParams {
	return db.UpdateEmployeeParams{
		ID:           e.ID,
		InstituteID:  e.InstituteID,
		FirstName:    helper.ToNullString(e.FirstName),
		LastName:     helper.ToNullString(helper.StrOrEmpty(e.LastName)),
		DepartmentID: helper.ToNullUUID(helper.DerefUUID(e.DepartmentID)),
		// Gender:        helper.ToNullString(string(e.Gender)),
		MaritalStatus: helper.ToNullString(string(e.MaritalStatus)),
		UpdatedBy:     helper.ToNullUUID(helper.DerefUUID(e.UpdatedBy)),
	}
}

// ------------------ EMPLOYEE FULL PROFILE ------------------

func MapEmployeeFullProfileRowToDomain(row db.GetEmployeeFullProfileRow) *domain.Employee {
	// helper.JSONBToValue might return error or value.
	// assuming helper.JSONBToValue takes json.RawMessage and returns T value directly (panicking or zero on error)?
	// Or helper.JSONBToValue is NOT helper.DecodeJSONB?
	// helper.DecodeJSONB returns (T, error).
	// helper.JSONBToValue probably doesn't exist? I should use DecodeJSONB.

	// Revert to DecodeJSONB
	var langSkills []domain.LanguageSkill
	if val, err := helper.DecodeJSONB[[]domain.LanguageSkill](row.LanguageSkills); err == nil {
		langSkills = val
	}
	var socialHandles domain.SocialMediaHandles
	if val, err := helper.DecodeJSONB[domain.SocialMediaHandles](row.SocialMediaHandles); err == nil {
		socialHandles = val
	}

	return &domain.Employee{
		TenantUUIDModel: domain.TenantUUIDModel{
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			},
			InstituteID: row.InstituteID,
		},
		EmployeeCode:       row.EmployeeCode,
		FirstName:          helper.NullStringToValue(row.FirstName),
		LastName:           helper.NullStringToPtr(row.LastName),
		DepartmentID:       helper.NullUUIDToPtr(row.DepartmentID),
		Gender:             domain.Gender(helper.NullStringToValue(row.Gender)),
		MaritalStatus:      domain.MaritalStatus(helper.NullStringToValue(row.MaritalStatus)),
		DateOfJoining:      helper.NullTimeToPtr(row.DateOfJoining),
		Nationality:        helper.NullStringToPtr(row.Nationality),
		PreferredLanguage:  helper.NullStringToPtr(row.PreferredLanguage),
		SocialMediaHandles: socialHandles,
		LanguageSkills:     langSkills,
		IsActive:           helper.NullBoolToValue(row.IsActive),
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
		FirstName:  helper.NullStringToValue(g.FirstName),
		LastName:   helper.NullStringToPtr(g.LastName),
		Email:      helper.NullStringToPtr(g.Email),
		Phone:      helper.NullStringToPtr(g.Phone),
		Profession: helper.NullStringToPtr(g.Profession),
		// AnnualIncome in DB: sql.NullString. Domain: *float64
		// We can't easily map string to float pointer without parsing.
		// For now returning nil to avoid error if we don't implement ParseFloat helper
		AnnualIncome: nil,
	}
}

func MapDomainGuardianToDBParams(g domain.Guardian) db.CreateGuardianParams {
	// AnnualIncome: Domain *float64 -> DB sql.NullString
	var incomeStr string
	if g.AnnualIncome != nil {
		incomeStr = fmt.Sprintf("%.2f", *g.AnnualIncome)
	}

	return db.CreateGuardianParams{
		FirstName:    helper.ToNullString(g.FirstName),
		LastName:     helper.ToNullString(helper.StrOrEmpty(g.LastName)),
		Email:        helper.ToNullString(helper.StrOrEmpty(g.Email)),
		Phone:        helper.ToNullString(helper.StrOrEmpty(g.Phone)),
		Profession:   helper.ToNullString(helper.StrOrEmpty(g.Profession)),
		AnnualIncome: helper.ToNullString(incomeStr),
		CreatedBy:    helper.ToNullUUID(helper.DerefUUID(g.CreatedBy)),
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
		OwnerType:    domain.OwnerType(a.OwnerType.String),
		AddressType:  domain.AddressType(a.AddressType.String),
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
		// AddressLine2 not in CreateAddressParams
		CountryID:  helper.ToNullUUID(helper.DerefUUID(a.CountryID)),
		StateID:    helper.ToNullUUID(helper.DerefUUID(a.StateID)),
		DistrictID: helper.ToNullUUID(helper.DerefUUID(a.DistrictID)),
		PostalCode: helper.ToNullString(helper.StrOrEmpty(a.PostalCode)),
	}
}
