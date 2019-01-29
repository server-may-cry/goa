// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// divider gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/error/design -o
// $(GOPATH)/src/goa.design/goa/examples/error

package client

import (
	"context"

	dividersvc "goa.design/goa/examples/error/gen/divider"
	"goa.design/goa/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildIntegerDivideFunc builds the remote method to invoke for "divider"
// service "integer_divide" endpoint.
func BuildIntegerDivideFunc(grpccli pb.DividerClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		return grpccli.IntegerDivide(ctx, reqpb.(*pb.IntegerDivideRequest), opts...)
	}
}

// EncodeIntegerDivideRequest encodes requests sent to divider integer_divide
// endpoint.
func EncodeIntegerDivideRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*dividersvc.IntOperands)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "*dividersvc.IntOperands", v)
	}
	return NewIntegerDivideRequest(payload), nil
}

// DecodeIntegerDivideResponse decodes responses from the divider
// integer_divide endpoint.
func DecodeIntegerDivideResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*pb.IntegerDivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "*pb.IntegerDivideResponse", v)
	}
	res := NewIntegerDivideResponse(message)
	return res, nil
}

// BuildDivideFunc builds the remote method to invoke for "divider" service
// "divide" endpoint.
func BuildDivideFunc(grpccli pb.DividerClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		return grpccli.Divide(ctx, reqpb.(*pb.DivideRequest), opts...)
	}
}

// EncodeDivideRequest encodes requests sent to divider divide endpoint.
func EncodeDivideRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*dividersvc.FloatOperands)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "divide", "*dividersvc.FloatOperands", v)
	}
	return NewDivideRequest(payload), nil
}

// DecodeDivideResponse decodes responses from the divider divide endpoint.
func DecodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*pb.DivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "divide", "*pb.DivideResponse", v)
	}
	res := NewDivideResponse(message)
	return res, nil
}