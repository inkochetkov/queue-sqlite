package queue

import (
	"context"
	"time"

	"github.com/inkochetkov/queue-sqlite/sqlite"
)

type Mes []byte

type Queue struct {
	DB     sqlite.DataBase
	MesIn  chan Mes
	MesOut chan Mes
}

type Queues interface {
	Consumer(ctx context.Context, timeOut time.Duration) error
	Producer(ctx context.Context) error
}

func New(path, fileName string) *Queue {

	db := sqlite.Start(path, fileName)

	mesIn := make(chan Mes)
	mesOut := make(chan Mes)

	queue := &Queue{DB: db, MesIn: mesIn, MesOut: mesOut}

	return queue
}
