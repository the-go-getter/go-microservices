package main

import (
	"fmt"
	"net/http"
)

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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
