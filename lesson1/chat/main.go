package main

import (
	"log"
	"net/http"
)

// main メイン関数
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<html>
			<head><title>チャット</title></head>
			<body>チャットしましょう！</body>
		</html>
		`))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenServer: %v", err)
	}
}
