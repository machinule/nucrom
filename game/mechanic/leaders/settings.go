package leaders

import (
	//"errors"
	"fmt"
	pb "github.com/machinule/nucrom/proto/gen"
)

type Settings struct {
	leaders      map[pb.ProvinceId]map[string]*LeaderSettings // Map of available leader settings
	init_leaders []*InitLeaderSettings                        // Leaders active at start
}

type InitLeaderSettings struct {
	name      string
	l_type    pb.LeaderType
	dissident bool
	province  pb.ProvinceId
	elected   bool
}

type LeaderSettings struct {
	// Name of the leader functions as UUID
	birth_year int32 // Year of birth
	start_year int32 // First year leader is available
	end_year   int32 // Year leader stops becoming available
	elected    bool  // Whether or not leader was elected or not
}

func validate(settingsProto *pb.GameSettings) error {
	return nil
}

func NewSettings(settingsProto *pb.GameSettings) (*Settings, error) {
	if err := validate(settingsProto); err != nil {
		return nil, fmt.Errorf("validating settings proto: %e", err)
	}
	leaders := make(map[pb.ProvinceId]map[string]*LeaderSettings)
	var init_leaders []*InitLeaderSettings
	for _, ps := range settingsProto.GetLeadersSettings().GetProvinceLeaderSettings() {
		p := ps.Id
		ls := ps.GetLeaderSettings()
		ils := ps.GetInitLeaderSettings()
		for _, l := range ls {
			// TODO: Move to validate
			if l.BirthYear == 0 || l.StartYear == 0 || l.EndYear == 0 {
				return nil, fmt.Errorf("%d missing leader settings", l.Name)
			}
			leaders[p] = make(map[string]*LeaderSettings)
			leaders[p][l.Name] = &LeaderSettings{
				birth_year: l.BirthYear,
				start_year: l.StartYear,
				end_year:   l.EndYear,
				elected:    l.Elected,
			}
		}
		for _, l := range ils {
			// TODO: Move to validate
			if l.Type == pb.LeaderType_NO_LEADER {
				return nil, fmt.Errorf("Failed to create initial leader %s - missing parameters", l.Name)
			}
			init_leaders = append(init_leaders, &InitLeaderSettings{
				name:      l.Name,
				l_type:    l.Type,
				dissident: l.Dissident,
				province:  p,
				elected:   l.Elected,
			})
		}
	}

	return &Settings{
		leaders:      leaders,
		init_leaders: init_leaders,
	}, nil
}

func (s *Settings) InitState() (*State, error) {
	ns := &State{
		settings:    s,
		leaders:     make(map[string]*LeaderState),
		active:      make(map[pb.ProvinceId]string),
		deactivated: make(map[pb.ProvinceId]string),
		dissidents:  make(map[pb.ProvinceId]string),
	}
	for _, ils := range s.init_leaders {
		ns.leaders[ils.name] = &LeaderState{
			l_type:     ils.l_type,
			elected:    ils.elected,
			birth_year: s.leaders[ils.province][ils.name].birth_year,
		}
		if ils.dissident {
			ns.dissidents[ils.province] = ils.name
		} else {
			ns.active[ils.province] = ils.name
		}
	}
	return ns, nil
}
