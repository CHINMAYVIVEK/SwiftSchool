package mapper

import (
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
)

func MapInstituteRowToDomain(row db.CoreInstitute) domain.Institute {
	return domain.Institute{
		BaseUUIDModel: domain.BaseUUIDModel{
			ID:        row.ID,
			CreatedAt: helper.NullTimeToValue(row.CreatedAt),
			UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
			DeletedAt: helper.NullTimeToPtr(row.DeletedAt),
			CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
			UpdatedBy: helper.NullUUIDToPtr(row.UpdatedBy),
		},
		Name:               row.Name,
		Code:               row.Code,
		CurrencyCode:       helper.NullStringToPtr(row.CurrencyCode),
		LogoURL:            helper.NullStringToPtr(row.LogoUrl),
		Website:            helper.NullStringToPtr(row.Website),
		Timezone:           helper.NullStringToValue(row.Timezone),
		FiscalYearStartMon: int(helper.NullInt32ToValue(row.FiscalYearStartMonth)),
		IsActive:           helper.NullBoolToValue(row.IsActive),
	}
}

func MapInstituteDomainToParams(inst domain.Institute) db.CreateInstituteParams {
	return db.CreateInstituteParams{
		Name:                 inst.Name,
		Code:                 inst.Code,
		CurrencyCode:         helper.ToNullString(helper.StrOrEmpty(inst.CurrencyCode)),
		LogoUrl:              helper.ToNullString(helper.StrOrEmpty(inst.LogoURL)),
		Website:              helper.ToNullString(helper.StrOrEmpty(inst.Website)),
		Timezone:             helper.ToNullString(inst.Timezone),
		FiscalYearStartMonth: helper.ToNullInt32(int32(inst.FiscalYearStartMon)),
		IsActive:             helper.ToNullBool(inst.IsActive),
		CreatedBy:            helper.ToNullUUID(helper.DerefUUID(inst.CreatedBy)),
	}
}

func MapUpdateInstituteParams(inst domain.Institute) db.UpdateInstituteParams {
	return db.UpdateInstituteParams{
		ID:   inst.ID,
		Name: inst.Name,
		// CurrencyCode:         helper.ToNullString(helper.StrOrEmpty(inst.CurrencyCode)),
		LogoUrl:              helper.ToNullString(helper.StrOrEmpty(inst.LogoURL)),
		Website:              helper.ToNullString(helper.StrOrEmpty(inst.Website)),
		Timezone:             helper.ToNullString(inst.Timezone),
		FiscalYearStartMonth: helper.ToNullInt32(int32(inst.FiscalYearStartMon)),
		UpdatedBy:            helper.ToNullUUID(helper.DerefUUID(inst.UpdatedBy)),
	}
}
