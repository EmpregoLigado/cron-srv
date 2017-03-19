package conf

import "github.com/spf13/viper"

var (
	CRON_SRV_DB   string
	CRON_SRV_PORT string
)

func init() {
	viper.AutomaticEnv()

	CRON_SRV_DB = viper.GetString("CRON_SRV_DB")
	CRON_SRV_PORT = viper.GetString("CRON_SRV_PORT")
}
