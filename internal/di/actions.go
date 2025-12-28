package di

import someAction "go-template/internal/app/action/some-action"

func (di *DI) SomeAction() *someAction.SomeAction {
	return someAction.New()
}
