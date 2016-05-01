package points

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto           *pb.GameSettings
	usa_pol_init    int32
	usa_pol_income  int32
	usa_mil_init    int32
	usa_mil_income  int32
	usa_cov_init    int32
	usa_cov_income  int32
	ussr_pol_init   int32
	ussr_pol_income int32
	ussr_mil_init   int32
	ussr_mil_income int32
	ussr_cov_init   int32
	ussr_cov_income int32
	err             bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto:           &pb.GameSettings{},
			usa_pol_init:    5,
			usa_pol_income:  5,
			usa_mil_init:    5,
			usa_mil_income:  5,
			usa_cov_init:    3,
			usa_cov_income:  3,
			ussr_pol_init:   5,
			ussr_pol_income: 5,
			ussr_mil_init:   5,
			ussr_mil_income: 5,
			ussr_cov_init:   3,
			ussr_cov_income: 3,
			err:             false,
		},
		{
			proto: &pb.GameSettings{
				PointsSettings: &pb.PointsSettings{
					UsaSettings: &pb.PointSettings{
						PoliticalStoreInit:  1,
						PoliticalIncomeBase: 2,
						MilitaryStoreInit:   3,
						MilitaryIncomeBase:  4,
						CovertStoreInit:     5,
						CovertIncomeBase:    6,
					},
					UssrSettings: &pb.PointSettings{
						PoliticalStoreInit:  4,
						PoliticalIncomeBase: 4,
						MilitaryStoreInit:   3,
						MilitaryIncomeBase:  3,
						CovertStoreInit:     2,
						CovertIncomeBase:    2,
					},
				},
			},
			usa_pol_init:    1,
			usa_pol_income:  2,
			usa_mil_init:    3,
			usa_mil_income:  4,
			usa_cov_init:    5,
			usa_cov_income:  6,
			ussr_pol_init:   4,
			ussr_pol_income: 4,
			ussr_mil_init:   3,
			ussr_mil_income: 3,
			ussr_cov_init:   2,
			ussr_cov_income: 2,
			err:             false,
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
		if got, want := s.usa.pol_init, tc.usa_pol_init; got != want {
			t.Errorf("usa_pol_init: got %d, want %d", got, want)
		}
		if got, want := s.usa.pol_income, tc.usa_pol_income; got != want {
			t.Errorf("usa_pol_income: got %d, want %d", got, want)
		}
		if got, want := s.usa.mil_init, tc.usa_mil_init; got != want {
			t.Errorf("usa_mil_init: got %d, want %d", got, want)
		}
		if got, want := s.usa.mil_income, tc.usa_mil_income; got != want {
			t.Errorf("usa_mil_income: got %d, want %d", got, want)
		}
		if got, want := s.usa.cov_init, tc.usa_cov_init; got != want {
			t.Errorf("usa_cov_init: got %d, want %d", got, want)
		}
		if got, want := s.usa.cov_income, tc.usa_cov_income; got != want {
			t.Errorf("usa_cov_income: got %d, want %d", got, want)
		}
		if got, want := s.ussr.pol_init, tc.ussr_pol_init; got != want {
			t.Errorf("ussr_pol_init: got %d, want %d", got, want)
		}
		if got, want := s.ussr.pol_income, tc.ussr_pol_income; got != want {
			t.Errorf("ussr_pol_income: got %d, want %d", got, want)
		}
		if got, want := s.ussr.mil_init, tc.ussr_mil_init; got != want {
			t.Errorf("ussr_mil_init: got %d, want %d", got, want)
		}
		if got, want := s.ussr.mil_income, tc.ussr_mil_income; got != want {
			t.Errorf("ussr_mil_income: got %d, want %d", got, want)
		}
		if got, want := s.ussr.cov_init, tc.ussr_cov_init; got != want {
			t.Errorf("ussr_cov_init: got %d, want %d", got, want)
		}
		if got, want := s.ussr.cov_income, tc.ussr_cov_income; got != want {
			t.Errorf("ussr_cov_income: got %d, want %d", got, want)
		}
	}
}

type InitStateCase struct {
	s        *Settings
	usa_pol  int32
	usa_mil  int32
	usa_cov  int32
	ussr_pol int32
	ussr_mil int32
	ussr_cov int32
	err      bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			s: &Settings{
				usa: PointSettings{
					pol_init: 9,
					mil_init: 7,
					cov_init: 5,
				},
				ussr: PointSettings{
					pol_init: 2,
					mil_init: 5,
					cov_init: 8,
				},
			},
			usa_pol:  9,
			usa_mil:  7,
			usa_cov:  5,
			ussr_pol: 2,
			ussr_mil: 5,
			ussr_cov: 8,
			err:      false,
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
		if got, want := s.usa.pol, tc.usa_pol; got != want {
			t.Errorf("usa pol: got %d, want %d", got, want)
		}
		if got, want := s.usa.mil, tc.usa_mil; got != want {
			t.Errorf("usa mil: got %d, want %d", got, want)
		}
		if got, want := s.usa.cov, tc.usa_cov; got != want {
			t.Errorf("usa cov: got %d, want %d", got, want)
		}
		if got, want := s.ussr.pol, tc.ussr_pol; got != want {
			t.Errorf("ussr pol: got %d, want %d", got, want)
		}
		if got, want := s.ussr.mil, tc.ussr_mil; got != want {
			t.Errorf("ussr mil: got %d, want %d", got, want)
		}
		if got, want := s.ussr.cov, tc.ussr_cov; got != want {
			t.Errorf("ussr cov: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.s; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}
