package heat

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto *pb.GameSettings
	init  int32
	min   int32
	mxm   int32
	decay int32
	err   bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto: &pb.GameSettings{},
			init:  0, // No default
			min:   0,
			mxm:   100,
			decay: 0, // No default
			err:   false,
		},
		{
			proto: &pb.GameSettings{
				HeatSettings: &pb.HeatSettings{
					Init:  50,
					Min:   0,
					Mxm:   100,
					Decay: 5,
				},
			},
			init:  50,
			min:   0,
			mxm:   100,
			decay: 5,
			err:   false,
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
		if got, want := s.init, tc.init; got != want {
			t.Errorf("init: got %d, want %d", got, want)
		}
		if got, want := s.min, tc.min; got != want {
			t.Errorf("min: got %d, want %d", got, want)
		}
		if got, want := s.mxm, tc.mxm; got != want {
			t.Errorf("mxm: got %d, want %d", got, want)
		}
		if got, want := s.min, tc.min; got != want {
			t.Errorf("min: got %d, want %d", got, want)
		}
		if got, want := s.decay, tc.decay; got != want {
			t.Errorf("decay: got %d, want %d", got, want)
		}
	}
}

type InitStateCase struct {
	settings *Settings
	heat     int32
	err      bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			settings: &Settings{
				init:  60,
				min:   10,
				mxm:   90,
				decay: 5,
			},
			heat: 60,
			err:  false,
		},
	}
	for _, tc := range cases {
		s, err := tc.settings.InitState()
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
		if got, want := s.heat, tc.heat; got != want {
			t.Errorf("heat: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.settings; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}
