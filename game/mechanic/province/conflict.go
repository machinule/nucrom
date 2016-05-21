// This maintains and keeps track of the conflicts that occur in the world
package province

import (
	pb "github.com/machinule/nucrom/proto/gen"
)

// HELPERS

// QUERIES

// Checks if province is the location of a conflict
func (s *State) IsSiteOfConflict(id pb.ProvinceId) bool {
	for _, c := range s.Conflicts {
		for _, l := range c.Locations() {
			if l == id {
				return true
			}
		}
	}
	return false
}

// Checks if province is at war (not necessarily location of conflict)
func (s *State) IsAtWar(id pb.ProvinceId) bool {
	for _, c := range s.Conflicts {
		for _, a := range c.Attackers() {
			if a == id {
				return true
			}
		}
		for _, d := range c.Defenders() {
			if d == id {
				return true
			}
		}
	}
	return false
}

// ACTIONS

// Processes a conflict (rolls to determine progress, handles outcome)
func (c *Conflict) Process() {
	// TODO: Pseudorandom mechanic
	c.length = c.length + 1
}
