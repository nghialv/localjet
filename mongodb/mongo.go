package mongodb

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/mgo.v2"
)

func initClient() error {
	mongoURL := os.Getenv("MONGODB_URL")
	_, err := mgo.DialWithTimeout(mongoURL, 15*time.Second)
	if err != nil {
		fmt.Println("Failed to dial mongo. ", err.Error())
		return err
	}
	return nil
}
