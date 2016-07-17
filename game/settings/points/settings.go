// Package points implements game settings for the points mechanic.
package points

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

var (
	defaultSettings = &pb.PointsSettings{
		Usa: &pb.PointStoresSettings{
			Political: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
			Military: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
			Covert: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
		},
		Ussr: &pb.PointStoresSettings{
			Political: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
			Military: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
			Covert: &pb.PointStoreSettings{
				Init:   1,
				Income: 1,
			},
		},
	}
)

// Settings contains all the settings for the Points mechanic.
type Settings struct {
	pb.PointsSettings
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
	msg.PointsSettings = &pb.PointsSettings{}
	proto.Merge(msg.PointsSettings, &s.PointsSettings)
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	s.Reset()
	if msg.PointsSettings != nil {
		proto.Merge(&s.PointsSettings, msg.PointsSettings)
	} else {
		proto.Merge(&s.PointsSettings, defaultSettings)
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	state.PointsState = &pb.PointsState{
		Usa: &pb.PointStoresState{
			Political: &pb.PointStoreState{
				Count: s.Usa.Political.Init,
			},
			Military: &pb.PointStoreState{
				Count: s.Usa.Military.Init,
			},
			Covert: &pb.PointStoreState{
				Count: s.Usa.Covert.Init,
			},
		},
		Ussr: &pb.PointStoresState{
			Political: &pb.PointStoreState{
				Count: s.Ussr.Political.Init,
			},
			Military: &pb.PointStoreState{
				Count: s.Ussr.Military.Init,
			},
			Covert: &pb.PointStoreState{
				Count: s.Ussr.Covert.Init,
			},
		},
	}
	return nil
}
