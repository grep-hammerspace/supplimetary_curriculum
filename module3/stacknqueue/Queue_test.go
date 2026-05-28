package stacknqueue

import "testing"

func TestEnqueueAddsOneElement(t *testing.T) {
	q := Queue[int]{}

	q.Enqueue(1)
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Elements.Size())
	}

	data, found := q.Peek()
	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
}

func TestDequeueRemovesOneElement(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	data, found := q.Deque()

	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}
}

func TestQueuePeekDoesntRemove(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)

	data, found := q.Peek()
	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
	if q.Size() != 1 {
		t.Errorf("Expected size 1, got %d", q.Size())
	}
}

func TestDequeueGetsElementsInFifoAndDecreasesSize(t *testing.T) {
	q := Queue[int]{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	var returnedData []int
	firstVar, found := q.Deque()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, firstVar)
	}

	secondVar, found := q.Deque()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, secondVar)
	}

	thirdVar, found := q.Deque()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, thirdVar)
	}

	if returnedData[0] != 1 || returnedData[1] != 2 || returnedData[2] != 3 {
		t.Errorf("Expected data to be [1,2,3], got %v", returnedData)
	}

	if q.Size() != 0 {
		t.Errorf("Expected size 0, got %d", q.Size())
	}

}
