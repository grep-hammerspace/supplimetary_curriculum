package stacknqueue

import "testing"

func TestPushAddsOneElement(t *testing.T) {
	s := Stack[int]{}

	s.Push(1)
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}

	data, found := s.Peek()
	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
}

func TestPopRemovesOneElement(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	data, found := s.Pop()

	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestPopRemovesAllElements(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Pop()
	s.Pop()
	z
	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d", s.Size())
	}
}

func TestStackPeekDoesntRemove(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)

	data, found := s.Peek()
	if !found {
		t.Error("Expected element to be found")
	}
	if data != 1 {
		t.Errorf("Expected data to be 1, got %d", data)
	}
	if s.Size() != 1 {
		t.Errorf("Expected size 1, got %d", s.Size())
	}
}

func TestPopGetsElementsInLifo(t *testing.T) {
	s := Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	var returnedData []int
	firstVar, found := s.Pop()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, firstVar)
	}

	secondVar, found := s.Pop()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, secondVar)
	}

	thirdVar, found := s.Pop()
	if !found {
		t.Error("Expected element to be found")
	} else {
		returnedData = append(returnedData, thirdVar)
	}

	if returnedData[0] != 3 || returnedData[1] != 2 || returnedData[2] != 1 {
		t.Errorf("Expected data to be [3,2,1], got %v", returnedData)
	}

}
