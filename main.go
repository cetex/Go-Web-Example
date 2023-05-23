package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
)

// Simple hello world
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!\n")
}

// A little bit more complex. Handle POST of form data like 'name=test&address=someaddress'
func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "FORM POST request successful\n")
	fmt.Printf("All form data: %s\n", r.Form)

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

// Same as above, but accept JSON and decode it into a struct.
// https://pkg.go.dev/encoding/json#Marshal - for info abouts `json:...` struct field tags
// Here we map the json string 'name' so it's stored in the struct field 'Name'
type my_struct struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// create new json decoder of Body
	decoder := json.NewDecoder(r.Body)

	// create new copy of struct my_struct named "t"
	var t my_struct
	// Decode json body into struct "t"
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "JSON POST request successful\n")
	fmt.Printf("Json data: %s\n", t)

	fmt.Fprintf(w, "Name = %s\n", t.Name)
	fmt.Fprintf(w, "Address = %s\n", t.Address)
}

// Example getting an IP, json post
// https://pkg.go.dev/net#Resolver

type domain struct {
	Domain string `json:"domain"`
}

func dnsJsonHandler(w http.ResponseWriter, r *http.Request) {
	// create new json decoder of Body
	decoder := json.NewDecoder(r.Body)

	// create new copy of struct my_struct named "t"
	var t domain
	// Decode json body into struct "t"
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "JSON POST request successful\n")
	fmt.Printf("Json data: %s\n", t)

	fmt.Fprintf(w, "Domain = %s\n", t.Domain)

	// Some dns resolution thing stolen from stackoverflow
	// https://stackoverflow.com/questions/59889882/specifying-dns-server-for-lookup-in-go
	resolver := &net.Resolver{
		PreferGo: true,
		//Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
		//    d := net.Dialer{
		//        Timeout: time.Millisecond * time.Duration(10000),
		//    }
		//    return d.DialContext(ctx, network, "8.8.8.8:53")
		//},
	}
	ip, _ := resolver.LookupHost(context.Background(), t.Domain)

	fmt.Fprintf(w, "%s, %s\n", ip, ip[0])
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/dns", dnsJsonHandler)

	fmt.Printf("Starting server at port 8081\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
