package goMetadata

import (
	"context"
	"errors"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	ADMIN     = "admin"
	ActorId   = "x-actor-id"
	ActorType = "x-actor-type"
)

func CheckMetaData(ctx context.Context) (string, string, error) {
	userID := ""
	userType := ""
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if md.Get(ActorType) == nil || md.Get(ActorId) == nil {
			return "", "", errors.New("headers nil")
		}
		switch md.Get(ActorType)[0] {
		case ADMIN:
			userType = ADMIN
		default:
			return "", "", errors.New("access denied")
		}

		uID := md.Get(ActorId)[0]
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

func CustomMatcherHeader(key string) (string, bool) {
	switch strings.ToLower(key) {
	case ActorId:
		return ActorId, true
	case ActorType:
		return ActorType, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
