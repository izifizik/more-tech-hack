package repository

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"more-tech-hack/internal/config"
	"more-tech-hack/internal/model"
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