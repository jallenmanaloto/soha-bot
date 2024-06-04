package utils

import (
	"fmt"

	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/models"
	"github.com/jallenmanaloto/soha-bot/pkg/logger"
)

func GenerateKey(keyType string, val string, uid string) (string, string) {
	var pk string
	var sk string

	if keyType == constants.ServerPK {
		pk = constants.ServerPK
		sk = fmt.Sprintf(constants.ServerManhwaSK, val, uid)
	}

	return pk, sk
}

func FormServerManhwa(manhwa models.Manhwa, gid string, chId string) models.ServerManhwa {
	pk, sk := GenerateKey("SERVER", gid, manhwa.ID)
	id, err := GenerateId()
	if err != nil {
		logger.Log.Errorf(constants.ErrorGenerateId, err)
	}

	serverManhwa := &models.ServerManhwa{
		PK:         pk,
		SK:         sk,
		ID:         id,
		ChanId:     chId,
		ServerId:   gid,
		TitleId:    manhwa.ID,
		Title:      manhwa.Title,
		TitleCh:    manhwa.Chapters,
		TitleImage: manhwa.Image,
		TitleUrl:   manhwa.Url,
	}
	return *serverManhwa
}
