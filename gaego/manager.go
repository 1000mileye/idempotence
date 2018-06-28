package idempotence

import (
	"context"
	"net/http"

	"google.golang.org/appengine/datastore"
)

// GAEKeyManager manage idempotence in Google App Engine.
// Pass context to the method because we don't want to create appengine context twice.
type GAEKeyManager interface {
	CreateKey(ctx context.Context, req *http.Request) (*datastore.Key, error)
}
