package config

import "os"

type Config struct {
	Addr string
}
var AppConfig Config
func Load(){
	AppConfig.Addr = os.Getenv("PORT")
}