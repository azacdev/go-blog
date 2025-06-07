package config

type Config struct {
	App         App
	Server      Server
	DB          DB
	JWTAccess   JWTAccess
	JWTRefresh  JWTRefresh
	GoogleOAuth GoogleOAuth
}

type App struct {
	Name string
}

type JWTAccess struct {
	Secret string
}

type JWTRefresh struct {
	Secret string
}
type Server struct {
	Host string
	Port string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
	Sll      string
}

type GoogleOAuth struct {
	ClientID     string
	ClientSecret string
}
