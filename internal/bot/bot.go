package bot

import (
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/bot/utils"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
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
	case constants.Bury:
		Bury(s, m, message)
	case constants.Command, constants.Tricks:
		Tricks(s, m)
	case constants.Fetch:
		Fetch(s, m)
	case constants.Hello:
		Hello(s, m)
	case constants.Look:
		Look(s, m, message)
	case constants.Watch:
		Watch(s, m, message)
	default:
		Default(s, m)
	}
}

func (b *DiscordBot) SendUpdate(manhwaId string, titleCh string, bot *DiscordBot) {
	serverManhwas, err := database.SearchSubscribedToManhwa(manhwaId)
	if err != nil {
		logger.Log.Error(err)
	}

	for _, server := range serverManhwas {
		keys := constants.Keys{
			PK: server.PK,
			SK: server.SK,
		}
		_, err := database.UpdateServerManhwaCh(keys, titleCh)
		if err != nil {
			logger.Log.Errorf(constants.ErrorUpdateItem, err)
		}

		chanId := server.ChanId
		image := utils.ExtractUrl(server.TitleImage)
		thumbnail := utils.EmbedThumbnail(image)
		embed := utils.EmbedManhwa(server.TitleId, titleCh, server.Title, server.TitleUrl, thumbnail)

		_, err = bot.Session.ChannelMessageSend(chanId, constants.MessageFoundNewCh)
		if err != nil {
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}

		_, err = bot.Session.ChannelMessageSendEmbed(chanId, &embed)
		if err != nil {
			_, err := bot.Session.ChannelMessageSend(chanId, constants.ErrorDiscordMessageSend)
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}
	}
}
