package servicecore

import (
	"errors"
	"strings"

	"github.com/narsilworks/servicecore/ifcs"
)

type ServiceGetter struct {
	setter *serviceSetter
}

var ErrSetterNotSet error = errors.New(`setter not set`)
var ErrConfigNotSet error = errors.New(`configuration function not set`)
var ErrLoggerNotSet error = errors.New(`logger function not set`)
var ErrRouterNotSet error = errors.New(`router function not set`)
var ErrCacheNotSet error = errors.New(`cache function not set`)
var ErrQueueNotSet error = errors.New(`queue function not set`)
var ErrDataNotSet error = errors.New(`data function not set`)
var ErrLocalDataNotSet error = errors.New(`local data function not set`)
var ErrCORSNotSet error = errors.New(`cors function not set`)
var ErrNotificationNotSet error = errors.New(`notification function not set`)

func (g *ServiceGetter) Config() (ifcs.IConfiguration, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.config == nil {
		return nil, ErrConfigNotSet
	}
	return g.setter.config, nil
}

func (g *ServiceGetter) Logger() (ifcs.ILogger, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.logger == nil {
		return nil, ErrLoggerNotSet
	}
	return g.setter.logger, nil
}

func (g *ServiceGetter) Router() (ifcs.IRouter, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.router == nil {
		return nil, ErrRouterNotSet
	}
	return g.setter.router, nil
}

func (g *ServiceGetter) Cache() (ifcs.ICache, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.cache == nil {
		return nil, ErrCacheNotSet
	}
	return g.setter.cache, nil
}

func (g *ServiceGetter) Queue() (ifcs.IQueue, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.queue == nil {
		return nil, ErrQueueNotSet
	}
	return g.setter.queue, nil
}

func (g *ServiceGetter) Data(id string) (ifcs.IData, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.data == nil {
		return nil, ErrDataNotSet
	}

	for _, d := range g.setter.data {
		if strings.EqualFold(id, d.ID()) {
			return d, nil
		}
	}

	return nil, errors.New("data id not found")
}

func (g *ServiceGetter) LocalData() (ifcs.ILocalData, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.localData == nil {
		return nil, ErrLocalDataNotSet
	}
	return g.setter.localData, nil
}

func (g *ServiceGetter) CORS() (ifcs.ICORS, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.cors == nil {
		return nil, ErrCORSNotSet
	}
	return g.setter.cors, nil
}

func (g *ServiceGetter) Notification() (ifcs.INotification, error) {
	if g.setter == nil {
		return nil, ErrSetterNotSet
	}
	if g.setter.notification == nil {
		return nil, ErrNotificationNotSet
	}
	return g.setter.notification, nil
}
