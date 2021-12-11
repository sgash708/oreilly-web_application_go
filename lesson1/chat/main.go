package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// templateHandler テンプレートを扱う構造体
type templateHandler struct {
	once     sync.Once
	filename string
	temp1    *template.Template
}

/*
	REF: https://pkg.go.dev/net/http#Handler
	HTTPリクエストの処理のために、ServerHTTPメソッドだけ存在していれば良い。
*/
// ServerHTTP HTTPリクエストの処理
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// REF: https://qiita.com/nirasan/items/2160be0a1d1c7ccb5e65
	t.once.Do(func() {
		t.temp1 = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	if err := t.temp1.Execute(w, nil); err != nil {
		log.Fatalln(err)
	}
}

// main メイン関数
func main() {
	// [WARN] メソッド呼び出しする際に、sync.Onceは同じものを使う必要がある
	// → レシーバはポインタである必要がある
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// WEBサーバの開始(port: 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("ListenServe: %v", err)
	}
}
