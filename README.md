## packageライブラリの作り方
- Githubでの操作
  - リポジトリを作る
  - URLをコピーしておく
- VScodeでの操作  
  - フォルダを作る。何も影響はないので、名前は何でもいい
  - リポジトリを初期化する。一番左のアクティビティバーから作る
  - ターミナルで `go mod init github.com/username/repo` (ここきちんと書かないとダウンロードできない。`.git`はつけない)
  - go mod tidy
  - コードを書く(関数の頭は大文字)初めはmasterで、それ以降の変更はブランチで作って、マージする
  - モジュールはmain.goがないのでテストを書く
  - git add (ステージングは + を押す)
  - git commit (コメントを入れて ctrl + enter)
  - コマンドパレットから git add remote　(リモートの追加。gitリポジトリを作っておかないと出てこない)
  - GithubのリポジトリのURLを貼り付ける。リモート名は\[origin]にしておく
  - git push ソース管理ビューの...をクリックしてメニュー先の\[プッシュ先...]を選択して[origin]を選択する
- ライブラリの使用をやめる
  - ファイルのimportからgithubのライブラリを削除する
  - `go mod tidy`をするとgo modファイルから削除される

## Githubライブラリの使い方
- メインのコードを書く
- go mod init モジュール名
- モジュールをダウンロードする方法は２通りあるが、`go mod tidy`を使っておけばいい
  - go mod tidy 
  - ~go get github.com/username/repo~
- コードの実行

```go
package main
 
import (
   "fmt"
   "github.com/rw5jkk6/culc"
)

func main(){
	n := culc.Add(1, 2)
	fmt.Println(n)
}
```
