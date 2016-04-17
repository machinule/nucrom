package year

import (
	"github.com/machinule/nucrom/proto/gen"
)

type state struct {
	Settings *settings
	year     int32
}

func NewState(stateProto *pb.GameState, settings *settings) (*state, error) {
	return &state{
		Settings: settings,
	}, nil
}

func (s *state) Marshal(stateProto *pb.GameState) error {
	return nil
}
