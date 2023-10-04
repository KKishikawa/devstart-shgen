package main

import (
	"os"
)

func main() {
	// コマンドライン引数からパスを取得し、generate_url関数を呼び出す
	// 指定がない場合はカレントディレクトリを使用する
	root_path := "."
	if len(os.Args) > 1 {
		root_path = os.Args[1]
	}
	path, folder := generate_url(root_path)
	writeCodeWorkspace(path, folder)
	println("Done!")
}
