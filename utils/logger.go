package utils

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func LoggerConfig() (con logger.Config) {
	if viper.GetBool("logger.enable") {
		accessLogName := viper.GetString("logger.output_path") + "." + time.Now().Format("20060102")
		file, err := os.OpenFile(accessLogName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			file, _ = os.Create(accessLogName)
		}

		con = logger.Config{
			Format:     "{\"pid\":${pid},\"time\":\"${time}\",\"latency\":\"${latency}\",\"status\":\"${status}\",\"host\":\"${host}\",\"method\":\"${method}\",\"path\":\"${path}\",\"request_body\":${body},\"response_body\":${resBody}}\n",
			Output:     file,
		}
		return
	}
	return
}


