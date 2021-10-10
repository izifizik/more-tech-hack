package repository

import (
	"log"
	"more-tech-hack/internal/config"
)

func InsertUserByUserId(userId string) error {
	insertStmt := `insert into "users"("userId") values($1)`
	_, err := config.Db.Exec(insertStmt, userId)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//func InsertUserAccessToModel()