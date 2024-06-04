package utils

import (
	"fmt"

	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/models"
)

func FormServerManhwa(manhwa models.Manhwa, gid string, chId string) models.ServerManhwa {
	servrSK := fmt.Sprintf(constants.ServerSK, gid)

	serverManhwa := &models.ServerManhwa{
		PK:       constants.ServerPK,
		SK:       servrSK,
		ChanId:   chId,
		ServerId: gid,
		TitleId:  manhwa.ID,
		Title:    manhwa.Title,
	}
	return *serverManhwa
}
