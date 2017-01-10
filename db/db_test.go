package db

import (
    "testing"
    "bytes"
)

func TestNew(t *testing.T) {

    d, err := New("../testdata/bolt.db", []byte("new"))
    if err != nil {
        t.Fatalf("E! test fail %v", err)
    }
    defer d.Close()
}

func TestSetAndGetAndDelete(t *testing.T) {

    testKey, testValue, testBucket := []byte("testKey"), []byte("testValue"), []byte("testBucket")
    d, err := New("../testdata/bolt.db", testBucket)
    if err != nil {
        t.Fatalf("E! new fail %v", err)
    }
    defer d.Close()

    err = d.Set(testKey, testValue, testBucket)
    if err != nil {
        t.Errorf("E! set fail %v", err)
    }

    valueInBolt, err := d.Get(testKey, testBucket)
    if err != nil {
        t.Errorf("E! get fail %v", err)
    }
    if !bytes.Equal(testValue, valueInBolt) {
        t.Errorf("E! get value %s didn't equal %s", valueInBolt, testValue)
    }

    err = d.Delete(testKey, testBucket)
    if err != nil {
        t.Errorf("E! delete fail %v", err)
    }

    // make sure key is deleted
    _, err = d.Get(testKey, testBucket)
    if err.Error() != ErrKeyNotFound.Error() {
        t.Error("E! key not deleted")
    }
}