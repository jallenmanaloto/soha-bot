package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func Default(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, constants.MessageDefault)
	if err != nil {
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}

func Hello(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, constants.MessageHello)
	if err != nil {
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}

func Look(s *discordgo.Session, m *discordgo.MessageCreate, param []string) {
	p := param[2:]
	title := strings.Join(p, " ")
	database.SearchManhwa(title)

	manhwas, err := database.SearchManhwa(title)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.DiscordUnexpectedHandler)
		if err != nil {
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}
	}
	logger.Log.Infof("Manhwa values: %v\n", manhwas)
}

func Tricks(s *discordgo.Session, m *discordgo.MessageCreate) {
	look := "`!soha look <title>:` Soha to look for a manhwa with that title."
	fetch := "`!soha fetch:` Soha will fetch all the titles he is watching for any latest chapters."
	watch := "`!soha watch <title>:` Soha will watch out for new chapters for the title."
	tricks := fmt.Sprintf(
		"Soha's tricks or command displays things he can do.\nYou can call out to Soha with `!soha` followed by your command\n\n%s\n\n%s\n\n%s",
		look,
		fetch,
		watch,
	)

	embed := discordgo.MessageEmbed{
		Title:       "**Soha's tricks and quirks**",
		Description: tricks,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}
