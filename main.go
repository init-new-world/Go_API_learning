package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/init-new-world/Go_API_learning/model"
	"github.com/init-new-world/Go_API_learning/router/middleware"
	"github.com/spf13/viper"

	"github.com/init-new-world/Go_API_learning/config"

	ver "github.com/init-new-world/Go_API_learning/pkg/version"
	"github.com/spf13/pflag"

	"github.com/init-new-world/Go_API_learning/handler/sd"

	"github.com/gin-gonic/gin"
	"github.com/init-new-world/Go_API_learning/router"
	"github.com/lexkong/log"
)

var (
	cfg     = pflag.StringP("config", "c", "", "Go_API_learning Server config file path.")
	version = pflag.BoolP("version", "v", false, "Show version info.")
)

func main() {
	pflag.Parse()

	if *version {
		v := ver.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}

	gin.SetMode(viper.GetString("runmode"))
	server := gin.Default()

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	middlewares := []gin.HandlerFunc{
		middleware.RequestId(),
		middleware.Logging(),
	}

	model.DB.Init()
	defer model.DB.Close()

	server = router.Load(server, middlewares...)

	go func() {
		if err := sd.PingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start TLS-Server at Port %s", viper.GetString("tls.port"))
			log.Info(server.RunTLS(":"+viper.GetString("tls.port"), cert, key).Error())
		}()
	}

	log.Infof("Start Server at Port %s", viper.GetString("port"))
	log.Info(server.Run(":" + viper.GetString("port")).Error())
}
