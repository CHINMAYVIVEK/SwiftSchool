package domain

import (
	"github.com/casbin/casbin/v2"
	"github.com/google/uuid"
)

// PermissionService handles RBAC and ABAC checks
type PermissionService struct {
	enforcer *casbin.Enforcer
}

func NewPermissionService(e *casbin.Enforcer) *PermissionService {
	return &PermissionService{enforcer: e}
}

// EnforceRBAC checks if a user (sub) in an institute (dom) can perform action (act) on resource (obj)
func (s *PermissionService) EnforceRBAC(userID uuid.UUID, instituteID uuid.UUID, resource string, action string) (bool, error) {
	// Casbin treats everything as strings
	// Format: Enforce(UserUUID, InstituteUUID, "finance/invoices", "read")
	return s.enforcer.Enforce(userID.String(), instituteID.String(), resource, action)
}

// EnforceABAC checks ownership or specific attribute rules
// Example: Can this student view THIS specific invoice?
func (s *PermissionService) EnforceInvoiceAccess(user User, invoice Invoice, action string) bool {
	// 1. First, check basic RBAC (Does this user have 'read' access to invoices generally?)
	hasRoleAccess, _ := s.EnforceRBAC(user.ID, *user.InstituteID, "finance_invoices", action)
	if !hasRoleAccess {
		return false
	}

	// 2. Super Admins and Accountants can access ANY invoice in their institute
	if user.RoleType == RoleSuperAdmin || user.RoleType == RoleAccountant {
		return true
	}

	// 3. ABAC Rule: Students/Guardians can ONLY see their own invoices
	if user.RoleType == RoleStudent {
		return invoice.StudentID == user.LinkedEntityID
	}

	// 4. Guardians need to look up if they are mapped to this student (requires extra lookup logic usually)
	// For strict ABAC here, we'd need the mapping passed in.

	return false
}

// AddPolicy adds a permission dynamically (e.g., when a new Institute is created)
// Example: AddPolicy("teacher", "institute-123", "academics_classes", "read")
func (s *PermissionService) AddPolicy(role string, instituteID string, resource string, action string) error {
	_, err := s.enforcer.AddPolicy(role, instituteID, resource, action)
	return err
}

// AssignRole assigns a user to a role within an institute
// Example: AssignRole("user-uuid-123", "teacher", "institute-123")
func (s *PermissionService) AssignRole(userID string, role string, instituteID string) error {
	_, err := s.enforcer.AddGroupingPolicy(userID, role, instituteID)
	return err
}
