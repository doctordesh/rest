package main

import (
	"errors"
	"fmt"

	"github.com/doctordesh/rest/models"
)

func authInvite(code string) (*models.Invite, error) {
	fmt.Printf("Code: %+v\n", code)

	// invite := models.Invite{}
	// if invite.findByCode(code) {
	// 	return &invite, nil
	// }
	return nil, errors.New("Invite auth failed")
}
