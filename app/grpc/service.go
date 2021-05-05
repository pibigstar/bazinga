package baziga

import (
	"context"
	"github.com/pibigstar/bazinga/app/grpc/proto/pb"
)

type BazingaService struct {
	*pb.UnimplementedBazingaServiceServer
}

func (*BazingaService) RandomStory(ctx context.Context, req *pb.RandomStoryReq) (*pb.RandomStoryResp, error) {
	return nil, nil
}
