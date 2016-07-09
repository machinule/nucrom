package points

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestMechanic(t *testing.T) {
	testProto := &pb.GameSettings{
		PointsSettings: &pb.PointsSettings{
			UsaSettings: &pb.PointSettings{
				PoliticalStoreInit:  5,
				PoliticalIncomeBase: 5,
				MilitaryStoreInit:   4,
				MilitaryIncomeBase:  4,
				CovertStoreInit:     2,
				CovertIncomeBase:    2,
			},
			UssrSettings: &pb.PointSettings{
				PoliticalStoreInit:  5,
				PoliticalIncomeBase: 5,
				MilitaryStoreInit:   4,
				MilitaryIncomeBase:  4,
				CovertStoreInit:     2,
				CovertIncomeBase:    2,
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
	stateProto := &pb.GameState{}
	err = m.Marshal(stateProto)
	if err != nil {
		t.Fatalf("Marshal: unexpected error: %e", err)
	}
	n, err := NewState(stateProto, m.settings)
	if err != nil {
		t.Fatalf("NewState: unexpected error: %e", err)
	}
	if got, want := n.POL(pb.Player_USSR), int32(5); got != want {
		t.Fatalf("USSR pol: got %d, want %d", got, want)
	}
	if got, want := n.COV(pb.Player_USA), int32(2); got != want {
		t.Fatalf("USA cov: got %d, want %d", got, want)
	}
	if got, want := n.MIL(pb.Player_USSR), int32(4); got != want {
		t.Fatalf("USSR mil: got %d, want %d", got, want)
	}
	n.ChngPOL(pb.Player_USSR, int32(2))
	if got, want := n.POL(pb.Player_USSR), int32(7); got != want {
		t.Fatalf("USSR pol w/ change: got %d, want %d", got, want)
	}
	n.ChngCOV(pb.Player_USA, int32(5))
	if got, want := n.COV(pb.Player_USA), int32(7); got != want {
		t.Fatalf("USA cov w/ change: got %d, want %d", got, want)
	}
	n.ChngMIL(pb.Player_USSR, int32(-2))
	if got, want := n.MIL(pb.Player_USSR), int32(2); got != want {
		t.Fatalf("USSR mil w/ change: got %d, want %d", got, want)
	}
}
