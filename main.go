package main

import (
	"fmt"
	ripframework "ripframework/ripFramework"
)

func getHandler(uri string) interface{} {
	fmt.Println("Handler para:", uri)

	return ""
}

func main() {
	app := ripframework.New()
	port := "2023"

	app.Get("/uri", getHandler)

	app.Listen(port)
}
