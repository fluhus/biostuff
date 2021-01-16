package sam

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	input := "@a\n@b\n" +
		"c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\n" +
		"f\t6\tg\t10\t60\t4D\th\t70\t80\tTCTC\t!!!!\n"
	r := NewReader(bytes.NewBuffer([]byte(input)))

	h, err := r.NextHeader()
	if err != nil {
		t.Fatalf("NextHeader() failed: %v", err)
	}
	if h != "@a" {
		t.Fatalf("NextHeader()=%v, want @a", h)
	}

	h, err = r.NextHeader()
	if err != nil {
		t.Fatalf("NextHeader() failed: %v", err)
	}
	if h != "@b" {
		t.Fatalf("NextHeader()=%v, want @b", h)
	}

	_, err = r.NextHeader()
	if err != io.EOF {
		t.Fatalf("NextHeader() error=%v, want EOF", err)
	}

	want := &SAM{
		"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
		map[string]interface{}{},
	}
	got, err := r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next()=%v, want %v", got, want)
	}

	want = &SAM{
		"f", 6, "g", 10, 60, "4D", "h", 70, 80, "TCTC", "!!!!",
		map[string]interface{}{},
	}
	got, err = r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next()=%v, want %v", got, want)
	}

	_, err = r.Next()
	if err != io.EOF {
		t.Fatalf("Next() error=%v, want EOF", err)
	}
}

func TestReader_skipHeader(t *testing.T) {
	input := "@a\n@b\n" +
		"c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\n" +
		"f\t6\tg\t10\t60\t4D\th\t70\t80\tTCTC\t!!!!\n"
	r := NewReader(bytes.NewBuffer([]byte(input)))

	want := &SAM{
		"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
		map[string]interface{}{},
	}
	got, err := r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next()=%v, want %v", got, want)
	}

	want = &SAM{
		"f", 6, "g", 10, 60, "4D", "h", 70, 80, "TCTC", "!!!!",
		map[string]interface{}{},
	}
	got, err = r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next()=%v, want %v", got, want)
	}

	_, err = r.Next()
	if err != io.EOF {
		t.Fatalf("Next() error=%v, want EOF", err)
	}
}
