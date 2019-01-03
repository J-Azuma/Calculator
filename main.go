package main

import (
	"fmt"
	"html" //HTMLパッケージのインポートを追加
	"net/http"
)

//ServeHTTPメソッド用の構造体
type Server struct{}

//httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//フォームの入力値を取得
	v := r.FormValue("input_value")

	h := `
	<html><head><title>HTMLのフォーム</title></head><body>
	 <form action= "/" method= "post">
	  <input type= "text" name= "input_value">
	  <input type= "submit" name= "送信"><br/>
	  入力値: ` + html.EscapeString(v) + ` 
	 </form>
	</body></html>
	`
	//クライアント＝ブラウザにHTMLを送信
	fmt.Fprint(w, h)
}

func main() {
	//Webサーバを起動
	http.ListenAndServe(":4000", Server{})
}
