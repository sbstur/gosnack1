package tdd

import "testing"

func TestInput1ShouldBe1(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"
	if actual != expected {
		t.Error("expected ", expected, "but got ", actual)
	}
	//	t.Error("fails")
}

func TestInput2ShouldBe2(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"
	if actual != expected {
		t.Error("expected ", expected, "but got ", actual)
	}
}

func TestInput3ShouldBeFizz(t *testing.T) {
	actual := FizzBuzz(3)
	expected := "Fizz"
	if actual != expected {
		t.Error("expected ", expected, "but got ", actual)
	}
}

func TestInput4ShouldBe4(t *testing.T) {
	actual := FizzBuzz(4)
	expected := "4"
	if actual != expected {
		t.Error("expected ", expected, "but got ", actual)
	}
}

func TestInput5ShouldBeBuzz(t *testing.T) {
	actual := FizzBuzz(5)
	expected := "Buzz"
	if actual != expected {
		t.Error("expected ", expected, "but got ", actual)
	}
}

func TestInput6ShouldBeFizz(t *testing.T) {

}
