package sqlite

func (s *DB) Get() ([]*Entity, error) {

	var entity []*Entity

	err := s.db.Order("id asc").Find(&entity).Error
	if err != nil {
		return nil, err
	}

	return entity, nil
}
