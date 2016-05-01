package points

import (
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

type PointSettings struct {
	pol_init   int32 // Political points at start
	pol_income int32 // Political points base income
	mil_init   int32 // Military points at start
	mil_income int32 // Military points base income
	cov_init   int32 // Covert points at start
	cov_income int32 // Covert points base income
}

// Settings contains parameters that do not change over the course of a game.
type Settings struct {
	usa  PointSettings // USA Settings
	ussr PointSettings // USSR Settings
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
	if settingsProto.GetPointsSettings() == nil {
		return &Settings{
			usa: PointSettings{
				pol_init:   5,
				pol_income: 5,
				mil_init:   5,
				mil_income: 5,
				cov_init:   3,
				cov_income: 3,
			},
			ussr: PointSettings{
				pol_init:   5,
				pol_income: 5,
				mil_init:   5,
				mil_income: 5,
				cov_init:   3,
				cov_income: 3,
			},
		}, nil
	}
	return &Settings{
		usa: PointSettings{
			pol_init:   settingsProto.GetPointsSettings().GetUsaSettings().PoliticalStoreInit,
			pol_income: settingsProto.GetPointsSettings().GetUsaSettings().PoliticalIncomeBase,
			mil_init:   settingsProto.GetPointsSettings().GetUsaSettings().MilitaryStoreInit,
			mil_income: settingsProto.GetPointsSettings().GetUsaSettings().MilitaryIncomeBase,
			cov_init:   settingsProto.GetPointsSettings().GetUsaSettings().CovertStoreInit,
			cov_income: settingsProto.GetPointsSettings().GetUsaSettings().CovertIncomeBase,
		},
		ussr: PointSettings{
			pol_init:   settingsProto.GetPointsSettings().GetUssrSettings().PoliticalStoreInit,
			pol_income: settingsProto.GetPointsSettings().GetUssrSettings().PoliticalIncomeBase,
			mil_init:   settingsProto.GetPointsSettings().GetUssrSettings().MilitaryStoreInit,
			mil_income: settingsProto.GetPointsSettings().GetUssrSettings().MilitaryIncomeBase,
			cov_init:   settingsProto.GetPointsSettings().GetUssrSettings().CovertStoreInit,
			cov_income: settingsProto.GetPointsSettings().GetUssrSettings().CovertIncomeBase,
		},
	}, nil
}

// InitState creates the initial state for the points mechanic.
func (s *Settings) InitState() (*State, error) {
	return &State{
		settings: s,
		usa: PointState{
			pol: s.usa.pol_init,
			mil: s.usa.mil_init,
			cov: s.usa.cov_init,
		},
		ussr: PointState{
			pol: s.ussr.pol_init,
			mil: s.ussr.mil_init,
			cov: s.ussr.cov_init,
		},
	}, nil
}
