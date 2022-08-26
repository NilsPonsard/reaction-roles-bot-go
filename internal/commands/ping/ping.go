package ping

import (
	"github.com/bwmarrin/discordgo"
	"github.com/nilsponsard/reaction-roles-bot-go/internal/commands/common"
)

var command = &discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Ping the bot",
}

func execute(s *discordgo.Session, i *discordgo.InteractionCreate) {
}

var Summary = common.Summary{
	Name:    "ping",
	Command: command,
	Execute: execute,
}
