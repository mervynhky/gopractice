package main

import (
	"net/http"
)

func login(resp http.ResponseWriter, req *http.Request) {
	var redirectTarget = "/"
	name := req.FormValue("username")
	password := req.FormValue("password")
	if name != "" && password != ""{
		setSession(name,resp)
		redirectTarget="/templates/hometest.html"
	}
	http.Redirect(resp, req, redirectTarget, 302)
}

func logout(){

}