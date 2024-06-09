package logerr

import "github.com/charmbracelet/log"

func Error(err error, keyvals ...interface{}) error {
	log.Error(err, keyvals...)

	return err
}
