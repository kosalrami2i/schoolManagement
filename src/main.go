package main

import (
	"datasource"
	"fmt"
)

func main() {
	datasource.InitDB()
	fmt.Println("Hello there")
}
