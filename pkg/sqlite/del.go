package sqlite

func (s *DB) Del(id int64) error {

	err := s.db.Delete(&Entity{}, id).Error
	if err != nil {
		return err
	}

	return nil
}
