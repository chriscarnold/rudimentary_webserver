package main

import(
	"fmt"
	"log"
	"net/http"
)
//function for the form page
func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err !=nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful ")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}
//function for the /hello page
func helloHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method !="GET"{ //making sure that users are not able to submit POST
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Howdy doody!")
}
//main function of the webserver, depending on the url entered it will route the user to the form or the home page
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n") // The heart of the web server using the http import
	if err :=http.ListenAndServe(":8080", nil); err !=nil { 
		log.Fatal(err)
	}
}