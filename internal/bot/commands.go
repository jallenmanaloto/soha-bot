package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/bot/utils"
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

	for _, manhwa := range manhwas {
		imgSplit := strings.SplitAfter(manhwa.Image, "(webp)/")
		image := imgSplit[1]
		thumbnail := utils.EmbedThumbnail(image)
		result := utils.EmbedManhwa(manhwa.Chapters, manhwa.Title, manhwa.Url, thumbnail)
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &result)
		if err != nil {
			_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessage)
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}
	}
}

func Tricks(s *discordgo.Session, m *discordgo.MessageCreate) {
	tricks := fmt.Sprintf(
		constants.MessageTricks,
		constants.MessageLook,
		constants.MessageFetch,
		constants.MessageWatch,
	)

	embed := discordgo.MessageEmbed{
		Title:       constants.MessageTricksEmbedTitle,
		Description: tricks,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessage)
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}
