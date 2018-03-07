package nss_cache

import (
	"errors"
	"fmt"
	"io/ioutil"
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
