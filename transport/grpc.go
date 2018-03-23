package transport

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/yanndr/aura"
	"github.com/yanndr/aura/pb"
	"github.com/yanndr/temperature"
)

type Server struct {
	service aura.Service
}

func New(s aura.Service) *Server {
	return &Server{
		service: s,
	}
}

func (s Server) GetTemperature(context.Context, *pb.TemperatureRequest) (*pb.TemperatureReply, error) {
	t := s.service.GetTemperature()
	pbUnit, err := getPbUnit(t.Unit())
	if err != nil {
		return nil, fmt.Errorf("could not get pbUnit: %v", err)
	}
	return &pb.TemperatureReply{Value: t.Value(), Unit: pbUnit}, nil
}

func (s Server) UpdateTemperature(ctx context.Context, r *pb.UpdateTemperatureRequest) (*pb.UpdateTemperatureReply, error) {

	u, err := getUnit(r.Unit)
	if err != nil {
		return nil, fmt.Errorf("could not get pbUnit: %v", err)
	}
	t := temperature.New(r.Value, u)
	s.service.UpdateTemperature(t)
	return &pb.UpdateTemperatureReply{}, nil
}

func getPbUnit(unit temperature.Convertible) (pb.Unit, error) {

	switch unit.(type) {
	case temperature.CelsiusUnit:
		return pb.Unit_CELSIUS, nil
	case temperature.FahrenheitUnit:
		return pb.Unit_FAHRENHEIT, nil

	case temperature.KelvinUnit:
		return pb.Unit_KELVIN, nil
	default:
		return 0, fmt.Errorf("unit not found")
	}
}

func getUnit(unit pb.Unit) (temperature.Convertible, error) {

	switch unit {
	case pb.Unit_CELSIUS:
		return temperature.Celsius, nil
	case pb.Unit_FAHRENHEIT:
		return temperature.Fahrenheit, nil

	case pb.Unit_KELVIN:
		return temperature.Kelvin, nil
	default:
		return nil, fmt.Errorf("unit not found")
	}
}
