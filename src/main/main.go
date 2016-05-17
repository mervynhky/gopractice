package main

import (
	"net/http"
//	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
//	"github.com/julienschmidt/httprouter"
//	"html/template"
	"io/ioutil"
	"strings"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"encoding/json"
)
var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))
//var router = mux.NewRouter()

///////////////////////////////////////////////////////////////////////////
//in this case, MyHandler is private, if I wanted it to be public,	///
//i'd use myHandler. Notice the case of the first character.		///
///////////////////////////////////////////////////////////////////////////
type MyHandler struct {
}

type db struct {
	*sql.DB
}

///////////////////////////////////////////////////////////////////////////////////////////
//attaches the method ServeHTTP to the MyHandler struct.				///
//if written like func ServeHTTP(blabla){} , it is a function,				///
// but with the (this *MyHandler) it becomes a method of any instance of  MyHandler	///
// so, MyHandler receives the method ServeHTTP						///
///////////////////////////////////////////////////////////////////////////////////////////
func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}

		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		http.Redirect(w, r, "/templates/errorpage.html", http.StatusMovedPermanently)
//		w.WriteHeader(404)
//		w.Write([]byte("404 Mi amigo - " + http.StatusText(404)))
	}
}

/////////////////////////////////////////////////////////////////////////////////////
/// loginHandler reads the name and password from the submitted form,             ///
/// then if the credientials pass the sophisticated check,                        ///
/// the username is stored in a session, then a redirect to the homepage is sent. ///
/// if it fails the check, clear any existing session and redirect to login page. ///
/// Will add a "loginfailed" page in future.                                      ///
/////////////////////////////////////////////////////////////////////////////////////
func loginHandler (resp http.ResponseWriter, req *http.Request) {

	name := req.FormValue("username")
	password := req.FormValue("password")
 	log.Println("name is " + name)
	log.Println("pw is " + password)


	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/gotest")
	if err != nil {
		log.Println(err)
	}

	rows, err := db.Query("SELECT * FROM usertable WHERE USERNAME='" + name + "' AND PASSWORD='" + password + "'")
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var nodeid int
		var username string
		var upassword string
		err = rows.Scan(&nodeid, &username, &upassword)
		if err != nil {
			log.Println(err)
		}
		log.Println(nodeid)
		log.Println(username)
		log.Println(upassword)

		if name == username && password == upassword {
			log.Println("it came inside")
			setSession(name, resp)
			//http.Redirect(resp, req, "/templates/hometest.html", http.StatusFound)
			//http.Redirect(resp, req, "/templates/hometest.html", http.StatusMovedPermanently)
			http.Redirect(resp, req, "http://localhost:7998/templates/hometest.html", 302)
		} else {
			http.Redirect(resp, req, "/templates/login.html", 302)
		}
	}
	db.Close()
}

////////////////////////////////////////
/// 		Testing              ///
////////////////////////////////////////
type Label struct {
	Name string
	Posts []Post
}

type Post struct {
	Start int
	Length int
}

func testHandler (resp http.ResponseWriter, req *http.Request) {
	var objVar []Label
	log.Println(json.NewDecoder(req.Body).Decode(&objVar))
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&objVar)
	if err != nil {
		log.Println("error")
	}
	log.Println(objVar)
}

//func rdrLink (resp http.ResponseWriter, req *http.Request) {
//	http.Redirect(resp, req, "/templates/login.html", 302)
//}

///////////////////////////////////////////////////////////////////////////////////
/// setSession puts the username into a simple string map.                      ///
/// then use securecookie(cookieHandler) to ENCODE the value map,               ///
/// then encrypted session value is stored in a standard http.Cookie instance.  ///
///////////////////////////////////////////////////////////////////////////////////
func setSession (username string, resp http.ResponseWriter){
	value := map[string]string{
		"username" : username,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil{
		cookie := &http.Cookie{
			Name: "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(resp, cookie)
	}
}

///////////////////////////////////////////////////////////////////////////////////
/// getUserName implements sequence other way around. cookie read from request, ///
/// then use securecookie(cookieHandler) to DECODE the cookie value,            ///
/// then result is string mapped and username returned                          ///
///////////////////////////////////////////////////////////////////////////////////
func getUserName(req *http.Request) (userName string){
	if cookie, err := req.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil{
			userName = cookieValue["username"]
		}
	}
	return userName
}
////////////////////////////////////////////////////////////////////
/// clearSession is to delete the current session by             ///
/// setting a negative value for maxage.                         ///
/// this is to delete session information from client            ///
////////////////////////////////////////////////////////////////////
func clearSession (resp http.ResponseWriter){
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(resp, cookie)
}



func main() {
	//database//
	//database end//

//	router := httprouter.New()
//	router.GET("/templates/login", loginHandler)
//  /templates/loginverify.html
//	http.HandleFunc("/hometest", rdrLink)
	http.HandleFunc("/loginverify", testHandler)
	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":7998", nil)
}

/*
//    router.HandleFunc("/home", internalPageHandler)
//    router.HandleFunc("/logout", logoutHandler).Methods("POST")
//    router.HandleFunc("/templates/login", loginHandler).Methods("POST")
*/