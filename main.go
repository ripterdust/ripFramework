package main

import ripframework "ripframework/ripFramework"

func main() {
	app := ripframework.New()
	port := "2023"

	app.Listen(port)
}
