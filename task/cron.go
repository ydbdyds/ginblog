package task

import (
	"ginblog/model"
	"github.com/robfig/cron/v3"
)

func ExecuteCron() {
	c := cron.New()
	c.AddFunc("@every 1h", func() {
		model.CheckAndUpdate()
	})
	c.Start()
}
