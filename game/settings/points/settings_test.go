package points

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestUnmarshalDefault(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if !proto.Equal(&s.PointsSettings, defaultSettings) {
		t.Errorf("Default not applied: got %v, want %v", s, defaultSettings)
	}
}

func TestUnmarshal(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{
		PointsSettings: &pb.PointsSettings{
			Usa: &pb.PointStoresSettings{
				Political: &pb.PointStoreSettings{
					Init:   1,
					Income: 2,
				},
				Military: &pb.PointStoreSettings{
					Init:   3,
					Income: 4,
				},
				Covert: &pb.PointStoreSettings{
					Init:   5,
					Income: 6,
				},
			},
			Ussr: &pb.PointStoresSettings{
				Political: &pb.PointStoreSettings{
					Init:   7,
					Income: 8,
				},
				Military: &pb.PointStoreSettings{
					Init:   9,
					Income: 10,
				},
				Covert: &pb.PointStoreSettings{
					Init:   11,
					Income: 12,
				},
			},
		},
	})
	if got, want := s.Usa.Political.Init, int32(1); got != want {
		t.Errorf("Usa.Political.Init: got %v, want %v", got, want)
	}
	if got, want := s.Usa.Political.Income, int32(2); got != want {
		t.Errorf("Usa.Political.Income: got %v, want %v", got, want)
	}
	if got, want := s.Usa.Military.Init, int32(3); got != want {
		t.Errorf("Usa.Military.Init: got %v, want %v", got, want)
	}
	if got, want := s.Usa.Military.Income, int32(4); got != want {
		t.Errorf("Usa.Military.Income: got %v, want %v", got, want)
	}
	if got, want := s.Usa.Covert.Init, int32(5); got != want {
		t.Errorf("Usa.Covert.Init: got %v, want %v", got, want)
	}
	if got, want := s.Usa.Covert.Income, int32(6); got != want {
		t.Errorf("Usa.Covert.Income: got %v, want %v", got, want)
	}

	if got, want := s.Ussr.Political.Init, int32(7); got != want {
		t.Errorf("Ussr.Political.Init: got %v, want %v", got, want)
	}
	if got, want := s.Ussr.Political.Income, int32(8); got != want {
		t.Errorf("Ussr.Political.Income: got %v, want %v", got, want)
	}
	if got, want := s.Ussr.Military.Init, int32(9); got != want {
		t.Errorf("Ussr.Military.Init: got %v, want %v", got, want)
	}
	if got, want := s.Ussr.Military.Income, int32(10); got != want {
		t.Errorf("Ussr.Military.Income: got %v, want %v", got, want)
	}
	if got, want := s.Ussr.Covert.Init, int32(11); got != want {
		t.Errorf("Ussr.Covert.Init: got %v, want %v", got, want)
	}
	if got, want := s.Ussr.Covert.Income, int32(12); got != want {
		t.Errorf("Ussr.Covert.Income: got %v, want %v", got, want)
	}

}

func TestValidateDefaultIsValid(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if err := s.Validate(); err != nil {
		t.Errorf("Validate(): got %v, want nil", err)
	}
}

func TestMarshal(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	msg := &pb.GameSettings{}
	s.Marshal(msg)

	if got, want := msg.PointsSettings.Usa.Political.Init, defaultSettings.Usa.Political.Init; got != want {
		t.Errorf("Init: got %v, want %v", got, want)
	}
}

func TestInitialize(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{
		PointsSettings: &pb.PointsSettings{
			Usa: &pb.PointStoresSettings{
				Political: &pb.PointStoreSettings{
					Init:   1,
					Income: 2,
				},
				Military: &pb.PointStoreSettings{
					Init:   3,
					Income: 4,
				},
				Covert: &pb.PointStoreSettings{
					Init:   5,
					Income: 6,
				},
			},
			Ussr: &pb.PointStoresSettings{
				Political: &pb.PointStoreSettings{
					Init:   7,
					Income: 8,
				},
				Military: &pb.PointStoreSettings{
					Init:   9,
					Income: 10,
				},
				Covert: &pb.PointStoreSettings{
					Init:   11,
					Income: 12,
				},
			},
		},
	})
	state := &pb.GameState{}
	s.Initialize(state)

	if got, want := state.PointsState.Usa.Political.Count, int32(1); got != want {
		t.Errorf("Usa.Political.Count: got %v, want %v", got, want)
	}
	if got, want := state.PointsState.Usa.Military.Count, int32(3); got != want {
		t.Errorf("Usa.Military.Count: got %v, want %v", got, want)
	}
	if got, want := state.PointsState.Usa.Covert.Count, int32(5); got != want {
		t.Errorf("Usa.Covert.Count: got %v, want %v", got, want)
	}

	if got, want := state.PointsState.Ussr.Political.Count, int32(7); got != want {
		t.Errorf("Ussr.Political.Count: got %v, want %v", got, want)
	}
	if got, want := state.PointsState.Ussr.Military.Count, int32(9); got != want {
		t.Errorf("Ussr.Military.Count: got %v, want %v", got, want)
	}
	if got, want := state.PointsState.Ussr.Covert.Count, int32(11); got != want {
		t.Errorf("Ussr.Covert.Count: got %v, want %v", got, want)
	}
}
