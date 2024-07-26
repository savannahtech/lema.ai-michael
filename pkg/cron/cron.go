package cron

import (
	"github.com/go-co-op/gocron"
	"strconv"

	"time"
)

var schedular *gocron.Scheduler

// time in minutes (15 minutes)
var timeDuration int64 = 15

// GetTimeDuration returns the time delay for the cron job
func GetTimeDuration(duration string) int64 {
	intDuration, err := strconv.Atoi(duration)
	if err != nil {
		return timeDuration
	}
	return int64(intDuration)
}

// InitCron initializes the gocron scheduler
func InitCron() {

	s := gocron.NewScheduler(time.UTC)
	schedular = s
}

// StopCronJob stops the cron jobs
func StopCronJob() {
	schedular.Stop()
}

// StartCronJob starts the cron jobs
func StartCronJob() {
	schedular.StartAsync()
}

// SetCronJob sets the cron job
func SetCronJob(f func() error, interval int64) {
	timeD := time.Duration(interval) * time.Minute
	_, err := schedular.Every(timeD).Do(f)
	if err != nil {
		return
	}
}
