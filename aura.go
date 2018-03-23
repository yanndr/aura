package aura

import (
	"github.com/yanndr/temperature"
)

type Service interface {
	GetTemperature() temperature.Temperature
	UpdateTemperature(t temperature.Temperature)
}

type service struct {
	temperature temperature.Temperature
}

func New(t temperature.Temperature) Service {
	return &service{temperature: t}
}

func (s service) GetTemperature() temperature.Temperature {
	return s.temperature
}

func (s service) UpdateTemperature(t temperature.Temperature) {
	if s.temperature == nil {
		s.temperature = t
		return
	}
	s.temperature.SetTemperature(t)
}
