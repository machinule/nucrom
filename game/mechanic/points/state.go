// The points mechanic keeps track of the point stores for each superpower
package points

import (
	"fmt"
	"github.com/machinule/nucrom/proto/gen"
)

// The state of the points mechanic.
type State struct {
	settings *Settings
	usa      PointState
	ussr     PointState
}

// The state of each superpower's points
type PointState struct {
	pol int32
	mil int32
	cov int32
}

// NewState creates a new state from the GameState message and the previous state.
func NewState(stateProto *pb.GameState, settings *Settings) (*State, error) {
	if settings == nil {
		return nil, fmt.Errorf("received nil Settings, unable to continue.")
	}
	if stateProto.GetPointsState() == nil {
		return nil, fmt.Errorf("received nil PointsState, unable to continue.")
	}
	return &State{
		settings: settings,
		usa: PointState{
			pol: stateProto.GetPointsState().GetUsaState().Political,
			mil: stateProto.GetPointsState().GetUsaState().Military,
			cov: stateProto.GetPointsState().GetUsaState().Covert,
		},
		ussr: PointState{
			pol: stateProto.GetPointsState().GetUssrState().Political,
			mil: stateProto.GetPointsState().GetUssrState().Military,
			cov: stateProto.GetPointsState().GetUssrState().Covert,
		},
	}, nil
}

// Marshal fills in the GameState proto with the current state.
func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	stateProto.PointsState = &pb.PointsState{
		UsaState: &pb.PointState{
			Political: s.usa.pol,
			Military:  s.usa.mil,
			Covert:    s.usa.cov,
		},
		UssrState: &pb.PointState{
			Political: s.ussr.pol,
			Military:  s.ussr.mil,
			Covert:    s.ussr.cov,
		},
	}
	return nil
}

func (s *State) ApplyBaseIncome() {
	s.usa.pol += s.settings.usa.pol_income
	s.usa.mil += s.settings.usa.mil_income
	s.usa.cov += s.settings.usa.cov_income

	s.ussr.pol += s.settings.ussr.pol_income
	s.ussr.mil += s.settings.ussr.mil_income
	s.ussr.cov += s.settings.ussr.cov_income
}

// GETTERS

func (s *State) POL(player pb.Player) int32 {
	if player == pb.Player_USSR {
		return s.ussr.pol
	} else {
		return s.usa.pol
	}
}

func (s *State) MIL(player pb.Player) int32 {
	if player == pb.Player_USSR {
		return s.ussr.mil
	} else {
		return s.usa.mil
	}
}

func (s *State) COV(player pb.Player) int32 {
	if player == pb.Player_USSR {
		return s.ussr.cov
	} else {
		return s.usa.cov
	}
}
