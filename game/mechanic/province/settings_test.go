package province

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto           *pb.GameSettings
	id              pb.ProvinceId
	label           string
	adjacency       []pb.ProvinceId
	stability_base  int32
	region          pb.Region
	coastal         bool
	init_influence  int32
	init_government pb.Government
	init_leader     string
	err             bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto: &pb.GameSettings{
				ProvincesSettings: &pb.ProvincesSettings{
					ProvinceSettings: []*pb.ProvinceSettings{
						&pb.ProvinceSettings{
							Id:             pb.ProvinceId_ROMANIA,
							Label:          "Romania",
							Adjacency:      []pb.ProvinceId{pb.ProvinceId_P_USSR, pb.ProvinceId_HUNGARY},
							StabilityBase:  2,
							Region:         pb.Region_EASTERN_EUROPE,
							Coastal:        false,
							InitInfluence:  -1,
							InitGovernment: pb.Government_COMMUNISM,
							InitLeader:     "David Mihai",
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
			},
			err:   false,
			id:    pb.ProvinceId_ROMANIA,
			label: "Romania",
			adjacency: []pb.ProvinceId{
				pb.ProvinceId_P_USSR, pb.ProvinceId_HUNGARY,
			},
			stability_base:  2,
			region:          pb.Region_EASTERN_EUROPE,
			coastal:         false,
			init_influence:  -1,
			init_government: pb.Government_COMMUNISM,
			init_leader:     "David Mihai",
		},
	}
	for _, tc := range cases {
		s, err := NewSettings(tc.proto)
		if got, want := err != nil, tc.err; got != want {
			msg := map[bool]string{
				true:  "error",
				false: "no error",
			}
			t.Errorf("err: got %s, want %s", msg[got], msg[want])
			continue
		}
		if tc.err {
			continue
		}
		if got, want := s.Get(tc.id).Id(), tc.id; got != want {
			t.Errorf("id: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).Label(), tc.label; got != want {
			t.Errorf("label: got %d, want %d", got, want)
		}
		/*if got, want := s.Get(tc.id).Adjacencies(), tc.adjacency; got != want {
			t.Errorf("adjacency: got %d, want %d", got, want)
		}*/
		if got, want := s.Get(tc.id).BaseStability(), tc.stability_base; got != want {
			t.Errorf("base stability: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).Region(), tc.region; got != want {
			t.Errorf("region: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).isCoastal(), tc.coastal; got != want {
			t.Errorf("coastal: got %d, want %d", got, want)
		}
	}
}

type InitStateCase struct {
	s          *Settings
	id         pb.ProvinceId
	influence  int32
	government pb.Government
	leader     string
	err        bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			s: &Settings{
				Provinces: map[pb.ProvinceId]*ProvSettings{
					pb.ProvinceId_ROMANIA: &ProvSettings{
						id:              pb.ProvinceId_ROMANIA,
						init_influence:  -1,
						init_government: pb.Government_COMMUNISM,
						init_leader:     "David Mihai",
					},
				},
			},
			id:         pb.ProvinceId_ROMANIA,
			influence:  -1,
			government: pb.Government_COMMUNISM,
			leader:     "David Mihai",
			err:        false,
		},
	}
	for _, tc := range cases {
		s, err := tc.s.InitState()
		if got, want := err != nil, tc.err; got != want {
			msg := map[bool]string{
				true:  "error",
				false: "no error",
			}
			t.Errorf("err: got %s, want %s", msg[got], msg[want])
			continue
		}
		if tc.err {
			continue
		}
		if got, want := s.Get(tc.id).Id(), tc.id; got != want {
			t.Errorf("id: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.s; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).Infl(), tc.influence; got != want {
			t.Errorf("influence: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).Gov(), tc.government; got != want {
			t.Errorf("government: got %d, want %d", got, want)
		}
		if got, want := s.Get(tc.id).Leader(), tc.leader; got != want {
			t.Errorf("leader: got %d, want %d", got, want)
		}
	}
}
