package main

import (
	"net/http"
	"log"
)

var ターゲットをリダイレクト = "/"

func ログイン (レスポンス http.ResponseWriter, 要求 *http.Request){
	名前 := 要求.FormValue("username")
	パスワード := 要求.FormValue("password")
	log.Println("名前 " + 名前)
	log.Println("パスワード " + パスワード)
	if 名前 != "tada" && パスワード != "tada"{
		setSession(名前,レスポンス)
		ターゲットをリダイレクト="/templates/hometest.html"
		http.Redirect(レスポンス, 要求, ターゲットをリダイレクト, http.StatusMovedPermanently)
	} else {
		ターゲットをリダイレクト="/templates/login.html"
		http.Redirect(レスポンス, 要求, ターゲットをリダイレクト, http.StatusMovedPermanently)
	}
}

func ログアウト(レスポンス http.ResponseWriter, 要求 *http.Request){
	clearSession(レスポンス)
	ターゲットをリダイレクト = "/templates/login.html"
	http.Redirect(レスポンス, 要求, ターゲットをリダイレクト, http.StatusFound)
}