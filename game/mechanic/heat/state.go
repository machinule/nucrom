package heat

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/machinule/nucrom/proto/gen"
)

type State struct {
	settings *Settings
	heat     int32
}

func NewState(stateProto *pb.GameState, prev *State) (*State, error) {
	if prev == nil {
		return nil, fmt.Errorf("recieved nil previous state, unable to propogate settings.")
	}
	return &State{
		settings: prev.settings,
		heat:     stateProto.GetHeatState().GetHeat(),
	}, nil
}

func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetHeatState() == nil {
		stateProto.HeatState = &pb.HeatState{}
	}
	stateProto.GetHeatState().Heat = proto.Int32(s.heat)
	return nil
}

func (s *State) Heat() int32 {
	return s.heat
}

func (s *State) Chng(mag int32) {
	s.heat += mag
	if s.heat > s.settings.mxm {
		// The world ends, ya fucked up
		s.heat = s.settings.mxm
	}
	if s.heat < s.settings.min {
		s.heat = s.settings.min
	}
}

func (s *State) Decay() {
	s.Chng(-s.settings.decay)
}
