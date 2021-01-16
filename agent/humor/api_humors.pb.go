// Code generated by protoc-gen-go-humors. DO NOT EDIT.
// protoc-gen-go-humors version: 0.1.0

package humor

import (
	context "context"
	humors "github.com/wilenceyao/humors"
	proto "google.golang.org/protobuf/proto"
)

// AgentServiceClient is the client API for AgentService service.
type AgentServiceClient interface {
	Tts(ctx context.Context, clientID string, in *TtsRequest) (*TtsResponse, error)
	Weather(ctx context.Context, clientID string, in *WeatherRequest) (*WeatherResponse, error)
	TakePhoto(ctx context.Context, clientID string, in *TakePhotoRequest) (*TakePhotoResponse, error)
}

type agentServiceClient struct {
	adaptor *humors.HumorAdaptor
}

func NewAgentServiceClient(h *humors.Humors) AgentServiceClient {
	adaptor := h.NewAdaptor("AgentService")
	return &agentServiceClient{
		adaptor: adaptor,
	}
}

func (c *agentServiceClient) Tts(ctx context.Context, clientID string, in *TtsRequest) (*TtsResponse, error) {
	out := new(TtsResponse)
	err := c.adaptor.Call(ctx, clientID, "Tts", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Weather(ctx context.Context, clientID string, in *WeatherRequest) (*WeatherResponse, error) {
	out := new(WeatherResponse)
	err := c.adaptor.Call(ctx, clientID, "Weather", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) TakePhoto(ctx context.Context, clientID string, in *TakePhotoRequest) (*TakePhotoResponse, error) {
	out := new(TakePhotoResponse)
	err := c.adaptor.Call(ctx, clientID, "TakePhoto", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

type AgentServiceServer interface {
	Tts(context.Context, *TtsRequest) (*TtsResponse, error)
	Weather(context.Context, *WeatherRequest) (*WeatherResponse, error)
	TakePhoto(context.Context, *TakePhotoRequest) (*TakePhotoResponse, error)
}

func RegisterAgentServiceServer(h *humors.Humors, impl AgentServiceServer) {
	servant := h.NewServant("AgentService")
	servant.RegisterMethod("Tts", TtsRequest{}, func(ctx context.Context, req proto.Message) (proto.Message, error) {
		reqObj := req.(*TtsRequest)
		return impl.Tts(ctx, reqObj)
	})
	servant.RegisterMethod("Weather", WeatherRequest{}, func(ctx context.Context, req proto.Message) (proto.Message, error) {
		reqObj := req.(*WeatherRequest)
		return impl.Weather(ctx, reqObj)
	})
	servant.RegisterMethod("TakePhoto", TakePhotoRequest{}, func(ctx context.Context, req proto.Message) (proto.Message, error) {
		reqObj := req.(*TakePhotoRequest)
		return impl.TakePhoto(ctx, reqObj)
	})
}
