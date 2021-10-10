package repository

import (
	"log"
	"more-tech-hack/intern/config"
)

func UpdateBalance(userId string) float64 {
	updateStmt := `update "users" set "balance"="balance"-$1 where "id"=$2`
	_, err := config.Db.Exec(updateStmt, 10, userId)
	if err != nil {
		log.Println(err)
	}
	user, _ := GetUser(userId)
	return user.Balance
}
