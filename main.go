package main

import (
	"fmt"
	"time"

	cmd "money-transfer/cmd"
	"money-transfer/internal/config"
)

func main() {
	if tz := config.GetString("TZ"); tz != "" {
		var err error
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			fmt.Printf("error loading location '%s': %v\n", tz, err)
		} else {
			fmt.Printf("location loaded '%s'\n", tz)
		}
	}
	cmd.Run()
}
