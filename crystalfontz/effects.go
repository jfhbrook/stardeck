package crystalfontz

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

const deviceWidth int = 16

type Marquee struct {
	row     byte
	text    string
	tick    time.Duration
	pause   time.Duration
	shift   int
	running bool
	client  *Crystalfontz
}

func NewMarquee(row byte, text string, client *Crystalfontz) (*Marquee, error) {
	if row < 0 || row > 1 {
		return nil, errors.New(fmt.Sprintf("Invalid row: %d", row))
	}

	tick := time.Duration(0.3 * float64(time.Second))
	pause := time.Duration(0.3 * float64(time.Second))
	text = fmt.Sprintf("%-16s", text)

	marquee := Marquee{
		row:     row,
		text:    text,
		tick:    tick,
		pause:   pause,
		shift:   0,
		running: false,
		client:  client,
	}

	return &marquee, nil
}

func (m *Marquee) line() []byte {
	left := m.text[m.shift:]
	right := m.text[0:m.shift]
	middle := strings.Repeat(" ", max(16-len(m.text), 1))
	assembled := []byte(left + middle + right)
	return assembled[0:16]
}

func (m *Marquee) send() {
	err := m.client.SendData(m.row, 0, m.line(), -1.0, -1)

	if err != nil {
		log.Error().Err(err).Msg("Error while rendering marquee")
	}
}

func (m *Marquee) Start() {
	if m.running {
		return
	}
	m.running = true

	m.send()
	m.shift += 1

	if !m.running {
		return
	}

	time.Sleep(m.pause)

	for {
		if !m.running {
			return
		}

		m.send()
		m.shift += 1

		if m.shift > 16 {
			m.shift = 0
		}

		time.Sleep(m.tick)
	}
}

func (m *Marquee) Stop() {
	m.running = false
}
