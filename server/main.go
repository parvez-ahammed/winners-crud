package main

import (
	"fmt"
	"log"
)

func main() {

	app := SetupApp()

	port := 4000

	fmt.Printf("Serving on port %d...\n", port)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
