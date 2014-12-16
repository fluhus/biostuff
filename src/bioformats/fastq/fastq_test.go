package fastq

import (
	"testing"
)

func Test_Trim(t *testing.T) {
	s := "\r\n\r\namitamit\n\ramit\n\n\r\n"
	ss := string( trimNewLines([]byte(s)) )
	
	if ss != "amitamit\n\ramit" {
		t.Errorf("Bad trimming result. Expected '%s', got '%s'.",
				"amitamit\n\ramit", ss)
	}
}

