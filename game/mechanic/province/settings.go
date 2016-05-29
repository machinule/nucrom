package province

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

// Settings contains parameters that do not change over the course of a game.
type Settings struct {
	Provinces map[pb.ProvinceId]*ProvSettings
	Conflict  ConflictSettings
}

// Conflict settings
type ConflictSettings struct {
	init_active   []*Conflict
	init_dormant  []*Conflict
	init_possible []*Conflict
    base_chance_civil int32
    base_chance_conventional int32
    base_chance_action int32
    base_chance_colonial int32
    goal_civil int32
    goal_conventional int32
    goal_action int32
    goal_colonial int32
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
        Conflict: ConflictSettings{
            base_chance_civil: settingsProto.GetProvincesSettings().GetConflictsSettings().BaseChanceCivil,
            base_chance_conventional: settingsProto.GetProvincesSettings().GetConflictsSettings().BaseChanceConventional,
            base_chance_action: settingsProto.GetProvincesSettings().GetConflictsSettings().BaseChanceAction,
            base_chance_colonial: settingsProto.GetProvincesSettings().GetConflictsSettings().BaseChanceColonial,
            goal_civil: settingsProto.GetProvincesSettings().GetConflictsSettings().GoalCivil,
            goal_conventional: settingsProto.GetProvincesSettings().GetConflictsSettings().GoalConventional,
            goal_action: settingsProto.GetProvincesSettings().GetConflictsSettings().GoalAction,
            goal_colonial: settingsProto.GetProvincesSettings().GetConflictsSettings().GoalColonial,
        },
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
		Conflicts: make(map[pb.ProvinceId]*Conflict),
		Dormant:   make(map[pb.ProvinceId]*Conflict),
		Possible:  make(map[pb.ProvinceId]*Conflict),
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

func (s *Settings) GetConflictBaseChance(conflict_type pb.ConflictType) int32 {
    switch conflict_type {
        case pb.ConflictType_CIVIL_WAR:
            return s.Conflict.base_chance_civil
        case pb.ConflictType_CONVENTIONAL_WAR:
            return s.Conflict.base_chance_conventional
        case pb.ConflictType_MILITARY_ACTION:
            return s.Conflict.base_chance_action
        case pb.ConflictType_COLONIAL_WAR:
            return s.Conflict.base_chance_colonial
    }
    return -1
}

func (s *Settings) GetConflictGoal(conflict_type pb.ConflictType) int32 {
    switch conflict_type {
        case pb.ConflictType_CIVIL_WAR:
            return s.Conflict.goal_civil
        case pb.ConflictType_CONVENTIONAL_WAR:
            return s.Conflict.goal_conventional
        case pb.ConflictType_MILITARY_ACTION:
            return s.Conflict.goal_action
        case pb.ConflictType_COLONIAL_WAR:
            return s.Conflict.goal_colonial
    }
    return -1
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
