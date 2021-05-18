package main

import (
	"errors"
)

type Config struct {
	AggrIntervalSec     int `toml:"aggrIntervalSec"`
	Addr                string
	Source              string
	SendBatchSize       int
	SendMaxIntervaSec   int
	SendMaxRetryCount   int
	SendTimeoutSec      int
	MetricGCIntervalSec int
	MetricGCSec         int
}

func (c *Config) ConfigValidate() error {
	if c == nil {
		return errors.New("config of tracer can't be nil")
	}

	if c.Addr == "" {
		return errors.New("addr of tracer can't be nil")
	}

	if c.AggrIntervalSec == 0 {
		c.AggrIntervalSec = 60
	}

	if c.Source == "" {
		c.Source = "init etracer"
	}

	if c.SendBatchSize == 0 {
		c.SendBatchSize = 1000
	}

	if c.SendMaxIntervaSec == 0 {
		c.SendMaxIntervaSec = 10
	}

	if c.SendMaxRetryCount == 0 {
		c.SendMaxRetryCount = 3
	}

	if c.SendTimeoutSec == 0 {
		c.SendTimeoutSec = 5
	}

	if c.MetricGCSec == 0 {
		c.MetricGCSec = 180 // 3min by default
	}

	if c.MetricGCIntervalSec == 0 {
		c.MetricGCIntervalSec = 60 // 1min by default
	}

	return nil
}
