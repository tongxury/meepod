package runner

import (
	"gitee.com/meepo/backend/kit/components/sdk/helper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type IService interface {
	Start() error
	Stop() error
}

func New() *Runner {
	return &Runner{}
}

type Runner struct {
	services []IService
	wg       sync.WaitGroup
}

func (t *Runner) Service(service IService) *Runner {
	t.services = append(t.services, service)
	return t
}

func (t *Runner) listenSignal() {

	c := make(chan os.Signal)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT,
		syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGPIPE, syscall.SIGABRT)

	go func() {
		defer helper.DeferFunc()

		for sig := range c {
			log.Printf("received signal: %s", sig)
			switch sig {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				log.Printf("exiting ...")

				go func() {
					to := 10 * time.Second
					time.Sleep(to)

					//t.wg.Done()
					os.Exit(1)
				}()

				t.Stop()
				//atomic.StoreInt32(&r.stopped, 1)
			}
		}
	}()
}

func (t *Runner) Stop() {
	for _, x := range t.services {
		if err := x.Stop(); err != nil {
			log.Println(err)
		}
	}
}

func (t *Runner) Run() {

	t.listenSignal()

	t.wg.Add(1)

	for _, x := range t.services {
		go func(s IService) {
			defer helper.DeferFunc()

			if err := s.Start(); err != nil {
				log.Fatal(err)
			}
		}(x)
	}

	t.wg.Wait()
}
