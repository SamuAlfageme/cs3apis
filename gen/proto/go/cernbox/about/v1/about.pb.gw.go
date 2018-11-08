// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: cernbox/about/v1/about.proto

/*
Package aboutv1pb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package aboutv1pb

import (
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_AboutService_ListMembers_0(ctx context.Context, marshaler runtime.Marshaler, client AboutServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq ListMembersRequest
	var metadata runtime.ServerMetadata

	msg, err := client.ListMembers(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func request_AboutService_GetDocumentation_0(ctx context.Context, marshaler runtime.Marshaler, client AboutServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetDocumentationRequest
	var metadata runtime.ServerMetadata

	msg, err := client.GetDocumentation(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

// RegisterAboutServiceHandlerFromEndpoint is same as RegisterAboutServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterAboutServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterAboutServiceHandler(ctx, mux, conn)
}

// RegisterAboutServiceHandler registers the http handlers for service AboutService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterAboutServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterAboutServiceHandlerClient(ctx, mux, NewAboutServiceClient(conn))
}

// RegisterAboutServiceHandlerClient registers the http handlers for service AboutService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "AboutServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "AboutServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "AboutServiceClient" to call the correct interceptors.
func RegisterAboutServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client AboutServiceClient) error {

	mux.Handle("GET", pattern_AboutService_ListMembers_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AboutService_ListMembers_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AboutService_ListMembers_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_AboutService_GetDocumentation_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_AboutService_GetDocumentation_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_AboutService_GetDocumentation_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_AboutService_ListMembers_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "about", "members"}, ""))

	pattern_AboutService_GetDocumentation_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2}, []string{"v1", "about", "documentation"}, ""))
)

var (
	forward_AboutService_ListMembers_0 = runtime.ForwardResponseMessage

	forward_AboutService_GetDocumentation_0 = runtime.ForwardResponseMessage
)