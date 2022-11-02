package webapi

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pangpanglabs/echoswagger/v2"
	"go.uber.org/dig"

	"github.com/iotaledger/wasp/packages/webapi"
	v1 "github.com/iotaledger/wasp/packages/webapi/v1"
	v2 "github.com/iotaledger/wasp/packages/webapi/v2"

	"github.com/iotaledger/hive.go/core/app"
	"github.com/iotaledger/hive.go/core/app/pkg/shutdown"
	"github.com/iotaledger/inx-app/httpserver"
	"github.com/iotaledger/wasp/packages/chains"
	"github.com/iotaledger/wasp/packages/dkg"
	"github.com/iotaledger/wasp/packages/metrics"
	"github.com/iotaledger/wasp/packages/parameters"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/users"
	"github.com/iotaledger/wasp/packages/wal"
	"github.com/iotaledger/wasp/packages/wasp"
)

func init() {
	Plugin = &app.Plugin{
		Component: &app.Component{
			Name:           "WebAPI",
			DepsFunc:       func(cDeps dependencies) { deps = cDeps },
			Params:         params,
			InitConfigPars: initConfigPars,
			Provide:        provide,
			Run:            run,
		},
		IsEnabled: func() bool {
			return ParamsWebAPI.Enabled
		},
	}
}

var (
	Plugin *app.Plugin
	deps   dependencies
)

type dependencies struct {
	dig.In

	EchoSwagger echoswagger.ApiRoot `name:"webapiServer"`
}

func initConfigPars(c *dig.Container) error {
	type cfgResult struct {
		dig.Out
		WebAPIBindAddress string `name:"webAPIBindAddress"`
	}

	if err := c.Provide(func() cfgResult {
		return cfgResult{
			WebAPIBindAddress: ParamsWebAPI.BindAddress,
		}
	}); err != nil {
		Plugin.LogPanic(err)
	}

	return nil
}

func provide(c *dig.Container) error {
	type webapiServerDeps struct {
		dig.In

		ShutdownHandler              *shutdown.ShutdownHandler
		WAL                          *wal.WAL
		APICacheTTL                  time.Duration `name:"apiCacheTTL"`
		PublisherPort                int           `name:"publisherPort"`
		Chains                       *chains.Chains
		Metrics                      *metrics.Metrics `optional:"true"`
		DefaultRegistry              registry.Registry
		DefaultNetworkProvider       peering.NetworkProvider       `name:"defaultNetworkProvider"`
		DefaultTrustedNetworkManager peering.TrustedNetworkManager `name:"defaultTrustedNetworkManager"`
		DefaultNode                  *dkg.Node                     `name:"defaultNode"`
		UserManager                  *users.UserManager
	}

	type webapiServerResult struct {
		dig.Out

		EchoSwagger echoswagger.ApiRoot `name:"webapiServer"`
	}

	if err := c.Provide(func(deps webapiServerDeps) webapiServerResult {
		e := httpserver.NewEcho(
			Plugin.Logger(),
			nil,
			ParamsWebAPI.DebugRequestLoggerEnabled,
		)

		e.HideBanner = true
		e.HidePort = true
		e.HTTPErrorHandler = webapi.CompatibilityHTTPErrorHandler

		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
			AllowMethods: []string{"*"},
		}))

		echoSwagger := echoswagger.New(e, "/doc", &echoswagger.Info{
			Title:       "Wasp API",
			Description: "REST API for the Wasp node",
			Version:     wasp.Version,
		})

		echoSwagger.SetRequestContentType(echo.MIMEApplicationJSON)
		echoSwagger.SetResponseContentType(echo.MIMEApplicationJSON)

		v1.Init(
			Plugin.App().NewLogger("WebAPI"),
			echoSwagger,
			deps.DefaultNetworkProvider,
			deps.DefaultTrustedNetworkManager,
			deps.UserManager,
			func() registry.Registry {
				return deps.DefaultRegistry
			},
			func() *chains.Chains {
				return deps.Chains
			},
			func() *dkg.Node {
				return deps.DefaultNode
			},
			func() {
				deps.ShutdownHandler.SelfShutdown("wasp was shutdown via API", false)
			},
			deps.Metrics,
			deps.WAL,
			ParamsWebAPI.Auth,
			ParamsWebAPI.NodeOwnerAddresses,
			deps.APICacheTTL,
			deps.PublisherPort,
		)

		v2.Init(
			Plugin.App().NewLogger("WebAPI/v2"),
			echoSwagger,
			Plugin.App().Config(),
			func() *chains.Chains {
				return deps.Chains
			},
			func() *dkg.Node {
				return deps.DefaultNode
			},
			deps.Metrics,
			deps.DefaultNetworkProvider,
			func() registry.Registry {
				return deps.DefaultRegistry
			},
			deps.ShutdownHandler,
			deps.DefaultTrustedNetworkManager,
			deps.WAL,
		)

		return webapiServerResult{
			EchoSwagger: echoSwagger,
		}
	}); err != nil {
		Plugin.LogPanic(err)
	}

	return nil
}

func run() error {
	Plugin.LogInfof("Starting %s server ...", Plugin.Name)
	if err := Plugin.Daemon().BackgroundWorker(Plugin.Name, func(ctx context.Context) {
		Plugin.LogInfof("Starting %s server ... done", Plugin.Name)

		go func() {
			Plugin.LogInfof("You can now access the WebAPI using: http://%s", ParamsWebAPI.BindAddress)
			if err := deps.EchoSwagger.Echo().Start(ParamsWebAPI.BindAddress); err != nil && !errors.Is(err, http.ErrServerClosed) {
				Plugin.LogWarnf("Stopped %s server due to an error (%s)", Plugin.Name, err)
			}
		}()

		<-ctx.Done()
		Plugin.LogInfof("Stopping %s server ...", Plugin.Name)

		shutdownCtx, shutdownCtxCancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer shutdownCtxCancel()

		//nolint:contextcheck // false positive
		if err := deps.EchoSwagger.Echo().Shutdown(shutdownCtx); err != nil {
			Plugin.LogWarn(err)
		}

		Plugin.LogInfof("Stopping %s server ... done", Plugin.Name)
	}, parameters.PriorityWebAPI); err != nil {
		Plugin.LogPanicf("failed to start worker: %s", err)
	}

	return nil
}
