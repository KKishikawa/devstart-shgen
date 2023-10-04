package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

func generate_url(root_path string) (path string, folder string) {
	// root_pathから絶対パスを取得する
	abs_path, err := filepath.Abs(root_path)
	if err != nil {
		panic(err)
	}
	wfolder := find_workspace_folder(abs_path)
	// abs_pathを16進数に変換後、パスを結合する
	remote_path := "dev-container+" + fmt.Sprintf("%x", abs_path)
	// abs_pathをURIエンコード後、パスを結合する
	return remote_path, wfolder
}

func find_workspace_folder(root_path string) string {
	// .devcontainer/devcontainer.jsonがあるかどうかを確認する
	var path = filepath.Join(root_path, "/.devcontainer/devcontainer.json")
	if !fileExists(path) {
		// なければエラーを出力して終了する
		panic("Error: .devcontainer/devcontainer.json not found.")
	}
	// あれば、その中のworkspaceFolderを取得する
	rawFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	jsonStr := string(rawFile)
	if err != nil {
		panic(err)
	}
	var workspaceFolder = gjson.Get(jsonStr, "workspaceFolder").String()
	if workspaceFolder == "" {
		workspaceFolder = "/workspaces/${localWorkspaceFolder}"
	}
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	localWorkspaceFolder := filepath.Base(currentDir)
	workspaceFolder = strings.Replace(workspaceFolder, "${localWorkspaceFolder}", localWorkspaceFolder, 1)
	return workspaceFolder

}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
