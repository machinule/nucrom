package settings

import (
	"github.com/golang/protobuf/proto"
	pb "github.com/machinule/nucrom/proto/gen"
	"testing"
)

func TestDefaultIsValid(t *testing.T) {
	s := New()
	if err := s.Validate(); err != nil {
		t.Errorf("Default settings not valid: %s", err)
	}
}

func TestDefaultMarshalUnmarshal(t *testing.T) {
	s1 := New()
	a := &pb.GameSettings{}
	s1.Marshal(a)
	s2 := New()
	s2.Unmarshal(a)
	b := &pb.GameSettings{}
	s2.Marshal(b)
	if !proto.Equal(a, b) {
		t.Errorf("b = Mashal(Unmarshal(a)) a != b: a = %v, b = %v", a, b)
	}
}
