package trie

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestTrie(t *testing.T) {
	tr := New()
	tr.Add([]byte("amit"))
	tr.Add([]byte("amut"))
	tr.Add([]byte("lavon"))

	has := []string{"", "a", "am", "ami", "amit", "amu", "amut",
		"l", "la", "lav", "lavo", "lavon"}
	hasnt := []string{"A", "aa", "amm", "amitt", "amat", "-"}

	for _, s := range has {
		if !tr.Has([]byte(s)) {
			t.Fatalf("has(%q)=false, want true", s)
		}
	}
	for _, s := range hasnt {
		if tr.Has([]byte(s)) {
			t.Fatalf("has(%q)=true, want false", s)
		}
	}
}

func TestDelete(t *testing.T) {
	tr := New()
	tr.Add([]byte("amit"))
	tr.Add([]byte("amut"))
	tr.Add([]byte("lavon"))

	tests := []struct {
		del     string
		wantDel bool
		want    []string
	}{
		{"amam", false, []string{"amit", "amut", "lavon"}},
		{"lavon", true, []string{"amit", "amut"}},
		{"am", true, []string{}},
	}

	for _, test := range tests {
		if tr.Delete([]byte(test.del)) != test.wantDel {
			t.Fatalf("Del(%q)=%v, want %v", test.del, !test.wantDel, test.wantDel)
		}
		if tr.Has([]byte(test.del)) {
			t.Fatalf("Del(%q)=true, want false", test.del)
		}
		for _, want := range test.want {
			if !tr.Has([]byte(want)) {
				t.Fatalf("Has(%q)=false, want true", want)
			}
		}
	}
}

func TestJSON(t *testing.T) {
	tr := New()
	tr.Add([]byte("Hello"))
	tr.Add([]byte("Henno"))
	j, err := json.Marshal(tr)
	if err != nil {
		t.Fatalf("Marshal(%v) failed: %v", tr, err)
	}
	got := New()
	if err := json.Unmarshal(j, got); err != nil {
		t.Fatalf("Unmarshal(%s) failed: %v", j, err)
	}
	if !reflect.DeepEqual(got, tr) {
		jj, _ := json.Marshal(got)
		t.Fatalf("Unmarshal(%s)=%s, want %v", j, jj, tr)
	}
}

func TestForEach(t *testing.T) {
	tr := New()
	tr.Add([]byte("alice"))
	tr.Add([]byte("alicey"))
	tr.Add([]byte("alicie"))
	tr.Add([]byte("bob"))
	tr.Add([]byte("boris"))
	tr.Add([]byte("charles"))

	want := map[string]struct{}{
		"alicey":  {},
		"alicie":  {},
		"bob":     {},
		"boris":   {},
		"charles": {},
	}
	got := map[string]struct{}{}
	tr.ForEach(func(b []byte) bool {
		got[string(b)] = struct{}{}
		return true
	})
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Trie.ForEach()=%v, want %v", got, want)
	}
}

func TestForEach_empty(t *testing.T) {
	tr := New()
	tr.ForEach(func(b []byte) bool {
		t.Errorf("Trie.ForEach yields %v, want nothing", b)
		return true
	})
}

func TestForEach_false(t *testing.T) {
	tr := New()
	tr.Add([]byte("alice"))
	tr.Add([]byte("alicey"))
	tr.Add([]byte("alicie"))
	tr.Add([]byte("bob"))
	tr.Add([]byte("boris"))
	tr.Add([]byte("charles"))

	for i := 0; i < 10; i++ {
		var got string
		tr.ForEach(func(b []byte) bool {
			got = string(b)
			return got != "bob"
		})
		if got != "bob" {
			t.Errorf("Trie.ForEach=%q, want %q", got, "bob")
		}
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, k := range []int{10, 20} {
		b.Run(fmt.Sprint(k), func(b *testing.B) {
			tr := New()
			data := make([]byte, k*b.N)
			rand.Read(data)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				tr.Add(data[i*k : (i+1)*k])
			}
		})
	}
}

func BenchmarkHas(b *testing.B) {
	tr := New()
	text := []byte("aaaaaaaaaaaaaaaaaaa")
	tr.Add(text)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tr.Has(text)
	}
}
