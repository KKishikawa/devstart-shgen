package main

import "os"

func writeSh(remote_path string) {
	sh := `#!/bin/sh
code --folder-uri ` + remote_path
	file, err := os.Create("code.sh")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(sh)
	if err != nil {
		panic(err)
	}
	// shファイルに実行権限を付与する
	err = os.Chmod("code.sh", 0755)
	if err != nil {
		panic(err)
	}
}
func writeBat(remote_path string) {
	bat := `@echo off
code --folder-uri ` + remote_path
	file, err := os.Create("code.bat")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(bat)
	if err != nil {
		panic(err)
	}
}
