package util

type Connection struct {
	From *State
	To *State

	Fun func(interface{})
}

type State struct {
	ID 		int
}

type StateMachine struct {
	currentState *State

	States 	[]*State
	Conn 	[]*Connection
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		currentState: nil,
		States:       nil,
		Conn:         nil,
	}
}

func (s *StateMachine) SwitchToState(state *State, args interface{}) {
	for _, ele := range s.Conn {
		if state.ID == ele.To.ID {
			ele.Fun(args)
			s.currentState = ele.To
			break
		}
	}
}

func (s *StateMachine) AppendState(state *State) {
	if s.currentState == nil {
		s.currentState = state
	}

	s.States = append(s.States, state)
}

func (s *StateMachine) AppendRoute(from *State, to *State, fp func(interface{})) {
	s.Conn = append(s.Conn, &Connection{
		From: from,
		To:   to,
		Fun:  fp,
	})
}
