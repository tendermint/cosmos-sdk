package db

import "errors"

var (
	// ErrBatchClosed is returned when a closed or written batch is used.
	ErrBatchClosed = errors.New("batch has been written or closed")

	// ErrKeyEmpty is returned when attempting to use an empty or nil key.
	ErrKeyEmpty = errors.New("key cannot be empty")

	// ErrValueNil is returned when attempting to set a nil value.
	ErrValueNil = errors.New("value cannot be nil")

	// // ErrVersionDoesNotExist is returned when a DB version does not exist
	// ErrVersionDoesNotExist = errors.New("version does not exist")
)

// Connection represents a connection to a versioned database.
// Records are accessed via transaction objects, and must be safe for concurrent creation
// and read and write access.
// Past versions are only accessible read-only.
type DB interface {
	// Opens a read-only transaction at the current working version.
	Reader() DBReader

	// Opens a read-only transaction at a specified version.
	// Returns nil for invalid versions.
	ReaderAt(uint64) DBReader

	// Opens a read-write transaction at the current version.
	ReadWriter() DBReadWriter

	// Opens a write-only transaction at the current version.
	Writer() DBWriter

	// Returns all saved versions
	Versions() VersionSet

	// Saves the current version of the database and returns its ID.
	// Waits for any pending Writer transactions to be discarded. TODO: is an error preferred?
	SaveVersion() uint64

	// Close closes the database connection.
	Close() error

	// Stats returns a map of property values for all keys and the size of the cache.
	Stats() map[string]string
}

// TODO: rename to batch/transaction?
// DBReader is a read-only transaction interface. It is safe for concurrent access.
// Callers must call Discard when done with the transaction.
//
// Keys cannot be nil or empty, while values cannot be nil. Keys and values should be considered
// read-only, both when returned and when given, and must be copied before they are modified.
type DBReader interface {
	// Get fetches the value of the given key, or nil if it does not exist.
	// CONTRACT: key, value readonly []byte
	Get([]byte) ([]byte, error)

	// Has checks if a key exists.
	// CONTRACT: key, value readonly []byte
	Has(key []byte) (bool, error)

	// Iterator returns an iterator over a domain of keys, in ascending order. The caller must call
	// Close when done. End is exclusive, and start must be less than end. A nil start iterates
	// from the first key, and a nil end iterates to the last key (inclusive). Empty keys are not
	// valid.
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// CONTRACT: start, end readonly []byte
	Iterator(start, end []byte) (Iterator, error)

	// ReverseIterator returns an iterator over a domain of keys, in descending order. The caller
	// must call Close when done. End is exclusive, and start must be less than end. A nil end
	// iterates from the last key (inclusive), and a nil start iterates to the first key (inclusive).
	// Empty keys are not valid.
	// CONTRACT: No writes may happen within a domain while an iterator exists over it.
	// CONTRACT: start, end readonly []byte
	ReverseIterator(start, end []byte) (Iterator, error)

	// Discards the transaction, invalidating any future operations on it.
	Discard()
}

// DBWriter is a write-only transaction interface.
// It is safe for concurrent writes, following an optimistic (OCC) strategy, detecting any write
// conflicts and returning an error on commit, rather than locking the DB.
// This can be used to wrap a write-optimized batch object if provided by the backend implementation.
type DBWriter interface {
	// Set sets the value for the given key, replacing it if it already exists.
	// CONTRACT: key, value readonly []byte
	Set([]byte, []byte) error

	// Delete deletes the key, or does nothing if the key does not exist.
	// CONTRACT: key readonly []byte
	Delete([]byte) error

	// Flushes pending writes and discards the transaction.
	// TODO: maybe change to Flush() and follow WriteBatch semantics (ie. don't discard)
	Commit() error

	// Discards the transaction, invalidating any future operations on it.
	Discard()
}

// DBReadWriter is a transaction interface that allows both reading and writing.
type DBReadWriter interface {
	DBReader
	DBWriter
}

// Iterator represents an iterator over a domain of keys. Callers must call Close when done.
// No writes can happen to a domain while there exists an iterator over it, some backends may take
// out database locks to ensure this will not happen.
//
// Callers must make sure the iterator is valid before calling any methods on it, otherwise
// these methods will panic. This is in part caused by most backend databases using this convention.
//
// As with DB, keys and values should be considered read-only, and must be copied before they are
// modified.
//
// Typical usage:
//
// var itr Iterator = ...
// defer itr.Close()
//
// for ; itr.Valid(); itr.Next() {
//   k, v := itr.Key(); itr.Value()
//   ...
// }
// if err := itr.Error(); err != nil {
//   ...
// }
type Iterator interface {
	// Domain returns the start (inclusive) and end (exclusive) limits of the iterator.
	// CONTRACT: start, end readonly []byte
	Domain() (start []byte, end []byte)

	// Valid returns whether the current iterator is valid. Once invalid, the Iterator remains
	// invalid forever.
	Valid() bool

	// Next moves the iterator to the next key in the database, as defined by order of iteration.
	// If Valid returns false, this method will panic.
	Next()

	// Key returns the key at the current position. Panics if the iterator is invalid.
	// CONTRACT: key readonly []byte
	Key() (key []byte)

	// Value returns the value at the current position. Panics if the iterator is invalid.
	// CONTRACT: value readonly []byte
	Value() (value []byte)

	// Error returns the last error encountered by the iterator, if any.
	Error() error

	// Close closes the iterator, relasing any allocated resources.
	Close() error
}

// VersionSet specifies a set of existing versions
type VersionSet interface {
	// Last returns the most recent saved version, or 0 if none
	Last() uint64
	// Initial returns the first saved version, or 0 if none
	Initial() uint64
	// Count returns the number of saved versions
	Count() int
	// All returns all saved versions as an ordered slice
	All() []uint64
}
