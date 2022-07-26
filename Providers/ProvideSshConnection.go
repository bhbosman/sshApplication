package Providers

import (
	"fmt"
	common2 "github.com/bhbosman/goCommsSshListener/common"
	"github.com/bhbosman/goCommsSshListener/netListener"
	"github.com/bhbosman/gocommon/messages"
	"github.com/bhbosman/gocomms/common"
	"github.com/cskr/pubsub"
	"go.uber.org/fx"
	"net/url"
)

func SshConnectionManager(
	maxConnections int,
	urlAsText string,
) fx.Option {
	sshConnectionUrl, err := url.Parse(urlAsText)
	if err != nil {
		return fx.Error(err)
	}
	Name := fmt.Sprintf("SSHConnection:%v", sshConnectionUrl.Port())
	return fx.Provide(
		fx.Annotated{
			Group: "Apps",
			Target: func(params struct {
				fx.In
				PubSub             *pubsub.PubSub `name:"Application"`
				NetAppFuncInParams common.NetAppFuncInParams
				SshChannelSettings common2.ISshChannelSettings `optional:"true"`
			}) (messages.CreateAppCallback, error) {
				f := netListener.NewSshListenApp(
					Name,
					Name,
					false,
					nil,
					sshConnectionUrl,
					params.SshChannelSettings,
					common.MaxConnectionsSetting(maxConnections),
				)
				return f(params.NetAppFuncInParams), nil
			},
		},
	)
}
