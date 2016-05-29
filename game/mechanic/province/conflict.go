// This maintains and keeps track of the conflicts that occur in the world
package province

import (
	pseudo "github.com/machinule/nucrom/game/mechanic/pseudorandom"
	pb "github.com/machinule/nucrom/proto/gen"
)

type WarResult string

// State of war
const (
	ATTACKER WarResult = "Attacker"
	DEFENDER WarResult = "Defender"
	ONGOING  WarResult = "Ongoing"
)

// HELPERS

// Gets the attacker chance after accounting for modifiers
func (c *Conflict) GetModAttackerChance() int32 {
	// TODO: Return modified attacker chance
	return c.BaseChance()
}

// Gets the defender chance after accounting for modifiers
func (c *Conflict) GetModDefenderChance() int32 {
	// TODO: Return modifier defender chance
	return c.BaseChance()
}

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

// Gets a conflict by location
func (s *State) GetConflict(location pb.ProvinceId) *Conflict {
	return s.Conflicts[location]
}

// ACTIONS

// Creates a new conventional war
func (s *State) NewConventionalWar(defenders []pb.ProvinceId, attackers []pb.ProvinceId, locations []pb.ProvinceId) bool { // TODO: Error return
	for _, d := range defenders {
		if s.IsAtWar(d) || s.IsSiteOfConflict(d) {
			return false
		}
	}
	for _, a := range attackers {
		if s.IsAtWar(a) || s.IsSiteOfConflict(a) {
			return false
		}
	}
	for _, l := range locations {
		if s.IsAtWar(l) || s.IsSiteOfConflict(l) {
			return false
		}
	}
	// TODO: Logic for joining wars?
	c := &Conflict{
		name:   "War!", // TODO
		length: 0,
		attackers: Faction{
			members:  attackers,
			progress: 0,
		},
		defenders: Faction{
			members:  defenders,
			progress: 0,
		},
		goal:        2,      // TODO: Settings
		base_chance: 150000, // TODO: Settings
		locations:   locations,
	}
	// For now it maps only to the first location
	s.Conflicts[locations[0]] = c
	return true
}

// Creates a new civil war
// TODO

// Creates a new colonial war
// TODO

// Creates a new military intervention
// TODO

// Processes a conflict (rolls to determine progress, handles outcome)
// Returns false if conflict resolves
func (c *Conflict) Process(p *pseudo.State) WarResult {
	c.length = c.length + 1
	def_prog := p.Happens(c.GetModDefenderChance())
	att_prog := p.Happens(c.GetModAttackerChance())
	if att_prog {
		// Attackers progress
		c.attackers.progress++
	}
	if def_prog {
		// Defenders progress
		c.defenders.progress++
	}
	if att_prog && def_prog {
		c.goal++
	}
	if c.attackers.progress >= c.goal {
		return ATTACKER
	} else if c.defenders.progress >= c.goal {
		return DEFENDER
	}
	return ONGOING
}

// Processes all conflicts
func (s *State) ResolveConflicts(p *pseudo.State) {
	for _, c := range s.Conflicts {
		if c.Process(p) != ONGOING {
			delete(s.Conflicts, c.Defenders()[0])
		}
	}
}
