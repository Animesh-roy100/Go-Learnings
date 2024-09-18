package main

type CommandBus interface {
	apply(command Command)
}

type defaultCommandBus struct{}

func (c defaultCommandBus) apply(command Command) {

}

func NewCommandBus() CommandBus {
	return &defaultCommandBus{}
}
