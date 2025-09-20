package service

import (
	"bytes"

	"github.com/godbus/dbus/v5"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
)

type crystalfontzSender func(text string)

func makeCrystalfontzSender(systemConn *dbus.Conn) crystalfontzSender {
	lcd := crystalfontz.NewClient(systemConn)

	var row byte = 0
	var column byte = 0

	return func(text string) {
		data := []byte(text)
		if len(data) > 16 {
			data = data[0:16]
		} else {
			data = append(data, bytes.Repeat([]byte{' '}, 16 - len(data))...)
		}

		log.Debug().
			Int("row", int(row)).
			Int("column", int(column)).
			Str("data", string(data)).
			Msg("Sending data to LCD")

		err := lcd.SendData(row, column, data, -1.0, -1)

		if err != nil {
			log.Error().
				Int("row", int(row)).
				Int("column", int(column)).
				Str("data", string(data)).
				Err(errors.Wrap(err, err.Error())).
				Msg("Error while sending data to LCD")
		}
	}
}
