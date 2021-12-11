package main

import (
	"log"
	"net/http"
)

// main メイン関数
func main() {
	// HTTPリクエスト
	/*
		[WARN]
		"/"以下のリクエストにもマッチする。例) /favicon.ico
		ブラウザからアクセスすると、2回アクセスされたように見える
	*/
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
<html>
	<head><title>チャット</title></head>
	<body>チャットしましょう！</body>
</html>
`))
	})

	// WEBサーバの開始(port: 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenServer: %v", err)
	}
}
