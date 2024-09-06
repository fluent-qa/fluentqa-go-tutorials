package kg

import (
	"sync"
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"
)

func TestMergeCorrectlyMergesTwoMapsOfIntToBool(t *testing.T) {
	t.Parallel()
	inputs := []map[int]bool{
		{
			1: false,
			2: false,
			3: false,
		},
		{
			3: true,
			5: true,
		},
	}
	want := map[int]bool{
		1: false,
		2: false,
		3: true,
		5: true,
	}
	got := Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeCorrectlyMergesThreeMapsOfStringToAny(t *testing.T) {
	t.Parallel()
	inputs := []map[string]any{
		{
			"a": nil,
		},
		{
			"b": "hello, world",
			"c": 0,
		},
		{
			"a": 6 + 2i,
		},
	}
	want := map[string]any{
		"a": 6 + 2i,
		"b": "hello, world",
		"c": 0,
	}
	got := Merge(inputs...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func positive(p int) bool {
	return p > 0
}

func TestContainsFunc_IsTrueForPositiveOnInputContainingPositiveInts(t *testing.T) {
	t.Parallel()
	input := []int{-2, 0, 1, -1, 5}
	got := ContainsFunc(input, positive)
	if !got {
		t.Fatalf("%v: want true for 'contains positive', got false", input)
	}
}

func TestContainsFunc_IsFalseForIsUpperOnInputContainingNoUppercaseRunes(t *testing.T) {
	t.Parallel()
	input := []rune("hello")
	got := ContainsFunc(input, unicode.IsUpper)
	if got {
		t.Fatalf("%q: want false for 'contains uppercase', got true", input)
	}
}

func TestSendFollowedByReceiveGivesOriginalValue(t *testing.T) {
	t.Parallel()
	c := NewChannel[int](1)
	want := 99
	c.Send(want)
	got := c.Receive()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestSendsReturns3AfterThreeSendOperations(t *testing.T) {
	t.Parallel()
	c := NewChannel[float64](3)
	c.Send(1.0)
	c.Send(2.0)
	c.Send(3.0)
	want := 3
	got := c.Sends()
	if want != got {
		t.Fatalf("want %d sends, got %d", want, got)
	}
}

func TestReceivesReturns2AfterTwoSendOperations(t *testing.T) {
	t.Parallel()
	c := NewChannel[struct{}](1)
	c.Send(struct{}{})
	_ = c.Receive()
	c.Send(struct{}{})
	_ = c.Receive()
	want := 2
	got := c.Receives()
	if want != got {
		t.Fatalf("want 2 receives, got %d", got)
	}
}

func TestChannelDoesNotPanicOnManyConcurrentSendsAndReceives(t *testing.T) {
	t.Parallel()
	c := NewChannel[string](10)
	want := 100
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < want; i++ {
			c.Send("hello")
			_ = c.Receives()
		}
		wg.Done()
	}()
	for i := 0; i < want; i++ {
		_ = c.Receive()
		_ = c.Sends()
	}
	wg.Wait()
	got := c.Sends()
	if want != got {
		t.Errorf("want %d sends, got %d", want, got)
	}
	got = c.Receives()
	if want != got {
		t.Errorf("want %d receives, got %d", want, got)
	}
}

func TestPushOneValueToEmptyStackLeavesTheStackWithLength1(t *testing.T) {
	t.Parallel()
	s := Stack[int]{}
	s.Push(0)
	if s.Len() != 1 {
		t.Fatal("Push didn't add value to stack")
	}
}

func TestPushTwiceThenPopTwiceOnEmptyStackLeavesTheStackEmpty(t *testing.T) {
	t.Parallel()
	s := Stack[string]{}
	s.Push("a", "b")
	if s.Len() != 2 {
		t.Fatal("Push didn't add all values to stack")
	}
	s.Pop()
	want := "a"
	got, ok := s.Pop()
	if !ok {
		t.Fatal("Pop returned not ok on non-empty stack")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestPopReturnsNotOKOnEmptyStack(t *testing.T) {
	t.Parallel()
	s := Stack[int]{}
	_, ok := s.Pop()
	if ok {
		t.Fatal("Pop returned ok on empty stack")
	}
}

func TestFuncMap_AppliesDoubleTo2AndReturns4(t *testing.T) {
	t.Parallel()
	fm := FuncMap[int, int]{
		"double": func(i int) int {
			return i * 2
		},
		"addOne": func(i int) int {
			return i + 1
		},
	}
	want := 4
	got := fm.Apply("double", 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFuncMap_AppliesUpperToUppercaseInputAndReturnsTrue(t *testing.T) {
	t.Parallel()
	fm := FuncMap[rune, bool]{
		"upper": unicode.IsUpper,
		"lower": unicode.IsLower,
	}
	got := fm.Apply("upper", 'A')
	if !got {
		t.Fatal("upper('A'): want true, got false")
	}
}
