module github.com/bhbosman/sshApplication

go 1.18

require (
	github.com/bhbosman/goCommsSshListener v0.0.0-20220614205207-b3d88b2f7ef8
	github.com/bhbosman/goConnectionManager v0.0.0-20230302065222-d613f6fe8f80
	github.com/bhbosman/goFxApp v0.0.0-20220621055337-d5d1bfa131ee
	github.com/bhbosman/goUi v0.0.0-20230302065227-24c3cb06165e
	github.com/bhbosman/gocommon v0.0.0-20230303055326-d59d6b655a59
	github.com/bhbosman/gocomms v0.0.0-20220628074707-e93417aaaed2
	github.com/cskr/pubsub v1.0.2
	github.com/gdamore/tcell/v2 v2.5.1
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rivo/tview v0.0.0-20220610163003-691f46d6f500
	go.uber.org/fx v1.19.2
	go.uber.org/multierr v1.6.0
	go.uber.org/zap v1.23.0
)

require (
	github.com/bhbosman/goCommsDefinitions v0.0.0-20230302063431-887458f90947 // indirect
	github.com/bhbosman/goCommsNetListener v0.0.0-20220614200404-f1193638b13b // indirect
	github.com/bhbosman/goCommsStacks v0.0.0-20220611141421-a7d405cadbfa // indirect
	github.com/bhbosman/goFxAppManager v0.0.0-20220730103022-02c299931769 // indirect
	github.com/bhbosman/goerrors v0.0.0-20220623084908-4d7bbcd178cf // indirect
	github.com/bhbosman/gomessageblock v0.0.0-20220617132215-32f430d7de62 // indirect
	github.com/bhbosman/goprotoextra v0.0.2-0.20210817141206-117becbef7c7 // indirect
	github.com/cenkalti/backoff/v4 v4.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/emirpasic/gods v1.12.0 // indirect
	github.com/gdamore/encoding v1.0.0 // indirect
	github.com/golang/mock v1.6.0 // indirect
	github.com/icza/gox v0.0.0-20220321141217-e2d488ab2fbc // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/stretchr/testify v1.8.0 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	golang.org/x/crypto v0.7.0 // indirect
	golang.org/x/net v0.8.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.6.0 // indirect
	golang.org/x/term v0.6.0 // indirect
	golang.org/x/text v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/gdamore/tcell/v2 => ../tcell

replace github.com/bhbosman/gocomms => ../gocomms

replace github.com/bhbosman/gocommon => ../gocommon

replace github.com/golang/mock => ../gomock

replace github.com/bhbosman/goCommsNetListener => ../goCommsNetListener

replace github.com/bhbosman/goCommsSshListener => ../goCommsSshListener

replace github.com/bhbosman/goCommsDefinitions => ../goCommsDefinitions

replace github.com/bhbosman/goCommsStacks => ../goCommsStacks

replace github.com/bhbosman/goCommsSSHProtocols => ../goCommsSSHProtocols

replace github.com/bhbosman/goCommsSSH => ../goCommsSSH

replace github.com/bhbosman/goFxApp => ../goFxApp

replace github.com/bhbosman/goUi => ../goUi

replace github.com/bhbosman/goerrors => ../goerrors

replace github.com/bhbosman/goFxAppManager => ../goFxAppManager

replace github.com/bhbosman/goConnectionManager => ../goConnectionManager

replace github.com/rivo/tview => ../tview

replace github.com/bhbosman/goprotoextra => ../goprotoextra

replace github.com/cskr/pubsub => github.com/bhbosman/pubsub v1.0.3-0.20220802200819-029949e8a8af
