package config

type Deploy struct {
	Token string            `json:"token"`
	Cmds  map[string]string `json:"cmds"`
}

func GetDeploy() Deploy {
	token := gViper.GetString("deploy.token")
	cmds := gViper.GetStringMapString("deploy.cmds")

	return Deploy{
		Token: token,
		Cmds:  cmds,
	}
}
