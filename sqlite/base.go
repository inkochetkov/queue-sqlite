package sqlite

type Entity struct {
	ID      int64 `gorm:"primaryKey;AUTO_INCREMENT;not null;"`
	Message []byte
}

type DataBase interface {
	Ping() error
	Set(entity *Entity) error
	Get() ([]*Entity, error)
	Del(id int64) error
}
