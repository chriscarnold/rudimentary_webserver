package main

import (
	"fmt"
	"log"
	"net/http"
)

//function for the form page
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful ")
	name := r.FormValue("walkname")
	address := r.FormValue("date")
	start := r.FormValue("start")
	end := r.FormValue("end")
	meetingpt := r.FormValue("meetingpt")
	distance := r.FormValue("distance")
	leader := r.FormValue("lead")
	comments := r.FormValue("Comments")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Date = %s\n", address)
	fmt.Fprintf(w, "Start time = %s\n", start)
	fmt.Fprintf(w, "End time = %s\n", end)
	fmt.Fprintf(w, "Meeting point = %s\n", meetingpt)
	fmt.Fprintf(w, "Leader = %s\n", leader)
	fmt.Fprintf(w, "Distance = %s km\n", distance)
	fmt.Fprintf(w, "Comments = %s\n", comments)
}

//function for the /hello page
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" { //making sure that users are not able to submit POST
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Howdy doody!")
}

//main function of the webserver, depending on the url entered it will route the user to the form or the home page
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n") // The heart of the web server using the http import
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
