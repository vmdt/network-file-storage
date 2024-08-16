package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func teardown(t *testing.T, s *Store) {
	if err := s.Clear(); err != nil {
		t.Error(err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)
	defer teardown(t, store)

	key := "folder_name"
	data := []byte("some data")

	_, err := store.writeStream(key, bytes.NewBuffer(data))
	if err != nil {
		t.Error(err)
	}

	if ok := store.Has(key); !ok {
		t.Errorf("expected to have key: %s", key)
	}

	_, r, err := store.Read(key)
	if err != nil {
		t.Error(err)
	}

	b, _ := io.ReadAll(r)
	fmt.Printf("read data in bytes: %v\n", b)
	if string(b) != string(data) {
		t.Errorf("want %s have %s\n", b, data)
	}

	store.Delete(key)
}

func TestDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	store := NewStore(opts)
	data := bytes.NewReader([]byte("some data"))

	key := "folder_name"

	_, err := store.writeStream(key, data)
	if err != nil {
		t.Error(err)
	}

	if err := store.Delete(key); err != nil {
		t.Error(err)
	}
}
