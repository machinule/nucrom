package leaders

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

/*
type NewStateLeaderCase struct {
    name string
    l_type pb.LeaderType
    elected bool
    birth_year int32
    province pb.ProvinceId
}

type NewStateCase struct {
	proto *pb.GameState
	prev  *State
    leaders []NewStateLeaderCase
	err   bool
}

func TestNewState(t *testing.T) {
	cases := []NewStateCase{
		{
			proto: &pb.GameState{
				LeadersState: &pb.LeadersState{
                    ProvinceLeaderStates: []*pb.ProvinceLeaderState {
                        &pb.ProvinceLeaderState {
                            Id: pb.ProvinceId_CUBA,
                            Dissident: &pb.LeaderState {
                                Name: "Fidel Castro",
                                Type: pb.LeaderType_DISSIDENT,
                                Elected: false,
                                BirthYear: 1926,
                            },
                        },
                        &pb.ProvinceLeaderState {
                            Id:pb.ProvinceId_ALBANIA,
                            Active: &pb.LeaderState {
                                Name: "Enver Hoxha",
                                Type: pb.LeaderType_ROGUE,
                                Elected: false,
                                BirthYear: 1908,
                            },
                        },
                        &pb.ProvinceLeaderState {
                            Id: pb.ProvinceId_ARGENTINA,
                            Active: &pb.LeaderState {
                                Name: "Juan Peron",
                                Type: pb.LeaderType_NORMAL,
                                Elected: true,
                                BirthYear: 1895,
                            },
                        },
                        &pb.ProvinceLeaderState {
                            Id: pb.ProvinceId_ISRAEL,
                            Active: &pb.LeaderState {
                                Name: "David Ben-Gurion",
                                Type: pb.LeaderType_STRONG,
                                // Skipping elected,
                                BirthYear: 1886,
                            },
                        },
                    },
                },
			},
			prev: &State{
				settings: &Settings{},
			},
			leaders: []NewStateLeaderCase {
                NewStateLeaderCase {
                    name: "Fidel Castro",
                    l_type: pb.LeaderType_DISSIDENT,
                    elected: false,
                    birth_year: 1926,
                    province: pb.ProvinceId_CUBA,
                },
                NewStateLeaderCase {
                    name: "Enver Hoxha",
                    l_type: pb.LeaderType_ROGUE,
                    elected: false,
                    birth_year: 1908,
                    province: pb.ProvinceId_ALBANIA,
                },
                NewStateLeaderCase {
                    name: "Juan Peron",
                    l_type: pb.LeaderType_NORMAL,
                    elected: true,
                    birth_year: 1895,
                    province: pb.ProvinceId_ARGENTINA,
                },
                NewStateLeaderCase {
                    name: "David Ben-Gurion",
                    l_type: pb.LeaderType_STRONG,
                    birth_year: 1886,
                    province: pb.ProvinceId_ISRAEL,
                },
            },
			err:  false,
		},
		{
			proto: &pb.GameState{
				LeadersState: &pb.LeadersState{
                    ProvinceLeaderStates: []*pb.ProvinceLeaderState {
                        &pb.ProvinceLeaderState {
                            Id: pb.ProvinceId_CUBA,
                            Dissident: &pb.LeaderState {
                                Name: "Fidel Castro",
                                Type: pb.LeaderType_DISSIDENT,
                                Elected: false,
                                BirthYear: 1926,
                            },
                        },
                    },
			    },
            },
            leaders: nil,
			prev: nil,
			err:  true,
		},
    }
	for index, tc := range cases {
		s, err := NewState(tc.proto, tc.prev)
		if got, want := err != nil, tc.err; got != want {
			msg := map[bool]string{
				true:  "error",
				false: "no error",
			}
			t.Errorf("err: got %s, want %s", msg[got], msg[want])
			continue
		}
        l2 := tc.leaders[index]
        if tc.err {
			continue
		}
		if l2.l_type == pb.LeaderType_DISSIDENT {
            if got, want := s.dissidents[tc.leaders[index].province], l2.name; got != want {
			    t.Errorf("leader: got %d, want %d", got, want)
            }
        } else {
            if got, want := s.active[tc.leaders[index].province], l2.name; got != want {
			    t.Errorf("leader: got %d, want %d", got, want)
            }
        }
        l1 := s.leaders[l2.name]
		if got, want := l1.l_type, l2.l_type; got != want {
			t.Errorf("type: got %d, want %d", got, want)
		}
		if got, want := l1.elected, l2.elected; got != want {
			t.Errorf("elected: got %d, want %d", got, want)
		}
		if got, want := l1.birth_year, l2.birth_year; got != want {
			t.Errorf("birth year: got %d, want %d", got, want)
		}
	}
}
*/
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
				leaders: map[string]*LeaderState{
					"Fidel Castro": &LeaderState{
						l_type:     pb.LeaderType_DISSIDENT,
						elected:    false,
						birth_year: 1926,
					},
					"Enver Hoxha": &LeaderState{
						l_type:     pb.LeaderType_ROGUE,
						elected:    false,
						birth_year: 1908,
					},
					"Juan Peron": &LeaderState{
						l_type:     pb.LeaderType_NORMAL,
						elected:    true,
						birth_year: 1895,
					},
					"David Ben-Gurion": &LeaderState{
						l_type: pb.LeaderType_STRONG,
						// Skipping elected,
						birth_year: 1886,
					},
				},
				dissidents: map[pb.ProvinceId]string{
					pb.ProvinceId_CUBA: "Fidel Castro",
				},
				active: map[pb.ProvinceId]string{
					pb.ProvinceId_ALBANIA:   "Enver Hoxha",
					pb.ProvinceId_ARGENTINA: "Juan Peron",
					pb.ProvinceId_ISRAEL:    "David Ben-Gurion",
				},
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				LeadersState: &pb.LeadersState{
					ProvinceLeaderStates: []*pb.ProvinceLeaderState{
						&pb.ProvinceLeaderState{
							Id: pb.ProvinceId_ALBANIA,
							Active: &pb.LeaderState{
								Name:      "Enver Hoxha",
								Type:      pb.LeaderType_ROGUE,
								Elected:   false,
								BirthYear: 1908,
							},
						},
						&pb.ProvinceLeaderState{
							Id: pb.ProvinceId_ARGENTINA,
							Active: &pb.LeaderState{
								Name:      "Juan Peron",
								Type:      pb.LeaderType_NORMAL,
								Elected:   true,
								BirthYear: 1895,
							},
						},
						&pb.ProvinceLeaderState{
							Id: pb.ProvinceId_ISRAEL,
							Active: &pb.LeaderState{
								Name: "David Ben-Gurion",
								Type: pb.LeaderType_STRONG,
								// Skipping elected,
								BirthYear: 1886,
							},
						},
						&pb.ProvinceLeaderState{
							Id: pb.ProvinceId_CUBA,
							Dissident: &pb.LeaderState{
								Name:      "Fidel Castro",
								Type:      pb.LeaderType_DISSIDENT,
								Elected:   false,
								BirthYear: 1926,
							},
						},
					},
				},
			},
			err: false,
		},
		{
			s: State{
				leaders: map[string]*LeaderState{
					"Fidel Castro": &LeaderState{
						l_type:     pb.LeaderType_DISSIDENT,
						elected:    false,
						birth_year: 1926,
					},
					"Enver Hoxha": &LeaderState{
						l_type:     pb.LeaderType_ROGUE,
						elected:    false,
						birth_year: 1908,
					},
					"Juan Peron": &LeaderState{
						l_type:     pb.LeaderType_NORMAL,
						elected:    true,
						birth_year: 1895,
					},
					"David Ben-Gurion": &LeaderState{
						l_type: pb.LeaderType_STRONG,
						// Skipping elected,
						birth_year: 1886,
					},
				},
				dissidents: map[pb.ProvinceId]string{
					pb.ProvinceId_CUBA: "Fidel Castro",
				},
				active: map[pb.ProvinceId]string{
					pb.ProvinceId_ALBANIA:   "Enver Hoxha",
					pb.ProvinceId_ARGENTINA: "Juan Peron",
					pb.ProvinceId_ISRAEL:    "David Ben-Gurion",
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
