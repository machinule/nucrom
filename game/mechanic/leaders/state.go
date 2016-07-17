package leaders

import (
	"fmt"
	"github.com/machinule/nucrom/proto/gen"
)

type State struct {
	settings    *Settings
	leaders     map[string]*LeaderState
	active      map[pb.ProvinceId]string
	deactivated map[pb.ProvinceId]string
	dissidents  map[pb.ProvinceId]string
}

type LeaderState struct {
	l_type     pb.LeaderType
	elected    bool
	birth_year int32
}

func NewState(stateProto *pb.GameState, prev *State) (*State, error) {
	if prev == nil {
		return nil, fmt.Errorf("received nil previous state, unable to propogate settings.")
	}
	if stateProto.GetLeadersState() == nil {
		return nil, fmt.Errorf("received nil LeadersState, unable to continue.")
	}

	return &State{
		settings:    prev.settings,
		leaders:     prev.leaders,
		active:      prev.active,
		deactivated: prev.deactivated,
		dissidents:  prev.dissidents,
	}, nil
}

func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetLeadersState() == nil {
		stateProto.LeadersState = &pb.LeadersState{}
	}
	p_l_state_holder := make(map[pb.ProvinceId]*pb.ProvinceLeaderState)
	for p, l := range s.deactivated {
		l_state := &pb.LeaderState{
			Name:      l,
			Type:      s.leaders[l].l_type,
			Elected:   s.leaders[l].elected,
			BirthYear: s.leaders[l].birth_year,
		}
		if p_l_state_holder[p] != nil {
			p_l_state_holder[p].Deactivated = []*pb.LeaderState{
				l_state,
			}
			p_l_state_holder[p].Id = p
		} else {
			p_l_state_holder[p].Deactivated = append(p_l_state_holder[p].Deactivated, l_state)
		}
	}
	for p, l := range s.active {
		l_state := &pb.LeaderState{
			Name:      l,
			Type:      s.leaders[l].l_type,
			Elected:   s.leaders[l].elected,
			BirthYear: s.leaders[l].birth_year,
		}
		if p_l_state_holder[p] != nil {
			p_l_state_holder[p].Active = l_state
		} else {
			p_l_state_holder[p] = &pb.ProvinceLeaderState{
				Active: l_state,
				Id:     p,
			}
		}
	}
	for p, l := range s.dissidents {
		l_state := &pb.LeaderState{
			Name:      l,
			Type:      s.leaders[l].l_type,
			Elected:   s.leaders[l].elected,
			BirthYear: s.leaders[l].birth_year,
		}
		if p_l_state_holder[p] != nil {
			p_l_state_holder[p].Dissident = l_state
		} else {
			p_l_state_holder[p] = &pb.ProvinceLeaderState{
				Dissident: l_state,
				Id:        p,
			}
		}
	}
	var leaders_state []*pb.ProvinceLeaderState
	for _, a := range p_l_state_holder {
		leaders_state = append(leaders_state, a)
	}
	stateProto.LeadersState = &pb.LeadersState{
		ProvinceLeaderStates: leaders_state,
	}
	return nil
}

// GETTERS/HELPERS

func (s *State) GetActiveLeaderType(id pb.ProvinceId) pb.LeaderType {
	return s.leaders[s.active[id]].l_type
}

func (s *State) HasLeader(id pb.ProvinceId) bool {
	ret := s.GetActiveLeaderType(id)
	if ret == pb.LeaderType_NO_LEADER {
		return false
	} else {
		return true
	}
}

// ACTIONS

/*
// Form a new leader - will either look in the settings or, if one is not available, generates a new one
func (s *State) NewLeader(id pb.ProvinceId, year int32, elected bool) {
    for _, l := range s.settings.leaders[id] {
        if l.start_year <= year && l.end_year >= year {
            s.active[id] = &LeaderState{
                l_type: l.Type,
            }
            return
        }
    }
    s.active[id] = GenerateLeader(id, elected)
}

// Generates a new leader
func (s *State) GenerateLeader(id pb.ProvinceId, elected bool) {
    s.active[id] = "Louis Francis Victor Albert Charles Mountbatten"
}
*/
