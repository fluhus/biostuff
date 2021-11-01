package newick

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestMarshalText(t *testing.T) {
	tests := []struct {
		input *Node
		want  string
	}{
		{&Node{"hi", 0, nil}, "hi;"},
		{&Node{"", 0, []*Node{
			{"aaa", 11, nil},
			{"", 22, []*Node{
				{"bb", 23, []*Node{
					{"B", 0, nil},
				}},
				{"bbbb", 25, nil},
			}},
			{"c", 33, nil},
		}}, "(aaa:11,((B)bb:23,bbbb:25):22,c:33);"},
	}
	for _, test := range tests {
		if got, err := test.input.MarshalText(); err != nil ||
			string(got) != test.want {
			t.Errorf("%v.Newick()=%q,%v, want %q", test.input, got, err, test.want)
		}
	}
}

func TestNextToken(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{"(", []string{"("}},
		{"(((", []string{"(", "(", "("}},
		{"( (  (", []string{"(", "(", "("}},
		{"AAAA", []string{"AAAA"}},
		{"AA(BB", []string{"AA", "(", "BB"}},
		{
			"(a:3.14,B,:6)GGG;", []string{
				"(", "a", ":", "3.14", ",", "B", ",", ":", "6", ")", "GGG", ";"},
		},
		{
			" (   a:3.14, B, :6)  GGG ; ", []string{
				"(", "a", ":", "3.14", ",", "B", ",", ":", "6", ")", "GGG", ";"},
		},
	}

	for _, test := range tests {
		var got []string
		var err error
		var token string
		r := NewReader(strings.NewReader(test.input))
		for token, err = r.nextToken(); err == nil; token, err = r.nextToken() {
			got = append(got, token)
		}
		if err != io.EOF {
			t.Fatalf("Reader(%q).nextToken() failed: %v",
				test.input, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Reader(%q).nextToken()=%v, want %v",
				test.input, got, test.want)
		}
	}
}

func TestReader(t *testing.T) {
	tests := []struct {
		input string
		want  *Node
	}{
		{";", &Node{"", 0, nil}},
		{"();", &Node{"", 0, []*Node{{}}}},
		{":4;", &Node{"", 4, nil}},
		{"AAA;", &Node{"AAA", 0, nil}},
		{"AAA:1.23;", &Node{"AAA", 1.23, nil}},
		{"(A,(B,C))D;", &Node{"D", 0, []*Node{
			{"A", 0, nil}, {"", 0, []*Node{{"B", 0, nil}, {"C", 0, nil}}},
		}}},
		{"(A,BB,CCC);", &Node{"", 0,
			[]*Node{{"A", 0, nil}, {"BB", 0, nil}, {"CCC", 0, nil}}}},
		{"  (aaa:11,(    ( B)bb: 23,  bbbb:25):22,  c:33  )  ;  ",
			&Node{"", 0, []*Node{
				{"aaa", 11, nil},
				{"", 22, []*Node{
					{"bb", 23, []*Node{
						{"B", 0, nil},
					}},
					{"bbbb", 25, nil},
				}},
				{"c", 33, nil},
			}}},
	}
	for _, test := range tests {
		got, err := NewReader(strings.NewReader(test.input)).Read()
		if err != nil {
			t.Fatalf("Read(%q) failed: %v", test.input, err)
		}
		if !reflect.DeepEqual(got, test.want) {
			t.Fatalf("Read(%q)=%v, want %v", test.input, got, test.want)
		}
	}
}

func TestReader_multi(t *testing.T) {
	input := "AAA;BB:123;:321;"
	want := []*Node{
		{"AAA", 0, nil},
		{"BB", 123, nil},
		{"", 321, nil},
	}
	var got []*Node
	var node *Node
	var err error
	r := NewReader(strings.NewReader(input))
	for node, err = r.Read(); err == nil; node, err = r.Read() {
		got = append(got, node)
	}
	if err != io.EOF {
		t.Fatalf("Read(%q) failed: %v", input, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Read(%q)=%v, want %v", input, got, want)
	}
}

func TestReader_bad(t *testing.T) {
	inputs := []string{
		"AAA",
		");",
		"());",
		"AAA:;",
		"(AAA:);",
		"AAA::123;",
		"AAA AAA;",
		"AAA:123 AAA;",
	}
	for _, input := range inputs {
		got, err := NewReader(strings.NewReader(input)).Read()
		if err == nil {
			t.Errorf("Read(%q)=%v, want error", input, got)
		}
	}
}
