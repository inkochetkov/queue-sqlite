package test

import (
	"context"
	"testing"
	"time"

	"github.com/inkochetkov/queue-sqlite/pkg/queue"
	"github.com/inkochetkov/queue-sqlite/pkg/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestSqlite(t *testing.T) {

	q := queue.New("./", "db.sqlite3")

	err := q.DB.Ping()
	assert.NoError(t, err)

	mes := []byte("test")

	m := &sqlite.Entity{Message: mes}
	err = q.DB.Set(m)
	assert.NoError(t, err)

	ms, err := q.DB.Get()
	assert.NoError(t, err)

	err = q.DB.Del(ms[0].ID)
	assert.NoError(t, err)

	t.Log("sqlite current")

}

func TestQueueProducer(t *testing.T) {

	mes := []byte("test")

	ctx := context.Background()
	q := queue.New("./", "db.sqlite3")

	t.Log("get queue interface current")

	go func() {

		t.Log("start producer")
		err := q.Producer(ctx)
		assert.NoError(t, err)
	}()
	go func() {
		t.Log("send  mes")
		q.MesIn <- mes
	}()

	time.Sleep(20 * time.Second)

	t.Log("test  producer success")

}

func TestQueueConsumer(t *testing.T) {

	ctx := context.Background()
	q := queue.New("./", "db.sqlite3")

	t.Log("get queue interface current")
	t.Log("test  consumer start")

	go func() {
		err := q.Consumer(ctx, 10*time.Second)
		assert.NoError(t, err)
	}()

	var mNew []byte
	go func() {
		mNew = <-q.MesOut
	}()

	time.Sleep(20 * time.Second)

	t.Logf("%s", mNew)

	t.Log("test  consumer success")

}
