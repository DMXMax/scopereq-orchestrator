package orchestrator

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/DMXMax/evaluator-subsystem"
	jira "github.com/DMXMax/jira-subsystem"
	_ "github.com/DMXMax/jira-subsystem/region"
	"github.com/DMXMax/jira-subsystem/system"
)

type Orchestrator struct{}

func New() *Orchestrator {
	log.Info().Msg("Orchestrator Created")
	o := new(Orchestrator)

	//o.

	return o
}

func (o *Orchestrator) OnTicket(ticket jira.Ticket) {
	log.Info().Str("recieved ticket", ticket.Id).Send()
	vt, err := system.ValidateAndRetrieveTicket(ticket)

	if err != nil {
		HandleError(err)
	}

	_ = vt

	results, err := evaluator.EvaluateTicket(vt)

	if err != nil {
		HandleError(err)
	}

	log.Info().Interface("Evaluation Results", results).Send()
	log.Info().Int("Request Results", int(results.GetAction())).Send()
}

func HandleError(err error) {
	log.Error().Err(err).Msg("failure")
	os.Exit(-1)

}
