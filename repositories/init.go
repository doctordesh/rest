package repositories

import (
	"database/sql"
)

type Serializeable interface {
	Serialize() map[string]string
}

type Unserializeable interface {
	Unserialize(map[string]string) error
}

type Model interface {
	Serializeable
	Unserializeable
}

type repository struct {
	DB *sql.DB
}
