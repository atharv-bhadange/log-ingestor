package service

import (
	"github.com/atharv-bhadange/log-ingestor/db"
	"github.com/atharv-bhadange/log-ingestor/model"
)

func AddLog(log model.Log) (model.Log, error) {

	tx := db.Db.Create(&log)

	if tx.Error != nil {
		return model.Log{}, tx.Error
	}

	return log, nil
}

func GetLog() ([]model.Log, error) {

	var logs []model.Log

	tx := db.Db.Order("id").Find(&logs)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return logs, nil
}

func GetLogById(id int) (model.Log, error) {
	
	var log model.Log

	tx := db.Db.Where("id = ?", id).First(&log)

	if tx.Error != nil {
		return model.Log{}, tx.Error
	}

	return log, nil
}
