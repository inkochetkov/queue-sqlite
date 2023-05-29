package queue

import (
	"context"

	"github.com/inkochetkov/queue-sqlite/sqlite"
)

// Producer ...
func (q *Queue) Producer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case mes := <-q.MesIn:
			err := q.DB.Set(&sqlite.Entity{Message: mes})
			if err != nil {
				return err
			}
		}
	}

}
