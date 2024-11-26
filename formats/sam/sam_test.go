package sam

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func TestReaderHeader(t *testing.T) {
	input := "@a\n@b\n" +
		"c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\n" +
		"f\t6\tg\t10\t60\t4D\th\t70\t80\tTCTC\t!!!!\n"
	wantHeader1 := "@a"
	wantHeader2 := "@b"
	wantSAM1 := &SAM{
		"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
		map[string]interface{}{},
	}
	wantSAM2 := &SAM{
		"f", 6, "g", 10, 60, "4D", "h", 70, 80, "TCTC", "!!!!",
		map[string]interface{}{},
	}

	want := []SAMOrHeader{
		{H: &wantHeader1},
		{H: &wantHeader2},
		{S: wantSAM1},
		{S: wantSAM2},
	}

	var got []SAMOrHeader
	for sh, err := range ReaderHeader(bytes.NewBufferString(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, sh)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestReader(t *testing.T) {
	input := "@a\n@b\n" +
		"c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\n" +
		"f\t6\tg\t10\t60\t4D\th\t70\t80\tTCTC\t!!!!\n"

	want := []*SAM{
		{
			"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
			map[string]interface{}{},
		}, {
			"f", 6, "g", 10, 60, "4D", "h", 70, 80, "TCTC", "!!!!",
			map[string]interface{}{},
		},
	}

	var got []*SAM
	for sm, err := range Reader(bytes.NewBufferString(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, sm)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestDecoder_tags(t *testing.T) {
	input := "c\t2\td\t5\t30\t32M\te\t40\t50\tAAAA\tFFFF\t" +
		"BC:Z:barcode\tAS:i:123\tZF:f:3.1415\tZH:H:1234abcd"

	want := []*SAM{{
		"c", 2, "d", 5, 30, "32M", "e", 40, 50, "AAAA", "FFFF",
		map[string]interface{}{
			"BC": "barcode",
			"AS": 123,
			"ZF": 3.1415,
			"ZH": []byte{18, 52, 171, 205},
		},
	}}

	var got []*SAM
	for sm, err := range Reader(bytes.NewBufferString(input)) {
		if err != nil {
			t.Fatalf("Reader(%q) failed: %v", input, err)
		}
		got = append(got, sm)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Reader(%q)=%v, want %v", input, got, want)
	}
}

func TestText(t *testing.T) {
	input := "GGCGTT\t0\tbvu:BVU_3729\t38\t255\t24M\t*\t0\t0\t" +
		"FADFNAKNNKKNLHDCNEYMNNDE\t*AS:i:44\tMD:Z:14G3C5\tNM:i:2\t" +
		"ZE:f:1.07e-05\tZF:i:-3\tZI:i:91\tZL:i:116\tZR:i:104\tZS:i:72\n"

	var sm *SAM
	for s, err := range Reader(bytes.NewBufferString(input)) {
		if err != nil {
			t.Fatalf("Next() failed: %v", err)
		}
		sm = s
	}
	if got, err := sm.MarshalText(); err != nil || string(got) != input {
		t.Fatalf("%v.Text()=%v,%v want %v", sm, got, err, input)
	}
}

func TestFlag(t *testing.T) {
	s := &SAM{}
	if got := s.Flag.Multiple(); got {
		t.Errorf("SAM{}.Multiple()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Each(); got {
		t.Errorf("SAM{}.Each()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Unmapped(); got {
		t.Errorf("SAM{}.Unmapped()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Unmapped2(); got {
		t.Errorf("SAM{}.Unmapped2()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.ReverseComplement(); got {
		t.Errorf("SAM{}.ReverseComplement()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.ReverseComplement2(); got {
		t.Errorf("SAM{}.ReverseComplement2()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.First(); got {
		t.Errorf("SAM{}.First()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Last(); got {
		t.Errorf("SAM{}.Last()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Secondary(); got {
		t.Errorf("SAM{}.Secondary()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.NotPassing(); got {
		t.Errorf("SAM{}.NotPassing()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Duplicate(); got {
		t.Errorf("SAM{}.Duplicate()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Supplementary(); got {
		t.Errorf("SAM{}.Supplementary()=true, want false. flag=%v", s.Flag)
	}

	s.Flag.SetEach(true)
	s.Flag.SetNotPassing(false)
	s.Flag.SetUnmapped2(true)
	s.Flag.SetSupplementary(false)

	if got := s.Flag.Multiple(); got {
		t.Errorf("SAM{}.Multiple()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Each(); !got {
		t.Errorf("SAM{}.Each()=false, want true. flag=%v", s.Flag)
	}
	if got := s.Flag.Unmapped(); got {
		t.Errorf("SAM{}.Unmapped()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Unmapped2(); !got {
		t.Errorf("SAM{}.Unmapped2()=false, want true. flag=%v", s.Flag)
	}
	if got := s.Flag.ReverseComplement(); got {
		t.Errorf("SAM{}.ReverseComplement()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.ReverseComplement2(); got {
		t.Errorf("SAM{}.ReverseComplement2()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.First(); got {
		t.Errorf("SAM{}.First()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Last(); got {
		t.Errorf("SAM{}.Last()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Secondary(); got {
		t.Errorf("SAM{}.Secondary()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.NotPassing(); got {
		t.Errorf("SAM{}.NotPassing()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Duplicate(); got {
		t.Errorf("SAM{}.Duplicate()=true, want false. flag=%v", s.Flag)
	}
	if got := s.Flag.Supplementary(); got {
		t.Errorf("SAM{}.Supplementary()=true, want false. flag=%v", s.Flag)
	}
}

func BenchmarkWrite(b *testing.B) {
	text := "GGCGTT\t0\tbvu:BVU_3729\t38\t255\t24M\t*\t0\t0\t" +
		"FADFNAKNNKKNLHDCNEYMNNDE\t*AS:i:44\tMD:Z:14G3C5\tNM:i:2\t" +
		"ZE:f:1.07e-05\tZF:i:-3\tZI:i:91\tZL:i:116\tZR:i:104\tZS:i:72\n"
	var sm *SAM
	for s, err := range Reader(bytes.NewBufferString(text)) {
		if err != nil {
			b.Fatal(err)
		}
		sm = s
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sm.Write(io.Discard)
	}
}
