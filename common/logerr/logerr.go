package logerr

import (
	"fmt"

	"github.com/charmbracelet/log"
)

func Error(err error, keyvals ...interface{}) error {
	log.Error(err, keyvals...)

	return err
}

func Errorf(format string, err error, keyvals ...interface{}) error {
	err = fmt.Errorf(format, err)

	log.Error(err, keyvals...)

	return err
}
