package utils

import (
	"fmt"

	"github.com/jallenmanaloto/soha-bot/internal/constants"
	"github.com/jallenmanaloto/soha-bot/models"
)

func GenerateKey(keyType string, val string) (string, string) {
	var pk string
	var sk string

	if keyType == constants.ServerPK {
		pk = constants.ServerPK
		sk = fmt.Sprintf(constants.ServerSK, val)
	}

	return pk, sk
}

func FormServerManhwa(manhwa models.Manhwa, gid string, chId string) models.ServerManhwa {
	pk, sk := GenerateKey("SERVER", gid)

	serverManhwa := &models.ServerManhwa{
		PK:       pk,
		SK:       sk,
		ChanId:   chId,
		ServerId: gid,
		TitleId:  manhwa.ID,
		Title:    manhwa.Title,
	}
	return *serverManhwa
}
