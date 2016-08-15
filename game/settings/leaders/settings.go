package leaders

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

var (
	defaultSettings = &pb.LeadersSettings{}
)

// Settings contains all the settings for the Leaders mechanic.
type Settings struct {
	initYear int32
	pb.LeadersSettings
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
	msg.LeadersSettings = &pb.LeadersSettings{}
	proto.Merge(msg.GetLeadersSettings(), &s.LeadersSettings)
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	s.Reset()
	if msg.LeadersSettings != nil {
		proto.Merge(&s.LeadersSettings, msg.GetLeadersSettings())
	}
	if msg.YearSettings != nil {
		s.initYear = msg.GetYearSettings().InitYear
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	// TODO
	// - positions don't overlap
	// - unique leader names
	// - positions are after the birth date
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	state.LeadersState = &pb.LeadersState{}
	for _, leader := range s.LeaderSettings {
		leaderState := &pb.LeaderState{
			Name: leader.Name,
		}
		// Set the leaders current position to the one currently occupied.
		// If there is no current position, start the leader as a passive leader in the last known active country.
		maxStart := int32(0)
		noCurrentPosition := &pb.LeaderPosition{
			Title:     pb.LeaderTitle_NO_TITLE,
			StartYear: s.initYear,
			EndYear:   0,
			Province:  leader.BirthProvince,
			Elected:   false,
			Dissident: false,
			Type:      pb.LeaderType_NO_LEADER,
		}
		foundCurrent := false
		for _, position := range leader.Positions {
			if position.StartYear <= s.initYear && (position.EndYear == 0 || position.EndYear >= s.initYear) {
				leaderState.CurrentPosition = position
				foundCurrent = true
				break
			}
			if position.StartYear > maxStart {
				maxStart = position.StartYear
				noCurrentPosition.StartYear = position.EndYear
				noCurrentPosition.Province = position.Province
			}
		}
		if !foundCurrent {
			leaderState.CurrentPosition = noCurrentPosition
		}

		state.LeadersState.LeaderStates = append(state.LeadersState.LeaderStates, leaderState)
	}
	return nil
}
