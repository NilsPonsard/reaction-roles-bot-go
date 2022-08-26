package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/commands/common"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/commands/ping"
)

var summaries = []common.Summary{
	ping.Summary,
}

func Configure(session *discordgo.Session) (commandsHandler *CommandsHandler) {

	commandsHandler = New(summaries)
	commandsHandler.Register(session)
	session.AddHandler(commandsHandler.Handle)

	return
}
