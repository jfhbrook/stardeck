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
	id      uint
	text    string
}

type notificationManager struct {
	line     *lcdLine
	running  bool
	commands chan notificationCommand
	length   int
}

func newNotificationManager(line *lcdLine) *notificationManager {
	return &notificationManager{
		line:     line,
		running:  false,
		commands: make(chan notificationCommand),
		length:   0,
	}
}

func (m *notificationManager) timeout() time.Duration {
	return time.Duration(
		viper.GetFloat64("notifications.timeout") * float64(time.Second),
	)
}

func (m *notificationManager) minWait() time.Duration {
	return time.Duration(
		viper.GetFloat64("notifications.min_wait") * float64(time.Second),
	)
}

func (m *notificationManager) maxQueueLength() int {
	return viper.GetInt("notifications.max_queue_length")
}

func notificationText(info *notifications.NotificationInfo) string {
	text := info.Summary

	if len(info.Body) > 0 {
		text += (" - " + info.Body)
	}

	return text
}

func (m *notificationManager) update(info *notifications.NotificationInfo) {
	m.length += 1
	text := notificationText(info)

	m.commands <- notificationCommand{displayNotificationCommand, 0, text}
}

func (m *notificationManager) display(id uint, text string) {
	log.Trace().
		Uint("id", id).
		Str("text", text).
		Msg("Displaying notification")

	m.line.update(text)

	// Clear the notification after a timeout
	go func() {
		time.Sleep(m.timeout())
		m.commands <- notificationCommand{clearNotificationCommand, id, ""}
	}()
}

func (m *notificationManager) wait(id uint) {
	minWait := m.minWait()

	log.Trace().
		Uint("id", id).
		Float64("min_wait", minWait.Seconds()).
		Msg("Waiting before displaying next notification")
	time.Sleep(m.minWait())
}

func (m *notificationManager) clear() {
	log.Trace().Msg("Clearing notification")
	m.line.update("")
}

func (m *notificationManager) loop() {
	var id uint = 0
	m.running = true

	for {
		log.Trace().Msg("Waiting for command")
		command := <-m.commands
		log.Trace().Any("command", command).Msg("Received command")

		switch command.command {
		case displayNotificationCommand:
			m.length -= 1
			if m.length >= m.maxQueueLength() {
				continue
			}
			id++
			m.display(id, command.text)
			m.wait(id)
		case clearNotificationCommand:
			// Clear the notification if it's the current notification
			if command.id == id {
				m.clear()
			}
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
	m.commands <- notificationCommand{stopNotifyingCommand, 0, ""}
}
