package main

import (
	"github.com/axellelanca/urlshortener/cmd"

	// 👇 C'est ICI la clé !
	// Le "_" dit à Go : "Charge ces dossiers pour exécuter leur fonction init(), même si je n'utilise pas leurs variables."
	_ "github.com/axellelanca/urlshortener/cmd/cli"
	_ "github.com/axellelanca/urlshortener/cmd/server"
)

func main() {
	cmd.Execute()
}
