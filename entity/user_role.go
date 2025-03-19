package entity


type UserRole struct {
	UserID int `json:"user_id"`
	RoleID int `json:"role_id"`
	User User
	Role Role
}