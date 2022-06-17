// Code generated by Kitex v0.2.1. DO NOT EDIT.

package aggregateddiscoveryservice

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	discoveryv3 "github.com/ppzqh/kitex_xds_api/discoveryv3"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return aggregatedDiscoveryServiceServiceInfo
}

var aggregatedDiscoveryServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "AggregatedDiscoveryService"
	handlerType := (*discoveryv3.AggregatedDiscoveryService)(nil)
	methods := map[string]kitex.MethodInfo{
		"StreamAggregatedResources": kitex.NewMethodInfo(streamAggregatedResourcesHandler, newStreamAggregatedResourcesArgs, newStreamAggregatedResourcesResult, false),
		"DeltaAggregatedResources":  kitex.NewMethodInfo(deltaAggregatedResourcesHandler, newDeltaAggregatedResourcesArgs, newDeltaAggregatedResourcesResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "envoy.service.discovery.v3",
	}
	extra["streaming"] = true
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.2.1",
		Extra:           extra,
	}
	return svcInfo
}

func streamAggregatedResourcesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &aggregatedDiscoveryServiceStreamAggregatedResourcesServer{st}
	return handler.(discoveryv3.AggregatedDiscoveryService).StreamAggregatedResources(stream)
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesClient struct {
	streaming.Stream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Send(m *discoveryv3.DiscoveryRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*discoveryv3.DiscoveryResponse, error) {
	m := new(discoveryv3.DiscoveryResponse)
	return m, x.Stream.RecvMsg(m)
}

type aggregatedDiscoveryServiceStreamAggregatedResourcesServer struct {
	streaming.Stream
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Send(m *discoveryv3.DiscoveryResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*discoveryv3.DiscoveryRequest, error) {
	m := new(discoveryv3.DiscoveryRequest)
	return m, x.Stream.RecvMsg(m)
}

func newStreamAggregatedResourcesArgs() interface{} {
	return &StreamAggregatedResourcesArgs{}
}

func newStreamAggregatedResourcesResult() interface{} {
	return &StreamAggregatedResourcesResult{}
}

type StreamAggregatedResourcesArgs struct {
	Req *discoveryv3.DiscoveryRequest
}

func (p *StreamAggregatedResourcesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in StreamAggregatedResourcesArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *StreamAggregatedResourcesArgs) Unmarshal(in []byte) error {
	msg := new(discoveryv3.DiscoveryRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var StreamAggregatedResourcesArgs_Req_DEFAULT *discoveryv3.DiscoveryRequest

func (p *StreamAggregatedResourcesArgs) GetReq() *discoveryv3.DiscoveryRequest {
	if !p.IsSetReq() {
		return StreamAggregatedResourcesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *StreamAggregatedResourcesArgs) IsSetReq() bool {
	return p.Req != nil
}

type StreamAggregatedResourcesResult struct {
	Success *discoveryv3.DiscoveryResponse
}

var StreamAggregatedResourcesResult_Success_DEFAULT *discoveryv3.DiscoveryResponse

func (p *StreamAggregatedResourcesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in StreamAggregatedResourcesResult")
	}
	return proto.Marshal(p.Success)
}

func (p *StreamAggregatedResourcesResult) Unmarshal(in []byte) error {
	msg := new(discoveryv3.DiscoveryResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *StreamAggregatedResourcesResult) GetSuccess() *discoveryv3.DiscoveryResponse {
	if !p.IsSetSuccess() {
		return StreamAggregatedResourcesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *StreamAggregatedResourcesResult) SetSuccess(x interface{}) {
	p.Success = x.(*discoveryv3.DiscoveryResponse)
}

func (p *StreamAggregatedResourcesResult) IsSetSuccess() bool {
	return p.Success != nil
}

func deltaAggregatedResourcesHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st := arg.(*streaming.Args).Stream
	stream := &aggregatedDiscoveryServiceDeltaAggregatedResourcesServer{st}
	return handler.(discoveryv3.AggregatedDiscoveryService).DeltaAggregatedResources(stream)
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesClient struct {
	streaming.Stream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *discoveryv3.DeltaDiscoveryRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*discoveryv3.DeltaDiscoveryResponse, error) {
	m := new(discoveryv3.DeltaDiscoveryResponse)
	return m, x.Stream.RecvMsg(m)
}

type aggregatedDiscoveryServiceDeltaAggregatedResourcesServer struct {
	streaming.Stream
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *discoveryv3.DeltaDiscoveryResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *aggregatedDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*discoveryv3.DeltaDiscoveryRequest, error) {
	m := new(discoveryv3.DeltaDiscoveryRequest)
	return m, x.Stream.RecvMsg(m)
}

func newDeltaAggregatedResourcesArgs() interface{} {
	return &DeltaAggregatedResourcesArgs{}
}

func newDeltaAggregatedResourcesResult() interface{} {
	return &DeltaAggregatedResourcesResult{}
}

type DeltaAggregatedResourcesArgs struct {
	Req *discoveryv3.DeltaDiscoveryRequest
}

func (p *DeltaAggregatedResourcesArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in DeltaAggregatedResourcesArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *DeltaAggregatedResourcesArgs) Unmarshal(in []byte) error {
	msg := new(discoveryv3.DeltaDiscoveryRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var DeltaAggregatedResourcesArgs_Req_DEFAULT *discoveryv3.DeltaDiscoveryRequest

func (p *DeltaAggregatedResourcesArgs) GetReq() *discoveryv3.DeltaDiscoveryRequest {
	if !p.IsSetReq() {
		return DeltaAggregatedResourcesArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *DeltaAggregatedResourcesArgs) IsSetReq() bool {
	return p.Req != nil
}

type DeltaAggregatedResourcesResult struct {
	Success *discoveryv3.DeltaDiscoveryResponse
}

var DeltaAggregatedResourcesResult_Success_DEFAULT *discoveryv3.DeltaDiscoveryResponse

func (p *DeltaAggregatedResourcesResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in DeltaAggregatedResourcesResult")
	}
	return proto.Marshal(p.Success)
}

func (p *DeltaAggregatedResourcesResult) Unmarshal(in []byte) error {
	msg := new(discoveryv3.DeltaDiscoveryResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *DeltaAggregatedResourcesResult) GetSuccess() *discoveryv3.DeltaDiscoveryResponse {
	if !p.IsSetSuccess() {
		return DeltaAggregatedResourcesResult_Success_DEFAULT
	}
	return p.Success
}

func (p *DeltaAggregatedResourcesResult) SetSuccess(x interface{}) {
	p.Success = x.(*discoveryv3.DeltaDiscoveryResponse)
}

func (p *DeltaAggregatedResourcesResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) StreamAggregatedResources(ctx context.Context) (AggregatedDiscoveryService_StreamAggregatedResourcesClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "StreamAggregatedResources", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &aggregatedDiscoveryServiceStreamAggregatedResourcesClient{res.Stream}
	return stream, nil
}

func (p *kClient) DeltaAggregatedResources(ctx context.Context) (AggregatedDiscoveryService_DeltaAggregatedResourcesClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "DeltaAggregatedResources", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &aggregatedDiscoveryServiceDeltaAggregatedResourcesClient{res.Stream}
	return stream, nil
}
