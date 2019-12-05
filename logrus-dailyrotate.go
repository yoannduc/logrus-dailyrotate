package logrusdailyrotate

import (
	"github.com/sirupsen/logrus"
	"github.com/yoannduc/dailyrotate"
)

type Hook struct {
	rotateWriter *dailyrotate.RotateWriter
	formatter    logrus.Formatter
	minLevel     logrus.Level
}

func New(
	p string,
	ma int,
	f logrus.Formatter,
	ml logrus.Level,
) (*Hook, error) {
	rw, err := dailyrotate.New(p, ma)
	if err != nil {
		return nil, err
	}

	return &Hook{
		rotateWriter: rw,
		formatter:    f,
		minLevel:     ml,
	}, nil
}

func NewWithDefaults() (*Hook, error) {
	rw, err := dailyrotate.NewWithDefaults()
	if err != nil {
		return nil, err
	}

	return &Hook{
		rotateWriter: rw,
		formatter:    &logrus.TextFormatter{},
		minLevel:     logrus.InfoLevel,
	}, nil
}

func (h *Hook) Fire(e *logrus.Entry) error {
	m, err := h.formatter.Format(e)
	if err != nil {
		return err
	}

	if _, err = h.rotateWriter.RotateWrite([]byte(m)); err != nil {
		return err
	}

	return nil
}

func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.minLevel+1]
}
