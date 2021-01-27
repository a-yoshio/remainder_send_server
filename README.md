# Remainder Send Server
リマインドを配信するバッチサーバー

## How to prepare enviroment?
1. get json key file for FCM from Firebase console. And set it under root folder.

2. set file path.
```
$ export GOOGLE_APPLICATION_CREDENTIALS={please write your json key file path from process No.1}
```

## How to execute?
use go command
```
$ go run main.go
```
