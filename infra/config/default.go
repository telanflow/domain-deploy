package config

type Default struct {
}

func GetDefault() Default {
	return Default{}
}

func GetAddress() (string, int) {
	ip := gViper.GetString("default.ip")
	port := gViper.GetInt("default.port")
	return ip, port
}
