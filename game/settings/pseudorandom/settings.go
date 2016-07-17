// Package pseudorandom implements game settings for the pseudorandom mechanic.
package pseudorandom

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

var (
	defaultSettings = &pb.PseudorandomSettings{
		InitSeed: 0,
	}
)

// Settings contains all the settings for the Pseudorandom mechanic.
type Settings struct {
	pb.PseudorandomSettings
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
	msg.PseudorandomSettings = &pb.PseudorandomSettings{}
	proto.Merge(msg.PseudorandomSettings, &s.PseudorandomSettings)
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	s.Reset()
	if msg.PseudorandomSettings != nil {
		proto.Merge(&s.PseudorandomSettings, msg.PseudorandomSettings)
	} else {
		proto.Merge(&s.PseudorandomSettings, defaultSettings)
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	state.PseudorandomState = &pb.PseudorandomState{
		Seed: s.InitSeed,
	}
	return nil
}
