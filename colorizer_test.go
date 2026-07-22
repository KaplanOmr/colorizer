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

func TestBrightForegroundColors(t *testing.T) {
	tests := []struct {
		name  string
		color Color
		code  int
	}{
		{name: "black", color: BrightBlack, code: 90},
		{name: "red", color: BrightRed, code: 91},
		{name: "green", color: BrightGreen, code: 92},
		{name: "yellow", color: BrightYellow, code: 93},
		{name: "blue", color: BrightBlue, code: 94},
		{name: "magenta", color: BrightMagenta, code: 95},
		{name: "cyan", color: BrightCyan, code: 96},
		{name: "white", color: BrightWhite, code: 97},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := fmt.Sprintf("\x1b[%dmtext\x1b[0m", tt.code)
			if got := Paint("text", tt.color); got != want {
				t.Fatalf("Paint() = %q, want %q", got, want)
			}
		})
	}
}

func TestBrightBackgroundColors(t *testing.T) {
	tests := []struct {
		name       string
		background Background
		code       int
	}{
		{name: "black", background: BgBrightBlack, code: 100},
		{name: "red", background: BgBrightRed, code: 101},
		{name: "green", background: BgBrightGreen, code: 102},
		{name: "yellow", background: BgBrightYellow, code: 103},
		{name: "blue", background: BgBrightBlue, code: 104},
		{name: "magenta", background: BgBrightMagenta, code: 105},
		{name: "cyan", background: BgBrightCyan, code: 106},
		{name: "white", background: BgBrightWhite, code: 107},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := fmt.Sprintf("\x1b[%dmtext\x1b[0m", tt.code)
			if got := Paint("text", tt.background); got != want {
				t.Fatalf("Paint() = %q, want %q", got, want)
			}
		})
	}
}

func TestCombinedBrightStyle(t *testing.T) {
	const want = "\x1b[93;104;1mnotice\x1b[0m"
	if got := Paint("notice", BrightYellow, BgBrightBlue, Bold); got != want {
		t.Fatalf("Paint() = %q, want %q", got, want)
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
