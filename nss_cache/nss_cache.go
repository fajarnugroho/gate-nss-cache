package nss_cache

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func UpdatePasswdCacheFile(filePath string, users []User) error {
	if len(users) == 0 {
		return errors.New("Empty Users")
	}

	var userData string
	for _, user := range users {
		singleUser := fmt.Sprintf("%s:%s:%d:%d:%s:%s:%s\n", user.Name, user.Password, user.Uid, user.Gid, user.Gecos, user.Directory, user.Shell)
		userData += singleUser
	}

	err := ioutil.WriteFile(filePath, []byte(userData), 0644)
	return err
}

func UpdateGroupCacheFile(filePath string, groups []Group) error {
	if len(groups) == 0 {
		return errors.New("Empty Groups")
	}

	var groupData string
	for _, group := range groups {
		singleGroup := fmt.Sprintf("%s:%s:%d:%s\n", group.Name, group.Password, group.Gid, strings.Join(group.Members, ","))
		groupData += singleGroup
	}

	err := ioutil.WriteFile(filePath, []byte(groupData), 0644)
	return err
}
