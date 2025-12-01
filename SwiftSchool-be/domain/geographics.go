package domain

import "github.com/google/uuid"

type Country struct {
	BaseUUIDModel
	Name                string  `json:"name" db:"name"`
	Code                *string `json:"code,omitempty" db:"code"`
	IsoCode2            *string `json:"iso_code_2,omitempty" db:"iso_code_2"`
	IsoCode3            *string `json:"iso_code_3,omitempty" db:"iso_code_3"`
	DefaultCurrencyCode *string `json:"default_currency_code,omitempty" db:"default_currency_code"`
}

type State struct {
	BaseUUIDModel
	CountryID uuid.UUID `json:"country_id" db:"country_id"`
	Name      string    `json:"name" db:"name"`
	Code      *string   `json:"code,omitempty" db:"code"`
}

type District struct {
	BaseUUIDModel
	StateID uuid.UUID `json:"state_id" db:"state_id"`
	Name    string    `json:"name" db:"name"`
	Code    *string   `json:"code,omitempty" db:"code"`
}
