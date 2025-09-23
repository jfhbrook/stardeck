package service

import (
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/jfhbrook/stardeck/notifications"
)

type notificationCommandType int

const (
	displayNotificationCommand notificationCommandType = 0
	clearNotificationCommand                           = 1
	stopNotifyingCommand                               = 2
)

type notificationCommand struct {
	command notificationCommandType
	text    string
}

type notificationManager struct {
	commands chan notificationCommand
	timeout  time.Duration
	running  bool
	line     *lcdLine
}

func newNotificationManager(line *lcdLine) *notificationManager {
	timeout := time.Duration(
		viper.GetFloat64("notifications.timeout") * float64(time.Second),
	)

	return &notificationManager{
		line:     line,
		timeout:  timeout,
		running:  false,
		commands: make(chan notificationCommand),
	}
}

func notificationText(info *notifications.NotificationInfo) string {
	text := info.Summary

	if len(info.Body) > 0 {
		text += (" - " + info.Body)
	}

	return text
}

func (m *notificationManager) update(info *notifications.NotificationInfo) {
	text := notificationText(info)

	m.commands <- notificationCommand{
		command: displayNotificationCommand,
		text:    text,
	}
}

func (m *notificationManager) display(text string) {
	log.Trace().Str("text", text).Msg("Displaying notification")

	m.line.update(text)

	go func() {
		time.Sleep(m.timeout)
		m.commands <- notificationCommand{clearNotificationCommand, ""}
	}()
}

func (m *notificationManager) clear() {
	log.Trace().Msg("Clearing notification")
	m.line.update("")
}

func (m *notificationManager) loop() {
	m.running = true

	for {
		log.Trace().Msg("Waiting for command")
		command := <-m.commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.command {
		case displayNotificationCommand:
			m.display(command.text)
		case clearNotificationCommand:
			m.clear()
		case stopNotifyingCommand:
			m.running = false
			return
		}
	}
}

func (m *notificationManager) start() {
	if m.running {
		return
	}

	go m.loop()
}

func (m *notificationManager) stop() {
	m.commands <- notificationCommand{stopNotifyingCommand, ""}
}
