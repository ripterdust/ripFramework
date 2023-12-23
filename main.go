package main

import (
	ripframework "ripframework/ripFramework"
	"ripframework/ripFramework/requester"
	"ripframework/ripFramework/responser"
)

func getHandler(request *requester.Requester, response *responser.Responser) {

	response.Json("Request")
}

func main() {
	app := ripframework.New()
	port := "2023"

	app.Get("/uri", getHandler)

	app.Listen(port)
}
