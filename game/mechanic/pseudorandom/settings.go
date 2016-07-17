package pseudorandom

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
	"math/rand"
)

// Settings contains parameters that do not change over the course of a game.
type Settings struct {
	init int64 // Initial seed
}

func validate(settingsProto *pb.GameSettings) error {
	// There cannot be any errors in configuring the seed mechanic, as the default is appropriate.
	return nil
}

// NewSettings creates a new settings struct from the GameSettings message.
func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	if settingsProto.GetPseudorandomSettings() == nil {
		return &Settings{
			init: 1,
		}, nil
	}
	return &Settings{
		init: settingsProto.GetPseudorandomSettings().InitSeed,
	}, nil
}

// InitState creates the initial state for the year mechanic.
func (s *Settings) InitState() (*State, error) {
	return &State{
		settings: s,
		seed:     s.init,
		r:        rand.New(rand.NewSource(s.init)),
	}, nil
}
