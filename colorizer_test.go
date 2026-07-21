package colorizer

import (
	"fmt"
	"sync"
	"testing"
)

func TestPaint(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		options []interface{}
		want    string
	}{
		{
			name: "no options",
			text: "plain",
			want: "plain",
		},
		{
			name:    "foreground",
			text:    "error",
			options: []interface{}{Red},
			want:    "\x1b[31merror\x1b[0m",
		},
		{
			name:    "foreground background and attributes",
			text:    "warning",
			options: []interface{}{Yellow, BgRed, Bold, Underline},
			want:    "\x1b[33;41;1;4mwarning\x1b[0m",
		},
		{
			name:    "empty text",
			options: []interface{}{Green},
			want:    "\x1b[32m\x1b[0m",
		},
		{
			name:    "unsupported options are ignored",
			text:    "text",
			options: []interface{}{struct{}{}, "red"},
			want:    "text",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Paint(tt.text, tt.options...); got != tt.want {
				t.Fatalf("Paint() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStyleBuilder(t *testing.T) {
	style := New().
		WithColor(Cyan).
		WithBackground(BgBlack).
		WithAttribute(Dim)

	const want = "\x1b[36;40;2mmessage\x1b[0m"
	if got := style.Paint("message"); got != want {
		t.Fatalf("Style.Paint() = %q, want %q", got, want)
	}
}

func TestReusableStyleConcurrentPaint(t *testing.T) {
	style := New(Green, Bold)
	const want = "\x1b[32;1mok\x1b[0m"

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if got := style.Paint("ok"); got != want {
				t.Errorf("Style.Paint() = %q, want %q", got, want)
			}
		}()
	}
	wg.Wait()
}

func ExamplePaint() {
	fmt.Println(Paint("Success!", Green, Bold))
	// Output: [32;1mSuccess![0m
}
