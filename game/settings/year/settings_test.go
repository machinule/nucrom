package year

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestUnmarshalDefault(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	if !proto.Equal(&s.YearSettings, defaultSettings) {
		t.Errorf("Default not applied: got %v, want %v", s, defaultSettings)
	}
}

func TestUnmarshal(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{
		YearSettings: &pb.YearSettings{
			InitYear: 202012,
		},
	})
	if got, want := s.InitYear, int32(202012); got != want {
		t.Errorf("InitYear: got %v, want %v", got, want)
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

	if got, want := msg.YearSettings.InitYear, defaultSettings.InitYear; got != want {
		t.Errorf("InitYear: got %v, want %v", got, want)
	}
}

func TestInitialize(t *testing.T) {
	s := Settings{}
	s.Unmarshal(&pb.GameSettings{})
	state := &pb.GameState{}
	s.Initialize(state)

	if got, want := state.YearState.Year, defaultSettings.InitYear; got != want {
		t.Errorf("Year: got %v, want %v", got, want)
	}
}
