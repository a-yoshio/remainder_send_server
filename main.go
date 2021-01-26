package main // mainパッケージであることを宣言

import (
	"context"
	"fmt" // fmtモジュールをインポート
	"log"

	firebase "firebase.google.com/go"
)

func sendToToken(app *firebase.App) {
	ctx := context.Background()
	_, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}
	fmt.Println("OK!!")
}

func main() { // 最初に実行されるmain()関数を定義
	// opt := option.WithCredentialsFile("remainder-send-server-2ca730bad69e.json")
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		fmt.Errorf("ERROR!", err)
	}
	sendToToken(app)
}
