package main

import (
	"fmt"

	"github.com/the-go-getter/go-microservices/rectangle"
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
	fmt.Println(rectangle.Area(2, 3))

}
