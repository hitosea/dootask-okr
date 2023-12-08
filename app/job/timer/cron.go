package timer

import (
	"dootask-okr/app/service"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New(cron.WithSeconds())
	Crontab(c)
	c.Start()
}

// Crontab 定时任务
func Crontab(c *cron.Cron) {
	_, _ = c.AddFunc("*/10 * * * * *", service.NewOkrService().OkrNotice) // 10秒检查OKR提醒
}
