package bot

import (
	"fmt"

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
	title := utils.ExtractTitle(param)
	manhwas, err := database.SearchManhwas("Title", title, constants.CONTAINS)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		if err != nil {
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}
	}

	if manhwas == nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.MessageManhwaNotExist)
		if err != nil {
			logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		}
		return
	}

	for _, manhwa := range manhwas {
		image := utils.ExtractUrl(manhwa.Image)
		thumbnail := utils.EmbedThumbnail(image)
		result := utils.EmbedManhwa(manhwa.ID, manhwa.Chapters, manhwa.Title, manhwa.Url, thumbnail)
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &result)
		if err != nil {
			_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
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
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}

func Watch(s *discordgo.Session, m *discordgo.MessageCreate, param []string) {
	pk, sk := utils.GenerateKey("SERVER", m.GuildID)
	keys := constants.Keys{
		PK: pk,
		SK: sk,
	}

	uid := utils.ExtractTitle(param)
	serverManhwa, err := database.SearchServerManhwas(keys, "TitleId", uid, constants.EQUALTO)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Error(err)
		return
	}

	if serverManhwa != nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, constants.WatchAlreadyExist)
		return
	}

	manhwa, err := database.SearchManhwas("ID", uid, constants.EQUALTO)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Error(err)
	}

	servrManhwa := utils.FormServerManhwa(manhwa[0], m.GuildID, m.ChannelID)
	err = database.CreateServerManhwa(servrManhwa)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
	_, err = s.ChannelMessageSend(m.ChannelID, constants.MessageWatchSuccess)
}
