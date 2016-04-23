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
	id             pb.ProvinceId   // Province id enum
	label          string          // Text label of the country
	adjacencies    []pb.ProvinceId // List of adjacent countries
	stability_base int32           // Base stability
	region         pb.Region       // Region
	coastal        bool            // Coastal

	// Initialization
	init_influence  int32         // Influence
	init_government pb.Government // Government enum
	init_leader     string        // Leader
	// init_dissidents TYPE // Dissidents
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

	// Contains map of all individual province settings
	provs := make(map[pb.ProvinceId]*ProvSettings)
	for _, p := range settingsProto.GetProvincesSettings().GetProvinceSettings() {
		provs[p.Id] = &ProvSettings{
			id:              p.Id,
			label:           p.Label,
			adjacencies:     p.Adjacency,
			stability_base:  p.StabilityBase,
			region:          p.Region,
			coastal:         p.Coastal,
			init_influence:  p.InitInfluence,
			init_government: p.InitGovernment,
			init_leader:     p.InitLeader,
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
			id:         p.Id(),
			influence:  p.initInfl(),
			government: p.initGov(),
			leader:     p.initLeader(),
		}
	}

	return &State{
		settings:  s,
		Provinces: provs,
	}, nil
}

// GETTERS

func (s *Settings) Get(id pb.ProvinceId) *ProvSettings {
	return s.Provinces[id]
}

func (s *ProvSettings) Id() pb.ProvinceId {
	return s.id
}

func (s *ProvSettings) Label() string {
	return s.label
}

func (s *ProvSettings) Adjacencies() []pb.ProvinceId {
	return s.adjacencies
}

func (s *ProvSettings) BaseStability() int32 {
	return s.stability_base
}

func (s *ProvSettings) Region() pb.Region {
	return s.region
}

func (s *ProvSettings) isCoastal() bool {
	return s.coastal
}

// Initialization

func (s *ProvSettings) initInfl() int32 {
	return s.init_influence
}

func (s *ProvSettings) initGov() pb.Government {
	return s.init_government
}

func (s *ProvSettings) initLeader() string {
	return s.init_leader
}
