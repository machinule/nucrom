package year

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewStateCase struct {
	proto    *pb.GameState
	settings *Settings
	year     int32
	err      bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				YearState: &pb.YearState{
					Year: 34,
				},
			},
			settings: &Settings{},
			year:     34,
			err:      false,
		},
		{
			proto: &pb.GameState{
				YearState: &pb.YearState{
					Year: 34,
				},
			},
			settings: nil,
			err:      true,
		},
	}
	for _, tc := range cases {
		s, err := NewState(tc.proto, tc.settings)
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
		if got, want := s.settings, tc.settings; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}

type MarshalCase struct {
	s     State
	proto *pb.GameState
	want  *pb.GameState
	err   bool
}

func TestMarshal(t *testing.T) {
	cases := []MarshalCase{
		{
			s: State{
				year: 364,
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				YearState: &pb.YearState{
					Year: 364,
				},
			},
			err: false,
		},
		{
			s: State{
				year: 364,
			},
			proto: nil,
			err:   true,
		},
	}
	for _, tc := range cases {
		err := tc.s.Marshal(tc.proto)
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
		if got, want := tc.proto, tc.want; !proto.Equal(got, want) {
			t.Errorf("GameState: got %v, want %v", got, want)
		}
	}
}

func TestMechanic(t *testing.T) {
	settingsProto := &pb.GameSettings{
		YearSettings: &pb.YearSettings{
			InitYear: 43,
		},
	}
	set, err := NewSettings(settingsProto)
	if err != nil {
		t.Fatalf("NewSettings: unexpected error: %e", err)
	}
	s, err := set.InitState()
	if err != nil {
		t.Fatalf("NewSettings: unexpected error: %e", err)
	}
	if got, want := s.Year(), int32(43); got != want {
		t.Fatalf("year: got %d, want %d", got, want)
	}
	s.Incr()
	if got, want := s.Year(), int32(44); got != want {
		t.Fatalf("year: got %d, want %d", got, want)
	}
	stateProto := &pb.GameState{}
	err = s.Marshal(stateProto)
	if err != nil {
		t.Fatalf("Marshal: unexpected error: %e", err)
	}
	newState, err := NewState(stateProto, s.settings)
	if err != nil {
		t.Fatalf("NewState: unexpected error: %e", err)
	}
	if got, want := newState.Year(), int32(44); got != want {
		t.Fatalf("year: got %d, want %d", got, want)
	}
	newState.Incr()
	if got, want := newState.Year(), int32(45); got != want {
		t.Fatalf("year: got %d, want %d", got, want)
	}
}
