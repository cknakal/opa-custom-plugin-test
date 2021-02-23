package opa

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/open-policy-agent/opa/plugins"
	"github.com/open-policy-agent/opa/plugins/logs"
	"github.com/open-policy-agent/opa/util"
)

// PluginName --
const PluginName = "persistent_decision_logger"

// Config --
type Config struct {
	Stderr bool `json:"stderr"`
}

// Factory --
type Factory struct{}

// PrintlnLogger --
type PrintlnLogger struct {
	manager *plugins.Manager
	config  Config
	mtx     sync.Mutex
}

// New --
func (Factory) New(m *plugins.Manager, config interface{}) plugins.Plugin {

	m.UpdatePluginStatus(PluginName, &plugins.Status{State: plugins.StateNotReady})

	return &PrintlnLogger{
		manager: m,
		config:  config.(Config),
	}
}

// Validate --
func (Factory) Validate(_ *plugins.Manager, config []byte) (interface{}, error) {
	parsedConfig := Config{}
	return parsedConfig, util.Unmarshal(config, &parsedConfig)
}

// Start --
func (p *PrintlnLogger) Start(ctx context.Context) error {
	return nil
}

// Stop --
func (p *PrintlnLogger) Stop(ctx context.Context) {
}

// Reconfigure --
func (p *PrintlnLogger) Reconfigure(ctx context.Context, config interface{}) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.config = config.(Config)
}

// Log --
func (p *PrintlnLogger) Log(ctx context.Context, event logs.EventV1) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	w := os.Stdout
	if p.config.Stderr {
		w = os.Stderr
	}
	_, err := fmt.Fprintln(w, "You did it!")
	if err != nil {
		p.manager.UpdatePluginStatus(PluginName, &plugins.Status{State: plugins.StateErr})
	}
	return nil
}
