package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/crystalfontz"
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

func (l *lcdLine) width() int {
	return viper.GetInt("crystalfontz.width")
}

func (l *lcdLine) pause() time.Duration {
	return time.Duration(
		viper.GetFloat64("crystalfontz.pause") * float64(time.Second),
	)
}

func (l *lcdLine) tick() time.Duration {
	return time.Duration(
		viper.GetFloat64("crystalfontz.tick") * float64(time.Second),
	)
}

func (l *lcdLine) update(text string) {
	width := l.width()
	if text == "" {
		text = l.defaultText
	}

	if len(text) > width {
		// Allow marquee text to scroll out of the screen
		text += strings.Repeat(" ", width)
	} else {
		// Pad the text to at least the LCD's width
		text += strings.Repeat(" ", max(width-len(text), 0))
	}

	if l.text != text {
		l.shift = 0
	}

	l.text = text
}

func (l *lcdLine) data() []byte {
	width := l.width()
	left := l.text[l.shift:]
	right := l.text[0:l.shift]
	data := []byte(left + right)
	// In case characters are multi-byte
	return data[0:width]
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
	width := l.width()
	// Only scroll text if it's wider than the LCD
	if len(l.text) > width {
		l.shift += 1
	}
	if l.shift >= len(l.text) {
		l.shift = 0
	}
}

func (l *lcdLine) loop() {
	l.running = true

	l.send(l.data())
	l.scroll()

	if !l.running {
		return
	}

	time.Sleep(l.pause())

	for {
		if !l.running {
			return
		}

		// TODO: Debounce this signal
		l.send(l.data())
		l.scroll()

		time.Sleep(l.tick())
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
