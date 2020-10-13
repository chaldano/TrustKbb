package main

import "github.com/gorilla/mux"
import "net/http"
import "strconv"
import "fmt"

var names = make(map[int]string, 3)

func init() {
	fmt.Println("This will get called on main initialization")
	names[1] = "Name1"
	names[2] = "Name2"
	names[3] = "Name3"
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/contacts", contacts)
	r.HandleFunc("/contacts/{id:[0-9]+}", getContact).Methods("GET")
	http.ListenAndServe(":8080", r)

}

func contacts(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Astrid, Kalle, Penny, Kbb"))
}
func getContact(writer http.ResponseWriter, request *http.Request) {
	v := mux.Vars(request)
	id, _ := strconv.Atoi(v["id"])
	writer.Write([]byte(names[id]))
}
