//package main
//
//import (
//	"net/http"
////"github.com/gorilla/mux"
//	"github.com/gorilla/securecookie"
//	"io/ioutil"
//	"strings"
//	"log"
//)
//
//var cookieHandler = securecookie.New(
//	securecookie.GenerateRandomKey(64),
//	securecookie.GenerateRandomKey(32))
//var redirectTarget = ""
////var router = mux.NewRouter()
////var redirectTarget string = ""
////in this case, MyHandler is private, if I wanted it to be public,
////i'd use myHandler. Notice the case of the first character.
//type MyHandler struct {
//}
////
////type userData struct{
////    username string
////    password string
////}
//
////attaches the method ServeHTTP to the MyHandler struct.
////if written like func ServeHTTP(blabla){} , it is a function,
//// but with the (this *MyHandler) it becomes a method of any instance of  MyHandler
//// so, MyHandler receives the method ServeHTTP
//func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	path := r.URL.Path[1:]
//	log.Println(path)
//	data, err := ioutil.ReadFile(string(path))
//
//	if err == nil {
//		var contentType string
//
//		if strings.HasSuffix(path, ".css") {
//			contentType = "text/css"
//		} else if strings.HasSuffix(path, ".html") {
//			contentType = "text/html"
//		} else if strings.HasSuffix(path, ".js") {
//			contentType = "application/javascript"
//		} else if strings.HasSuffix(path, ".png") {
//			contentType = "image/png"
//		} else if strings.HasSuffix(path, ".svg") {
//			contentType = "image/svg+xml"
//		} else {
//			contentType = "text/plain"
//		}
//
//		w.Header().Add("Content Type", contentType)
//		w.Write(data)
//	} else {
//		w.WriteHeader(404)
//		w.Write([]byte("404 Mi amigo - " + http.StatusText(404)))
//	}
//}
////need to add logouthandler too, and all it does is clear existing sesssion and redirect to login
///////////////////////////////////////////////////////////////////////////////////////
///// loginHandler reads the name and password from the submitted form,             ///
///// then if the credientials pass the sophisticated check,                        ///
//// the username is stored in a session, then a redirect to the homepage is sent.  ///
///// if it fails the check, clear any existing session and redirect to login page. ///
///// Will add a "loginfailed" page in future.                                      ///
///////////////////////////////////////////////////////////////////////////////////////
//func loginHandler (resp http.ResponseWriter, req *http.Request){
//	//    _ = req.ParseForm()
//	//
//	//    name := req.Form.Get("userName")
//	//    password := req.Form.Get("passWord")
//
//	name := req.FormValue("username")
//	password := req.FormValue("password")
//	if name != "" && password != ""{
//		setSession(name,resp)
//		redirectTarget="/templates/home.html"
//	}
//	http.Redirect(resp, req, redirectTarget, 302)
//}
//
/////////////////////////////////////////////////////////////////////////////////////
///// setSession puts the username into a simple string map.                      ///
///// then use securecookie(cookieHandler) to ENCODE the value map,               ///
///// then encrypted session value is stored in a standard http.Cookie instance.  ///
/////////////////////////////////////////////////////////////////////////////////////
//func setSession (username string, resp http.ResponseWriter){
//	value := map[string]string{
//		"username" : username,
//	}
//	if encoded, err := cookieHandler.Encode("session", value); err == nil{
//		cookie := &http.Cookie{
//			Name: "session",
//			Value: encoded,
//			Path: "/",
//		}
//		http.SetCookie(resp, cookie)
//	}
//}
//
/////////////////////////////////////////////////////////////////////////////////////
///// getUserName implements sequence other way around. cookie read from request, ///
///// then use securecookie(cookieHandler) to DECODE the cookie value,            ///
///// then result is string mapped and username returned                          ///
/////////////////////////////////////////////////////////////////////////////////////
//func getUserName(req *http.Request) (userName string){
//	if cookie, err := req.Cookie("session"); err == nil {
//		cookieValue := make(map[string]string)
//		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil{
//			userName = cookieValue["username"]
//		}
//	}
//	return userName
//}
//////////////////////////////////////////////////////////////////////
///// clearSession is to delete the current session by             ///
///// setting a negative value for maxage.                         ///
///// this is to delete session information from client            ///
//////////////////////////////////////////////////////////////////////
//func clearSession (resp http.ResponseWriter){
//	cookie := &http.Cookie{
//		Name: "session",
//		Value: "",
//		Path: "/",
//		MaxAge: -1,
//	}
//	http.SetCookie(resp, cookie)
//}
//
//
//
//func main() {
//
//	http.HandleFunc("/templates/login", loginHandler)
//	//    router.HandleFunc("/templates/login", loginHandler).Methods("POST")
//	//http.HandleFunc("/templates/login.html", loginHandler)
//	//    http.Handle("/", new(MyHandler))
//	http.Handle("/", new(MyHandler))
//	http.ListenAndServe(":7998", nil)
//}
//
///*
////    router.HandleFunc("/home", internalPageHandler)
////    router.HandleFunc("/logout", logoutHandler).Methods("POST")
////    router.HandleFunc("/templates/login", loginHandler).Methods("POST")
//*/