package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	// Run the app
	app := app.Generate()
	_, erro := fmt.Println(app.Run(os.Args))
	if erro != nil {
		log.Fatal(erro)
	}
}
