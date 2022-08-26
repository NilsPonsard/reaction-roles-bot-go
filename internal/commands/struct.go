package commands

import (
	"github.com/NilsPonsard/verbosity"
	"github.com/bwmarrin/discordgo"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/commands/common"
)

type CommandsHandler struct {
	summaries map[string]*common.Summary
}

func New(summaries []common.Summary) *CommandsHandler {
	summariesMap := make(map[string]*common.Summary)

	for _, summary := range summaries {
		summariesMap[summary.Command.Name] = &summary
	}

	return &CommandsHandler{
		summaries: summariesMap,
	}
}

func (h *CommandsHandler) Register(session *discordgo.Session) (registeredCommands []*discordgo.ApplicationCommand) {
	registeredCommands = make([]*discordgo.ApplicationCommand, len(h.summaries))

	for _, summary := range h.summaries {
		cmd, err := session.ApplicationCommandCreate(session.State.User.ID, "", summary.Command)
		if err != nil {
			verbosity.Fatal(err)

		}
		registeredCommands = append(registeredCommands, cmd)

	}

	return
}

func (h *CommandsHandler) Handle(session *discordgo.Session, interaction *discordgo.InteractionCreate) {

	name := interaction.ApplicationCommandData().Name

	summary, ok := h.summaries[name]
	if !ok {
		verbosity.Debug("Unknown command:", name)
		return
	}
	summary.Execute(session, interaction)

}
