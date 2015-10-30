package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/doctordesh/rest/middlewares"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	resp := middlewares.NewResponder(w)

	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 999999))

	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	code, ok := data["code"].(string)

	if ok == false {
		resp.SendError(http.StatusBadRequest, "Wrong datatype of code (string required)")
		return
	}

	invite, err := authInvite(code)

	if err != nil {
		resp.SendError(http.StatusUnauthorized, err.Error())
		return
	}

	resp.SendJson(invite)
}
