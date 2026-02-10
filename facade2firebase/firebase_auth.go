package facade2firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

func getFirebaseAuthClient(ctx context.Context) (client *auth.Client, err error) {
	var app *firebase.App
	if app, err = getFirebaseApp(ctx); err != nil {
		return nil, fmt.Errorf("failed to initializing *firebase.App: %w", err)
	}
	if client, err = app.Auth(ctx); err != nil {
		return nil, fmt.Errorf("failed to initializing Firebase *auth.Client: %w", err)
	}
	return
}
