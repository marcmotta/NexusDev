// cmd/nexusdev/main.go
package main

import (
	"flag"
	"log"
	"os"

	"nexusdev/internal/nexusdev"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := nexusdev.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
