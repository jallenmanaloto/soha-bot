package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/bot/utils"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/internal/database"
	"github.com/jallenmanaloto/soha-bot/models"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func Bury(s *discordgo.Session, m *discordgo.MessageCreate, param []string) {
	gid := m.GuildID
	titleId := utils.ExtractTitle(param)
	keys := constants.Keys{
		PK: "SERVER",
		SK: fmt.Sprintf("SERVER#%s##MANHWA#%s", gid, titleId),
	}

	err := database.DeleteServerManhwa(keys)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		return
	}

	_, err = s.ChannelMessageSend(m.ChannelID, constants.MessageDeleteSuccess)
}

func Default(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, constants.MessageDefault)
	if err != nil {
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
	}
}

func Fetch(s *discordgo.Session, m *discordgo.MessageCreate) {
	var serverManhwas []models.ServerManhwa
	var err error
	var message string

	gid := m.GuildID
	keys := constants.Keys{
		PK: "SERVER",
		SK: fmt.Sprintf("SERVER#%s", gid),
	}

	serverManhwas, err = database.SearchServerManhwas(keys)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
		logger.Log.Errorf("%s: %v\n", constants.ErrorDiscordMessage, err)
		return
	}

	if serverManhwas == nil {
		_, _ = s.ChannelMessageSend(m.ChannelID, constants.MessageEmptyWatchList)
		return
	}

	for idx, manhwa := range serverManhwas {
		n := idx + 1
		title := fmt.Sprintf("%d. %s", n, strings.ToTitle(manhwa.Title))
		deets := fmt.Sprintf(
			constants.EmbedManhwaWatchList,
			title,
			manhwa.TitleId,
			manhwa.TitleCh,
			manhwa.TitleUrl,
		)
		message = message + deets
	}

	embed := discordgo.MessageEmbed{
		Title:       constants.MessageWatchListTitle,
		Description: message,
	}
	_, err = s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, constants.ErrorDiscordMessageSend)
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
		constants.MessageBury,
		constants.MessageAlert,
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
	uid := utils.ExtractTitle(param)
	pk, sk := utils.GenerateKey("SERVER", m.GuildID, uid)
	keys := constants.Keys{
		PK: pk,
		SK: sk,
	}

	serverManhwa, err := database.SearchServerManhwasByTitle(keys, "TitleId", uid, constants.EQUALTO)
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
