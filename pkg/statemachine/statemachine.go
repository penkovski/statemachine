package statemachine

type StateMachine struct {
	state   string
	actions chan func()
}

func New() *StateMachine {
	return &StateMachine{
		state:   "initial",
		actions: make(chan func()),
	}
}

func (s *StateMachine) Run(quit chan chan struct{}) {
	for {
		select {
		case f := <-s.actions:
			f()
		case q := <-quit:
			close(q)
			return
		}
	}
}

func (s *StateMachine) NewState(state string) {
	c := make(chan struct{})
	s.actions <- func() {
		s.state = state
		c <- struct{}{}
	}
	<-c
	return
}

func (s *StateMachine) State() string {
	c := make(chan string)
	s.actions <- func() {
		c <- s.state
		return
	}
	return <-c
}
