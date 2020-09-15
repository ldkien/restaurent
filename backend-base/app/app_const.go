package app

var JWT_KEY = []byte("ldKienyHG1911@#1307!")

const (
	ConfigPath  = "CONFIG_PATH"
	ApiPublic   = "/public/"
	ApiRoot     = "/api/"
	ApiLogin    = ApiPublic + "login"
	ApiRegister = ApiPublic + "register"
	ApiOrder    = ApiRoot + "order"

	DEFAULT_GROUP_USER = "1"
)
