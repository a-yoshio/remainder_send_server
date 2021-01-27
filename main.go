package main // mainパッケージであることを宣言

import (
	"context"
	"fmt" // fmtモジュールをインポート
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	// This registration token comes from the client FCM SDKs.
	registrationToken := ""

	// See documentation on defining a message payload.
	message := &messaging.Message{
		Data: map[string]string{
			"title":    "Remainder",
			"contents": "起きる時間だよ！",
			"tag":      "家",
		},
		Token: registrationToken,
	}

	// Send a message to the device corresponding to the provided
	// registration token.
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	// Response is a message ID string.
	fmt.Println("Successfully sent message:", response)
}

func main() { // 最初に実行されるmain()関数を定義
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		fmt.Errorf("ERROR!", err)
	}
	sendToToken(app)
}
