package example

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// Plugin is the interface that we're exposing as a plugin.
type Plugin interface {
	Init() string
	UI() string
}

// Here is an implementation that talks over RPC
type PluginRPC struct{ client *rpc.Client }

func (g *PluginRPC) Init() string {
	var resp string
	err := g.client.Call("Plugin.Init", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

func (g *PluginRPC) UI() string {
	var resp string
	err := g.client.Call("Plugin.UI", new(interface{}), &resp)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}

	return resp
}

// Here is the RPC server that GreeterRPC talks to, conforming to
// the requirements of net/rpc
type PluginRPCServer struct {
	// This is the real implementation
	Impl Plugin
}

func (s *PluginRPCServer) Init(args interface{}, resp *string) error {
	*resp = s.Impl.Init()
	return nil
}

func (s *PluginRPCServer) UI(args interface{}, resp *string) error {
	*resp = s.Impl.UI()
	return nil
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a GreeterRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return GreeterRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type PluginPlugin struct {
	// Impl Injection
	Impl Plugin
}

func (p *PluginPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &PluginRPCServer{Impl: p.Impl}, nil
}

func (PluginPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &PluginRPC{client: c}, nil
}
