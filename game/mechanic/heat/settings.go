package heat

import (
	"errors"
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
	if settingsProto.GetHeatSettings() == nil {
		return errors.New("Could not find heat settings in game settings")
	}
	heatSettings := settingsProto.HeatSettings
	if heatSettings.Min >= heatSettings.Mxm {
		return errors.New(fmt.Sprintf("Heat minimum larger than maximum - min: ", heatSettings.Min, ", mxm: ", heatSettings.Mxm))
	}
	if heatSettings.Init >= heatSettings.Mxm {
		return errors.New(fmt.Sprintf("Heat init set to at or above maximum - mxm: ", heatSettings.Mxm, ", init: ", heatSettings.Init))
	}
	if heatSettings.Init < heatSettings.Min {
		return errors.New(fmt.Sprintf("Heat init set to below minimum - min: ", heatSettings.Min, ", init: ", heatSettings.Init))
	}
	return nil
}

func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	if settingsProto.GetHeatSettings() == nil {
		return &Settings{
			init:  0,
			min:   0,
			mxm:   100,
			decay: 0,
		}, nil
	}
	return &Settings{
		init:  settingsProto.GetHeatSettings().Init,
		min:   settingsProto.GetHeatSettings().Min,
		mxm:   settingsProto.GetHeatSettings().Mxm,
		decay: settingsProto.GetHeatSettings().Decay,
	}, nil
}

func (s *Settings) InitState() (*State, error) {
	return &State{
		settings: s,
		heat:     s.init,
	}, nil
}
