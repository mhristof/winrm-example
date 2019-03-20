package main

import (
	"github.com/masterzen/winrm"
	"os"
)

func main() {
	var (
		user     string = "Administrator"
		password string = os.Args[2]
		host     string = os.Args[1]
		port     int    = 5985
	)

	endpoint := winrm.NewEndpoint(host, port, false, false, nil, nil, nil, 0)
	client, err := winrm.NewClient(endpoint, user, password)
	if err != nil {
		panic(err)
	}
	cmd := "ipconfig /all"
	_, err = client.Run(cmd, os.Stdout, os.Stderr)

	if err != nil {
		panic(err)
	}

	return
}
