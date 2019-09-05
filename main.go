package main

import (
	"time"
	"github.com/heptiolabs/healthcheck"
)

func main() {
	healthcheck.DatabasePingCheck(db, 5*time.Second)
}
