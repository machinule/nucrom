// This maintains and keeps track of the conflicts that occur in the world
package province

import (
	pb "github.com/machinule/nucrom/proto/gen"
    pseudo "github.com/machinule/nucrom/game/mechanic/pseudorandom"
)

// HELPERS

func (c *Conflict) GetModAttackerChance() int32 {
    // TODO: Return modified attacker chance
    return c.BaseChance()
}

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

// ACTIONS

// TODO: UNTESTED
// Processes a conflict (rolls to determine progress, handles outcome)
func (c *Conflict) Process(p *pseudo.State) {
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
    if c.attackers.progress == c.goal {
        // Attackers win
    } else if c.defenders.progress == c.goal {
        // Defenders win
    }
}
