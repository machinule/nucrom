// Package mechanic manages the state and settings of mechanics.
package mechanic

import (
	"github.com/machinule/nucrom/game/mechanic/heat"
	"github.com/machinule/nucrom/game/mechanic/points"
	"github.com/machinule/nucrom/game/mechanic/province"
	"github.com/machinule/nucrom/game/mechanic/pseudorandom"
	"github.com/machinule/nucrom/game/mechanic/year"
	pb "github.com/machinule/nucrom/proto/gen"
	"reflect"
	"sort"
)

// An Initializer initializes itself from GameSettings.
type Initializer interface {
	Initialize(settings *pb.GameSettings) error
}

// A StateHolder maintains some portion of the GameState.
type StateHolder interface {
	SetState(state *pb.GameState) error
	GetState(state *pb.GameState) error
}

// Mechanics provides accessors to settings and state registries of various mechanics.
type Mechanics struct {
	// Mechanics are listed here. They will be accessible directly by name, or indirectly by the interfaces they implement.
	Heat         heat.Mechanic
	Points       points.Mechanic
	Province     province.Mechanic
	Pseudorandom pseudorandom.Mechanic
	Year         year.Mechanic

	initializers []Initializer
	stateHolders []StateHolder
}

func New() *Mechanics {
	m := &Mechanics{}
	m.buildAccessors()
	return m
}

// buildAccessors will use reflection to add every field in Mechanics to internal lists based on which interfaces it implements.
func (m *Mechanics) buildAccessors() {
	mType := reflect.TypeOf(*m)
	mValue := reflect.Indirect(reflect.ValueOf(m))

	initializerType := reflect.TypeOf((*Initializer)(nil)).Elem()
	stateHolderType := reflect.TypeOf((*StateHolder)(nil)).Elem()

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
		if reflect.PtrTo(f.Type).Implements(initializerType) {
			m.initializers = append(m.initializers, mValue.FieldByName(fieldName).Addr().Interface().(Initializer))
		}
		if reflect.PtrTo(f.Type).Implements(stateHolderType) {
			m.stateHolders = append(m.stateHolders, mValue.FieldByName(fieldName).Addr().Interface().(StateHolder))
		}
	}
}

func (m *Mechanics) Initialize(settings *pb.GameSettings) error {
	for _, i := range m.initializers {
		if err := i.Initialize(settings); err != nil {
			return err
		}
	}
	return nil
}

func (m *Mechanics) GetState(state *pb.GameState) error {
	for _, h := range m.stateHolders {
		if err := h.GetState(state); err != nil {
			return err
		}
	}
	return nil
}

func (m *Mechanics) SetState(state *pb.GameState) error {
	for _, h := range m.stateHolders {
		if err := h.SetState(state); err != nil {
			return err
		}
	}
	return nil
}
