package setup

import (
    "github.com/golang/protobuf/proto"
    "github.com/machinule/nucrom/proto/gen"
)

// ===== Game State/Settings =====

func CreateGameSettings() *pb.GameSettings {
    ret := &pb.GameSettings {
        PseudorandomSysSettings: createPseudorandomSystemSettings(),
        //VictorySysSettings:
        //HeatSysSettings:

        //ProvinceMechSettings:
        //LeaderMechSettings:
    }

    return ret
}

/*
func CreateGameState() pb.GameState {
    return nil
}
*/

// ===== Systems =====

func createPseudorandomSystemSettings() *pb.PseudorandomSystemSettings {
    return &pb.PseudorandomSystemSettings {
        InitSeed: proto.Int64(1),
    }
}

func createVictorySystemSettings() *pb.VictorySystemSettings {
    return &pb.VictorySystemSettings {
    }
}

// ===== Mechanics =====
