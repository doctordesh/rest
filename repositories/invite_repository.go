package repositories

import (
	"database/sql"
	"fmt"
)

type invite_repository struct {
	repository
}

func NewInviteRepository(db *sql.DB) *invite_repository {
	i := new(invite_repository)
	i.DB = db
	return i
}

func (r *repository) Save(model Serializeable) error {
	data := model.Serialize()
	if data["id"] == "0" {
		return r.create(data)
	}
	return r.update(data)
}

func (r *repository) update(data map[string]string) error {
	fmt.Printf("DB: %v\n", r.DB)
	r.DB.Exec(`
    UPDATE
      invites
    SET
      code = ?,
      token = ?
    WHERE
      id = ?
    `, data["code"], data["token"], data["id"])
	return nil
}

func (r *repository) create(data map[string]string) error {
	return nil
}
