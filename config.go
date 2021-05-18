package main

import (
	"errors"
)

type Config struct {
	AggerIntervalSec	int `toml:"aggrIntervalSec"`
}