package setup

import (
	"github.com/machinule/nucrom/proto/gen"
)

// TODO(anyone please): consistent naming.

func CreateGameSettings() *pb.GameSettings {
	ret := &pb.GameSettings{
		HeatSettings: &pb.HeatSettings{
			Init:  50,
			Min:   0,
			Mxm:   100,
			Decay: 5,
		},
		PointsSettings: &pb.PointsSettings{
			UsaSettings: &pb.PointSettings{
				PoliticalStoreInit:  1,
				PoliticalIncomeBase: 1,
				MilitaryStoreInit:   1,
				MilitaryIncomeBase:  1,
				CovertStoreInit:     1,
				CovertIncomeBase:    1,
			},
			UssrSettings: &pb.PointSettings{
				PoliticalStoreInit:  1,
				PoliticalIncomeBase: 1,
				MilitaryStoreInit:   1,
				MilitaryIncomeBase:  1,
				CovertStoreInit:     1,
				CovertIncomeBase:    1,
			},
		},
		PseudorandomSettings: &pb.PseudorandomSettings{
			InitSeed: 0,
		},
		ProvincesSettings: &pb.ProvincesSettings{
			ProvinceSettings: []*pb.ProvinceSettings{
				&pb.ProvinceSettings{
					Id:             pb.ProvinceId_IRAN,
					Label:          "Iran",
					Adjacency:      []pb.ProvinceId{pb.ProvinceId_P_USSR, pb.ProvinceId_IRAQ, pb.ProvinceId_AFGHANISTAN, pb.ProvinceId_PAKISTAN},
					StabilityBase:  2,
					Region:         pb.Region_MIDDLE_EAST,
					Coastal:        true,
					InitInfluence:  3,
					InitGovernment: pb.Government_AUTOCRACY,
					InitLeader:     "Reza Pahlavi",
				},
				&pb.ProvinceSettings{ //TODO: Test conflict
					Id:             pb.ProvinceId_AFGHANISTAN,
					Label:          "Afghanistan",
					Adjacency:      []pb.ProvinceId{pb.ProvinceId_P_USSR, pb.ProvinceId_IRAN, pb.ProvinceId_PAKISTAN},
					StabilityBase:  1,
					Region:         pb.Region_SOUTH_ASIA,
					Coastal:        false,
					InitInfluence:  -1,
					InitGovernment: pb.Government_WEAK,
				},
				&pb.ProvinceSettings{
					Id:             pb.ProvinceId_PAKISTAN,
					Label:          "Pakistan",
					Adjacency:      []pb.ProvinceId{pb.ProvinceId_IRAN, pb.ProvinceId_INDIA, pb.ProvinceId_AFGHANISTAN},
					StabilityBase:  2,
					Region:         pb.Region_SOUTH_ASIA,
					Coastal:        true,
					InitInfluence:  3,
					InitGovernment: pb.Government_WEAK,
				},
				&pb.ProvinceSettings{
					Id:             pb.ProvinceId_INDIA,
					Label:          "India",
					Adjacency:      []pb.ProvinceId{pb.ProvinceId_BANGLADESH, pb.ProvinceId_BURMA, pb.ProvinceId_PAKISTAN, pb.ProvinceId_CHINA},
					StabilityBase:  2,
					Region:         pb.Region_SOUTH_ASIA,
					Coastal:        true,
					InitInfluence:  0,
					InitGovernment: pb.Government_WEAK,
				},
			},
			ConflictsSettings: &pb.ConflictsSettings{
				BaseChanceCivil:        150000,
				BaseChanceConventional: 150000,
				BaseChanceAction:       150000,
				BaseChanceColonial:     150000,
				GoalCivil:              2,
				GoalConventional:       2,
				GoalAction:             2,
				GoalColonial:           2,
			},
		},
		YearSettings: &pb.YearSettings{
			InitYear: 1948,
		},
	}
	return ret
}
