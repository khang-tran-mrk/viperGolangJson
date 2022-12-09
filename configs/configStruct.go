package configs

type (
	Config struct {
		Server   Server   `json:"server"`
		Database Database `json:"database"`
	}
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	}
	Database struct {
		Host       string `json:"host"`
		DbUserName string `json:"dbUserName"`
		DbPass     string `json:"dbPass"`
	}
)
