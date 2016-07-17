package heat

import (
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestUnmarshalDefault(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if got, want := s.Init, defaultInit; got != want {
		t.Errorf("Init: got %v, want %v", got, want)
	}
	if got, want := s.Min, defaultMin; got != want {
		t.Errorf("Min: got %v, want %v", got, want)
	}
	if got, want := s.Max, defaultMax; got != want {
		t.Errorf("Max: got %v, want %v", got, want)
	}
	if got, want := s.Decay, defaultDecay; got != want {
		t.Errorf("Decay: got %v, want %v", got, want)
	}
}

func TestUnmarshal(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{
		HeatSettings: &pb.HeatSettings{
			Init: 50,
		},
	})
	if got, want := s.Init, int32(50); got != want {
		t.Errorf("Init: got %v, want %v", got, want)
	}
	if got, want := s.Min, int32(0); got != want {
		t.Errorf("Min: got %v, want %v", got, want)
	}
	if got, want := s.Max, int32(0); got != want {
		t.Errorf("Max: got %v, want %v", got, want)
	}
	if got, want := s.Decay, int32(0); got != want {
		t.Errorf("Decay: got %v, want %v", got, want)
	}
}

func TestValidateDefaultIsValid(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if err := s.Validate(); err != nil {
		t.Errorf("Validate(): got %v, want nil", err)
	}
}

type InvalidTestCase *pb.GameSettings

func TestValidateInvalid(t *testing.T) {
	cases := []InvalidTestCase{
		{
			HeatSettings: &pb.HeatSettings{
				Init:  50,
				Min:   55,
				Max:   100,
				Decay: 5,
			},
		},
		{
			HeatSettings: &pb.HeatSettings{
				Init:  150,
				Min:   0,
				Max:   100,
				Decay: 5,
			},
		},
		{
			HeatSettings: &pb.HeatSettings{
				Init:  125,
				Min:   155,
				Max:   100,
				Decay: 5,
			},
		},
	}
	for _, c := range cases {
		s := Settings{}
		s.Unmarshal(c)
		if err := s.Validate(); err == nil {
			t.Errorf("Validate(): got %v, want error", err)
		}
	}
}

func TestMarshal(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{
		HeatSettings: &pb.HeatSettings{
			Init: 50,
		},
	})
	msg := &pb.GameSettings{
		HeatSettings: &pb.HeatSettings{
			Init: 60,
		},
	}
	s.Marshal(msg)

	if got, want := msg.HeatSettings.Init, int32(50); got != want {
		t.Errorf("Init: got %v, want %v", got, want)
	}
	if got, want := msg.HeatSettings.Min, int32(0); got != want {
		t.Errorf("Min: got %v, want %v", got, want)
	}
	if got, want := msg.HeatSettings.Max, int32(0); got != want {
		t.Errorf("Max: got %v, want %v", got, want)
	}
	if got, want := msg.HeatSettings.Decay, int32(0); got != want {
		t.Errorf("Decay: got %v, want %v", got, want)
	}
}

func TestInitialize(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	state := &pb.GameState{}
	s.Initialize(state)

	if got, want := state.HeatState.Heat, defaultInit; got != want {
		t.Errorf("Init: got %v, want %v", got, want)
	}
}
