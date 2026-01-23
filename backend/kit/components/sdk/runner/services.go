package runner

import (
	"gitee.com/meepo/backend/kit/components/sdk/redisstream"
	"github.com/robfig/cron/v3"
)

func (t *Runner) CronService(cr *cron.Cron) *Runner {
	t.Service(cronService{cr: cr})
	return t
}

type cronService struct {
	cr *cron.Cron
}

func (s cronService) Start() error {
	s.cr.Start()
	return nil
}
func (s cronService) Stop() error {
	s.cr.Stop()
	return nil
}

func (t *Runner) RedisStreamService(stream *redisstream.RedisStream) *Runner {
	t.Service(redisStreamService{stream: stream})
	return t
}

type redisStreamService struct {
	stream *redisstream.RedisStream
}

func (t redisStreamService) Start() error {
	t.stream.Start()
	return nil
}
func (t redisStreamService) Stop() error {
	t.stream.Stop()
	return nil
}
