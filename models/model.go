package models

import (
	"strconv"
	"time"
)

type model struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *model) idToString() string {
	return strconv.Itoa(m.ID)
}

func (m *model) createdAtToString() string {
	return m.CreatedAt.String()
}

func (m *model) updatedAtToString() string {
	return m.UpdatedAt.String()
}

func (m *model) serialize() map[string]string {
	return map[string]string{
		"id":         m.idToString(),
		"created_at": m.createdAtToString(),
		"updated_at": m.updatedAtToString(),
	}
}

func (m *model) unserialize(data map[string]string) {
	m.ID, _ = strconv.Atoi(data["id"])
	m.CreatedAt, _ = time.Parse(time.RFC3339, data["created_at"])
	m.UpdatedAt, _ = time.Parse(time.RFC3339, data["updated_at"])
}
