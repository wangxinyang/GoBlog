package utils

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	AppMode    string
	HttpPort   string
	Page       string
	Size       string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

// init()->main()
func init() {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Println("環境変数を読み込みがエラーになりました、ファイルパスをチェックしてください", err)
	}
	loadServer(cfg)
	loadDatabase(cfg)
}

// バックエンドの環境変数を取得する
func loadServer(cfg *ini.File) {
	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString(":3100")
	Page = cfg.Section("server").Key("Page").MustString("1")
	Size = cfg.Section("server").Key("Size").MustString("5")
}

// データベースの環境変数を取得する
func loadDatabase(cfg *ini.File) {
	DbHost = cfg.Section("database").Key("DbHost").MustString("localhost")
	DbPort = cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = cfg.Section("database").Key("DbUser").MustString("root")
	DbPassword = cfg.Section("database").Key("DbPassword").MustString("123456")
	DbName = cfg.Section("database").Key("DbName").MustString("goblog")
}
