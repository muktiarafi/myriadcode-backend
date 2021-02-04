package logs

import (
	"bytes"
	"regexp"
	"testing"
)

func TestNewLogger(t *testing.T) {
	regex, _ := regexp.Compile(`([A-Z])\w+`)

	t.Run("info log", func(t *testing.T) {
		myBuffer := new(bytes.Buffer)
		defaultWriter = myBuffer

		logger := NewLogger()
		logger.InfoLog.Println("anjay")

		if len(myBuffer.String()) == 0 {
			t.Error("should have a string")
		}

		want := "INFO"
		got := regex.FindString(myBuffer.String())

		if want != got {
			t.Errorf("Should print %s got %s instead", want, got)
		}
	})

	t.Run("warning log", func(t *testing.T) {
		myBuffer := new(bytes.Buffer)
		defaultWriter = myBuffer

		logger := NewLogger()
		logger.WarningLog.Println("anjay")

		if len(myBuffer.String()) == 0 {
			t.Error("should have a string")
		}

		want := "WARNING"
		got := regex.FindString(myBuffer.String())

		if want != got {
			t.Errorf("Should print %s got %s instead", want, got)
		}
	})

	t.Run("error log", func(t *testing.T) {
		myBuffer := new(bytes.Buffer)
		defaultWriter = myBuffer

		logger := NewLogger()
		logger.ErrorLog.Println("anjay")

		if len(myBuffer.String()) == 0 {
			t.Error("should have a string")
		}

		want := "ERROR"
		got := regex.FindString(myBuffer.String())

		if want != got {
			t.Errorf("Should print %s got %s instead", want, got)
		}
	})
}
