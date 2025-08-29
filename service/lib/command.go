package lib

import (
	"github.com/rs/zerolog/log"
)

func CommandRunner(events chan *Event) {
	for {
		event := <-events

		switch event.Type {
		case WindowEvent:
			log.Debug().Any("event", event).Msg("")
		case PlusdeckEvent:
			log.Debug().Any("event", event).Msg("")
		case KeyActivityReport:
			log.Debug().Any("event", event).Msg("")
		case Notification:
			log.Debug().Any("event", event).Msg("")
		}
	}
}
