package bot

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/database"
)

type DiscordBot struct {
	Session *discordgo.Session
	db      *database.Service
}

func New(db *database.Service) (*DiscordBot, error) {
	token := os.Getenv("DISCORD_BOT_TOKEN")
	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}

	bot := &DiscordBot{
		Session: sess,
		db:      db,
	}

	// put commands as handler
	// sess.AddHandler()

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	return bot, nil
}
