package orm

import (
	"log"
	"semantiker/orm/ent"
)

func GetConnection() *ent.Client {
	client, err := ent.Open("mysql", "root:nopass@tcp(localhost:3307)/sem?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	//defer client.Close()

	return client
}
