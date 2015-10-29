package main

import (
	"errors"
	"fmt"
)

func authInvite(code string) (*Invite, error) {
	fmt.Printf("Code: %+v\n", code)

	invite := Invite{}
	if invite.findByCode(code) {
		return &invite, nil
	}
	return nil, errors.New("Invite auth failed")
}
