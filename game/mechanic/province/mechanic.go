// The province mechanic keeps track of the provinces that make up the world map
package province

import (
	pb "github.com/machinule/nucrom/proto/gen"
)

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

func (s *State) GetAlly(id pb.ProvinceId) pb.Player {
	net_stab := s.GetNetStability(id)
	infl := s.Get(id).Infl()
	if infl > net_stab {
		return pb.Player_USA
	} else if infl < -net_stab {
		return pb.Player_USSR
	} else {
		return pb.Player_NEITHER
	}
}

func (s *State) IsAllied(id pb.ProvinceId, player pb.Player) bool {
	if s.GetAlly(id) == player {
		return true
	}
	return false
}
