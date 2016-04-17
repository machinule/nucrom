package year

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

type settings struct {
	init int32 // Starting year.
	incr int32 // Number of years to increment per turn.
}

func validate(settingsProto *pb.GameSettings) error {
	// There cannot be any errors in configuring the year mechanic, as the default is appropriate.
	return nil
}

func NewSettings(settingsProto *pb.GameSettings) (*settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	return &settings{
		init: settingsProto.GetYearSysSettings().GetInitYear(),
		incr: 1,
	}, nil
}

func (s *settings) InitState() (*state, error) {
	return &state{
		s:    s,
		year: s.init,
	}, nil
}
