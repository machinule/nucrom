// This is the state for the province mechanic, including the states of each individual province
package province

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

// The state of the province mechanic.
type State struct {
	settings  *Settings
	Provinces map[pb.ProvinceId]*ProvState
	Conflicts map[pb.ProvinceId]*Conflict // Keys = locations
	Dormant   map[pb.ProvinceId]*Conflict
	Possible  map[pb.ProvinceId]*Conflict
}

// The state of a single province
type ProvState struct {
	id         pb.ProvinceId // Province id enum
	influence  int32
	government pb.Government
	occupier   pb.ProvinceId
	leader     string
	dissidents Dissidents
}

// Conflict structure
type Conflict struct {
	name          string // Name of conflict
	conflict_type pb.ConflictType
	goal          int32 // Goal to reach to end conflict
	attackers     Faction
	defenders     Faction
	length        int32
	base_chance   int32
	locations     []pb.ProvinceId
}

// Side of a conflict
type Faction struct {
	members   []pb.ProvinceId
	supporter pb.Player
	progress  int32
	rebels    Dissidents
}

// Dissidents
type Dissidents struct {
	gov    pb.Government
	leader string
}

// NewState creates a new state from the GameState message and the previous state.
func NewState(stateProto *pb.GameState, settings *Settings) (*State, error) {
	if settings == nil {
		return nil, fmt.Errorf("received nil Settings, unable to continue.")
	}
	return &State{
		settings: settings,
		// TODO(david): produce a State object from stateProto.
	}, nil
}

// Helper for marshaling conflicts
func MarshalConflict(c *Conflict) pb.Conflict {
	return pb.Conflict{
		Name: c.Name(),
		Type: c.Type(),
		Goal: c.Goal(),
		Attackers: &pb.Faction{
			Ids:       c.Attackers(),
			Supporter: c.Att_Supporter(),
			Progress:  c.Att_Progress(),
			Rebels: &pb.Dissidents{
				Gov:    c.Rebels().Gov(),
				Leader: c.Rebels().Leader(),
			},
		},
		Defenders: &pb.Faction{
			Ids:       c.Defenders(),
			Supporter: c.Def_Supporter(),
			Progress:  c.Def_Progress(),
			// REBELS CANNOT BE DEFENDERS
		},
		Length:     c.Length(),
		BaseChance: c.BaseChance(),
		Locations:  c.Locations(),
	}
}

// Marshal fills in the GameState proto with the current state.
func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetProvincesState() == nil {
		stateProto.ProvincesState = &pb.ProvincesState{}
	}
	var provs []*pb.ProvinceState
	var conflicts []*pb.Conflict
	var dormant []*pb.Conflict
	var possible []*pb.Conflict
	for _, p := range s.Provinces {
		provs = append(provs, &pb.ProvinceState{
			Id:        p.Id(),
			Influence: p.Infl(),
			Gov:       p.Gov(),
			Occupier:  p.Occupier(),
			Leader:    p.Leader(),
			Dissidents: &pb.Dissidents{
				Gov:    p.Dissidents().Gov(),
				Leader: p.Dissidents().Leader(),
			},
		})
	}
	for _, c := range s.Conflicts {
		mc := MarshalConflict(c)
		conflicts = append(conflicts, &mc)
	}
	for _, c := range s.Dormant {
		mc := MarshalConflict(c)
		dormant = append(dormant, &mc)
	}
	for _, c := range s.Possible {
		mc := MarshalConflict(c)
		possible = append(possible, &mc)
	}
	stateProto.ProvincesState = &pb.ProvincesState{
		ProvinceStates: provs,
		Conflicts: &pb.ConflictsState{
			Active:   conflicts,
			Dormant:  dormant,
			Possible: possible,
		},
	}
	return nil
}

// GETTERS

// Provinces

func (s *State) Settings() *Settings {
	return s.settings
}

func (s *State) Get(id pb.ProvinceId) *ProvState {
	return s.Provinces[id]
}

func (s *ProvState) Id() pb.ProvinceId {
	return s.id
}

func (s *ProvState) Infl() int32 {
	return s.influence
}

func (s *ProvState) Gov() pb.Government {
	return s.government
}

func (s *ProvState) Occupier() pb.ProvinceId {
	return s.occupier
}

func (s *ProvState) Leader() string {
	return s.leader
}

func (s *ProvState) Dissidents() *Dissidents {
	return &s.dissidents
}

// Dissidents

func (d *Dissidents) Gov() pb.Government {
	return d.gov
}

func (d *Dissidents) Leader() string {
	return d.leader
}

// Conflict

func (s *State) Conflict(id pb.ProvinceId) *Conflict {
	return s.Conflicts[id]
}

func (c *Conflict) Name() string {
	return c.name
}

func (c *Conflict) Type() pb.ConflictType {
	return c.conflict_type
}

func (c *Conflict) Goal() int32 {
	return c.goal
}

func (c *Conflict) Attackers() []pb.ProvinceId {
	return c.attackers.members
}

func (c *Conflict) Defenders() []pb.ProvinceId {
	return c.defenders.members
}

func (c *Conflict) Att_Progress() int32 {
	return c.attackers.progress
}

func (c *Conflict) Def_Progress() int32 {
	return c.defenders.progress
}

func (c *Conflict) Def_Supporter() pb.Player {
	return c.defenders.supporter
}

func (c *Conflict) Att_Supporter() pb.Player {
	return c.attackers.supporter
}

func (c *Conflict) Length() int32 {
	return c.length
}

func (c *Conflict) BaseChance() int32 {
	return c.base_chance
}

func (c *Conflict) Locations() []pb.ProvinceId {
	return c.locations
}

func (c *Conflict) Rebels() *Dissidents {
	return &(c.attackers).rebels
}

// SETTERS

func (s *ProvState) ApplyInfl(player pb.Player, magnitude int32) {
	delta := magnitude
	if player == pb.Player_USSR {
		delta = delta * -1
	}
	s.influence = s.influence + delta
}

func (s *ProvState) SetGov(gov pb.Government) {
	s.government = gov
}

func (s *ProvState) SetLeader(name string) {
	s.leader = name
}

func (s *ProvState) SetDissidents(gov pb.Government, ldr string) {
	s.dissidents = Dissidents{
		gov:    gov,
		leader: ldr,
	}
}

func (s *ProvState) RemoveDissidents() {
	s.dissidents = Dissidents{}
}
