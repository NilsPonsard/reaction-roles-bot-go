package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/NilsPonsard/verbosity"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/bot"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/commands"
)

// Version will be set by the script build.sh
var version string

func main() {
	verbosity.Debug("Starting reaction-roles-bot-go ", version)

	session, err := bot.Connect()

	if err != nil {
		verbosity.Fatal(err)
	}

	commands.Configure(session)

	defer session.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop
}
