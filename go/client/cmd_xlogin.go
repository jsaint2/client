package client

import (
	"errors"

	"github.com/keybase/cli"
	"github.com/keybase/client/go/libcmdline"
	"github.com/keybase/client/go/libkb"
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

func NewCmdXLogin(cl *libcmdline.CommandLine) cli.Command {
	return cli.Command{
		Name:         "xlogin",
		ArgumentHelp: "[username]",
		Usage:        "Establish a session with the keybase server",
		Action: func(c *cli.Context) {
			cl.ChooseCommand(&CmdXLogin{}, "xlogin", c)
		},
	}
}

type CmdXLogin struct {
	username string
}

func (c *CmdXLogin) Run() error {
	protocols := []rpc.Protocol{}
	if err := RegisterProtocols(protocols); err != nil {
		return err
	}
	client, err := GetLoginClient()
	if err != nil {
		return err
	}
	return client.XLogin(c.username)
}

func (c *CmdXLogin) ParseArgv(ctx *cli.Context) error {
	nargs := len(ctx.Args())
	if nargs > 1 {
		return errors.New("Invalid arguments.")
	}

	if nargs == 1 {
		c.username = ctx.Args()[0]
	}
	return nil
}

func (c *CmdXLogin) GetUsage() libkb.Usage {
	return libkb.Usage{
		Config:    true,
		KbKeyring: true,
		API:       true,
	}
}
