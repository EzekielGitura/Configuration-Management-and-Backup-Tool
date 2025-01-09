package utils

import (
    "bytes"
    "testing"
)

func TestPrintVersion(t *testing.T) {
    // Capture stdout
    old := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w

    // Call the function
    PrintVersion()

    // Restore stdout
    w.Close()
    os.Stdout = old

    // Read the output
    out, _ := ioutil.ReadAll(r)
    output := string(out)

    expectedOutput := "Networking Tool v1.0.0\n"
    if output != expectedOutput {
        t.Fatalf("Output mismatch: got %s, want %s", output, expectedOutput)
    }
}