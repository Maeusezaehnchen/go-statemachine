package gosm

import (
	"fmt"
	"testing"
)

const (
	START int = iota
	STATE1
	STATE2
	STATE3
	END
)

type TestInterface struct {}

func SwitchSTARTTo1(i interface{}) {
	fmt.Println("Switching from State START to State 2")
}

func Switch1To2(i interface{}) {
	fmt.Println("Switching from State 1 to State 2")
}

func Switch2To3(i interface{}) {
	fmt.Println("Switching from State 2 to State 3")
}

func Switch3ToEND(i interface{}) {
	fmt.Println("Switching from State 3 to State END")
}

func TestStateMachine_AppendRoute(t *testing.T) {
	sm := NewStateMachine()

	sm.AppendState(&State{ID:START})
	sm.AppendState(&State{ID:STATE1})
	sm.AppendState(&State{ID:STATE2})
	sm.AppendState(&State{ID:STATE3})
	sm.AppendState(&State{ID:END})

	sm.AppendRoute(&State{ID:START}, &State{ID:STATE1}, SwitchSTARTTo1)
	sm.AppendRoute(&State{ID:STATE1}, &State{ID:STATE2}, Switch1To2)
	sm.AppendRoute(&State{ID:STATE2}, &State{ID:STATE3}, Switch2To3)
	sm.AppendRoute(&State{ID:STATE3}, &State{ID:END}, Switch3ToEND)

	if len(sm.Conn) != 4 {
		_ = fmt.Errorf("failure appending states to Statemachine")
	}
}

func TestStateMachine_AppendState(t *testing.T) {
	sm := NewStateMachine()

	sm.AppendState(&State{ID:START})
	sm.AppendState(&State{ID:STATE1})
	sm.AppendState(&State{ID:STATE2})
	sm.AppendState(&State{ID:STATE3})
	sm.AppendState(&State{ID:END})

	if len(sm.States) != 5 {
		_ = fmt.Errorf("failure appending states to Statemachine")
	}

	if sm.currentState.ID != START {
		_ = fmt.Errorf("failure setting current state in Statemachine")
	}
}

func TestStateMachine_SwitchToState(t *testing.T) {
	sm := NewStateMachine()

	sm.AppendState(&State{ID:START})
	sm.AppendState(&State{ID:STATE1})
	sm.AppendState(&State{ID:STATE2})
	sm.AppendState(&State{ID:STATE3})
	sm.AppendState(&State{ID:END})

	sm.AppendRoute(&State{ID:START}, &State{ID:STATE1}, SwitchSTARTTo1)
	sm.AppendRoute(&State{ID:STATE1}, &State{ID:STATE2}, Switch1To2)
	sm.AppendRoute(&State{ID:STATE2}, &State{ID:STATE3}, Switch2To3)
	sm.AppendRoute(&State{ID:STATE3}, &State{ID:END}, Switch3ToEND)

	i := TestInterface{}

	sm.SwitchToState(&State{ID:STATE1}, &i)

	if sm.currentState.ID != STATE1 {
		_ = fmt.Errorf("failure switching states")
	}

	sm.SwitchToState(&State{ID:STATE2}, &i)

	if sm.currentState.ID != STATE2 {
		_ = fmt.Errorf("failure switching states")
	}

	sm.SwitchToState(&State{ID:STATE3}, &i)

	if sm.currentState.ID != STATE3 {
		_ = fmt.Errorf("failure switching states")
	}

	sm.SwitchToState(&State{ID:END}, &i)

	if sm.currentState.ID != END {
		_ = fmt.Errorf("failure switching states")
	}
}
