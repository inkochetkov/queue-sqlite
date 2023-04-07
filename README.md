## queue-sqlite

queue implemented on sqlite

## use module

 gorm
 sqlite

## Examle

### Consumer

	ctx := context.Background()
	q := queue.New("./", "any.sqlite3")


	go func() {
		err := q.Consumer(ctx, 10*time.Second)
		if err!=nil{
            log.Println(err)
        }
	}()

	var mNew []byte
	go func() {
		mNew = <-q.MesOut
	}()

    time.Sleep(10 * time.Second)

    log.Println(mNew)

### Producer

    mes := []byte("test")

	ctx := context.Background()
	q := queue.New("./", "any.sqlite3")

	go func() {
		err := q.Producer(ctx)
		if err!=nil{
            log.Println(err)
        }
	}()

	go func() {
		q.MesIn <- mes
	}()

	time.Sleep(10 * time.Second)