package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
)

const (
	lcdWidth int           = 16
	lcdPause time.Duration = time.Duration(1.0 * float64(time.Second))
	lcdTick                = time.Duration(0.3 * float64(time.Second))
)

type lcdLine struct {
	row         byte
	defaultText string
	text        string
	shift       int
	running     bool
	client      *crystalfontz.Client
}

func newLcdLine(row byte, defaultText string, client *crystalfontz.Client) *lcdLine {
	if row < 0 || row > 1 {
		panic(fmt.Sprintf("Invalid row: %d", row))
	}

	l := lcdLine{
		row:         row,
		defaultText: defaultText,
		text:        "",
		shift:       0,
		running:     false,
		client:      client,
	}

	l.update("")

	return &l
}

func (l *lcdLine) update(text string) {
	if text == "" {
		text = l.defaultText
	}

	if len(text) > lcdWidth {
		// Allow marquee text to scroll out of the screen
		text += strings.Repeat(" ", lcdWidth)
	} else {
		// Pad the text to at least the LCD's width
		text += strings.Repeat(" ", max(lcdWidth-len(text), 0))
	}

	if l.text != text {
		l.shift = 0
	}

	l.text = text
}

func (l *lcdLine) data() []byte {
	left := l.text[l.shift:]
	right := l.text[0:l.shift]
	data := []byte(left + right)
	// In case characters are multi-byte
	return data[0:lcdWidth]
}

func (l *lcdLine) send(data []byte) {
	log.Trace().
		Int("row", int(l.row)).
		Int("column", 0).
		Str("data", string(data)).
		Msg("Sending data to LCD")

	err := l.client.SendData(
		l.row,
		0,
		data,
		crystalfontz.NilFloat,
		crystalfontz.NilInt,
	)

	if err != nil {
		log.Error().
			Err(err).
			Int("row", int(l.row)).
			Int("column", 0).
			Str("data", string(data)).
			Msg("Error while sending data to LCD")
	}
}

func (l *lcdLine) scroll() {
	// Only scroll text if it's wider than the LCD
	if len(l.text) > lcdWidth {
		l.shift += 1
	}
	if l.shift >= len(l.text) {
		l.shift = 0
	}
}

func (l *lcdLine) loop() {
	text := l.text

	l.running = true

	l.send(l.data())
	l.scroll()

	if !l.running {
		return
	}

	time.Sleep(lcdPause)

	for {
		if !l.running {
			return
		}

		if text != l.text && len(l.text) > lcdWidth {
			l.send(l.data())
		}
		l.scroll()

		time.Sleep(lcdTick)
	}
}

func (l *lcdLine) start() {
	if l.running {
		return
	}

	go l.loop()
}

func (l *lcdLine) stop() {
	l.running = false
}
