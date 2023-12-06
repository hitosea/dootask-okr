package timer

import (
	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New(cron.WithSeconds())
	Crontab(c)
	c.Start()
}

// Crontab 定时任务
func Crontab(c *cron.Cron) {
	// c.AddFunc("*/5 * * * * *", func() {
	// 	fmt.Println("Run every 5 seconds 1212")
	// })
	// _, _ = c.AddFunc("*/10 * * * * *", service.NewOkrService().OkrNotice) // 10秒检查OKR提醒
}
