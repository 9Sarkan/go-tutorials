package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type myElement struct {
	Name    string
	Surname string
	ID      int
}

// FILEPATH data file path
var FILEPATH = "/tmp/go/datafile.gob"

// DATA variable for keep data
var DATA = make(map[string]myElement)

func save() error {
	fmt.Println("saving file ", FILEPATH)
	err := os.Remove(FILEPATH)
	if err != nil {
		fmt.Println(err)
		return err
	}
	file, err := os.Create(FILEPATH)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(DATA)
	if err != nil {
		fmt.Println("Can not save data", err)
		return err
	}
	return nil
}
func load() error {
	file, err := os.Open(FILEPATH)
	if err != nil {
		fmt.Println("can not open file, ", err)
		return err
	}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&DATA)
	if err != nil {
		fmt.Println("Can not decode file, ", err)
		return err
	}
	return nil
}
func lookup(k string) *myElement {
	if k == "" {
		return nil
	}
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	}
	return nil
}
func add(k string, element myElement) bool {
	if k == "" {
		return false
	}
	if lookup(k) == nil {
		DATA[k] = element
		return true
	}
	return false
}
func deleteElement(k string) bool {
	if lookup(k) != nil {
		delete(DATA, k)
		return true
	}
	return false
}
func change(k string, e myElement) bool {
	if k == "" {
		return false
	}
	DATA[k] = e
	return true
}
func print() {
	for k, v := range DATA {
		fmt.Printf("key: %s\tValue: %v\n", k, v)
	}
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("serving: %s for %s\n", r.Host, r.URL.Path)
	myTemplate := template.Must(template.ParseGlob("home.gohtml"))
	myTemplate.ExecuteTemplate(w, "home.gohtml", nil)
}
func listAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Listing the contents of the KV store!")
	fmt.Fprintf(w, "<a href=\"/\" style=\"margin-right: 20px;\">Home sweethome!</a>")
	fmt.Fprintf(w, "<a href=\"/list\" style=\"margin-right: 20px;\">Listall elements!</a>")
	fmt.Fprintf(w, "<a href=\"/change\" style=\"margin-right:20px;\">Change an element!</a>")
	fmt.Fprintf(w, "<a href=\"/insert\" style=\"margin-right:20px;\">Insert new element!</a>")
	fmt.Fprintf(w, "<h1>The contents of the KV store are:</h1>")
	fmt.Fprintf(w, "<ul>")
	for k, v := range DATA {
		fmt.Fprintf(w, "<li><strong>key: %s\t</strong>value: %v</ul>", k, v)
	}
	fmt.Fprintf(w, "</ul>")
}
func changeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving change page")
	tmpl := template.Must(template.ParseGlob("change.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	id, _ := strconv.Atoi(r.FormValue("ID"))
	element := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      id,
	}
	if !change(key, element) {
		fmt.Println("field to update data!")
	} else {
		err := save()
		if err != nil {
			fmt.Println("failed to save file!")
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}
func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("serving insert page")
	tmpl := template.Must(template.ParseGlob("insert.gohtml"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	key := r.FormValue("key")
	id, _ := strconv.Atoi(r.FormValue("ID"))
	element := myElement{
		Name:    r.FormValue("name"),
		Surname: r.FormValue("surname"),
		ID:      id,
	}
	if !add(key, element) {
		fmt.Println("field to insert data!")
	} else {
		err := save()
		if err != nil {
			fmt.Println("failed to save file!")
			return
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}
func main() {
	initial := myElement{
		Name:    "test",
		Surname: "test",
		ID:      1,
	}
	add("init", initial)
	save()
	err := load()
	if err != nil {
		fmt.Println(err)
		return
	}
	port := ":8080"
	http.HandleFunc("/", homePage)
	http.HandleFunc("/change", changeHandler)
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/list", listAll)

	err = http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
