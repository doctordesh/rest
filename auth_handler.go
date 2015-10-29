package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 999999))

	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	code, ok := data["code"].(string)

	if ok == false {
		sendError(http.StatusBadRequest, w, "Wrong datatype of code (string required)")
		return
	}

	invite, err := authInvite(code)

	if err != nil {
		sendError(http.StatusUnauthorized, w, err.Error())
		return
	}

	sendJson(w, invite)
}
