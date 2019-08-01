package oc

import (
	// stdlib
	"net"
	"net/http"

	// external
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"go.opencensus.io/zpages"

	// project
	"github.com/basvanbeek/opencensus-gokit-example/shared/network"
)

// ZPages handling setup
func ZPages(g run.Group, logger log.Logger) {
	var (
		bindIP, _   = network.HostIP()
		listener, _ = net.Listen("tcp", bindIP+":0") // dynamic port assignment
		addr        = listener.Addr().String()
	)

	g.Add(func() error {
		level.Info(logger).Log("msg", "zpages started", "addr", "http://"+addr)
		return http.Serve(listener, zpages.Handler)
	}, func(error) {
		listener.Close()
	})
}
