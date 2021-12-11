# chatアプリ
WebSockerの理解に努める

# 実装手順
## Webサーバの準備

```:go
http.HandleFunc("/", func(w htto.ResponseWriter, r *http.Request) {
    w.Write([]byte("<html></html>"))
})
```
上記のようにHTMLを埋め込むアプローチも可能だが、再利用性などが悪く見づらい。
クリーンな方法として、templateを使う。

## templateについて
メリット: 汎用的なテキストの中に固有のテキストを混在可能。

```
// 例) 文字の埋め込み
こんにちは、{{ .Name }}さん
```

### 種類
* <code>text/template</code>
* <code>html/template</code>
    * データ挿入時、コンテキストを認識している
        * 不正なスクリプト埋め込む攻撃を回避
        * URLで使用できない文字をエンコードする

### 基本的な使い方
* <code>templates/</code>でHTMLファイル作成
* テンプレートの読み込み
* テンプレートをコンパイル
    * データを埋め込める状態にする
    * 利用前に一度だけコンパイルすること
    * 再利用可能
* 出力を受け持つ型を定義

### 出力を受け持つ型
* ファイル名を受け取る
* テンプレートを一回だけコンパイル
* コンパイルされたテンプレートへの参照を保持
* HTTPリクエストに応答する

### templateHandlerについて
ServeHTTPメソッドは、<code>http.HandleFunc</code>に似ている。
入力元のファイルを読み込みテンプレートをコンパイル。
その結果を<code>http.ResponseWriter</code>オブジェクトに出力する。
<code>ServeHTTP</code>は、<code>http.Handler</code>インターフェースに適合しているので<code>http.Handle</code>に直接渡すことができる
