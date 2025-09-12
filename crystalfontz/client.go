package crystalfontz

import (
	"github.com/pkg/errors"
	"github.com/godbus/dbus/v5"
)

type KeyActivity = byte

const (
	KeyUpPress      KeyActivity = 1
	KeyDownPress                = 2
	KeyLeftPress                = 3
	KeyRightPress               = 4
	KeyEnterPress               = 5
	KeyExitPress                = 6
	KeyUpRelease                = 7
	KeyDownRelease              = 8
	KeyLeftRelease              = 9
	KeyRightRelease             = 10
	KeyEnterRelease             = 11
	KeyExitRelease              = 12
)

func AddKeyActivityReportMatchSignal(conn *dbus.Conn) error {
	err := conn.AddMatchSignal(
		dbus.WithMatchObjectPath("/"),
		dbus.WithMatchInterface("org.jfhbrook.crystalfontz.KeyActivityReports"),
	)

	if err != nil {
		return errors.Wrap(err, "Failed to match signal for key activity reports")
	}

	return nil
}

type Crystalfontz struct {
	object               dbus.BusObject
	DefaultContrast      float64
	DefaultLcdBrightness float64
}

func NewCrystalfontz(conn *dbus.Conn) *Crystalfontz {
	obj := conn.Object("org.jfhbrook.crystalfontz", "/")
	lcd := Crystalfontz{object: obj, DefaultContrast: 0.5, DefaultLcdBrightness: 0.1}
	return &lcd
}

func (lcd *Crystalfontz) SetContrast(contrast float64, timeout float64, retryTimes int64) error {
	call := lcd.object.Call("org.jfhbrook.crystalfontz.SetContrast", 0, contrast, timeout, retryTimes)

	if call.Err != nil {
		return call.Err
	}

	return nil
}

func (lcd *Crystalfontz) SetBacklight(lcdBrightness float64, keypadBrightness float64, timeout float64, retryTimes int64) error {
	call := lcd.object.Call("org.jfhbrook.crystalfontz.SetBacklight", 0, lcdBrightness, keypadBrightness, timeout, retryTimes)

	if call.Err != nil {
		return errors.Wrap(call.Err, "Failed to set backlight")
	}

	return nil
}

func (lcd *Crystalfontz) ClearScreen(timeout float64, retryTimes int64) error {
	call := lcd.object.Call("org.jfhbrook.crystalfontz.ClearScreen", 0, timeout, retryTimes)

	if call.Err != nil {
		return errors.Wrap(call.Err, "Failed to clear screen")
	}

	return nil
}

func (lcd *Crystalfontz) SendData(row byte, column byte, data []byte, timeout float64, retryTimes int64) error {
	call := lcd.object.Call("org.jfhbrook.crystalfontz.SendData", 0, row, column, data, timeout, retryTimes)

	if call.Err != nil {
		return errors.Wrap(call.Err, "Failed to send data to LCD")
	}

	return nil
}

func (lcd *Crystalfontz) Splash() error {
	if err := lcd.SendData(0, 0, []byte("YES THIS"), -1.0, -1); err != nil {
		return errors.Wrap(err, "Failed to display splash screen")
	}

	if err := lcd.SendData(1, 0, []byte("IS STARDECK"), -1.0, -1); err != nil {
		return errors.Wrap(err, "Failed to display splash screen")
	}

	return nil
}

func (lcd *Crystalfontz) Reset() error {
	if err := lcd.SetContrast(lcd.DefaultContrast, -1.0, -1); err != nil {
		return errors.Wrap(err, "Failed to reset LCD")
	}
	if err := lcd.SetBacklight(lcd.DefaultLcdBrightness, -1.0, -1.0, -1); err != nil {
		return errors.Wrap(err, "Failed to reset LCD")
	}
	if err := lcd.ClearScreen(-1.0, -1); err != nil {
		return errors.Wrap(err, "Failed to reset LCD")
	}

	return nil
}
