package gaego

import (
	"net/http"
	"time"

	netctx "golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

const done = "DONE"

type state struct {
	State      string    `datastore:"state"`
	StartedAt  time.Time `datastore:"startedAt"`
	FinishedAt time.Time `datastore:"finishedAt"`
}

// GAEIdempotenceManage manage idempotence in Google App Engine by Cloud Datastore Transaction.
func GAEIdempotenceManage(req *http.Request, km GAEKeyManager, options *datastore.TransactionOptions, offset int, op func() error) error {
	ctx := appengine.NewContext(req)

	key, err := km.CreateKey(ctx, req)
	if err != nil {
		return err
	}

	err = datastore.RunInTransaction(ctx, func(txCtx netctx.Context) error {
		s := &state{}
		if err := datastore.Get(ctx, key, s); err != nil && err != datastore.ErrNoSuchEntity {
			return err
		}

		if s.State == done {
			return nil
		}

		s.StartedAt = time.Now().Add(time.Hour * time.Duration(offset))

		if err := op(); err != nil {
			return err
		}

		s.FinishedAt = time.Now().Add(time.Hour * time.Duration(offset))

		if _, err := datastore.Put(ctx, key, s); err != nil {
			return err
		}
		return nil
	}, options)
	if err != nil {
		return err
	}
	return nil
}
