package year

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

// Settings contains parameters that do not change over the course of a game.
type Settings struct {
	init int32 // Starting year.
	incr int32 // Number of years to increment per turn.
}

func validate(settingsProto *pb.GameSettings) error {
	// There cannot be any errors in configuring the year mechanic, as the default is appropriate.
	return nil
}

// NewSettings creates a new settings struct from the GameSettings message.
func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	if settingsProto.GetYearSettings() == nil {
		return &Settings{
			init: 1948,
			incr: 1,
		}, nil
	}
	return &Settings{
		init: settingsProto.GetYearSettings().InitYear,
		incr: 1,
	}, nil
}

// InitState creates the initial state for the year mechanic.
func (s *Settings) InitState() (*State, error) {
	return &State{
		settings: s,
		year:     s.init,
	}, nil
}
