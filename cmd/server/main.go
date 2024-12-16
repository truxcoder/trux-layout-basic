package main

import (
	"fmt"
	"github.com/truxcoder/trux-layout-basic/cmd/server/wire"
	"github.com/truxcoder/trux-layout-basic/pkg/config"
	"github.com/truxcoder/trux-layout-basic/pkg/http"
	"github.com/truxcoder/trux-layout-basic/pkg/log"
	"go.uber.org/zap"
)

func main() {
	conf := config.NewConfig()
	logger := log.NewLog(conf)

	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
