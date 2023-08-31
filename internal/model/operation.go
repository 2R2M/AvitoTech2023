package model

import (
	"github.com/google/uuid"
)

type Operation struct {
	ID        string
	UserId    string
	Segment   []string
	ExpiredAt string
}

func NewOperation(user string, segment []string, expiredAt string) *Operation {
	return &Operation{
		ID:        uuid.New().String(),
		UserId:    user,
		Segment:   segment,
		ExpiredAt: expiredAt,
	}
}
