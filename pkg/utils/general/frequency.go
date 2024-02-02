package general

import (
	"time"
)

func MakeNextRemindDate(frequency int) string {
	t := time.Now()
	var nextRemindDate string
	switch frequency {
	case 0:
		nextRemindDate = t.AddDate(0, 0, 1).Format("2006-01-02 15:04:05")
	case 1:
		nextRemindDate = t.AddDate(0, 0, 7).Format("2006-01-02 15:04:05")
	case 2:
		nextRemindDate = t.AddDate(0, 0, 14).Format("2006-01-02 15:04:05")
	case 3:
		nextRemindDate = t.AddDate(0, 1, 0).Format("2006-01-02 15:04:05")
	}
	return nextRemindDate
}
