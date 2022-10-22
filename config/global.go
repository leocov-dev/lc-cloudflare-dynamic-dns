package config

var (
	IsDebug   bool
	Version   string
	UserAgent = "go/lc-cf-dns"
)

func GetVersion() string {
	if Version != "" {
		return Version
	}

	return "0.0.0-dev"
}
