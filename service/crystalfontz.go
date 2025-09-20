package service

import (
	"github.com/godbus/dbus/v5"
	"github.com/rs/zerolog/log"

	"github.com/jfhbrook/stardeck/crystalfontz"
)

func crystalfontzSender(systemConn *dbus.Conn) func(text string) {
	lcd := crystalfontz.NewClient(systemConn)

	var row byte = 0
	var column byte = 0

	return func(text string) {
		data := []byte(text)
		if len(data) > 16 {
			data = data[0:16]
		}

		err := lcd.SendData(row, column, data, -1.0, -1)

		if err != nil {
			log.Warn().
				Int("row", int(row)).
				Int("column", int(column)).
				Err(err).
				Msg("Error while sending data to LCD")
		}
	}
}
