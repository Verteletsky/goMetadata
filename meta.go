package goMetadata

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	ADMIN = "admin"
)

func CheckMetaData(ctx context.Context) (string, string, error) {
	userID := ""
	userType := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		switch md.Get("x-actor-type")[0] {
		case ADMIN:
			userType = ADMIN
		default:
			return "", "", errors.New("access denied")
		}

		uID := md.Get("x-actor-id")[0]
		if len(uID) > 0 {
			userID = uID
		} else {
			return "", "", errors.New("unauthorized")
		}
	} else {
		return "", "", errors.New("nil headers")
	}

	return userID, userType, nil
}

func CustomMatherHeader(key string) (string, bool) {
	switch strings.ToLower(key) {
	case "x-actor-id":
		return "x-actor-id", true
	case "x-actor-type":
		return "x-actor-type", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
