package pseudorandom

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewStateCase struct {
	proto *pb.GameState
	prev  *State
	seed  int64
	err   bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				PseudorandomState: &pb.PseudorandomState{
					Seed: 42,
				},
			},
			prev: &State{
				settings: &Settings{},
			},
			seed: 42,
			err:  false,
		},
		{
			proto: &pb.GameState{
				PseudorandomState: &pb.PseudorandomState{
					Seed: 34,
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
		if got, want := s.seed, tc.seed; got != want {
			t.Errorf("seed: got %d, want %d", got, want)
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
				seed: 9001,
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				PseudorandomState: &pb.PseudorandomState{
					Seed: 9001,
				},
			},
			err: false,
		},
		{
			s: State{
				seed: 9001,
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
		PseudorandomSettings: &pb.PseudorandomSettings{
			InitSeed: 42,
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
	if got, want := s.Get(), int64(42); got != want {
		t.Fatalf("seed: got %d, want %d", got, want)
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
	if got, want := newState.Get(), int64(42); got != want {
		t.Fatalf("seed: got %d, want %d", got, want)
	}
	// "Coin" flip (50/50)
	var c_pdf [2]int32
	for i := 0; i < 1000; i++ {
		if s.Happens(500000) {
			c_pdf[1]++
		} else {
			c_pdf[0]++
		}
	}
	if got, want := c_pdf[0], int32(484); got != want {
		t.Fatalf("coin flip false: got %d, want %d", got, want)
	}
	if got, want := c_pdf[1], int32(516); got != want {
		t.Fatalf("coin flip true: got %d, want %d", got, want)
	}
	// Weighted chances
	weights := []int32{4, 1, 2, 3} // Sum = 10
	var w_pdf [4]int32
	for i := 0; i < 1000; i++ {
		w_pdf[s.Roll(weights)]++
	}
	if got, want := w_pdf[0], int32(415); got != want {
		t.Fatalf("weighted 0: got %d, want %d", got, want)
	}
	if got, want := w_pdf[1], int32(95); got != want {
		t.Fatalf("weighted 1: got %d, want %d", got, want)
	}
	if got, want := w_pdf[2], int32(197); got != want {
		t.Fatalf("weighted 2: got %d, want %d", got, want)
	}
	if got, want := w_pdf[3], int32(293); got != want {
		t.Fatalf("weighted 3: got %d, want %d", got, want)
	}
}
