package domain

import (
	"time"

	"github.com/google/uuid"
)

// Corresponds to schema: cafeteria.meal_plans
type MealPlan struct {
	ID           uuid.UUID `json:"id" db:"id"`
	InstituteID  uuid.UUID `json:"institute_id" db:"institute_id"`
	Name         *string   `json:"name,omitempty" db:"name"`
	CostPerMonth *float64  `json:"cost_per_month,omitempty" db:"cost_per_month"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// Corresponds to schema: cafeteria.daily_menu
type DailyMenu struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	Date        *time.Time `json:"date,omitempty" db:"date"`
	MealType    MealType   `json:"meal_type,omitempty" db:"meal_type"`
	Items       *string    `json:"items,omitempty" db:"items"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

// Corresponds to schema: cafeteria.student_subscriptions
type StudentCafeteriaSubscription struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	InstituteID uuid.UUID  `json:"institute_id" db:"institute_id"`
	StudentID   *uuid.UUID `json:"student_id,omitempty" db:"student_id"`
	PlanID      *uuid.UUID `json:"plan_id,omitempty" db:"plan_id"`
	StartDate   *time.Time `json:"start_date,omitempty" db:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty" db:"end_date"`
	IsActive    *bool      `json:"is_active,omitempty" db:"is_active"`
}
