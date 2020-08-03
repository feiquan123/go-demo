//go:generate mockgen -source=db.go -destination=db_mock.go -package=db

package db

// DB is a interface of a database
type DB interface {
	Get(key string) (int, error)
}

// GetFromDB is get value from database by key
func GetFromDB(db DB, key string) (int, error) {
	v, err := db.Get(key)
	if err != nil {
		return -1, err
	}
	return v, nil
}
