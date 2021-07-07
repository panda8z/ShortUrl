package utils

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	JwtKey     string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	AccessKey  string
	SecretKey  string
	Bucket     string
	ImgUrl     string
)

func init() {
	os.Setenv("CONFIG_FILE", "./config/config.ini")
	log.Println(os.Getenv("CONFIG_FILE"))

	file, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println(err.Error() + os.Getenv("CONFIG_FILE"))
		log.Println("Config file not found or open it with error, pleas check src/config/config.ini ")
		panic(err)
	}
	loadApp(file)
	loadServer(file)
	loadDatabase(file)
	loadQiNiu(file)

}

func loadApp(file *ini.File) {
	AppMode = file.Section("app").Key("JwtKey").MustString("!@#$%^&*()qwertyuiop")
}
func loadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString("8081")
}

func loadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123456")
	DbName = file.Section("database").Key("DbName").MustString("shorturl")
}

func loadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	ImgUrl = file.Section("qiniu").Key("ImgUrl").String()
}
