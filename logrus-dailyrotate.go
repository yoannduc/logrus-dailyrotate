// Copyright 2019 Yoann Duc. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

// Package logrusdailyrotate is a daily rotating log file hook for logrus.
//
// Makes use of github.com/yoannduc/dailyrotate
// (documentation: https://godoc.org/github.com/yoannduc/dailyrotate)
//
// A trivial example is:
//
//  package main
//
//  import (
//      "github.com/sirupsen/logrus"
//      logrusdailyrotate "github.com/yoannduc/logrus-dailyrotate"
//  )
//
//  func main() {
//      log := logrus.New()
//
//      hook, err := logrusdailyrotate.New("", 5, logrus.InfoLevel, &logrus.TextFormatter{})
//      if err != nil {
//          // Handle error your way
//      }
//
//      log.Hooks.Add(hook)
//
//      log.Info("Ready to go ! :)")
//  }
//
package logrusdailyrotate

import (
	"github.com/sirupsen/logrus"
	"github.com/yoannduc/dailyrotate"
)

const (
	// DefaultLogPath is the default path that for the log file when using
	// NewWithDefaults.
	DefaultLogPath = dailyrotate.DefaultFilePath
	// DefaultMaxAge is the default number of days before cleaning
	// old log files when using NewWithDefaults.
	DefaultMaxAge = dailyrotate.DefaultMaxAge
	// DefaultMinLevel is the default minimum lever for which to log to
	// file or not when using NewWithDefaults.
	DefaultMinLevel = logrus.InfoLevel
)

// &logrus.TextFormatter{} is not a constant,
// therefore it can't be declared in const block.

// DefaultFormatter is the default text formatter that will be used to
// format text before sending to logrus when using NewWithDefaults.
var DefaultFormatter = &logrus.TextFormatter{}

// Hook is the logrus hook to log to the rotating log file. Make use
// internally of github.com/yoannduc/dailyrotate to rotate log file.
//
// Hook satisfies the logrus.Hook interface.
type Hook struct {
	rotateWriter *dailyrotate.RotateWriter
	formatter    logrus.Formatter
	minLevel     logrus.Level
}

// New instanciates a new *Hook with formatter and minimum log level inputed.
// It will internally instanciate a *dailyrotate.RotateWriter with
// inputed params path & maxAge.
//
// If you wonder how *dailyrotate.RotateWriter behaves, please read
// documentation at https://godoc.org/github.com/yoannduc/dailyrotate.
func New(
	path string,
	maxAge int,
	formatter logrus.Formatter,
	minLevel logrus.Level,
) (*Hook, error) {
	rw, err := dailyrotate.New(path, maxAge)
	if err != nil {
		return nil, err
	}

	return &Hook{
		rotateWriter: rw,
		formatter:    formatter,
		minLevel:     minLevel,
	}, nil
}

// NewWithDefaults instanciates a new *Hook with default params (see above).
func NewWithDefaults() (*Hook, error) {
	return New(DefaultLogPath, DefaultMaxAge, DefaultFormatter, DefaultMinLevel)
}

// Fire writes to the rotating log file after formatting input with formatter.
// It gets fired by logrus directly after been added to your logrus logger
// (see https://github.com/sirupsen/logrus#hooks).
//
// Fire satisfies the logrus.Hook interface.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	m, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}

	if _, err = hook.rotateWriter.RotateWrite([]byte(m)); err != nil {
		return err
	}

	return nil
}

// Levels returns the slice of logrus.Level for which the hook will write
// to the rotating log file.
//
// Levels satisfies the logrus.Hook interface.
func (hook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels[:hook.minLevel+1]
}
