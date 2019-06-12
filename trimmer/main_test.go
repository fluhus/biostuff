package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_Main(t *testing.T) {
	for i := range testCases {
		// Create input file
		in, err := ioutil.TempFile(".", "trimmer_test_")
		if err != nil {
			t.Fatalf("Could not open input file in test #%d.", i+1)
		}
		defer os.Remove(in.Name())

		_, err = in.WriteString(testCases[i].input)
		if err != nil {
			t.Fatalf("Could not write to input file in test #%d.", i+1)
		}
		in.Close()

		// Create output file
		out, err := ioutil.TempFile(".", "trimmer_test_")
		if err != nil {
			t.Fatalf("Could not open output file in test #%d.", i+1)
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
			t.Fatalf("Could not read output file in test #%d.", i+1)
		}

		if string(outputText) != testCases[i].output {
			t.Log("adapterEnd:", string(adapterEnd))
			t.Fatalf("Bad output in test #%d. Expected:\n%s\nActual:\n%s",
				i+1, testCases[i].output, string(outputText))
		}
	}
}

type testCase struct {
	args   []string // program arguments
	input  string   // fastq input
	output string   // expected fastq output
}

var testCases = []testCase{
	{
		[]string{"trimmer", "-l", "1"},
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
	{
		[]string{"trimmer", "-l", "1", "-q", "0", "-ae", "GGTTATGAC"},
		`@lalala
TCTCATCTGGTTGGTTA
+
**IIIIIIIIIIII***
`,
		`@lalala
TCTCATCTGGTT
+
**IIIIIIIIII
`,
	},
	{
		[]string{"trimmer", "-l", "1", "-q", "0", "-as", "AACCGTCTCA"},
		`@lalala
TCTCATCTGGTTGGTTA
+
**IIIIIIIIIIII***
`,
		`@lalala
TCTGGTTGGTTA
+
IIIIIIIII***
`,
	},
	{
		[]string{"trimmer", "-l", "1", "-q", "0", "-as", "AACCGGGTCA"},
		`@lalala
TCTCATCTGGTTGGTTA
+
**IIIIIIIIIIII***
`,
		`@lalala
TCTCATCTGGTTGGTTA
+
**IIIIIIIIIIII***
`,
	},
	{
		[]string{"trimmer", "-l", "1", "-q", "0", "-as", "AACCGTCTCA"},
		`@lalala
tctcatctggttggtta
+
**IIIIIIIIIIII***
`,
		`@lalala
tctggttggtta
+
IIIIIIIII***
`,
	},
}
