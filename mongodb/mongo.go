package mongodb

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
)

func initClient() error {
	fmt.Println("LOCALJET: Will dial mongo")
	_, err := mgo.DialWithTimeout("mongodb://mongodb:27017", 30*time.Second)
	if err != nil {
		fmt.Println("LOCALJET: Failed to dial mongo. ", err.Error())
		return err
	}
	fmt.Println("LOCALJET: Succeeded to dial mongo")
	return nil
}
