package model

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

// Test_Check_analyzeJSON tests JSON analysis, sorting, code generation for unmarshalling.
func Test_Check_analyzeJSON(t *testing.T) {
	// Unmarshal JSON into a map.
	var m map[string]interface{}
	err := json.Unmarshal(jsonTracingLog, &m)
	require.NoError(t, err)

	// Conducting research and analysis.
	an := Analyze{}
	an.NewAnalyze(m, "tData")

	// Generate code for unmarshalling.
	codes := generateUnmarshal(an)
	fmt.Println(codes)
}
