package some_action

type SomeAction struct{}

func New() *SomeAction {
	return &SomeAction{}
}

func (a *SomeAction) Do() {}
