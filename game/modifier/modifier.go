package modifier

import (
	"github.com/machinule/nucrom/game/mechanic"
	"github.com/machinule/nucrom/game/modifier/heat"
	"github.com/machinule/nucrom/game/modifier/year"
	pb "github.com/machinule/nucrom/proto/gen"
	"reflect"
	"sort"
)

// A Mover modifies a mechanics internal state based on player moves.
type Mover interface {
	Move(move *pb.GameMove, mechanics *mechanic.Mechanics) error
}

// A Turner modifies a mechanics internal state based on the advancement of turns.
type Turner interface {
	Turn(turn *pb.GameTurn, mechanics *mechanic.Mechanics) error
}

// Modifiers provides accessors to moving and turning for various modifiers.
type Modifiers struct {
	// Modifiers are listed here. They will be accessible directly by name, or indirectly by the interfaces they implement.
	Heat heat.Modifier
  Year year.Modifier

	movers  []Mover
	turners []Turner
}

func New() *Modifiers {
	m := &Modifiers{}
	m.buildAccessors()
	return m
}

// buildAccessors will use reflection to add every field in Modifiers to internal lists based on which interfaces it implements.
func (m *Modifiers) buildAccessors() {
	mType := reflect.TypeOf(*m)
	mValue := reflect.Indirect(reflect.ValueOf(m))

	moverType := reflect.TypeOf((*Mover)(nil)).Elem()
	turnerType := reflect.TypeOf((*Turner)(nil)).Elem()

	// The order of fields should be deterministic. This is not guaranteed by the compiler, so sort them here.
	fieldNames := make([]string, mType.NumField())
	for i := 0; i < mType.NumField(); i++ {
		fieldNames[i] = mType.Field(i).Name
	}
	sort.Sort(sort.StringSlice(fieldNames))
	for _, fieldName := range fieldNames {
		f, ok := mType.FieldByName(fieldName)
		if !ok {
			continue
		}
		if reflect.PtrTo(f.Type).Implements(moverType) {
			m.movers = append(m.movers, mValue.FieldByName(fieldName).Addr().Interface().(Mover))
		}
		if reflect.PtrTo(f.Type).Implements(turnerType) {
			m.turners = append(m.turners, mValue.FieldByName(fieldName).Addr().Interface().(Turner))
		}
	}
}

func (m *Modifiers) Move(move *pb.GameMove, mechanics *mechanic.Mechanics) error {
	for _, t := range m.movers {
		if err := t.Move(move, mechanics); err != nil {
			return err
		}
	}
	return nil
}

func (m *Modifiers) Turn(turn *pb.GameTurn, mechanics *mechanic.Mechanics) error {
	for _, t := range m.turners {
		if err := t.Turn(turn, mechanics); err != nil {
			return err
		}
	}
	return nil
}
