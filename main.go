package main

import (
	"fmt"
	"os"

	"github.com/CptOfEvilMinions/go-openssh-github-keys/pkg/api"
	"github.com/CptOfEvilMinions/go-openssh-github-keys/pkg/config"
)

func main() {
	var sshUsername string
	if len(os.Args) == 2 {
		sshUsername = os.Args[1]
	} else {
		fmt.Println("[-] - Exiting please enter a single username")
	}

	// Generate our config based on the config supplied
	cfg, err := config.NewConfig("/etc/go-openssh-github-keys/settings.yaml")
	if err != nil {
		os.Exit(3)
	}

	// Init HTTP client
	api.InitHTTPclient(cfg)

	// Check if username exists in Github team
	if !api.HttpClient.CheckListOfUsersInTeam(sshUsername, cfg) {
		os.Exit(3)
	}

	// get user SSH keys
	userSSHkeys, err := api.HttpClient.GetSSHKeysForUser(sshUsername)
	if err != nil {
		os.Exit(3)
	}

	// Return list of SSH keys to SSHD
	for _, userSSHkey := range userSSHkeys {
		fmt.Println(userSSHkey.SSHKey)
	}
}
