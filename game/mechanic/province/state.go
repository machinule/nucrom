// The province mechanic keeps track of the provinces that make up the world map
package province

import (
	"fmt"
	//TODO: "github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
)

// The state of the province mechanic.
type State struct {
	settings  *Settings
	Provinces map[pb.ProvinceId]*ProvState
}

// The state of a single province
type ProvState struct {
	id pb.ProvinceId // Province id enum
}

// NewState creates a new state from the GameState message and the previous state.
func NewState(stateProto *pb.GameState, prev *State) (*State, error) {
	if prev == nil {
		return nil, fmt.Errorf("recieved nil previous state, unable to propogate settings.")
	}
	return &State{
		settings:  prev.settings,
		Provinces: prev.Provinces,
	}, nil
}

// Marshal fills in the GameState proto with the current state.
func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetProvincesState() == nil {
		stateProto.ProvincesState = &pb.ProvincesState{}
	}
	// TODO: Populate fields into new GameState
	return nil
}

// GETTERS

func (s *State) Get(id pb.ProvinceId) *ProvState {
	return s.Provinces[id]
}

func (s *ProvState) Id() pb.ProvinceId {
	return s.id
}
