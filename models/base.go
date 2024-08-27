package models

import "time"

type Model interface {
	GetId() string
	GetCreated() time.Time
	GetUpdated() time.Time
	UpdateUpdatedAt()
}

type BaseModel struct {
	Id        string    `db:"id" json:"id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

// GetId returns the model id.
func (m *BaseModel) GetId() string {
	return m.Id
}

// GetCreated returns the model Created datetime.
func (m *BaseModel) GetCreated() time.Time {
	return m.CreatedAt
}

// GetUpdated returns the model Updated datetime.
func (m *BaseModel) GetUpdated() time.Time {
	return m.UpdatedAt
}

// UpdateUpdatedAt updates the UpdatedAt field with the current time.
func (m *BaseModel) UpdateUpdatedAt() {
	m.UpdatedAt = time.Now()
}

// Save method ?
