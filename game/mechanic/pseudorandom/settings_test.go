package pseudorandom

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto *pb.GameSettings
	init  int64
	err   bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto: &pb.GameSettings{},
			init:  1, // Default
			err:   false,
		},
		{
			proto: &pb.GameSettings{
				PseudorandomSettings: &pb.PseudorandomSettings{
					InitSeed: 42,
				},
			},
			init: 42,
			err:  false,
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
	}
}

type InitStateCase struct {
	s    *Settings
	seed int64
	err  bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			s: &Settings{
				init: 10,
			},
			seed: 10,
			err:  false,
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
		if got, want := s.seed, tc.seed; got != want {
			t.Errorf("seed: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.s; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}
