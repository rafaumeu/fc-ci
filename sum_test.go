package main
import "testing"
func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result!= 5 {
    t.Errorf("Expected 5, got %d", result)
  }
}
func TestSub(t *testing.T) {
	result := sub(3, 1)
	if result != 2 {
		t.Errorf("Expected 1, got %d", result)
	}
}
func TestDiv(t *testing.T) {
	result := div(20, 2)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

func TestTimes(t *testing.T) {
	result := times(5, 2)
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}