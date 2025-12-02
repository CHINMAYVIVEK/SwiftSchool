package auth

type UserRole string

const (
	RoleStudent UserRole = "student"
	RoleTeacher UserRole = "teacher"
	RoleParent  UserRole = "parent"
	RoleAdmin   UserRole = "admin"
)

type User struct {
	ID           string
	Email        string
	PasswordHash string
	Role         UserRole
	InstituteID  string
	ClassID      string
	OwnerID      string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type UserClaims struct {
	UserID      string
	Role        UserRole
	InstituteID string
	ClassID     string
	OwnerID     string
}
