// Package settings manages the settings of the game.
package settings

import (
	"github.com/machinule/nucrom/game/settings/heat"
	"github.com/machinule/nucrom/game/settings/points"
	"github.com/machinule/nucrom/game/settings/pseudorandom"
	"github.com/machinule/nucrom/game/settings/year"
	pb "github.com/machinule/nucrom/proto/gen"
	"reflect"
	"sort"
)

// A MechanicSettings contains enough information to produce an initial state, and can be de/serialized with pb.GameSettings.
type MechanicSettings interface {
	// Marshal will write this MechanicSettings values into the GameSettings.
	Marshal(*pb.GameSettings) error

	// Unmarshal will overwrite this MechanicSettings values with those in the GameSettings.
	Unmarshal(*pb.GameSettings) error

	// Validate returns any validation errors.
	Validate() error

	// Initialize initializes the GameState based on the settings.
	Initialize(*pb.GameState) error
}

// Settings provides accessors to settings of various mechanics.
type Settings struct {
	// Settings are listed here. They will be accessible directly by name, or indirectly by the MechanicSettings interface.
	Heat         heat.Settings
	Points       points.Settings
	Pseudorandom pseudorandom.Settings
	Year         year.Settings

	mechanicSettings []MechanicSettings
}

// New produces a new Settings with default values.
func New() *Settings {
	s := &Settings{}
	s.buildAccessors()
	s.Unmarshal(&pb.GameSettings{})
	return s
}

// buildAccessors will use reflection to add every field in Settings to an internal list as a MechanicSettings.
func (s *Settings) buildAccessors() {
	// The Type of the Settings receiver.
	sType := reflect.TypeOf(*s)
	// The Value of the Settings receiver.
	sValue := reflect.Indirect(reflect.ValueOf(s))
	// The Type of the MechanicSettings interface.
	iType := reflect.TypeOf((*MechanicSettings)(nil)).Elem()

	// The order of fields should be deterministic. This is not guaranteed by the compiler, so sort them here.
	fieldNames := make([]string, sType.NumField())
	for i := 0; i < sType.NumField(); i++ {
		fieldNames[i] = sType.Field(i).Name
	}
	sort.Strings(fieldNames)

	// Collect each field that implements MechanicSettings.
	for _, fieldName := range fieldNames {
		f, ok := sType.FieldByName(fieldName)
		if !ok {
			continue
		}
		if reflect.PtrTo(f.Type).Implements(iType) {
			s.mechanicSettings = append(s.mechanicSettings, sValue.FieldByName(fieldName).Addr().Interface().(MechanicSettings))
		}
	}
}

// Marshal marshals the settings into a GameSettings message.
func (s *Settings) Marshal(msg *pb.GameSettings) error {
	for _, m := range s.mechanicSettings {
		if err := m.Marshal(msg); err != nil {
			return err
		}
	}
	return nil
}

// Unmarshal unmarshals a GameSettings message into this Settings.
func (s *Settings) Unmarshal(msg *pb.GameSettings) error {
	for _, m := range s.mechanicSettings {
		if err := m.Unmarshal(msg); err != nil {
			return err
		}
	}
	return nil
}

// Validate validates the current Settings.
func (s *Settings) Validate() error {
	for _, m := range s.mechanicSettings {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Initialize initializes a GameState message from this Settings.
func (s *Settings) Initialize(state *pb.GameState) error {
	for _, m := range s.mechanicSettings {
		if err := m.Initialize(state); err != nil {
			return err
		}
	}
	return nil
}
