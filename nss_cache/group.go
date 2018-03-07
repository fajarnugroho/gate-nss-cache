package nss_cache

type Group struct {
	Name     string   `json:"gr_name"`
	Password string   `json:"gr_passwd"`
	Gid      int      `json:"gr_gid"`
	Members  []string `json:"gr_mem"`
}
