package queue

import (
	"context"
	"time"
)

// Consumer ...
func (q *Queue) Consumer(ctx context.Context, timeOut time.Duration) error {

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := q.nextSliceMessage(ctx, timeOut)
			if err != nil {
				return err
			}

		}
	}
}

func (q *Queue) nextSliceMessage(ctx context.Context, timeOut time.Duration) error {

	mess, err := q.DB.Get()
	if err != nil {
		return err
	}

	if len(mess) == 0 {
		timer := time.NewTimer(timeOut)
		select {
		case <-ctx.Done():
			return nil
		case <-timer.C:
			return nil
		}
	}

	for _, mes := range mess {
		q.MesOut <- mes.Message
		err := q.DB.Del(mes.ID)
		if err != nil {
			return err
		}
	}

	return nil
}
