// Package mechanic manages the state and settings of mechanics.
package mechanic

type Mechanic interface {
}

// Mechanics provides accessors to settings and state registries of various mechanics.
type Mechanics struct {
  // Mechanics are listed here. They will be accessible directly by name, or indirectly through the
  // Mechanic interface via reflection in New().
  Heat heat.Mechanic
  Points points.Mechanic
  Province province.Mechanic
  Pseudorandom pseudorandom.Mechanic
  Year year.Mechanic
  
  mechanics []Mechanic
}

func New() *Mechanics {
  m := &Mechanics{}
  mType := reflect.TypeOf(*m)
  mValue := reflect.Indirect(reflect.ValueOf(m))
  iType := reflect.TypeOf((*Mechanic)(nil)).Elem()
  for i := 0; i < mType.NumField(); i++ {
    f := mType.Field(i)
    if reflect.PtrTo(f.Type).Implements(iType) {
      m.mechanics = append(m.mechanics, mValue.Field(i).Addr().Interface().(Mechanic))
    }
  }
  return m
}



/*
func (m *Mechanics) mechanics() []Mechanic {
  reflect.ValueOf(m).
}

func (m *Mechanics) Set(settings *pb.GameSettings) error {
  for _, settings := range s.settings() {
    if err := settings.Validate(settingProto); err != nil {
      return err
    }
  }
	return nil  
}

// usage:
// m := mechanic.New(settingsProto)
// s := m.InitState()
// s := m.NewState(gameStateProto)
// 



type Settings interface {
  Set(*pb.GameSettings) error
  Validate(*pb.GameSettings) error
  Init(*pb.GameState) error
  NewState(*pb.GameState) State
}

type State interface {
  Set(Settings, *pb.GameState) error
  Marshal(*pb.GameState) error
}

type GameSettings struct {
  Heat heat.Settings
  Points points.Settings
  Province province.Settings
  Pseudorandom pseudorandom.Settings
  Year year.Settings
}

func (s *GameSettings) settings() []Settings {
  return []Settings{
    s.Heat,
    s.Points,
    s.Province,
    s.Pseudorandom,
    s.Year,
  }
}

func (s *GameSettings) Validate(settingsProto *pb.GameSettings) error {
  for _, settings := range s.settings() {
    if err := settings.Validate(settingProto); err != nil {
      return err
    }
  }
	return nil
}

func (s *GameSettings) Set(settingsProto *pb.GameSettings) error {
  for _, settings := range s.settings() {
    if err := settings.Set(settingProto); err != nil {
      return err
    }
  }
	return nil
}

func (s *GameSettings) Init(stateProto *pb.GameState) error {
  for _, settings := range s.settings() {
    if err := settings.Init(stateProto); err != nil {
      return err
    }
  }
  return nil
}

func (s *GameSettings) NewState(stateProto *pb.GameState) error {
  state := &GameState{
    settings: s,
  }
  state.Set(stateProto)
  for _, settings := range s.settings() {
    if err := settings.Init(stateProto); err != nil {
      return err
    }
  }
  return nil
}


type State struct {
  Heat heat.State
  Points points.State
  Province province.State
  Pseudorandom pseudorandom.State
  Year year.State
}

type mechanicState interface {
  Update(stateProto *pb.GameState) error
  Marshal(stateProto *pb.GameState) error
}

func (s *State) mechanicStates() []mechanicState {
  return []mechanicState{
    s.Heat,
    s.Points,
    s.Province,
    s.Pseudorandom,
    s.Year,
  }
}

func (s *State) Update(stateProto *pb.GameState) error {
  for _, state := range s.mechanicStates() {
    if err := state.Update(stateProto); err != nil {
      return err
    }
  }
  return nil
}

func (s *State) Marshal(stateProto *pb.GameState) error {
  for _, state := range s.mechanicStates() {
    if err := state.Marshal(stateProto); err != nil {
      return err
    }
  }
  return nil
}
*/
