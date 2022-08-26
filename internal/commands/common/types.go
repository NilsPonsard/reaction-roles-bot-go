package common

import "github.com/bwmarrin/discordgo"

type Summary struct {
	Name    string
	Command *discordgo.ApplicationCommand
	Execute func(s *discordgo.Session, i *discordgo.InteractionCreate)
}
