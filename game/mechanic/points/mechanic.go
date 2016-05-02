package points

import (
	pb "github.com/machinule/nucrom/proto/gen"
)

// ACTIONS

func (s *State) ChngPOL(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.pol = s.ussr.pol + magnitude
	} else if player == pb.Player_USA {
		s.usa.pol = s.usa.pol + magnitude
	}
}

func (s *State) ChngMIL(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.mil = s.ussr.mil + magnitude
	} else if player == pb.Player_USA {
		s.usa.mil = s.usa.mil + magnitude
	}
}

func (s *State) ChngCOV(player pb.Player, magnitude int32) {
	if player == pb.Player_USSR {
		s.ussr.cov = s.ussr.cov + magnitude
	} else if player == pb.Player_USA {
		s.usa.cov = s.usa.cov + magnitude
	}
}
