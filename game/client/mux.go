package client

import (
	pb "github.com/machinule/nucrom/proto/gen"
)

// A mux multiplexes over multiple clients, presenting a single current client. EndTurn cycles the client.
type mux struct {
	curr    int
	clients []Client
}

func NewMux(address string, num int) Client {
	m := &mux{
		curr:    0,
		clients: make([]Client, num),
	}
	for i := 0; i < num; i++ {
		m.clients[i] = New(address)
	}
	return m
}

func (m *mux) Connect() error {
	for _, c := range m.clients {
		if err := c.Connect(); err != nil {
			return err
		}
	}
	return nil
}

func (m *mux) Join() error {
	for _, c := range m.clients {
		if err := c.Join(); err != nil {
			return err
		}
	}
	return nil
}

func (m *mux) State() *pb.GameState {
  return m.clients[m.curr].State()
}

func (m *mux) StartTurn() error {
	return m.clients[m.curr].StartTurn()
}

func (m *mux) EndTurn() error {
	if err := m.clients[m.curr].EndTurn(); err != nil {
		return err
	}
	m.curr = (m.curr + 1) % len(m.clients)
	return nil
}

func (m *mux) Turn() int {
	return m.clients[m.curr].Turn()
}

func (m *mux) Player() pb.Player {
	return m.clients[m.curr].Player()
}

func (m *mux) GameOver() bool {
	return m.curr == len(m.clients)-1 && m.clients[m.curr].GameOver()
}
