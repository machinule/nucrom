package province

import (
	//"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestMechanic(t *testing.T) {
    // 1978-1981
    testProto := &pb.GameSettings{
        ProvincesSettings: &pb.ProvincesSettings{
        ProvinceSettings: []*pb.ProvinceSettings{
            &pb.ProvinceSettings{
                    Id:             pb.ProvinceId_IRAN,
		            Label:          "Iran",
		            Adjacency:      []pb.ProvinceId{pb.ProvinceId_P_USSR, pb.ProvinceId_IRAQ, pb.ProvinceId_AFGHANISTAN, pb.ProvinceId_PAKISTAN},
		            StabilityBase:  2,
		            Region:         pb.Region_MIDDLE_EAST,
		            Coastal:        true,
		            InitInfluence:  3,
		            InitGovernment: pb.Government_AUTOCRACY,
		            InitLeader:     "Reza Pahlavi",
		        },
		        &pb.ProvinceSettings{ //TODO: Test conflict
		            Id:             pb.ProvinceId_AFGHANISTAN,
		            Label:          "Afghanistan",
		            Adjacency:      []pb.ProvinceId{pb.ProvinceId_P_USSR, pb.ProvinceId_IRAN, pb.ProvinceId_PAKISTAN},
		            StabilityBase:  1,
		            Region:         pb.Region_SOUTH_ASIA,
		            Coastal:        false,
		            InitInfluence:  -1,
		            InitGovernment: pb.Government_WEAK,
		        },
		        &pb.ProvinceSettings{
		            Id:             pb.ProvinceId_PAKISTAN,
		            Label:          "Pakistan",
		            Adjacency:      []pb.ProvinceId{pb.ProvinceId_IRAN, pb.ProvinceId_INDIA, pb.ProvinceId_AFGHANISTAN},
		            StabilityBase:  2,
		            Region:         pb.Region_SOUTH_ASIA,
		            Coastal:        true,
		            InitInfluence:  3,
		            InitGovernment: pb.Government_WEAK,
		        },
		        &pb.ProvinceSettings{
		            Id:             pb.ProvinceId_INDIA,
		            Label:          "India",
		            Adjacency:      []pb.ProvinceId{pb.ProvinceId_BANGLADESH, pb.ProvinceId_BURMA, pb.ProvinceId_PAKISTAN, pb.ProvinceId_CHINA},
		            StabilityBase:  2,
	                Region:         pb.Region_SOUTH_ASIA,
	                Coastal:        true,
	                InitInfluence:  0,
	                InitGovernment: pb.Government_WEAK,
	            },
            },
        },
    }

	set, err := NewSettings(testProto)
	if err != nil {
		t.Fatalf("NewSettings: unexpected error: %e", err)
	}
	m, err := set.InitState()
	if err != nil {
		t.Fatalf("NewSettings: unexpected error: %e", err)
	}
    if got, want := m.GetNetStability(pb.ProvinceId_IRAN), int32(4); got != want {
        t.Fatalf("Iran net stability: got %d, want %d", got, want)
    }
	if got, want := m.GetAlly(pb.ProvinceId_IRAN), pb.Player_NEITHER; got != want {
		t.Fatalf("Iran ally #1: got %s, want %s", got, want)
	}
    m.Infl(pb.ProvinceId_IRAN, pb.Player_USA, 1)
	if got, want := m.GetAlly(pb.ProvinceId_IRAN), pb.Player_USA; got != want {
		t.Fatalf("Iran ally #2: got %s, want %s", got, want)
	}
	if got, want := m.GetAlly(pb.ProvinceId_AFGHANISTAN), pb.Player_USSR; got != want {
		t.Fatalf("Afghanistan ally: got %s, want %s", got, want)
	}
	if got, want := m.GetAlly(pb.ProvinceId_INDIA), pb.Player_NEITHER; got != want {
		t.Fatalf("India ally: got %s, want %s", got, want)
	}
    if got, want := m.GetAlly(pb.ProvinceId_PAKISTAN), pb.Player_USA; got != want {
        t.Fatalf("Pakistan ally #1: got %s, want %s", got, want)
    }
    m.SetGov(pb.ProvinceId_PAKISTAN, pb.Government_AUTOCRACY)
    if got, want := m.GetAlly(pb.ProvinceId_PAKISTAN), pb.Player_USA; got != want {
        t.Fatalf("Pakistan ally #2: got %s, want %s", got, want)
    }
    m.SetLeader(pb.ProvinceId_PAKISTAN, "Hosni Mubarak")
    if got, want := m.GetAlly(pb.ProvinceId_PAKISTAN), pb.Player_NEITHER; got != want {
        t.Fatalf("Pakistan ally #3: got %s, want %s", got, want)
    }

	stateProto := &pb.GameState{}
	err = m.Marshal(stateProto)
	if err != nil {
		t.Fatalf("Marshal: unexpected error: %e", err)
	}
	newState, err := NewState(stateProto, m)
	if err != nil {
		t.Fatalf("NewState: unexpected error: %e", err)
	}
	if got, want := newState.GetAlly(pb.ProvinceId_IRAN), pb.Player_USA; got != want {
		t.Fatalf("Iran ally #3: got %d, want %d", got, want)
	}
}
