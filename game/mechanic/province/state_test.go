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

type NewConflictCase struct {
	proto       *pb.GameState
	prev        *State
	id          pb.ProvinceId
	name        string
	att         pb.ProvinceId
	def         pb.ProvinceId
	goal        int32
	length      int32
	base_chance int32
	locations   pb.ProvinceId
	att_sup     pb.Player
	def_sup     pb.Player
	att_prog    int32
	def_prog    int32
	err         bool
}

func TestNewConflictState(t *testing.T) {
	cases := []NewConflictCase{
		{
			proto: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{},
					Conflicts: &pb.ConflictsState{
						Active: []*pb.Conflict{
							&pb.Conflict{
								Name: "American Revolution",
								Type: pb.ConflictType_COLONIAL_WAR,
								Goal: 5,
								Attackers: &pb.Faction{
									Ids: []pb.ProvinceId{ // TODO: Dissidents
										pb.ProvinceId_P_USA,
									},
									Supporter: pb.Player_USSR, // Basically the French
									Progress:  3,
								},
								Defenders: &pb.Faction{
									Ids: []pb.ProvinceId{
										pb.ProvinceId_GREAT_BRITAIN,
									},
									Supporter: pb.Player_NEITHER,
									Progress:  2,
								},
								Length:     6,
								BaseChance: 150000,
                                Locations: []pb.ProvinceId{
                                    pb.ProvinceId_P_USA,
                                },
							},
						},
						Dormant:  []*pb.Conflict{},
						Possible: []*pb.Conflict{},
					},
				},
			},
			prev: &State{
				settings: &Settings{},
				Conflicts: map[pb.ProvinceId]*Conflict{
					pb.ProvinceId_P_USA: &Conflict{
						name:          "American Revolution",
						conflict_type: pb.ConflictType_COLONIAL_WAR,
						goal:          6,
						attackers: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_P_USA,
							},
							supporter: pb.Player_USSR,
							progress:  4,
						},
						defenders: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_GREAT_BRITAIN,
							},
							supporter: pb.Player_NEITHER,
							progress:  3,
						},
						length:      7,
						base_chance: 150000,
                        locations: []pb.ProvinceId{
                            pb.ProvinceId_P_USA,
                        },
					},
				},
			},
			id:          pb.ProvinceId_P_USA,
			name:        "American Revolution",
			att:         pb.ProvinceId_P_USA,
			def:         pb.ProvinceId_GREAT_BRITAIN,
			goal:        6,
			length:      7,
			base_chance: 150000,
			locations:   pb.ProvinceId_P_USA,
			att_sup:     pb.Player_USSR,
			def_sup:     pb.Player_NEITHER,
			att_prog:    4,
			def_prog:    3,
			err:         false,
		},
		{
			proto: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{},
					Conflicts:      &pb.ConflictsState{},
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
		c := s.Conflict(tc.id)
		if got, want := c.Name(), tc.name; got != want {
			t.Errorf("name: got %d, want %d", got, want)
		}
		if got, want := c.Goal(), tc.goal; got != want {
			t.Errorf("goal: got %d, want %d", got, want)
		}
		if got, want := c.Defenders()[0], tc.def; got != want {
			t.Errorf("defenders: got %d, want %d", got, want)
		}
		if got, want := c.Attackers()[0], tc.att; got != want {
			t.Errorf("attackers: got %d, want %d", got, want)
		}
		if got, want := c.Attackers()[0], tc.att; got != want {
			t.Errorf("attackers: got %d, want %d", got, want)
		}
		if got, want := c.Length(), tc.length; got != want {
			t.Errorf("length: got %d, want %d", got, want)
		}
		if got, want := c.BaseChance(), tc.base_chance; got != want {
			t.Errorf("base_chance: got %d, want %d", got, want)
		}
		if got, want := c.Att_Supporter(), tc.att_sup; got != want {
			t.Errorf("att sup: got %d, want %d", got, want)
		}
		if got, want := c.Att_Progress(), tc.att_prog; got != want {
			t.Errorf("att prog: got %d, want %d", got, want)
		}
		if got, want := c.Def_Supporter(), tc.def_sup; got != want {
			t.Errorf("def sup: got %d, want %d", got, want)
		}
		if got, want := c.Def_Progress(), tc.def_prog; got != want {
			t.Errorf("def prog: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.prev.settings; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
        if got, want := s.IsAtWar(pb.ProvinceId_P_USA), true; got != want {
            t.Errorf("IsAtWar #1: got %d, want %d", got, want)
        }
        if got, want := s.IsAtWar(pb.ProvinceId_GREAT_BRITAIN), true; got != want {
            t.Errorf("IsAtWar #2: got %d, want %d", got, want)
        }
        if got, want := s.IsAtWar(pb.ProvinceId_P_USSR), false; got != want {
            t.Errorf("IsAtWar #3: got %d, want %d", got, want)
        }
        if got, want := s.IsSiteOfConflict(pb.ProvinceId_P_USA), true; got != want {
            t.Errorf("IsSiteOfConflict #1: got %d, want %d", got, want)
        }
        if got, want := s.IsSiteOfConflict(pb.ProvinceId_GREAT_BRITAIN), false; got != want {
            t.Errorf("IsSiteOfConflict #2: got %d, want %d", got, want)
        }

	}
}

type MarshalCase struct {
	s     State
	proto *pb.GameState
	want  *pb.GameState
	err   bool
}

func TestConflictMarshal(t *testing.T) {
	cases := []MarshalCase{
		{
			s: State{
				Conflicts: map[pb.ProvinceId]*Conflict{
					pb.ProvinceId_P_USA: &Conflict{
						name:          "American Revolution",
						conflict_type: pb.ConflictType_COLONIAL_WAR,
						goal:          6,
						attackers: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_P_USA,
							},
							supporter: pb.Player_USSR,
							progress:  4,
						},
						defenders: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_GREAT_BRITAIN,
							},
							supporter: pb.Player_NEITHER,
							progress:  3,
						},
						length:      7,
						base_chance: 150000,
                        locations: []pb.ProvinceId{
                            pb.ProvinceId_P_USA,
                        },
					},
				},
			},
			proto: &pb.GameState{},
			want: &pb.GameState{
				ProvincesState: &pb.ProvincesState{
					ProvinceStates: []*pb.ProvinceState{},
					Conflicts: &pb.ConflictsState{
						Active: []*pb.Conflict{
							&pb.Conflict{
								Name: "American Revolution",
								Type: pb.ConflictType_COLONIAL_WAR,
								Goal: 6,
								Attackers: &pb.Faction{
									Ids: []pb.ProvinceId{ // TODO: Dissidents
										pb.ProvinceId_P_USA,
									},
									Supporter: pb.Player_USSR, // Basically the French
									Progress:  4,
								},
								Defenders: &pb.Faction{
									Ids: []pb.ProvinceId{
										pb.ProvinceId_GREAT_BRITAIN,
									},
									Supporter: pb.Player_NEITHER,
									Progress:  3,
								},
								Length:     7,
								BaseChance: 150000,
                                Locations: []pb.ProvinceId{
                                    pb.ProvinceId_P_USA,
                                },
							},
						},
						Dormant:  []*pb.Conflict{},
						Possible: []*pb.Conflict{},
					},
				},
			},
			err: false,
		},
		{
			s: State{
				Conflicts: map[pb.ProvinceId]*Conflict{
					pb.ProvinceId_P_USA: &Conflict{
						name:          "American Revolution",
						conflict_type: pb.ConflictType_COLONIAL_WAR,
						goal:          6,
						attackers: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_P_USA,
							},
							supporter: pb.Player_USSR,
							progress:  4,
						},
						defenders: Faction{
							members: []pb.ProvinceId{
								pb.ProvinceId_GREAT_BRITAIN,
							},
							supporter: pb.Player_NEITHER,
							progress:  3,
						},
						length:      7,
						base_chance: 150000,
                        locations: []pb.ProvinceId{
                            pb.ProvinceId_P_USA,
                        },
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
					Conflicts: &pb.ConflictsState{},
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
