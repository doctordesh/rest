package main

import ()

type Invite struct {
	Model
	Code          string        `json:"code"`
	InviteMembers InviteMembers `json:"invite_members"`
	Token         string        `json:"token"`
}

type Invites []Invite

type InviteMember struct {
	Model
	InviteId  int    `json:"invite_id" sql:"index"`
	IsComing  bool   `json:"is_coming"`
	Name      string `json:"name"`
	Allergies string `json:"allergies"`
	Message   string `json:"message"`
}

type InviteMembers []InviteMember

func (i *Invite) findByCode(code string) bool {
	db, _ := newDb()
	db.Where("code = ?", code).First(i)

	if db.NewRecord(i) == false {
		i.InviteMembers = make(InviteMembers, 0)
		db.Where("invite_id = ?", i.ID).Find(&i.InviteMembers)
		return true
	}

	return false
}
