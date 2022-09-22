package goMetadata

import (
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
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
