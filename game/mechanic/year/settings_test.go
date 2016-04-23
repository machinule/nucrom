package year

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto *pb.GameSettings
	init  int32
	incr  int32
	err   bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto: &pb.GameSettings{},
			init:  1948,
			incr:  1,
			err:   false,
		},
		{
			proto: &pb.GameSettings{
				YearSettings: &pb.YearSettings{
					InitYear: 43,
				},
			},
			init: 43,
			incr: 1,
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
		if got, want := s.incr, tc.incr; got != want {
			t.Errorf("incr: got %d, want %d", got, want)
		}
	}
}

type InitStateCase struct {
	s    *Settings
	year int32
	err  bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			s: &Settings{
				init: 34,
				incr: 1,
			},
			year: 34,
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
		if got, want := s.year, tc.year; got != want {
			t.Errorf("year: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.s; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}
