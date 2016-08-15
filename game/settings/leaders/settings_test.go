package leaders

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

var (
	testSettings = &pb.GameSettings{
		YearSettings: &pb.YearSettings{
			InitYear: 1950,
		},
		LeadersSettings: &pb.LeadersSettings{
			LeaderSettings: []*pb.LeaderSettings{
				{
					Name:          "David Mihai",
					BirthYear:     1901,
					BirthProvince: pb.ProvinceId_ROMANIA,
					Positions: []*pb.LeaderPosition{
						{
							Title:     pb.LeaderTitle_PRESIDENT,
							StartYear: 1912,
							EndYear:   1945,
							Province:  pb.ProvinceId_ROMANIA,
							Elected:   true,
							Dissident: false,
							Type:      pb.LeaderType_ROGUE,
						},
						{
							Title:     pb.LeaderTitle_NO_TITLE,
							StartYear: 1946,
							EndYear:   1949,
							Province:  pb.ProvinceId_ROMANIA,
							Elected:   true,
							Dissident: false,
							Type:      pb.LeaderType_NO_LEADER,
						},
						{
							Title:     pb.LeaderTitle_PRESIDENT,
							StartYear: 1950,
							EndYear:   0,
							Province:  pb.ProvinceId_INDIA,
							Elected:   false,
							Dissident: true,
							Type:      pb.LeaderType_DISSIDENT,
						},
					},
				},
				{
					Name:          "William Hess",
					BirthYear:     1932,
					BirthProvince: pb.ProvinceId_MEXICO,
					Positions:     []*pb.LeaderPosition{},
				},
				{
					Name:          "Chris McKee",
					BirthYear:     1912,
					BirthProvince: pb.ProvinceId_SOUTH_AFRICA,
					Positions: []*pb.LeaderPosition{
						{
							Title:     pb.LeaderTitle_PRESIDENT,
							StartYear: 1940,
							EndYear:   1940,
							Province:  pb.ProvinceId_NIGERIA,
							Elected:   false,
							Dissident: false,
							Type:      pb.LeaderType_STRONG,
						},
					},
				},
			},
		},
	}
)

func TestUnmarshal(t *testing.T) {
	s := Settings{}
	if err := s.Unmarshal(testSettings); err != nil {
		t.Errorf("Unmarshal(): got %v, want nil", err)
	}
}

func TestValidateDefaultIsValid(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if err := s.Validate(); err != nil {
		t.Errorf("Validate(): got %v, want nil", err)
	}
}

func TestInitialize(t *testing.T) {
	s := Settings{}
	s.Unmarshal(testSettings)
	state := &pb.GameState{}
	s.Initialize(state)

	got := state.LeadersState
	want := &pb.LeadersState{
		LeaderStates: []*pb.LeaderState{
			{
				Name: "David Mihai",
				CurrentPosition: &pb.LeaderPosition{
					Title:     pb.LeaderTitle_PRESIDENT,
					StartYear: 1950,
					EndYear:   0,
					Province:  pb.ProvinceId_INDIA,
					Elected:   false,
					Dissident: true,
					Type:      pb.LeaderType_DISSIDENT,
				},
			},
			{
				Name: "William Hess",
				CurrentPosition: &pb.LeaderPosition{
					Title:     pb.LeaderTitle_NO_TITLE,
					StartYear: 1950,
					EndYear:   0,
					Province:  pb.ProvinceId_MEXICO,
					Elected:   false,
					Dissident: false,
					Type:      pb.LeaderType_NO_LEADER,
				},
			}, {
				Name: "Chris McKee",
				CurrentPosition: &pb.LeaderPosition{
					Title:     pb.LeaderTitle_NO_TITLE,
					StartYear: 1940,
					EndYear:   0,
					Province:  pb.ProvinceId_NIGERIA,
					Elected:   false,
					Dissident: false,
					Type:      pb.LeaderType_NO_LEADER,
				},
			},
		},
	}
	if !proto.Equal(got, want) {
		t.Errorf("Leader states: got %v, want %v", got, want)
	}
}
