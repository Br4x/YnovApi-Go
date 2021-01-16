package main

import (
	_ "ynov_immo/config"
	"ynov_immo/handlers"
	"ynov_immo/tasks"

	"github.com/spf13/viper"
)

func main() {
	if viper.GetBool("app.enable_cron") {
		go tasks.RunTasks()
	}
	defer handlers.Close()
	handlers.ServerRun()
}
