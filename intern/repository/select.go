package repository

import (
	"encoding/json"
	"more-tech-hack/intern/config"
	"more-tech-hack/intern/model"
)

func GetModels(userId string) ([]model.ModelsGet, error) {
	rows, err := config.Db.Query(`SELECT ua.model_id, m.name, m.urn, m.is_dataset FROM user_access ua inner join models m on ua.model_id = m.id where user_id=$1`, userId)

	if err != nil {
		return nil, err
	}

	var modelArr []model.ModelsGet

	var modelId int
	var modelName, modelUrn string
	var modelIsDataset bool

	for rows.Next() {
		err = rows.Scan(&modelId, &modelName, &modelUrn, &modelIsDataset)
		if err != nil {
			return nil, err
		}

		modelStruct := model.ModelsGet{
			ModelId:        modelId,
			ModelName:      modelName,
			ModelUrn:       modelUrn,
			ModelIsDataset: modelIsDataset,
		}

		modelArr = append(modelArr, modelStruct)
	}

	return modelArr, nil
}

func GetModel(id int) (m model.Model, err error) {
	rows, err := config.Db.Query(`SELECT id, urn, name, struct, is_dataset FROM models where id=$1`, id)
	if err != nil {
		return
	}

	var modelId int
	var modelName, modelUrn string
	var modelIsDataset bool
	var modelStruct model.Json
	var modelBytes []byte

	for rows.Next() {
		err = rows.Scan(&modelId, &modelUrn, &modelName, &modelBytes, &modelIsDataset)
		if err != nil {
			return
		}

		err = json.Unmarshal(modelBytes, &modelStruct)
		if err != nil {
			return
		}

		m = model.Model{
			Id:        modelId,
			Urn:       modelUrn,
			Name:      modelName,
			IsDataset: modelIsDataset,
			Struct:    modelStruct,
		}
	}

	return
}

func GetUsersByModelId(modelId int) (users []string, err error) {
	rows, err := config.Db.Query(`select user_id from user_access where model_id=$1`, modelId)
	if err != nil {
		return
	}

	var userID string

	for rows.Next() {
		err = rows.Scan(&userID)
		if err != nil {
			return
		}

		users = append(users, userID)
	}
	return
}
