package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strings"

// 	"github.com/gorilla/mux"
// )

// type callerRecord struct {
// 	hostname   string
// 	remoteAddr string
// 	requestURI string
// }

// // var registrationList = make([]callerRecord, 1)

// var callerRegister = make(map[string]callerRecord)

// func main() {
// 	r := mux.NewRouter()

// 	// add a host registration
// 	r.HandleFunc("/register/{hostname}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		caller := callerRecord{}
// 		response := ""
// 		caller.hostname = vars["hostname"]
// 		caller.remoteAddr = strings.Split(r.RemoteAddr, ":")[0]
// 		caller.requestURI = r.RequestURI
// 		callerRegister[caller.hostname] = caller
// 		//response = fmt.Sprintf("current caller: %s, %s, %s\n", caller.hostname, caller.remoteAddr, caller.requestURI)

// 		response += fmt.Sprintf("%+v", callerRegister)
// 		fmt.Printf("Test %+v\n", r)
// 		fmt.Fprintf(w, response)
// 	}).Methods("GET", "PUT", "POST")

// 	// remove a host registration
// 	r.HandleFunc("/register/{hostname}", func(w http.ResponseWriter, r *http.Request) {
// 		vars := mux.Vars(r)
// 		response := ""
// 		delete(callerRegister, vars["hostname"])
// 		response += fmt.Sprintf("%+v", callerRegister)
// 		fmt.Printf("Test %+v\n", r)
// 		fmt.Fprintf(w, response)
// 	}).Methods("DELETE")

// 	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("Test %+v\n", r)
// 		response := fmt.Sprintf("%+v", callerRegister)
// 		fmt.Fprintf(w, response)
// 	}).Methods("GET")

// 	// fs := http.FileServer(http.Dir("static/"))
// 	// http.Handle("/static/", http.StripPrefix("/static/", fs))

// 	http.ListenAndServe(":80", r)
// }
