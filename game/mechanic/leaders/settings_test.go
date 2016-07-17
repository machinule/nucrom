package leaders

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

type LeaderCase struct {
	name       string
	birth_year int32
	start_year int32
	end_year   int32
	province   pb.ProvinceId
	elected    bool
}

type NewSettingsCase struct {
	proto   *pb.GameSettings
	leaders []LeaderCase
	err     bool
}

func TestNewSettings(t *testing.T) {
	cases := []NewSettingsCase{
		/*{
			proto: &pb.GameSettings{},
			leaders:  []LeaderCase{},
			err:   false, //TODO: Currently only validates single leaders
		},*/
		{
			proto: &pb.GameSettings{
				LeadersSettings: &pb.LeadersSettings{
					ProvinceLeaderSettings: []*pb.ProvinceLeaderSettings{
						&pb.ProvinceLeaderSettings{
							Id: pb.ProvinceId_VIETNAM,
							LeaderSettings: []*pb.LeaderSettings{
								&pb.LeaderSettings{
									Name: "William Hess",
								},
							},
						},
					},
				},
			},
			err: true,
		},
		{
			proto: &pb.GameSettings{
				LeadersSettings: &pb.LeadersSettings{
					ProvinceLeaderSettings: []*pb.ProvinceLeaderSettings{
						&pb.ProvinceLeaderSettings{
							Id: pb.ProvinceId_VIETNAM,
							LeaderSettings: []*pb.LeaderSettings{
								&pb.LeaderSettings{
									Name:      "William Hess",
									BirthYear: 1930,
									StartYear: 1960,
									EndYear:   1980,
									// Elected left blank
								},
								&pb.LeaderSettings{
									Name:      "Ho Chi Minh",
									BirthYear: 1890,
									StartYear: 1948,
									EndYear:   1970,
									Elected:   true, // Just go with it
								},
							},
						},
						&pb.ProvinceLeaderSettings{
							Id: pb.ProvinceId_MALAYSIA,
							LeaderSettings: []*pb.LeaderSettings{
								&pb.LeaderSettings{
									Name:      "Paul Atriedes",
									BirthYear: 1920,
									StartYear: 1940,
									EndYear:   1990,
									Elected:   false,
								},
							},
						},
					},
				},
			},
			// TODO: Pointers?
			leaders: []LeaderCase{
				LeaderCase{
					name:       "William Hess",
					birth_year: 1930,
					start_year: 1960,
					end_year:   1980,
					province:   pb.ProvinceId_VIETNAM,
					//elected: left blank,
				},
				LeaderCase{
					name:       "Ho Chi Minh",
					birth_year: 1890,
					start_year: 1948,
					end_year:   1970,
					province:   pb.ProvinceId_VIETNAM,
					elected:    true, // Just go with it
				},
				LeaderCase{
					name:       "Paul Atriedes",
					birth_year: 1920,
					start_year: 1940,
					end_year:   1990,
					province:   pb.ProvinceId_MALAYSIA,
					elected:    false,
				},
			},
			err: false,
		},
	}
	for index, tc := range cases {
		s, err := NewSettings(tc.proto)
		if got, want := err != nil, tc.err; got != want {
			msg := map[bool]string{
				true:  "error",
				false: "no error",
			}
			t.Errorf("case %d - err: got %s, want %s", index, msg[got], msg[want])
			continue
		}
		if tc.err {
			continue
		}
		l2 := tc.leaders[index]
		l1 := s.leaders[l2.province][l2.name]
		// TODO
		/*if got, want := l1.name, l2.name; got != want {
			t.Errorf("case %d - name: got %d, want %d", index, got, want)
		}*/
		if got, want := l1.birth_year, l2.birth_year; got != want {
			t.Errorf("case %d - birth year: got %d, want %d", index, got, want)
		}
		if got, want := l1.start_year, l2.start_year; got != want {
			t.Errorf("case %d - start year: got %d, want %d", index, got, want)
		}
		if got, want := l1.end_year, l2.end_year; got != want {
			t.Errorf("case %d - end year: got %d, want %d", index, got, want)
		}
		if got, want := l1.elected, l2.elected; got != want {
			t.Errorf("case %d - elected: got %d, want %d", index, got, want)
		}
	}
}

/*

type InitLeaderCase struct {
    name string
    l_type pb.LeaderType
    elected bool
    province pb.ProvinceId
}

type InitStateCase struct {
	settings *Settings
	leaders []InitLeaderCase
    err      bool
}

func TestInitState(t *testing.T) {
	cases := []InitStateCase{
		{
			settings: &Settings{
				init_leaders: []*InitLeaderSettings {
                    &InitLeaderSettings {
                        name: "Enver Hoxha",
                        l_type: pb.LeaderType_ROGUE,
                        dissident false,
                        province pb.ProvinceId_ALBANIA,
                        elected false,
                    },
                    &InitLeaderSettings {
                        name: "Fidel Castro",
                        l_type: pb.LeaderType_NORMAL,
                        dissident true,
                        province pb.ProvinceId_CUBA,
                        elected false,
                    },
                    &InitLeaderSettings {
                        name: "Juan Peron",
                        l_type: pb.LeaderType_NORMAL,
                        dissident true,
                        province pb.ProvinceId_CUBA,
                        elected false,
                    },
                    &InitLeaderSettings {
                        name: "David Ben-Gurion",
                        l_type: pb.LeaderType_STRONG,
                        dissident false,
                        province pb.ProvinceId_ISRAEL,
                        // Skipping elected
                    }
                }
			},
            leaders: []InitLeaderCase{
                InitLeaderCase{
                    name: "Enver Hoxha",
                    l_type: pb.LeaderType_ROGUE,
                    dissident: false,
                    province pb.ProvinceId_ALBANIA,
                    elected false,
                },
                InitLeaderCase{
                    name: "Fidel Castro",
                    l_type: pb.LeaderType_DISSIDENT,
                },
                LeaderCase{
                    name: "Ho Chi Minh",
                    birth_year: 1890,
                    start_year: 1948,
                    end_year:1970,
                    province: pb.ProvinceId_VIETNAM,
                    elected: true, // Just go with it
                },
                LeaderCase{
                    name: "Paul Atriedes",
                    birth_year: 1920,
                    start_year: 1940,
                    end_year: 1990,
                    province: pb.ProvinceId_MALAYSIA,
                    elected: false,
                },
            },

			err:  false,
		},
	}
	for _, tc := range cases {
		s, err := tc.settings.InitState()
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
			t.Errorf("heat: got %d, want %d", got, want)
		}
		if got, want := s.settings, tc.settings; got != want {
			t.Errorf("settings: got %d, want %d", got, want)
		}
	}
}

*/
