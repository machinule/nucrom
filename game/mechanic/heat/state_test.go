package heat

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewStateCase struct {
	proto *pb.GameState
	prev  *State
	heat  int32
	err   bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				HeatState: &pb.HeatState{
					Heat: 45,
				},
			},
			prev: &State{
				settings: &Settings{},
			},
			heat: 45,
			err:  false,
		},
		{
			proto: &pb.GameState{
				HeatState: &pb.HeatState{
					Heat: 45,
				},
			},
			prev: nil,
			err:  true,
		},
	}
	for _, tc := range cases {
		s, err := NewState(tc.proto, tc.prev)
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
			t.Errorf("year: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.prev.settings; got != want {
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
				heat: 56,
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				HeatState: &pb.HeatState{
					Heat: 56,
				},
			},
			err: false,
		},
		{
			s: State{
				heat: 56,
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
		HeatSettings: &pb.HeatSettings{
			Init:  70,
			Min:   0,
			Mxm:   100,
			Decay: 5,
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
	if got, want := s.Heat(), int32(70); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	s.Chng(3)
	if got, want := s.Heat(), int32(73); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	s.Chng(-8)
	if got, want := s.Heat(), int32(65); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	s.Decay()
	if got, want := s.Heat(), int32(60); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	s.Chng(-100)
	if got, want := s.Heat(), int32(0); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	s.Chng(200)
	if got, want := s.Heat(), int32(100); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	stateProto := &pb.GameState{}
	err = s.Marshal(stateProto)
	if err != nil {
		t.Fatalf("Marshal: unexpected error: %e", err)
	}
	newState, err := NewState(stateProto, s)
	if err != nil {
		t.Fatalf("NewState: unexpected error: %e", err)
	}
	if got, want := newState.Heat(), int32(100); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
	newState.Decay()
	if got, want := newState.Heat(), int32(95); got != want {
		t.Fatalf("heat: got %d, want %d", got, want)
	}
}
