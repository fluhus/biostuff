package main

import (
	"testing"
	"io/ioutil"
	"os"
)

func Test_Main(t *testing.T) {
	for i := range testCases {
		// Create input file
		in, err := ioutil.TempFile(".", "trimmer_test_")
		if err != nil {
			t.Fatalf("Could not open input file in test #%d.", i)
		}
		defer os.Remove(in.Name())
		
		_, err = in.WriteString(testCases[i].input)
		if err != nil {
			t.Fatalf("Could not write to input file in test #%d.", i)
		}
		in.Close()
		
		// Create output file
		out, err := ioutil.TempFile(".", "trimmer_test_")
		if err != nil {
			t.Fatalf("Could not open output file in test #%d.", i)
		}
		defer os.Remove(out.Name())
		out.Close()
		
		// Set arguments
		os.Args = append(testCases[i].args, "-in", in.Name(),
				"-out", out.Name())
		
		// Execute program
		main()
		
		// Compare output
		outputText, err := ioutil.ReadFile(out.Name())
		if err != nil {
			t.Fatalf("Could not read output file in test #%d.", i)
		}
		
		if string(outputText) != testCases[i].output {
			t.Fatalf("Bad output in test #%d. Expected:\n%s\nActual:\n%s",
					i, testCases[i].output, string(outputText))
		}
	}
}

type testCase struct {
	args   []string    // program arguments
	input  string      // fastq input
	output string      // expected fastq output
}

var testCases = []testCase {
{
[]string{"trimmer"},
`@lalala
TCTCATCTGGTTGGTTA
+
**IIIIIIIIIIII***
`,
`@lalala
TCATCTGGTTGG
+
IIIIIIIIIIII
`,
},
}
