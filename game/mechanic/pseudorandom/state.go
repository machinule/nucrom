// The pseudorandom mechanic handles the rolling of all (pseudo)random events/actions.
package pseudorandom

import (
	"fmt"
	"github.com/machinule/nucrom/proto/gen"
	"math/rand"
	"sort"
)

// The state of the pseudorandom mechanic.
type State struct {
	settings *Settings
	seed     int64
	r        *rand.Rand
}

// NewState creates a new state from the GameState message and the previous state.
func NewState(stateProto *pb.GameState, settings *Settings) (*State, error) {
	if settings == nil {
		return nil, fmt.Errorf("received nil Settings, unable to continue.")
	}
	if stateProto.GetPseudorandomState() == nil {
		return nil, fmt.Errorf("received nil PseudorandomState, unable to continue.")
	}
	return &State{
		settings: settings,
		seed:     stateProto.GetPseudorandomState().Seed,
		r:        rand.New(rand.NewSource(stateProto.GetPseudorandomState().Seed)),
	}, nil
}

// Marshal fills in the GameState proto with the current state.
func (s *State) Marshal(stateProto *pb.GameState) error {
	if stateProto == nil {
		return fmt.Errorf("attempting to fill in nil GameState proto.")
	}
	stateProto.PseudorandomState = &pb.PseudorandomState{
		Seed: s.seed,
	}
	return nil
}

// ACTIONS

func (s *State) Reseed() {
	s.seed = s.r.Int63()
	s.r = rand.New(rand.NewSource(s.seed))
}

// Rolls for a particular value (rolls from 0 to 1,000,000)
func (s *State) Happens(chance int32) bool {
	if s.r.Int31n(1000000) < chance {
		return true
	}
	return false
}

// Weighted roll that Will is going to rip the fuck up
func (s *State) Roll(weights []int32) int32 {
	var pdf []int
	sum := int32(0)
	for _, w := range weights {
		if w <= 0 {
			continue
		}
		sum += w
		pdf = append(pdf, int(sum))
	}
	if sum == 0 {
		return -1
	}
	result := sort.SearchInts(pdf, int(s.r.Int31n(sum))+1)
	return int32(result)
}
