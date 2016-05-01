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
func NewState(stateProto *pb.GameState, prev *State) (*State, error) {
	if prev == nil {
		return nil, fmt.Errorf("received nil previous state, unable to propogate settings.")
	}
	if stateProto.GetPointsState() == nil {
		return nil, fmt.Errorf("received nil PointsState, unable to continue.")
	}
	return &State{
		settings: prev.settings,
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

// GETTERS
