package harvester

import (
	"time"
)

// Harvester struct
type Harvester struct {
	*Config
}

// Config Harvester config
type Config struct {
	Concurrency   int
	MaxAgentConns int
	User          string
	IgnoreHostKey bool
	Timeout       time.Duration

	Debug      bool
	IdentFiles []string
	// ciphers
	Kex     []string
	Ciphers []string
	Macs    []string
}

// Init Harvester
func (p *Harvester) Init() {
}

// Run is main process
func (p *Harvester) Run() int {
	return 0
}
