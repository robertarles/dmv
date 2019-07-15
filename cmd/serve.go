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
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "stand up a DMV http server",
	Long:  `the DMV server allows registration and deletion of host info`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Serve() {
	type callerRecord struct {
		hostname   string
		remoteAddr string
		requestURI string
	}

	// var registrationList = make([]callerRecord, 1)

	var callerRegister = make(map[string]callerRecord)

	r := mux.NewRouter()

	// add a host registration
	r.HandleFunc("/register/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		caller := callerRecord{}
		response := ""
		caller.hostname = vars["hostname"]
		caller.remoteAddr = strings.Split(r.RemoteAddr, ":")[0]
		caller.requestURI = r.RequestURI
		callerRegister[caller.hostname] = caller
		//response = fmt.Sprintf("current caller: %s, %s, %s\n", caller.hostname, caller.remoteAddr, caller.requestURI)

		response += fmt.Sprintf("%+v", callerRegister)
		fmt.Printf("Test %+v\n", r)
		fmt.Fprintf(w, response)
	}).Methods("GET", "PUT", "POST")

	// remove a host registration
	r.HandleFunc("/register/{hostname}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		response := ""
		delete(callerRegister, vars["hostname"])
		response += fmt.Sprintf("%+v", callerRegister)
		fmt.Printf("Test %+v\n", r)
		fmt.Fprintf(w, response)
	}).Methods("DELETE")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Test %+v\n", r)
		response := fmt.Sprintf("%+v", callerRegister)
		fmt.Fprintf(w, response)
	}).Methods("GET")

	// fs := http.FileServer(http.Dir("static/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", r)

}
