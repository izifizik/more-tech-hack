package repository

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"more-tech-hack/intern/config"
	"more-tech-hack/intern/model"
)

func InsertUser(userId string) error {
	insertStmt := `insert into users(id) values($1)`
	_, err := config.Db.Exec(insertStmt, userId)
	if err != nil {
		return err
	}

	return nil
}

func InsertModel(model *model.Model) error {
	bytes, err := json.Marshal(model.Struct)
	if err != nil {
		return err
	}

	insertSmth := `insert into models(urn, struct, name, is_dataset) values($1, $2, $3, $4)`
	_, err = config.Db.Exec(insertSmth, model.Urn, bytes, model.Name, model.IsDataset)
	if err != nil {
		return err
	}
	return nil
}

func InsertUserAccess(userId string, modelId int) error {
	insertSmth := `insert into user_access(user_id, model_id) values($1, $2)`
	_, err := config.Db.Exec(insertSmth, userId, modelId)
	if err != nil {
		return err
	}
	return nil
}