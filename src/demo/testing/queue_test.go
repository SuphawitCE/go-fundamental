package queue

import "testing"

func TestAddQueue(test *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		if len(q.items) != i {
			test.Errorf("Incorrect queue element count: %v, expect to: %v", len(q.items), i)
		}
		if !q.Append(i) {
			test.Errorf("Failed to append item %v to queue", i)
		}
	}

	if q.Append(4) {
		test.Errorf("should not be able to add to a full queue")
	}
}

func TestNext(test *testing.T) {
	q := New(3)

	for i := 0; i < 3; i++ {
		q.Append(i)
	}

	for i := 0; i < 3; i++ {
		item, ok := q.Next()
		if !ok {
			test.Errorf("Should be able to get item from queue")
		}
		if item != i {
			test.Errorf("Got item in wrong order: %v, expect to %v", item, i)
		}
	}

	// Queue is empty at this point
	item, ok := q.Next()
	if ok {
		test.Errorf("Should not be any more items in queue, got %v", item)
	}
}
