// Package heat implements game settings for the heat mechanic.
package heat

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

const (
	defaultInit  int32 = 20
	defaultMin   int32 = 0
	defaultMax   int32 = 100
	defaultDecay int32 = 5
)

// Settings contains all the settings for the Heat mechanic.
type Settings struct {
	pb.HeatSettings
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
  msg.HeatSettings = &pb.HeatSettings{}
	proto.Merge(msg.GetHeatSettings(), &s.HeatSettings)
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	s.Reset()
	if msg.HeatSettings != nil {
		proto.Merge(&s.HeatSettings, msg.GetHeatSettings())
	} else {
		s.Init = defaultInit
		s.Min = defaultMin
		s.Max = defaultMax
		s.Decay = defaultDecay
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	if s.Min > s.Max {
		return fmt.Errorf("Heat minimum greater than maximum. min: ", s.Min, ", max: ", s.Max)
	}
	if s.Init > s.Max {
		return fmt.Errorf("Heat init greater than maximum. init: ", s.Init, ", max: ", s.Max)
	}
	if s.Init < s.Min {
		return fmt.Errorf("Heat init less than minimum. init: ", s.Init, ", min: ", s.Min)
	}
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	state.HeatState = &pb.HeatState{
		Heat: s.Init,
	}
	return nil
}
