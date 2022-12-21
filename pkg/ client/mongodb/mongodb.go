package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewClient(ctx context.Context, host, pots, username, password, database, authDB string) (db *mongo.Database, err error) {
	//Если нет авторизации, то есть username и password пуствые
	var mongoDBURL string
	if username == "" && password == "" {
		mongoDBURL = "mongdb://&%s:%s"

	} else {
		mongoDBURL = "mongdb:/%s:%s@%s:%s"
	}
	fmt.Println(mongoDBURL)
	//Contex нам наверное передали с таймаутом,
	return db, err
}
