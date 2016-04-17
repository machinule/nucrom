package year

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/machinule/nucrom/proto/gen"
)

type state struct {
	s    *settings
	year int32
}

func NewState(stateProto *pb.GameState, prev *state) (*state, error) {
	if prev == nil {
		return nil, fmt.Errorf("recieved nil previous state, unable to propogate settings.")
	}
	return &state{
		s:    prev.s,
		year: stateProto.GetYearSysState().GetYear(),
	}, nil
}

func (s *state) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	if stateProto.GetYearSysState() == nil {
		stateProto.YearSysState = &pb.YearSystemState{}
	}
	stateProto.GetYearSysState().Year = proto.Int32(s.year)
	return nil
}

func (s *state) Year() int32 {
	return s.year
}

func (s *state) Incr() {
	s.year += s.s.incr
}
