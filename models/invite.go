package models

type Invite struct {
	model
	Code  string `json:"code"`
	Token string `json:"token"`
}

type Invites []Invite

func (i *Invite) Serialize() map[string]string {
	m := i.serialize()
	m["code"] = i.Code
	m["token"] = i.Token

	return m
}

func (i *Invite) Unserialize(data map[string]string) {
	i.unserialize(data)
	i.Code = data["code"]
	i.Token = data["token"]
}
