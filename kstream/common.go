package kstream

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc/metadata"
)

// GetFromContext 根据一个grpc的context获取出Session.
func getFromContext(ctx context.Context) (uint32, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	gids := md["gid"]
	if len(gids) == 0 {
		return uint32(0), errors.New("gid")
	}

	gid, _ := strconv.Atoi(gids[0])
	return uint32(gid), nil
}
