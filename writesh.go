package main

import (
	"os"

	"encoding/json"
)

type Workspace struct {
	Folders         []WorkspaceFolder `json:"folders"`
	RemoteAuthority string            `json:"remoteAuthority"`
}
type WorkspaceFolder struct {
	Uri string `json:"uri"`
}

func writeCodeWorkspace(remote_path string, workspaceFolder string) {
	ws := Workspace{
		Folders: []WorkspaceFolder{
			{
				Uri: "vscode-remote://" + remote_path + workspaceFolder,
			},
		},
		RemoteAuthority: remote_path,
	}
	file, err := os.Create("code.code-workspace")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonBytes, err := json.Marshal(ws)
	if err != nil {
		panic(err)
	}
	_, err = file.Write(jsonBytes)
	if err != nil {
		panic(err)
	}
}
