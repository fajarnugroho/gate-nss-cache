package nss_cache

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldUpdatePasswdCache(t *testing.T) {
	users := []User{
		User{
			Name:      "foobar_name",
			Password:  "foobar_passwd",
			Uid:       1,
			Gid:       2,
			Gecos:     "gecos",
			Directory: "dir",
			Shell:     "shell",
		},
		User{
			Name:      "sea_name",
			Password:  "sea_passwd",
			Uid:       2,
			Gid:       4,
			Gecos:     "gecos_sea",
			Directory: "dir_sea",
			Shell:     "shell_sea",
		},
	}

	filePath := "/tmp/passwd.cache"
	defer os.Remove(filePath)

	err := UpdatePasswdCacheFile(filePath, users)
	assert.NoError(t, err)

	expectedFileContent := "foobar_name:foobar_passwd:1:2:gecos:dir:shell\nsea_name:sea_passwd:2:4:gecos_sea:dir_sea:shell_sea\n"

	file, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedFileContent, string(file))
}

func TestShouldReturnErrorOnEmptyUsers(t *testing.T) {
	users := []User{}

	filePath := "/tmp/passwd.cache"
	defer os.Remove(filePath)

	err := UpdatePasswdCacheFile(filePath, users)
	assert.Equal(t, "Empty Users", err.Error())
}

func TestShouldUpdateGroupCache(t *testing.T) {
	groups := []Group{
		Group{
			Name:     "foo_name",
			Password: "foo_passwd",
			Gid:      2,
			Members:  []string{"gr_mem1", "gr_mem2", "gr_mem3"},
		},
		Group{
			Name:     "bar_name",
			Password: "bar_passwd",
			Gid:      3,
			Members:  []string{"gr_mem1", "gr_mem2", "gr_mem3"},
		},
	}

	filePath := "/tmp/group.cache"
	defer os.Remove(filePath)

	err := UpdateGroupCacheFile(filePath, groups)
	assert.NoError(t, err)

	expectedFileContent := "foo_name:foo_passwd:2:gr_mem1,gr_mem2,gr_mem3\nbar_name:bar_passwd:3:gr_mem1,gr_mem2,gr_mem3\n"

	file, err := ioutil.ReadFile(filePath)
	assert.NoError(t, err)
	assert.Equal(t, expectedFileContent, string(file))
}

func TestShouldReturnError(t *testing.T) {
	groups := []Group{}

	filePath := "/tmp/group.cache"
	defer os.Remove(filePath)

	err := UpdateGroupCacheFile(filePath, groups)
	assert.Equal(t, "Empty Groups", err.Error())
}
