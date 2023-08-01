package main

import (
	"fmt"
	"github.com/robfig/cron"
)

func main() {
	cronTab := cron.New()
	cronTab.Start()

	fmt.Println("First Program.")
}
