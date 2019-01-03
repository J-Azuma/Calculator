package main

import (
	"fmt"
	"html" //HTMLパッケージのインポートを追加
	"net/http"
	"strconv"
)

//ServeHTTPメソッド用の構造体
type Server struct{}

//httpリクエストを受け取るメソッド
func (Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//フォームの入力値を取得
	left := r.FormValue("left") //左項目
	right := r.FormValue("right") //右項目
	op := r.FormValue("op") //演算子(ラジオボタンの値)

	//文字列を整数に変換	
	leftInt, leftErr := stconv.Atoi(left)
	rightInt, rightErr := stconv.Atoi(right)

	var result string 
	if leftErr == nil && rightErr == nil {
		switch op {
		case "add" :
			result = stconv.Itoa(leftInt + rightInt)
		case "sub":
			result = stconv.Itoa(LeftInt - rightInt)
		case "multi":
			result = stconv.Itoa(leftInt * rightInt)
		case "div":
			result = stconv.Itoa(leftInt / rightInt)
		}
	}

	h := `
	<html><head><title>電卓アプリ</title></head><body>
	 <form action= "/" method= "post">
	  左項目:<input type= "text" name= "left">
	  右項目:<input type= "text" name= "right"><br/>
	  演算子:
	  <input type= "radio" name= "op" value= "add" checked> +
	  <input type= "radio" name= "op" value= "sub">-
	  <input type= "radio" name= "op" value= "multi"> +
	  <input type= "radio" name= "op" value= "div"> ÷
	  <br><input type= "submit" name="送信"><hr>

	  [フォームの入力値]<br>
	  左項目: `html.EscapeString(left) + `<br>
	  右項目: `html.EscapeString(right)+`<br>
	  演算子: `html.EscapeString(op)+` <hr>
	  演算結果: `+ EscapeString(result) + `
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

