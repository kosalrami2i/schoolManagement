package main

import (
	"configs"
	"datasource"
)

func main() {
	datasource.InitDB()
	configs.StartServer()
}
