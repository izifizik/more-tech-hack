package repository

import (
	"encoding/json"
	"fmt"
	"more-tech-hack/internal/config"
	"more-tech-hack/internal/model"
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
		_, err = fmt.Scan(&modelId, &modelName, &modelUrn, &modelIsDataset)
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

func GetModel(userId string) ([]model.Model, error) {
	rows, err := config.Db.Query(`select m.id, m.urn, m.name, m.struct, m.is_dataset from user_access ua inner join models m on m.id = ua.model_id where user_id=$1`, userId)
	if err != nil {
		return nil, err
	}

	var modelVar []model.Model

	var modelId int
	var modelName, modelUrn string
	var modelIsDataset bool
	var modelStruct model.Json
	var modelBytes []byte

	for rows.Next() {
		_, err = fmt.Scan(&modelId, &modelUrn, &modelName, &modelBytes, &modelIsDataset)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(modelBytes, &modelStruct)
		if err != nil {
			return nil, err
		}

		model_var := model.Model{
			Id:        modelId,
			Urn:       modelUrn,
			Name:      modelName,
			IsDataset: modelIsDataset,
			Struct:    modelStruct,
		}

		modelVar = append(modelVar, model_var)
	}

	return modelVar, nil
}
