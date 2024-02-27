package database

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/alejandrowaiz98/te-eme-backend-beta/models"
)

type Firestore struct {
	client *firestore.Client
}

type FirestoreImplementation interface {
	Register(User models.User) error
	Login(IncomingUser models.User) (models.User, error)
}

func New() (FirestoreImplementation, error) {

	ctx := context.Background()

	c, err := firestore.NewClient(ctx, os.Getenv("firestore_project"))

	if err != nil {
		return nil, fmt.Errorf("Err creating firestore client: %v", err)
	}

	return &Firestore{client: c}, nil
}
