package heat

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
