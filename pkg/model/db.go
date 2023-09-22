package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Tests struct {
	bun.BaseModel `bun:"table:public.tests,alias:t"`
	Id            int       `bun:",pk,autoincrement"`
	Topic         string    `bun:"column:topic" json:"topic"`
	Key           string    `bun:"column:key" json:"key"`
	Value         string    `bun:"column:value" json:"value"`
	Time          time.Time `bun:"column:time" json:"time"`
}
