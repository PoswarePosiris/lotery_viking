package models

import (
	"testing"
	"time"
)

func TestUpdateUpdatedAt(t *testing.T) {
    // Create a new BaseModel instance
    model := &BaseModel{
        Id:        "1",
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    // Get the initial UpdatedAt value
    initialUpdatedAt := model.UpdatedAt

    // Wait for a short period of time
    time.Sleep(1 * time.Second)

    // Update the UpdatedAt field
    model.UpdateUpdatedAt()

    // Check if the UpdatedAt field has been updated
    if model.UpdatedAt.Equal(initialUpdatedAt) {
        t.Errorf("Expected UpdatedAt to be updated, but it was not")
    }
}
