// The province mechanic keeps track of the provinces that make up the world map
package province

import (
	pb "github.com/machinule/nucrom/proto/gen"
)

type Mechanic struct {
	Settings *Settings
	State    *State
}

func (m *Mechanic) Initialize(settings *pb.GameSettings) error {
	var err error
	m.Settings, err = NewSettings(settings)
	if err != nil {
		return err
	}
	m.State, err = m.Settings.InitState()
	if err != nil {
		return err
	}
	return nil
}

func (m *Mechanic) SetState(state *pb.GameState) error {
	var err error
	m.State, err = NewState(state, m.Settings)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mechanic) GetState(state *pb.GameState) error {
	err := m.State.Marshal(state)
	if err != nil {
		return err
	}
	return nil
}

// HELPERS

func ToPlayer(id pb.ProvinceId) pb.Player {
	if id == pb.ProvinceId_P_USSR {
		return pb.Player_USSR
	} else if id == pb.ProvinceId_P_USA {
		return pb.Player_USA
	} else {
		return pb.Player_NEITHER
	}
}

// QUERIES

// Returns the net stability when accounting for modifiers
func (s *State) GetNetStability(id pb.ProvinceId) int32 {
	return s.Settings().Get(id).BaseStability() + s.GetStabilityMod(id)
}

// Get the stability modifiers that are affecting the province (independent of base stability)
func (s *State) GetStabilityMod(id pb.ProvinceId) int32 {
	var mod int32 = 0
	gov := s.Get(id).Gov()
	if gov == pb.Government_DEMOCRACY ||
		gov == pb.Government_AUTOCRACY ||
		gov == pb.Government_COMMUNISM {
		mod++
	}
	if s.Get(id).Leader() != "" {
		mod++
	}
	// TODO: Dissidents
	return mod
}

// Gets superpower ally of province; returns neither if applicable
func (s *State) GetAlly(id pb.ProvinceId) pb.Player {
	net_stab := s.GetNetStability(id)
	infl := s.Get(id).Infl()
	if infl >= net_stab {
		return pb.Player_USA
	} else if infl <= -net_stab {
		return pb.Player_USSR
	} else {
		return pb.Player_NEITHER
	}
}

// Checks if province currently has a superpower ally
func (s *State) IsAllied(id pb.ProvinceId, player pb.Player) bool {
	if s.GetAlly(id) == player {
		return true
	}
	return false
}

// ACTIONS

// Applys influence on a province
func (s *State) Infl(id pb.ProvinceId, player pb.Player, magnitude int32) {
	s.Get(id).ApplyInfl(player, magnitude)
}

// Sets the government of a province
func (s *State) SetGov(id pb.ProvinceId, gov pb.Government) {
	s.Get(id).SetGov(gov)
}

// Sets the leader of a province ("" string for no leader)
func (s *State) SetLeader(id pb.ProvinceId, name string) {
	s.Get(id).SetLeader(name)
}

// Adds dissidents to a province
func (s *State) SetDissidents(id pb.ProvinceId, gov pb.Government, ldr string) {
	s.Get(id).SetDissidents(gov, ldr)
}

// Removes dissidents
func (s *State) RemoveDissidents(id pb.ProvinceId) {
	s.Get(id).RemoveDissidents()
}
