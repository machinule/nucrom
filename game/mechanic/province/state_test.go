package province

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewStateCase struct {
	proto      *pb.GameState
	prev       *State
	id         pb.ProvinceId
	influence  int32
	government pb.Government
	occupier   pb.ProvinceId
	leader     string
	err        bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{
						&pb.ProvinceState{
							Id:        pb.ProvinceId_ROMANIA,
							Influence: -1,
							Gov:       pb.Government_COMMUNISM,
							// No occupier
							Leader: "David Mihai",
						},
					},
				},
			},
			prev: &State{
				settings: &Settings{},
				Provinces: map[pb.ProvinceId]*ProvState{
					pb.ProvinceId_ROMANIA: &ProvState{
						id:         pb.ProvinceId_ROMANIA,
						influence:  -2,
						government: pb.Government_COMMUNISM,
						// No occupier
						leader: "David Mihai",
					},
				}},
			id:         pb.ProvinceId_ROMANIA,
			influence:  -2,
			government: pb.Government_COMMUNISM,
			// TODO: occupier: nil,
			leader: "David Mihai",
			err:    false,
		},
		{
			proto: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{
						&pb.ProvinceState{
							Id:        pb.ProvinceId_ROMANIA,
							Influence: -1,
							Gov:       pb.Government_COMMUNISM,
							// No occupier
							Leader: "David Mihai",
						}},
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
		if got, want := s.Get(tc.id).Id(), tc.id; got != want {
			t.Errorf("province id: got %d, want %d", got, want)
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
				Provinces: map[pb.ProvinceId]*ProvState{
					pb.ProvinceId_ROMANIA: &ProvState{
						id:         pb.ProvinceId_ROMANIA,
						influence:  -1,
						government: pb.Government_COMMUNISM,
						// No occupier
						leader: "David Mihai",
					},
				},
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{
						&pb.ProvinceState{
							Id:        pb.ProvinceId_ROMANIA,
							Influence: -1,
							Gov:       pb.Government_COMMUNISM,
							// No occupier
							Leader: "David Mihai",
						},
					},
				},
			},
			err: false,
		},
		{
			s: State{
				Provinces: map[pb.ProvinceId]*ProvState{
					pb.ProvinceId_ROMANIA: &ProvState{
						id:         pb.ProvinceId_ROMANIA,
						influence:  -1,
						government: pb.Government_COMMUNISM,
						// No occupier
						leader: "David Mihai",
					},
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
