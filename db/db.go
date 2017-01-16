package db

import (
	"errors"

	"github.com/boltdb/bolt"
)

const (
	dbFileMode = 0600
	dbPath     = "/var/lib/mockr/bolt.db"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type DB struct {
	bolt *bolt.DB
}

func New(path string, bucketName []byte) (*DB, error) {

	b, err := bolt.Open(path, dbFileMode, nil)
	if err != nil {
		return nil, err
	}

	tx, err := b.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	if _, err = tx.CreateBucketIfNotExists(bucketName); err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &DB{b}, nil
}

func (d *DB) Close() error {
	return d.bolt.Close()
}

func (d *DB) Set(k, v, bucketName []byte) error {

	tx, err := d.bolt.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)
	if err = bucket.Put(k, v); err != nil {
		return err
	}

	return tx.Commit()
}

func (d *DB) Get(k, bucketName []byte) ([]byte, error) {

	tx, err := d.bolt.Begin(false)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket(bucketName)
	val := bucket.Get(k)

	if val == nil {
		return nil, ErrKeyNotFound
	}

	result := make([]byte, len(val))
	copy(result, val)

	return result, nil
}

func (d *DB) Delete(k, bucketName []byte) error {

	tx, err := d.bolt.Begin(true)
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
