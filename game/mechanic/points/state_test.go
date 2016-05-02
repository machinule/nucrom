package points

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewStateCase struct {
	proto    *pb.GameState
	prev     *State
	usa_pol  int32
	usa_mil  int32
	usa_cov  int32
	ussr_pol int32
	ussr_mil int32
	ussr_cov int32
	err      bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				PointsState: &pb.PointsState{
					UsaState: &pb.PointState{
						Political: 4,
						Military:  5,
						Covert:    3,
					},
					UssrState: &pb.PointState{
						Political: 5,
						Military:  6,
						Covert:    1,
					},
				},
			},
			prev: &State{
				settings: &Settings{},
			},
			usa_pol:  4,
			usa_mil:  5,
			usa_cov:  3,
			ussr_pol: 5,
			ussr_mil: 6,
			ussr_cov: 1,
			err:      false,
		},
		{
			proto: &pb.GameState{
				PointsState: &pb.PointsState{
					UsaState: &pb.PointState{
						Political: 4,
						Military:  5,
						Covert:    3,
					},
					UssrState: &pb.PointState{
						Political: 5,
						Military:  6,
						Covert:    1,
					},
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
				usa: PointState{
					pol: 4,
					mil: 5,
					cov: 3,
				},
				ussr: PointState{
					pol: 6,
					mil: 4,
					cov: 2,
				},
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				PointsState: &pb.PointsState{
					UsaState: &pb.PointState{
						Political: 4,
						Military:  5,
						Covert:    3,
					},
					UssrState: &pb.PointState{
						Political: 6,
						Military:  4,
						Covert:    2,
					},
				},
			},
			err: false,
		},
		{
			s: State{
				usa: PointState{
					pol: 4,
					mil: 5,
					cov: 3,
				},
				ussr: PointState{
					pol: 6,
					mil: 4,
					cov: 2,
				},
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
