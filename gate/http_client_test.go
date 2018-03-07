package gate

import (
	"encoding/json"
	"testing"

	"github.com/gate-nss-cache/config"
	"github.com/gate-nss-cache/nss_cache"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetResponse(t *testing.T) {
	url := config.UserURL()
	expectedUsers := []nss_cache.User{
		nss_cache.User{
			Name:      "foobar_name",
			Password:  "foobar_passwd",
			Uid:       1,
			Gid:       2,
			Gecos:     "gecos",
			Directory: "dir",
			Shell:     "shell",
		},
		nss_cache.User{
			Name:      "sea_name",
			Password:  "sea_passwd",
			Uid:       2,
			Gid:       4,
			Gecos:     "gecos_sea",
			Directory: "dir_sea",
			Shell:     "shell_sea",
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url,
		httpmock.NewStringResponder(200, `[{"pw_name": "foobar_name", "pw_passwd": "foobar_passwd", "pw_uid": 1, "pw_gid": 2, "pw_gecos": "gecos", "pw_dir": "dir", "pw_shell": "shell"},{"pw_name": "sea_name", "pw_passwd": "sea_passwd", "pw_uid": 2, "pw_gid": 4, "pw_gecos": "gecos_sea", "pw_dir": "dir_sea", "pw_shell": "shell_sea"}]`))

	usersResponse, err := GetResponse(url)
	assert.NoError(t, err)

	users := make([]nss_cache.User, 0)
	err = json.Unmarshal([]byte(usersResponse), &users)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestShouldReturnErrorMessageForEmptyResponse(t *testing.T) {
	url := config.UserURL()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, `null`))

	_, err := GetResponse(url)
	assert.Equal(t, "Empty Response Body", err.Error())
}

func TestShouldReturnErrorMessageForInternalServerError(t *testing.T) {
	url := config.UserURL()

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(500, `null`))

	_, err := GetResponse(url)
	assert.Equal(t, "Bad Response", err.Error())
}
