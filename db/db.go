package db

import (
	"errors"
	"github.com/boltdb/bolt"
)

type DB struct {
	*bolt.DB
}

const (
	dbFileMode = 0600
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

func NewDB(path string) (*DB, error) {

	db, err := bolt.Open(path, dbFileMode, nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) Close() error {

	return db.Close()
}

func (db *DB) Set(k, v, bucketName []byte) error {

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)
	if err := bucket.Put(k, v); err != nil {
		return err
	}

	return tx.Commit()
}

func (db *DB) Get(k, bucketName []byte) ([]byte, error) {

	tx, err := db.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)
	val := bucket.Get(k)

	if val == nil {
		return nil, ErrKeyNotFound
	}

	result := make([]int, len(val))
	copy(result, val)

	return result, nil
}

func (db *DB) Delete(k, bucketName []byte) error {

	tx, err := db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)

	err = bucket.Delete(k)
	if err != nil {
		return err
	}

	return tx.Commit()
}
