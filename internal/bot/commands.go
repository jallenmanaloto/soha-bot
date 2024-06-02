package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func Default(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSend(m.ChannelID, "I don't know that trick! Just give me treats, you piece of shit!")
	if err != nil {
		logger.Log.Errorf("ERROR unable to send message for command '%s': %v\n", constants.Default, err)
	}
}
func Hello(s *discordgo.Session, m *discordgo.MessageCreate) {
	message := "arf arf!"
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		logger.Log.Errorf("ERROR unable to send message for command '%s': %v\n", constants.Hello, err)
	}
}

func Tricks(s *discordgo.Session, m *discordgo.MessageCreate) {
	look := "`look:` this command followed by a title ex: `!soha look -title-`, will ask Soha to look for a manhwa."
	fetch := "`fetch:` Soha will fetch all the titles he is watching for any latest chapters."
	watch := "`watch:` this command followed by a title ex: `!soha watch -title-`, will ask Soha to watch out for new chapters."
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
		logger.Log.Errorf("ERROR unable to send message for command '%s': %v\n", constants.Tricks, err)
	}
}
