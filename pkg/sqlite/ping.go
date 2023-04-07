package sqlite

// Ping - check connect
func (s *DB) Ping() error {

	err := s.conn.Ping()
	if err != nil {
		return err
	}

	return nil
}
