package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type ConfigFile struct {
	Nss_Http struct {
		HostUrl string `yaml:"host_url"`
		ApiKey  string `yaml:"api_key"`
	}
}

type User struct {
	Name      string `json:"pw_name"`
	Password  string `json:"pw_passwd"`
	Uid       int    `json:"pw_uid"`
	Gid       int    `json:"pw_gid"`
	Gecos     string `json:"pw_gecos"`
	Directory string `json:"pw_dir"`
	Shell     string `json:"pw_shell"`
}

type Group struct {
	Name     string   `json:"gr_name"`
	Password string   `json:"gr_passwd"`
	Gid      int      `json:"gr_gid"`
	Members  []string `json:"gr_mem"`
}

func getData(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	return string(response)
}

func getConfiguration(fileName string) (string, string) {

	configuration := ConfigFile{}
	yamlConfiguration, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	yaml_err := yaml.Unmarshal([]byte(yamlConfiguration), &configuration)

	if yaml_err != nil {
		panic(yaml_err)
	}

	return configuration.Nss_Http.HostUrl, configuration.Nss_Http.ApiKey

}

func getUserUrl() string {
	hostUrl, apiKey := getConfiguration(os.Getenv("GATE_CONFIG_FILE"))
	return fmt.Sprintf("%s/passwd?token=%s", hostUrl, apiKey)
}

func getGroupUrl() string {
	hostUrl, apiKey := getConfiguration(os.Getenv("GATE_CONFIG_FILE"))
	return fmt.Sprintf("%s/group?token=%s", hostUrl, apiKey)

}
func getUsers() []User {
	usersData := []byte(getData(getUserUrl()))
	users := make([]User, 0)
	json.Unmarshal(usersData, &users)
	return users
}

func getGroups() []Group {
	groupsData := []byte(getData(getGroupUrl()))
	groups := make([]Group, 0)
	json.Unmarshal(groupsData, &groups)
	return groups
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	users := getUsers()
	var userData string
	for _, user := range users {
		singleUser := fmt.Sprintf("%s:%s:%d:%d:%s:%s:%s\n", user.Name, user.Password, user.Uid, user.Gid, user.Gecos, user.Directory, user.Shell)
		userData += singleUser
	}

	groups := getGroups()
	var groupData string
	for _, group := range groups {
		singleGroup := fmt.Sprintf("%s:%s:%d:%s\n", group.Name, group.Password, group.Gid, strings.Join(group.Members, ","))
		groupData += singleGroup
	}

	userErr := ioutil.WriteFile("/etc/passwd.cache", []byte(userData), 0644)
	check(userErr)

	groupErr := ioutil.WriteFile("/etc/group.cache", []byte(groupData), 0644)
	check(groupErr)
}
