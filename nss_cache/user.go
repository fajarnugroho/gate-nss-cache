package nss_cache

type User struct {
	Name      string `json:"pw_name"`
	Password  string `json:"pw_passwd"`
	Uid       int    `json:"pw_uid"`
	Gid       int    `json:"pw_gid"`
	Gecos     string `json:"pw_gecos"`
	Directory string `json:"pw_dir"`
	Shell     string `json:"pw_shell"`
}
