package heat

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

type Settings struct {
	init  int32 // Starting heat value
	min   int32 // Minimum value of heat
	mxm   int32 // Maximum value of heat
	decay int32 // Amount heat decays per turn
}

func validate(settingsProto *pb.GameSettings) error {
	// There cannot be any errors in configuring the heat mechanic, as the default is appropriate.
	return nil
}

func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	return &Settings{
		init:  settingsProto.GetHeatSysSettings().GetInit(),
		min:   settingsProto.GetHeatSysSettings().GetMin(),
		mxm:   settingsProto.GetHeatSysSettings().GetMxm(),
		decay: settingsProto.GetHeatSysSettings().GetDecay(),
	}, nil
}

func (s *Settings) InitState() (*State, error) {
	return &State{
		settings: s,
		heat:     s.init,
	}, nil
}
