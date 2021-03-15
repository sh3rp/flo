package event

import "errors"

type Channel string

type EventBus interface {
	Subscribe(Channel, func(Event))
	Publish(Channel, Event) error
}

type inmemEventBus struct {
	listeners map[Channel][]func(Event)
}

func (ceb inmemEventBus) Subscribe(c Channel, f func(Event)) {
	if _, ok := ceb.listeners[c]; !ok {
		ceb.listeners[c] = []func(Event){}
	}
	ceb.listeners[c] = append(ceb.listeners[c], f)
}

func (ceb inmemEventBus) Publish(c Channel, evt Event) error {
	if _, ok := ceb.listeners[c]; ok {
		for _, f := range ceb.listeners[c] {
			f(evt)
		}
	} else {
		return errors.New("No such channel")
	}
}
