package main

import (
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	example "github.com/plugin-poc/commons"
)

var (
	PluginName = "hello"
)

// Here is a real implementation of Greeter
type PluginHello struct {
	logger hclog.Logger
}

func (g *PluginHello) Init() string {
	g.logger.Debug(PluginName + ".INIT() message from GreeterHello.Greet")
	return PluginName + "Init stuff!"
}

func (g *PluginHello) UI() string {
	g.logger.Debug(PluginName + "UI() message from GreeterHello.Greet")
	return PluginName + "UI Retrun UI!"
}

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: PluginName,
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Level:      hclog.Trace,
		Output:     os.Stderr,
		JSONFormat: true,
	})

	greeter := &PluginHello{
		logger: logger,
	}
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		PluginName: &example.PluginPlugin{Impl: greeter},
	}

	message := "message from plugin " + PluginName
	logger.Debug(message, "foo", "bar")

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
