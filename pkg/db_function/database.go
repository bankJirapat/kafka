package database_function

import (
	"context"
	"kafka/pkg/model"
	"kafka/types"

	"github.com/sirupsen/logrus"
)

func (dbf DatabaseFunction) InsertValueToDatabase(ctx context.Context, data types.Dtests) error {
	logrus.Warn("InsertValueToDatabase")
	dTests := model.Tests{
		Topic: data.Topic,
		Key:   data.Key,
		Value: data.Value,
		Time:  data.Time,
	}
	_, err := dbf.db.NewInsert().Model(&dTests).On("conflict (id) do update").Exec(ctx)
	if err != nil {
		return nil
	}
	return nil
}
