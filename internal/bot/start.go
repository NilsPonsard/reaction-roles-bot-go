package bot

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
)

var (
	ErrNoToken = errors.New("no token in env")
)

func Connect() (session *discordgo.Session, err error) {

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		err = ErrNoToken
		return
	}

	session, err = discordgo.New("Bot " + token)
	if err != nil {
		return
	}
	err = session.Open()

	return
}
