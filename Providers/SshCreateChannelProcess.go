package Providers

import (
	"context"
	"fmt"
	"github.com/bhbosman/goCommsSshListener/common"
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
	"github.com/cskr/pubsub"
	"github.com/reactivex/rxgo/v2"
	"go.uber.org/zap"
)

type SshChannelSettings struct {
	m map[string]common.ISettings
}

func NewSshChannelSettings(settings ...common.ISettings) *SshChannelSettings {
	result := &SshChannelSettings{
		m: make(map[string]common.ISettings),
	}
	for _, setting := range settings {
		if setting == nil {
			continue
		}
		result.m[setting.Name()] = setting
	}
	return result

}

func (self *SshChannelSettings) GetSettingsFor(sessionType string) (common.ISettings, error) {
	if v, ok := self.m[sessionType]; ok {
		return v, nil
	}
	return nil, fmt.Errorf("not found")
}

type SshCreateChannelProcess struct {
	PubSub                   *pubsub.PubSub
	ConnectionHelper         goConnectionManager.IHelper
	ConnectionManagerIHelper goConnectionManager.IHelper
	UniqueReferenceService   interfaces.IUniqueReferenceService
}

func (self *SshCreateChannelProcess) Name() string {
	return "session"
}

func NewSshCreateChannelProcess(
	PubSub *pubsub.PubSub,
	ConnectionHelper goConnectionManager.IHelper,
	ConnectionManagerIHelper goConnectionManager.IHelper,
	UniqueReferenceService interfaces.IUniqueReferenceService,
) *SshCreateChannelProcess {
	return &SshCreateChannelProcess{
		PubSub:                   PubSub,
		ConnectionHelper:         ConnectionHelper,
		ConnectionManagerIHelper: ConnectionManagerIHelper,
		UniqueReferenceService:   UniqueReferenceService,
	}
}

func (self *SshCreateChannelProcess) UseDefault() (common.SshBuildInChannelProcess, error) {
	return common.UseCustomChannelProcess, nil
}

func (self *SshCreateChannelProcess) CreateChannelProcess(
	sshChannel common.IChannel,
	parentCtx context.Context,
	parentCancelFunc context.CancelFunc,
	onSend rxgo.NextFunc,
	logger *zap.Logger,
	GoFunctionCounter GoFunctionCounter.IService,
) (common.IChannelProcess, error) {
	return NewSshConnectionUi(
		sshChannel,
		parentCtx,
		parentCancelFunc,
		onSend,
		logger,
		self.PubSub,
		self.ConnectionManagerIHelper,
		self.UniqueReferenceService,
		GoFunctionCounter,
	)
}
