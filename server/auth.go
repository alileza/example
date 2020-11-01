package server

import (
	"context"
)

// exampleAuthFunc is used by a middleware to authenticate requests
func exampleAuthFunc(ctx context.Context) (context.Context, error) {
	// ALWAYS PASS
	return ctx, nil

	// token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	// if err != nil {
	// 	return nil, err
	// }

	// tokenInfo, err := parseToken(token)
	// if err != nil {
	// 	return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	// }

	// grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	// // WARNING: in production define your own type to avoid context collisions
	// newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)

	// return newCtx, nil
}

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}
