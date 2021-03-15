package event

import "github.com/sh3rp/flo/pkg/model"

type Event struct {
	Id   string
	Type EventType
	Data []model.DataObject
}

type EventType int

var (
	EventType_TaskComplete_String  = "complete"
	EventType_TaskFailed_String    = "failed"
	EventType_TaskSucceeded_String = "succeeded"
)

const (
	EventType_TaskComplete = iota
	EventType_TaskFailed
	EventType_TaskSucceeded
)

func (et EventType) String() {
	return []string{
		EventType_TaskComplete_String,
		EventType_TaskFailed_String,
		EventType_TaskSucceeded_String,
	}[et]
}

func ParseEventType(str string) EventType {
	switch str {
	case EventType_TaskComplete_String:
		return EventType(EventType_TaskComplete)
	case EventType_TaskFailed_String:
		return EventType(EventType_TaskFailed)
	case EventType_TaskSucceeded_String:
		return EventType(EventType_TaskSucceeded)
	default:
		return EventType(-1)
	}
}
