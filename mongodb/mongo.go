package mongodb

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func initClient() error {
	_, err := mgo.DialWithTimeout("mongodb://mongodb:27017", 30*time.Second)
	if err != nil {
		fmt.Println("Failed to dial mongo. ", err.Error())
		return err
	}
	return nil
}
