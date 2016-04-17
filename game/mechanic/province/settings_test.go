package province

import (
	//"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type NewSettingsCase struct {
	proto *pb.GameSettings
	id    pb.ProvinceId
	err   bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		{
			proto: &pb.GameSettings{
				ProvincesSettings: &pb.ProvincesSettings{
					ProvinceSettings: []*pb.ProvinceSettings{
						&pb.ProvinceSettings{
							Id: pb.ProvinceId_ROMANIA.Enum(),
						},
					},
				},
			},
			err: false,
			id:  pb.ProvinceId_ROMANIA,
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
		if got, want := s.Get(tc.id).Id(), tc.id; got != want {
			t.Errorf("id: got %d, want %d", got, want)
		}
	}
}

type InitStateCase struct {
	s   *Settings
	id  pb.ProvinceId
	err bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			s: &Settings{
				Provinces: map[pb.ProvinceId]*ProvSettings{
					pb.ProvinceId_ROMANIA: &ProvSettings{
						id: pb.ProvinceId_ROMANIA,
					},
				},
			},
			id:  pb.ProvinceId_ROMANIA,
			err: false,
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
		if got, want := s.Get(tc.id).Id(), tc.id; got != want {
			t.Errorf("id: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.s; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}
