package year

import (
	"github.com/machinule/nucrom/proto/gen"
)

type settings struct {
}

func NewSettings(settingsProto *pb.GameSettings) (*settings, error) {
	return &settings{}, nil
}

func (s *settings) InitState() (*state, error) {
	return &state{
		Settings: s,
	}, nil
}
