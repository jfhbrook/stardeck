package crystalfontz

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type EffectCommand int

const StopEffectCommand EffectCommand = 1

func (lcd *Crystalfontz) Marquee(row byte, text string) (*chan EffectCommand, error) {
	if row < 0 || row > 1 {
		return nil, errors.New(fmt.Sprintf("Invalid row: %d", row))
	}

	tick := time.Duration(0.3 * float64(time.Second))
	pause := tick
	text = fmt.Sprintf("%-16s", text)
	shift := 0

	commands := make(chan EffectCommand)

	line := func() []byte {
		left := text[shift:]
		right := text[0:shift]
		middle := strings.Repeat(" ", max(16-len(text), 1))
		assembled := []byte(left + middle + right)
		return assembled[0:16]
	}

	send := func() {
		err := lcd.SendData(row, 0, line(), -1.0, -1)

		if err != nil {
			log.Error().Err(err).Msg("Error while rendering marquee")
		}
	}

	loop := func() {
		send()
		time.Sleep(pause)

		shift += 1

		for {
			select {
			case cmd := <-commands:
				if cmd == StopEffectCommand {
					return
				}
			default:
				shift += 1

				if shift > 16 {
					shift = 0
				}

				time.Sleep(tick)
			}
		}
	}

	go loop()

	return &commands, nil
}
