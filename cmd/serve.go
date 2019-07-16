/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var httpPort uint = 80
var usingPort string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "stand up a DMV http server",
	Long:  `the DMV server allows registration and deletion of host info`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().UintVarP(&httpPort, "port", "p", 80, "Listen on this HTTP Port")
}

var callerRegister = make(map[string]map[string]string)

func serve() {

	// var registrationList = make([]callerRecord, 1)

	r := mux.NewRouter()

	// add a host registration
	r.HandleFunc("/register/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		caller := make(map[string]string)
		response := ""
		caller["hostname"] = vars["hostname"]
		caller["remoteAddr"] = strings.Split(r.RemoteAddr, ":")[0]
		caller["requestURI"] = r.RequestURI
		callerRegister[caller["hostname"]] = caller
		//response = fmt.Sprintf("current caller: %s, %s, %s\n", caller.hostname, caller.remoteAddr, caller.requestURI)

		response += fmt.Sprintf("%+v", callerRegister)
		fmt.Printf("Test %+v\n", r)
		fmt.Fprintf(w, response)
	}).Methods("GET", "PUT", "POST")

	// remove a host registration
	r.HandleFunc("/delete/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		response := ""
		fmt.Println(vars["hostname"])
		delete(callerRegister, vars["hostname"])
		response += fmt.Sprintf("%+v", callerRegister)
		fmt.Printf("Test %+v\n", r)
		fmt.Fprintf(w, response)
	}).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Test %+v\n", r)
		jsonString, err := json.Marshal(callerRegister)
		if err != nil {
			panic(err)
		}
		fmt.Printf("DEBUG -- jsonString %+v\n", string(jsonString))
		fmt.Fprintf(w, string(jsonString))
	}).Methods("GET")

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	usingPort = fmt.Sprintf(":%d", httpPort)
	fmt.Printf("starting server on port %s\n", usingPort)
	http.ListenAndServe(usingPort, r)

}
