package main

import (
	"strings"

	"github.com/schollz/croc/src/croc"
	"github.com/schollz/croc/src/utils"
)

func send(file string) string {
	sharedSecret := utils.GetRandomName()

	cr, err := croc.New(croc.Options{
		SharedSecret: sharedSecret,
		IsSender:     true,
		Debug:        false,
		NoPrompt:     true,
		RelayAddress: "142.93.177.120:9009",
		Stdout:       false,
		DisableLocal: false,
		RelayPorts:   strings.Split("", ","),
	})

	if err != nil {
		print(err)
		return "error"
	}

	a := []string{file}

	go cr.Send(croc.TransferOptions{
		PathToFiles:      a,
		KeepPathInRemote: false,
	})

	return sharedSecret
}
