package runner

import (
	"fmt"
	"testing"
	"time"
)

type service1 struct {
}

func (self *service1) Start() error {
	fmt.Println("service1 start")
	sum := 0
	for {
		sum++
		fmt.Println("service1 :", sum)
		time.Sleep(time.Second)
	}
}

func (self *service1) Stop() error {
	fmt.Println("service1 stop")
	return nil
}

type service2 struct {
}

func (self *service2) Start() error {
	fmt.Println("service2 start")
	sum := 0
	for {
		sum++
		fmt.Println("service2:", sum)
		time.Sleep(time.Second)
	}
}

func (self *service2) Stop() error {
	fmt.Println("service2 stop")
	return nil
}

func TestName(t *testing.T) {
	New().Service(&service1{}).Service(&service2{}).Run()
}
