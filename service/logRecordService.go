package service

import (
	"7D-admin/config/database"
	"7D-admin/model"
)

type LogRecordService struct {
}

func (l *LogRecordService) RecordLog(worldName string, action model.Action) {

	db := database.DB

	logRecord := model.LogRecord{}
	logRecord.WorldName = worldName
	logRecord.Action = action
	db.Save(&logRecord)

}

func (l *LogRecordService) GetLastLog(worldName string) *model.LogRecord {

	db := database.DB
	logRecord := model.LogRecord{}
	db.Last(&logRecord)

	return &logRecord
}