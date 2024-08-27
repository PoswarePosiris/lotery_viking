package models

type Users struct {
	BaseModel
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

// GetEmail returns the model email.
func (m *Users) GetEmail() string {
	return m.Email
}

// Check password
func (m *Users) CheckPassword(password string) bool {
	return m.Password == password
}
