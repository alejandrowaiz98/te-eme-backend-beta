package database

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/alejandrowaiz98/te-eme-backend-beta/models"
	"golang.org/x/crypto/bcrypt"
)

var ctx context.Context
var user models.User

func init() {

	ctx = context.Background()

}

func (f *Firestore) Register(User models.User) error {

	//Previous User.Hash value is the plain-text given by client in frontend

	hashed, _ := bcrypt.GenerateFromPassword([]byte(User.Hash), 12)

	User.Hash = string(hashed)

	_, err := f.client.Collection(os.Getenv("firestore_collection")).NewDoc().Set(context.Background(), User)

	if err != nil {

		return fmt.Errorf("firestore: %v", err)

	}

	return nil

}

func (f *Firestore) Login(IncomingUser models.User) (models.User, error) {

	query := f.client.Collection(os.Getenv("firestore_collection")).Where("Username", "==", IncomingUser.Username)

	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		return user, fmt.Errorf("firestore: %v", err)
	}

	if len(docs) == 0 {
		return user, fmt.Errorf("firestore: %v", err)
	}

	var userExists bool
	var userMap map[string]interface{}

	for _, doc := range docs {

		data := doc.Data()

		if err := bcrypt.CompareHashAndPassword([]byte(data["Hash"].(string)), []byte(IncomingUser.Hash)); err == nil {
			userExists = true
			userMap = data
			break
		}

	}

	if userExists {

		// Convert the map to JSON
		jsonData, err := json.Marshal(userMap)

		if err != nil {
			return user, fmt.Errorf("firestore: %v", err)
		}

		// Convert the JSON to a struct

		err = json.Unmarshal(jsonData, &user)

		if err != nil {
			return user, fmt.Errorf("firestore: %v", err)
		}

		return user, nil

	} else {

		return user, fmt.Errorf("firestore: %v", err)

	}

}
