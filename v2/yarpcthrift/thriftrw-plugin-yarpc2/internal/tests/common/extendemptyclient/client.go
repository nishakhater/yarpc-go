// Code generated by thriftrw-plugin-yarpc
// @generated

package extendemptyclient

import (
	"context"
	"go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc/v2"
	"go.uber.org/yarpc/v2/yarpcthrift"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/common"
	"go.uber.org/yarpc/v2/yarpcthrift/thriftrw-plugin-yarpc2/internal/tests/common/emptyserviceclient"
)

// Interface is a client for the ExtendEmpty service.
type Interface interface {
	emptyserviceclient.Interface

	Hello(
		ctx context.Context,
		opts ...yarpc.CallOption,
	) error
}

// New builds a new client for the ExtendEmpty service.
//
//  yarpcClient, err := yarpc.Provider.Client("extendempty")
//  if err != nil {
//	  return err
//  }
// 	client := extendemptyclient.New(yarpcClient)
func New(c yarpc.Client, opts ...yarpcthrift.ClientOption) Interface {
	return client{
		c: yarpcthrift.New(
			c,
			"ExtendEmpty",
			opts...),
		Interface: emptyserviceclient.New(c, opts...),
	}
}

type client struct {
	emptyserviceclient.Interface

	c yarpcthrift.Client
}

func (c client) Hello(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (err error) {

	args := common.ExtendEmpty_Hello_Helper.Args()

	var body wire.Value
	body, err = c.c.Call(ctx, args, opts...)
	if err != nil {
		return
	}

	var result common.ExtendEmpty_Hello_Result
	if err = result.FromWire(body); err != nil {
		return
	}

	err = common.ExtendEmpty_Hello_Helper.UnwrapResponse(&result)
	return
}