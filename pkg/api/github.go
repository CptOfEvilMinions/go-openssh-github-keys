package api

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/CptOfEvilMinions/go-openssh-github-keys/pkg/config"
)

type TeamMember struct {
	Username string `json:"login"`
	SSHKeys  []string
}

type UserSSHKeys struct {
	ID     int    `json:"id"`
	SSHKey string `json:"key"`
}

var TeamMembers map[string][]string

func init() {

}

func (c *Client) CheckListOfUsersInTeam(username string, cfg *config.Config) bool {
	resp, err := HttpClient.Get(fmt.Sprintf("/orgs/%s/teams/%s/members", cfg.Organization, cfg.Team))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create agent list struct
	var teamMembers []TeamMember

	// Extract JSON
	if err := json.NewDecoder(resp.Body).Decode(&teamMembers); err != nil {
		log.Fatal(err)
	}

	for _, teamMember := range teamMembers {
		if teamMember.Username == username {
			return true
		}
	}
	return false

}

func (c *Client) GetSSHKeysForUser(username string) ([]UserSSHKeys, error) {
	resp, err := HttpClient.Get(fmt.Sprintf("/users/%s/keys", username))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userSSHKeys []UserSSHKeys
	if err := json.NewDecoder(resp.Body).Decode(&userSSHKeys); err != nil {
		return nil, err
	}

	return userSSHKeys, nil
}
