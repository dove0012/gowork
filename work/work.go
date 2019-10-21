package work

import (
	"time"
)

type BeforeFunc func() error

type AfterFunc func() error

type Work interface {
	GetName() string
	GetUsage() string
	GetBeforeFunc() BeforeFunc
	GetAfterFunc() AfterFunc
	GetReboot() bool
	GetRebootTime() time.Duration
	Run()
}

type Base struct {
	name       string
	usage      string
	beforeFunc BeforeFunc
	afterFunc  AfterFunc
	reboot     bool
	rebootTime time.Duration
}

func (base *Base) GetName() string {
	return base.name
}

func (base *Base) GetUsage() string {
	return base.usage
}

func (base *Base) GetBeforeFunc() BeforeFunc {
	return base.beforeFunc
}

func (base *Base) GetAfterFunc() AfterFunc {
	return base.afterFunc
}

func (base *Base) GetReboot() bool {
	return base.reboot
}

func (base *Base) GetRebootTime() time.Duration {
	return base.rebootTime
}

func GetWorks() []Work {
	return []Work{
		newReckon(),
	}
}
