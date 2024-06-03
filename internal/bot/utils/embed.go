package utils

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func EmbedThumbnail(image string) discordgo.MessageEmbedThumbnail {
	thumbnail := discordgo.MessageEmbedThumbnail{
		URL:    image,
		Height: 450,
		Width:  286,
	}
	return thumbnail
}

func EmbedManhwa(chapter string, name string, url string, thumbnail discordgo.MessageEmbedThumbnail) discordgo.MessageEmbed {
	manhwa := discordgo.MessageEmbed{
		Title:       strings.ToTitle(name),
		URL:         url,
		Description: chapter,
		Image:       (*discordgo.MessageEmbedImage)(&thumbnail),
	}
	return manhwa
}

func ManhwaImage(url string) string {
	parts := strings.Split(url, "/")
	image := parts[len(parts)-1]
	return image
}
