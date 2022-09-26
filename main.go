package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	res, err := SendPushNotification("sad", "test", "ini test")
	fmt.Println(res, err)
}

func SendPushNotification(token string, titleMsg string, bodyMsg string) (response string, err error) {
	ctx := context.Background()

	serviceAccountKeyPath := "svcAccKey.json"

	opt := option.WithCredentialsFile(serviceAccountKeyPath)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return response, err
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return response, err
	}

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: titleMsg,
			Body:  bodyMsg,
		},
		Token: token,
		// Webpush: &messaging.WebpushConfig{
		// 	Notification: &messaging.WebpushNotification{
		// 		Actions: []*messaging.WebpushNotificationAction{

		// 		},
		// 	},
		// },
	}

	response, err = client.Send(ctx, message)
	if err != nil {
		return response, err
	}

	return response, nil
}
