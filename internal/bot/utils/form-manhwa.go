package utils

import (
	"fmt"
	"strings"

	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/models"
)

func FormServerManhwa(manhwa models.Manhwa, gid string, chId string) models.ServerManhwa {
	title := strings.ReplaceAll(manhwa.Title, " ", "_")
	servrSK := fmt.Sprintf(constants.ServerSK, title)

	serverManhwa := &models.ServerManhwa{
		PK:       constants.ServerPK,
		SK:       servrSK,
		ChanId:   chId,
		ServerId: gid,
	}
	return *serverManhwa
}
