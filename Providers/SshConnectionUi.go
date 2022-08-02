package Providers

import (
	"context"
	"github.com/bhbosman/goCommsSshListener/common"
	"github.com/bhbosman/goConnectionManager"
	"github.com/bhbosman/goUi/UiService"
	"github.com/bhbosman/goUi/UiSlides/cmSlide"
	"github.com/bhbosman/goUi/UiSlides/intoductionSlide"
	"github.com/bhbosman/gocommon/GoFunctionCounter"
	"github.com/bhbosman/gocommon/Services/interfaces"
	fx2 "github.com/bhbosman/gocommon/fx"
	"github.com/bhbosman/gocommon/fx/PubSub"
	"github.com/bhbosman/gocommon/services/Providers"
	"github.com/cskr/pubsub"
	"github.com/gdamore/tcell/v2"
	"github.com/reactivex/rxgo/v2"
	"github.com/rivo/tview"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"time"
)

type SshConnectionUi struct {
	common.BaseChannelProcess
	tty                      ITty
	logger                   *zap.Logger
	PubSub                   *pubsub.PubSub
	ConnectionManagerIHelper goConnectionManager.IHelper
	UniqueReferenceService   interfaces.IUniqueReferenceService
	GoFunctionCounter        GoFunctionCounter.IService
}

func (self *SshConnectionUi) Close() error {
	var err error
	err = self.tty.Close()
	err = multierr.Append(err, self.BaseChannelProcess.Close())
	return err
}

func (self *SshConnectionUi) SetSize(cols int, rows int) error {
	return self.tty.SetSize(cols, rows)
}

func (self *SshConnectionUi) RunHandler() error {
	fxAppOptions := fx2.NewFxApplicationOptions(
		time.Hour,
		time.Hour,
		fx.Provide(
			fx.Annotated{
				Target: func() (tcell.Screen, error) {
					screen, err := tcell.NewTerminfoScreenFromTtyTerminfo(self.tty, self.Terminfo)
					if err != nil {
						return nil, err
					}
					err = screen.Init()
					if err != nil {
						return nil, err
					}
					return screen, nil
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func() goConnectionManager.IHelper {
					return self.ConnectionManagerIHelper
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Target: func() *zap.Logger {
					return self.logger
				},
			},
		),
		Providers.ProvideUniqueReferenceServiceInstance(self.UniqueReferenceService),
		PubSub.ProvidePubSubInstance("Application", self.PubSub),
		fx.Provide(
			fx.Annotated{
				Target: func() GoFunctionCounter.IService {
					return self.GoFunctionCounter
				},
			},
		),
		fx.Provide(
			fx.Annotated{
				Name: "Application",
				Target: func() (context.Context, error) {
					return self.CancelCtx, nil
				},
			},
		),
		UiService.ProvideTerminalApplication(),
		intoductionSlide.ProvideCoverSlide(),
		cmSlide.ProvideConnectionManagerSlide(),
		cmSlide.InvokeConnectionManagerSlide(),
		UiService.InvokeTerminalApplication(),
	)
	var app *tview.Application
	var screen tcell.Screen
	fxApp := fx.New(
		fxAppOptions,
		fx.Populate(&app, &screen),
	)
	if fxApp.Err() != nil {
		self.logger.Error("Error in create SSH App", zap.Error(fxApp.Err()))
		return fxApp.Err()
	}
	err := fxApp.Start(context.Background())
	if err != nil {
		return err
	}
	go func() {
		// all the logging is for informational purposes. All steps need to complete
		if err := app.Run(); err != nil {
			self.logger.Error("error in RunHandler().App.Run", zap.Error(err))
		}
		if err := app.Close(); err != nil {
			self.logger.Error("error in RunHandler().App.Close", zap.Error(err))
		}
		if err := fxApp.Stop(context.Background()); err != nil {
			self.logger.Error("error in RunHandler().fxApp.Stop", zap.Error(err))
		}
		self.CancelFunc()
		if err := self.SshChannel.Close(); err != nil {
			self.logger.Error("error in RunHandler().SshChannel.Close", zap.Error(err))
		}

	}()
	return nil
}

func NewSshConnectionUi(
	sshChannel common.IChannel,
	parentCtx context.Context,
	parentCancelFunc context.CancelFunc,
	onSend rxgo.NextFunc,
	logger *zap.Logger,
	PubSub *pubsub.PubSub,
	ConnectionManagerIHelper goConnectionManager.IHelper,
	UniqueReferenceService interfaces.IUniqueReferenceService,
	GoFunctionCounter GoFunctionCounter.IService,

) (*SshConnectionUi, error) {
	baseChannelProcess := common.NewBaseChannelProcess(
		sshChannel,
		parentCtx,
		parentCancelFunc,
		onSend,
	)
	var ttyInstance ITty
	ttyInstance = newTty(baseChannelProcess.RwProxy)
	return &SshConnectionUi{
		BaseChannelProcess:       baseChannelProcess,
		tty:                      ttyInstance,
		logger:                   logger,
		PubSub:                   PubSub,
		ConnectionManagerIHelper: ConnectionManagerIHelper,
		UniqueReferenceService:   UniqueReferenceService,
		GoFunctionCounter:        GoFunctionCounter,
	}, nil
}
