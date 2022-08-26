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
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	})
}

var Summary = common.Summary{
	Command: command,
	Execute: execute,
}
