package models

import (
	"github.com/go-bongo/bongo"
	"github.com/google/uuid"
)

type Review struct {
	bongo.DocumentBase `bson:",inline"`
	Id                 uuid.UUID
	Reviewer           string
	Rating             uint
	Review             string
}
