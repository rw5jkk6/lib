## packageライブラリの作り方
- Githubでの操作
  - リポジトリを作る
  - URLをコピーしておく
- VScodeでの操作  
  - リポジトリを初期化する。一番左のアクティビティバーから作る
  - go mod init github.com/username/repo (ここきちんと書かないとダウンロードできない)
  - go mod tidy
  - コードを書く(関数は頭は大文字)
  - git add (ステージングは + を押す)
  - git commit (コメントを入れて ctrl + enter)
  - コマンドパレットから git add remote　(リモートの追加。gitを作っておかないと出てこない)
  - GithubのリポジトリのURLを貼り付ける。リモート名は\[origin]にしておく
  - git push ソース管理ビューの...をクリックしてメニュー先の\[プッシュ先...]を選択して[origin]を選択する

## Githubライブラリの使い方
- メインのコードを書く
- go mod init モジュール名
- go get github.com/username/repo
```go
package main
 
import (
   "fmt"
   "github.com/rw5jkk6/add"
)

func main(){
	n := add.Add(1, 2)
	fmt.Println(n)
}
```
