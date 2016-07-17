package year

import (
	"github.com/machinule/nucrom/game/mechanic"
	pb "github.com/machinule/nucrom/proto/gen"
)

type Modifier struct {
}

func (m *Modifier) Turn(turn *pb.GameTurn, mechanics *mechanic.Mechanics) error {
	mechanics.Year.State.Incr()
	return nil
}
