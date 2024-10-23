package grpc

import (
	"context"
	"git.solsynth.dev/hydrogen/passport/pkg/internal/services"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
)

func (v *Server) BroadcastEvent(ctx context.Context, request *proto.EventInfo) (*proto.EventResponse, error) {
	switch request.GetEvent() {
	case "ws.client.register":
		// No longer need update user online status
		// Based on realtime sever connection status
		break
	case "ws.client.unregister":
		// Update user last seen at
		data := nex.DecodeMap(request.GetData())
		_ = services.SetAccountLastSeen(uint(data["user"].(float64)))
	}

	return &proto.EventResponse{}, nil
}
