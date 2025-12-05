package mapper

import (
	"swiftschool/domain"
	"swiftschool/helper"
	"swiftschool/internal/db"
)

func MapStudentRowToDomain(row db.CoreStudent) domain.Student {
	return domain.Student{
		TenantUUIDModel: domain.TenantUUIDModel{
			InstituteID: row.InstituteID,
			BaseUUIDModel: domain.BaseUUIDModel{
				ID:        row.ID,
				CreatedAt: helper.NullTimeToValue(row.CreatedAt),
				UpdatedAt: helper.NullTimeToValue(row.UpdatedAt),
				DeletedAt: helper.NullTimeToPtr(row.DeletedAt),
				CreatedBy: helper.NullUUIDToPtr(row.CreatedBy),
				UpdatedBy: helper.NullUUIDToPtr(row.UpdatedBy),
			},
		},

		AdmissionNo:        row.AdmissionNo,
		FirstName:          helper.NullStringToValue(row.FirstName),
		LastName:           helper.NullStringToPtr(row.LastName),
		DOB:                helper.NullTimeToPtr(row.Dob),
		Gender:             domain.Gender(helper.NullStringToValue(row.Gender)),
		BloodGroup:         domain.BloodGroup(helper.NullStringToValue(row.BloodGroup)),
		SocialCategory:     domain.SocialCategory(helper.NullStringToValue(row.SocialCategory)),
		CurrentClassID:     helper.NullUUIDToPtr(row.CurrentClassID),
		Nationality:        helper.NullStringToPtr(row.Nationality),
		PreferredLanguage:  helper.NullStringToPtr(row.PreferredLanguage),
		SocialMediaHandles: helper.JSONBToValue[domain.SocialMediaHandles](row.SocialMediaHandles),
		LanguageSkills:     helper.JSONBToValue[[]domain.LanguageSkill](row.LanguageSkills),
	}
}

func MapStudentDomainToParams(s domain.Student) db.CreateStudentParams {
	return db.CreateStudentParams{
		InstituteID:        s.InstituteID,
		AdmissionNo:        s.AdmissionNo,
		FirstName:          helper.ToNullString(s.FirstName),
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
