package main

import (
	"encoding/json"

	"./config"
	"./gate"
	"./nss_cache"
)

const PasswdCacheFilePath = "/etc/passwd.cache"
const GroupCacheFilePath = "/etc/group.cache"

func main() {
	config.Load()

	usersData, err := gate.GetResponse(config.UserURL())
	handleError(err)

	users := make([]nss_cache.User, 0)
	err = json.Unmarshal([]byte(usersData), &users)
	handleError(err)

	err = nss_cache.UpdatePasswdCacheFile(PasswdCacheFilePath, users)
	handleError(err)

	groupsData, err := gate.GetResponse(config.GroupURL())
	handleError(err)

	groups := make([]nss_cache.Group, 0)
	err = json.Unmarshal([]byte(groupsData), &groups)
	handleError(err)

	err = nss_cache.UpdateGroupCacheFile(GroupCacheFilePath, groups)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
