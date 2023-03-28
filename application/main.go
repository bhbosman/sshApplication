package main

import (
	"github.com/bhbosman/goCommsSshListener/common"
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/goFxApp"
	"github.com/bhbosman/gocommon/services/interfaces"
	"github.com/bhbosman/sshApplication/Providers"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
)

func main() {
	app := goFxApp.NewFxMainApplicationServices(
		"SSH App",
		false,
		Providers.SshConnectionManager(
			512,
			"tcp4://localhost:8888"),
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						PubSub                   *pubsub.PubSub `name:"Application"`
						ConnectionHelper         goConnectionManager.IHelper
						ConnectionManagerIHelper goConnectionManager.IHelper
						UniqueReferenceService   interfaces.IUniqueReferenceService
					},
				) (common.ISshChannelSessionSettings, error) {
					return Providers.NewSshCreateChannelProcess(
						params.PubSub,
						params.ConnectionHelper,
						params.ConnectionManagerIHelper,
						params.UniqueReferenceService,
					), nil
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func(
					params struct {
						fx.In
						ISshChannelSessionSettings common.ISshChannelSessionSettings `optional:"true"`
					},
				) (common.ISshChannelSettings, error) {
					return Providers.NewSshChannelSettings(params.ISshChannelSessionSettings), nil
				},
			},
		),
	)
	if app.FxApp.Err() != nil {
		println(app.FxApp.Err().Error())
		return
	}
	app.RunTerminalApp()

}
