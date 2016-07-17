// Package year implements game settings for the year mechanic.
package year

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

var (
	defaultSettings = &pb.YearSettings{
		InitYear: 1948,
	}
)

// Settings contains all the settings for the Year mechanic.
type Settings struct {
	pb.YearSettings
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
	msg.YearSettings = &pb.YearSettings{}
	proto.Merge(msg.YearSettings, &s.YearSettings)
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	s.Reset()
	if msg.YearSettings != nil {
		proto.Merge(&s.YearSettings, msg.YearSettings)
	} else {
		proto.Merge(&s.YearSettings, defaultSettings)
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	state.YearState = &pb.YearState{
		Year: s.InitYear,
	}
	return nil
}
