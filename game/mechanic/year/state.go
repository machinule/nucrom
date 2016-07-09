// The year mechanic keeps track of the current game year.
package year

import (
	"fmt"
	"github.com/machinule/nucrom/proto/gen"
)

// The state of the year mechanic.
type State struct {
	settings *Settings
	year     int32
}

// NewState creates a new state from the GameState message and the previous state.
func NewState(stateProto *pb.GameState, settings *Settings) (*State, error) {
	if settings == nil {
		return nil, fmt.Errorf("received nil Settings, unable to continue.")
	}
	if stateProto.GetYearState() == nil {
		return nil, fmt.Errorf("received nil YearState, unable to continue.")
	}
	return &State{
		settings: settings,
		year:     stateProto.GetYearState().Year,
	}, nil
}

// Marshal fills in the GameState proto with the current state.
func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	stateProto.YearState = &pb.YearState{
		Year: s.year,
	}
	return nil
}

// Year gets the current year.
func (s *State) Year() int32 {
	return s.year
}

// Incr increments the year by the increment value in settings.
func (s *State) Incr() {
	s.year += s.settings.incr
}
