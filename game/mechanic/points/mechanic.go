package points

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

// ACTIONS

func (s *State) ChngPOL(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.pol = s.ussr.pol + magnitude
	} else if player == pb.Player_USA {
		s.usa.pol = s.usa.pol + magnitude
	}
}

func (s *State) ChngMIL(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.mil = s.ussr.mil + magnitude
	} else if player == pb.Player_USA {
		s.usa.mil = s.usa.mil + magnitude
	}
}

func (s *State) ChngCOV(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.cov = s.ussr.cov + magnitude
	} else if player == pb.Player_USA {
		s.usa.cov = s.usa.cov + magnitude
	}
}
