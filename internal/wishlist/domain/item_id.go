package domain

import (
	"github.com/google/uuid"
)

type ItemId string

func NewItemId() ItemId {
	return ItemId(uuid.New().String())
}

func (id ItemId) String() string {
	return string(id)
}

func (id ItemId) equals(otherId ItemId) bool {
	return id.String() == otherId.String()
}
