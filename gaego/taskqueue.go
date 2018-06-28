package gaego

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/appengine/datastore"
)

// TQKeyManager manage idempotent key of TaskQueue.
type TQKeyManager struct {
	kind string
}

// CreateKey creates idempotent key from task name.
func (m *TQKeyManager) CreateKey(ctx context.Context, req *http.Request) (*datastore.Key, error) {
	name := req.Header.Get("X-AppEngine-TaskName")
	if name == "" {
		return nil, errors.New("name of task is empty")
	}

	key := datastore.NewKey(ctx, m.kind, name, 0, nil)
	return key, nil
}
