package redis

import (
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

const Nil = redis.Nil

func HandleError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, redis.Nil) {
		return errors.Wrap(err, "NotFound")
	}
	if errors.Is(err, redis.TxFailedErr) {
		return errors.Wrap(err, "TxFailed")
	}
	return errors.Wrap(err, "InternalError")
}
