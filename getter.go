package servicecore

import (
	"errors"
	"strings"

	"github.com/narsilworks/servicecore/ifcs"
)

type serviceGetter struct {
	setter *serviceSetter
}

func (g *serviceGetter) Config() (ifcs.IConfiguration, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.config == nil {
		return nil, errors.New("configuration not set")
	}
	return g.setter.config, nil
}

func (g *serviceGetter) Logger() (ifcs.ILogger, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.logger == nil {
		return nil, errors.New("logger not set")
	}
	return g.setter.logger, nil
}

func (g *serviceGetter) Router() (ifcs.IRouter, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.router == nil {
		return nil, errors.New("router not set")
	}
	return g.setter.router, nil
}

func (g *serviceGetter) Cache() (ifcs.ICache, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.cache == nil {
		return nil, errors.New("Cache not set")
	}
	return g.setter.cache, nil
}

func (g *serviceGetter) Queue() (ifcs.IQueue, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.queue == nil {
		return nil, errors.New("queue not set")
	}
	return g.setter.queue, nil
}

func (g *serviceGetter) Data(id string) (ifcs.IData, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.data == nil {
		return nil, errors.New("data not set")
	}

	for _, d := range g.setter.data {
		if strings.EqualFold(id, d.ID()) {
			return d, nil
		}
	}

	return nil, errors.New("data id not found")
}

func (g *serviceGetter) LocalData() (ifcs.ILocalData, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.localData == nil {
		return nil, errors.New("local data not set")
	}
	return g.setter.localData, nil
}

func (g *serviceGetter) CORS() (ifcs.ICORS, error) {
	if g.setter == nil {
		return nil, errors.New("setter not set")
	}
	if g.setter.cors == nil {
		return nil, errors.New("cors not set")
	}
	return g.setter.cors, nil
}
