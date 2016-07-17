package heat

import (
	"fmt"
	"github.com/machinule/nucrom/proto/gen"
)

type State struct {
	settings *Settings
	heat     int32
}

func NewState(stateProto *pb.GameState, settings *Settings) (*State, error) {
	if settings == nil {
		return nil, fmt.Errorf("received nil Settings, unable to continue.")
	}
	if stateProto.GetHeatState() == nil {
		return nil, fmt.Errorf("received nil HeatState, unable to continue.")
	}
	return &State{
		settings: settings,
		heat:     stateProto.GetHeatState().Heat,
	}, nil
}

func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetHeatState() == nil {
		stateProto.HeatState = &pb.HeatState{}
	}
	stateProto.GetHeatState().Heat = s.heat
	return nil
}

func (s *State) Heat() int32 {
	return s.heat
}

func (s *State) Chng(mag int32) {
	s.heat += mag
	if s.heat > s.settings.max {
		// The world ends, ya messed up
		s.heat = s.settings.max
	}
	if s.heat < s.settings.min {
		s.heat = s.settings.min
	}
}

func (s *State) Decay() {
	s.Chng(-s.settings.decay)
}
