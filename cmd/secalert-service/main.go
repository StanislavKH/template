package main

import (
	"gitlab.qarea.org/jiraquality/cq-web-backend/secalert-service/pkg/config"
	"gitlab.qarea.org/jiraquality/cq-web-backend/secalert-service/internal/service"
)

func main() {
	localpath := "cmd/secalert-service/config/config.toml"
	commonpath := "../common-config/config.toml"

	cfg, err := config.Load(localpath, commonpath)
	checkErr(err)

	service.Forever(cfg)
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
