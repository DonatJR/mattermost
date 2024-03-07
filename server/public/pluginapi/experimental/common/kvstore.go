package common

import (
	"errors"
	"time"

	"github.com/mattermost/mattermost/server/public/pluginapi"
)

var ErrNotFound = errors.New("not found")

type KVStore interface {
	Set(key string, value interface{}, options ...pluginapi.KVSetOption) (bool, error)
	SetWithExpiry(key string, value interface{}, ttl time.Duration) error
	CompareAndSet(key string, oldValue, value interface{}) (bool, error)
	CompareAndDelete(key string, oldValue interface{}) (bool, error)
	Get(key string, o interface{}) error
	Delete(key string) error
	DeleteAll() error
	ListKeys(page, count int, options ...pluginapi.ListKeysOption) ([]string, error)
}
