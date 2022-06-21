package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path,
		r.URL.Query().Get("token"))
}

func main() {
	// x := 8
	// y, z := 5, 13
	// name := "Mercari"
	employeeIds := make(map[string]int)
	employeeIds["Vinay"] = 116
	fmt.Println(employeeIds)
	// fmt.Println(quote.Go())
	// fmt.Println(x, y, z, name)
	// fmt.Println(rectangle.Area(2, 3))
	http.HandleFunc("/", rootHandler)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Web server has started")
	http.ListenAndServe(":80", nil)
}
