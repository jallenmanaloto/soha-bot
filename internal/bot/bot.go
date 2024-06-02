package bot

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
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
	sess.AddHandler(bot.commands)

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	return bot, nil
}

func (b *DiscordBot) commands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	message := strings.Split(m.Content, " ")
	if message[0] != constants.Prefix {
		return
	}

	command := message[1]
	switch command {
	case constants.Command, constants.Tricks:
		Tricks(s, m)
	case constants.Fetch:
	case constants.Hello:
		Hello(s, m)
	case constants.Look:
	default:
		Default(s, m)
	}
}
