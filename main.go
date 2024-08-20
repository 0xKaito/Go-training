package main

import (
	processor "example/re/middleware"
	"example/re/router"
)

func main() {
	instance, auth := processor.Start()
	router.Start(instance, auth);
}
