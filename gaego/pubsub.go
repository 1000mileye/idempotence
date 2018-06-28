package gaego

import (
	"context"
	"encoding/json"
	"net/http"

	pubsub "google.golang.org/api/pubsub/v1"
	"google.golang.org/appengine/datastore"
)

type pubsubRequest struct {
	Message      *pubsub.PubsubMessage `json:"message"`
	Subscription string                `json:"subscription"`
}

// PubsubKeyManager manage idempotent key of Cloud Pub/Sub
type PubsubKeyManager struct {
	webhookID string
	kind      string
}

// CreateKey creates datastore key from both pubsub message id and webhook id.
func (m *PubsubKeyManager) CreateKey(ctx context.Context, req *http.Request) (*datastore.Key, error) {
	var pubsubReq pubsubRequest
	if err := json.NewDecoder(req.Body).Decode(&pubsubReq); err != nil {
		return nil, err
	}

	msgID := pubsubReq.Message.MessageId + "-" + m.webhookID
	key := datastore.NewKey(ctx, m.kind, msgID, 0, nil)
	return key, nil
}
