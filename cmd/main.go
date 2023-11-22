package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	jira "github.com/DMXMax/jira-subsystem"
	"github.com/DMXMax/orchestrator"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	log.Print("Hello World")

	ticket := jira.Ticket{
		Id: "SA-2222",
	}

	log.Print(ticket)

	orchestrator := orchestrator.New()

	orchestrator.OnTicket(ticket)

}
