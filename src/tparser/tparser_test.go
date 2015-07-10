package tparser

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
)

func TestScanner_simple(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader("chr1\t100\t200\t3.14")
	s := NewScanner(r, &thing{})
	s.Scan()
	obj, err := s.Obj()
	assert.Nil(err)
	assert.Equal(&thing{"chr1", 100, 200, 3.14}, obj)
}

func TestScanner_longLine(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader("chr1\t100\t200\t3.14\thello\tworld")
	s := NewScanner(r, &thing{})
	s.Scan()
	obj, err := s.Obj()
	assert.Nil(err)
	assert.Equal(&thing{"chr1", 100, 200, 3.14}, obj)
}

func TestScanner_shortLine(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader("chr1\t100\t200")
	s := NewScanner(r, &thing{})
	s.Scan()
	_, err := s.Obj()
	assert.NotNil(err)
}

func TestScanner_multiline(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader("chr1\t100\t200\t3.14\nchr2\t400\t500\t5.66\n")
	s := NewScanner(r, &thing{})
	s.Scan()
	obj, err := s.Obj()
	assert.Nil(err)
	assert.Equal(&thing{"chr1", 100, 200, 3.14}, obj)
	s.Scan()
	obj, err = s.Obj()
	assert.Nil(err)
	assert.Equal(&thing{"chr2", 400, 500, 5.66}, obj)
}

func TestScanner_scanAll(t *testing.T) {
	assert := assert.New(t)
	r := strings.NewReader("chr1\t100\t200\t3.14\nchr2\t400\t500\t5.66\n")
	a, err := ScanAll(r, &thing{})
	assert.Nil(err)
	arr := a.([]*thing)
	assert.Equal([]*thing{ &thing{"chr1", 100, 200, 3.14},
			&thing{"chr2", 400, 500, 5.66} }, arr)
}

type thing struct {
	S string
	I int
	J int
	F float64
}
