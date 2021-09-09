package sam

import (
	"bytes"
	"io"
	"reflect"
	"strings"
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

func TestDecoder_tags(t *testing.T) {
	input := "c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\t" +
		"BC:Z:barcode\tAS:i:123\tZF:f:3.1415\tZH:H:1234abcd"
	r := NewReader(bytes.NewBuffer([]byte(input)))

	want := &SAM{
		"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
		map[string]interface{}{
			"BC": "barcode",
			"AS": 123,
			"ZF": 3.1415,
			"ZH": []byte{18, 52, 171, 205},
		},
	}

	got, err := r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Next()=%v, want %v", got, want)
	}
}

func TestText(t *testing.T) {
	input := "GGCGTT\t0\tbvu:BVU_3729\t38\t255\t24M\t*\t0\t0\t" +
		"FADFNAKNNKKNLHDCNEYMNNDE\t*AS:i:44\tMD:Z:14G3C5\tNM:i:2\t" +
		"ZE:f:1.07e-05\tZF:i:-3\tZI:i:91\tZL:i:116\tZR:i:104\tZS:i:72\n"
	r := NewReader(strings.NewReader(input))
	sm, err := r.Next()
	if err != nil {
		t.Fatalf("Next() failed: %v", err)
	}
	if got := sm.Text(); got != input {
		t.Fatalf("%v.Text()=%v want %v", sm, got, input)
	}
}
