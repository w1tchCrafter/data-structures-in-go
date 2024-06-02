package hashtable

// Entry struct represents a key-value pair in the hash table
type Entry[T any] struct {
	Key   string
	Value T
}

// HashTable struct represents the hash table itself
type HashTable[T any] struct {
	Buckets [][]Entry[T]
	Size    int
}

// // NewHashTable creates a new hash table with a specified size
func NewHashTable[T any](size int) *HashTable[T] {
	buckets := make([][]Entry[T], size)
	return &HashTable[T]{
		Buckets: buckets,
		Size:    size,
	}
}
