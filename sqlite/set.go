package sqlite

func (s *DB) Set(entity *Entity) error {

	err := s.db.Create(entity).Error
	if err != nil {
		return err
	}

	return nil
}
