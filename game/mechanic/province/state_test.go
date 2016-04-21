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
							Id:        pb.ProvinceId_ROMANIA.Enum(),
							Influence: proto.Int32(-1),
							Gov:       pb.Government_COMMUNISM.Enum(),
							// No occupier
							Leader: proto.String("David Mihai"),
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
							Id:        pb.ProvinceId_ROMANIA.Enum(),
							Influence: proto.Int32(-1),
							Gov:       pb.Government_COMMUNISM.Enum(),
							// No occupier
							Leader: proto.String("David Mihai"),
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

/* TODO:
func TestMarshal(t *testing.T) {
	cases := []MarshalCase{
		{
			s: State{
				heat: 56,
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				HeatState: &pb.HeatState{
					Heat: proto.Int32(56),
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
			Init:  proto.Int32(70),
			Min:   proto.Int32(0),
			Mxm:   proto.Int32(100),
			Decay: proto.Int32(5),
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
*/
