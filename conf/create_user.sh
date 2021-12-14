#!/bin/bash

set -e

CONFIG_FILE="/etc/go-openssh-github-keys/settings.yaml"
GITHUB_TOKEN=$(/usr/local/bin/yq -r '.token' /etc/go-openssh-github-keys/settings.yaml)
GITHUB_ORG=$(/usr/local/bin/yq -r '.organization' /etc/go-openssh-github-keys/settings.yaml)
GITHUB_TEAM=$(/usr/local/bin/yq -r '.team' /etc/go-openssh-github-keys/settings.yaml)


teamMemberJson=$(curl -X GET -s "https://api.github.com/orgs/${GITHUB_ORG}/teams/${GITHUB_TEAM}/members" \
-H "Accept: application/vnd.github.v3+json" \
-H "Authorization: token ${GITHUB_TOKEN}")

for row in $(echo "${teamMemberJson}" | jq -r '.[] | @base64'); do
    _jq() {
        echo "${row}" | base64 --decode | jq -r "${1}"
    }
    # Extract username and html
    username=$(_jq '.login')

    if id "${username}" &>/dev/null; then
        echo "[+] - User ${username} already exists"
    else
        echo "[*] - Creating user $1"
        useradd -m ${username} -s /bin/bash

        # Make sure user can't modify the authorized_keys file
        mkdir -p /home/${username}/.ssh &&
        touch /home/${username}/.ssh/authorized_keys &&
        chown root:root /home/${username}/.ssh/authorized_keys && \
        chmod 600 /home/${username}/.ssh/authorized_keys && \
        chattr +i /home/${username}/.ssh/authorized_keys
    fi
done

