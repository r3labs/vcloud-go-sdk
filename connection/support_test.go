/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package connection

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func handleget(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func handlepost(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(201)
	w.Write(data)
}

func handleput(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.WriteHeader(202)
	w.Write(data)
}

func handledelete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"ok"}`))
}

func handleauth(w http.ResponseWriter, r *http.Request) {
	var req map[string]string

	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &req)

	if req["username"] != "user" && req["password"] != "pass" {
		w.WriteHeader(403)
		return
	}

	w.Write([]byte(`{"token":"test-token"}`))
}

func testhandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleget(w, r)
	case "POST":
		handlepost(w, r)
	case "PUT":
		handleput(w, r)
	case "DELETE":
		handledelete(w, r)
	}
}
