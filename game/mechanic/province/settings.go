package province

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

// Settings contains parameters that do not change over the course of a game.
type Settings struct {
	Provinces map[pb.ProvinceId]*ProvSettings
}

// Settings for an individual province
type ProvSettings struct {
	id pb.ProvinceId // Province id enum
}

func validate(settingsProto *pb.GameSettings) error {
	// TODO: Validate proper proto input
	return nil
}

// NewSettings creates a new settings struct from the GameSettings message.
func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}

	provs := make(map[pb.ProvinceId]*ProvSettings)
	// Contains map of all individual province settings
	for _, p := range settingsProto.GetProvincesSettings().GetProvinceSettings() {
		provs[p.GetId()] = &ProvSettings{
			id: p.GetId(),
		}
	}

	return &Settings{
		Provinces: provs,
	}, nil
}

// InitState creates the initial state for the province mechanic.
func (s *Settings) InitState() (*State, error) {
	provs := make(map[pb.ProvinceId]*ProvState)
	// Contains map of all individual province states
	for _, p := range s.Provinces {
		provs[p.Id()] = &ProvState{
			id: p.Id(),
		}
	}

	return &State{
		settings:  s,
		Provinces: provs,
	}, nil
}

// GETTERS

func (s *ProvSettings) Id() pb.ProvinceId {
	return s.id
}

func (s *Settings) Get(id pb.ProvinceId) *ProvSettings {
	return s.Provinces[id]
}
