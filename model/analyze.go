package model

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

// >>>>> >>>>> >>>>> Analysis Method

// Analyze represents a structure used for analyzing JSON data.
type Analyze struct {
	VarName string         // VarName stores the name of the variable used in JSON unmarshalling.
	Collect []SubAnalyze   // Collect holds a collection of SubAnalyze instances.
	Block   []AnalyzeBlock // Block stores a collection of Json Block.
}

// SubAnalyze represents a sub-analysis of JSON data.
type SubAnalyze struct {
	Key     string   // Key represents each key of the JSON object.
	Path    []string // Path is a list of strings representing the path to this JSON keys.
	VarType string   // VarType indicates each data type of the JSON object.
	VarPath string   // VarPath stores the path for accessing the element of the JSON variable.
	ShowKey bool     // ShowKey is a flag indicating whether to display the key.
}

// AnalyzeBlock represents block areas of JSON data.
type AnalyzeBlock struct {
	BlockName string
}

// >>>>> >>>>> >>>>> Sorting Method

// SubAnalyzeList is a custom type representing a list of SubAnalyze instances.
type SubAnalyzeList []SubAnalyze

// Len returns the length of the SubAnalyzeList.
func (s SubAnalyzeList) Len() int {
	return len(s)
}

// Less returns true if the SubAnalyze at index i is considered less than the SubAnalyze at index j.
func (s SubAnalyzeList) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

// Swap swaps the SubAnalyze instances at indices i and j within the SubAnalyzeList.
func (s SubAnalyzeList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// NewAnalyze is a method of the Analyze type used to initialize and analyze JSON data.
// It takes a map representing JSON data and a prefix string as input.
func (an *Analyze) NewAnalyze(data map[string]interface{}, prefix string) {
	// Create a new empty slice to collect SubAnalyze instances.
	collect := new([]SubAnalyze)

	// Initialize and populate the collect slice with SubAnalyze instances.
	NewSubAnalyze(data, prefix, collect)

	// Sort the collect slice based on the 'Key' field of SubAnalyze instances.
	sort.Sort(SubAnalyzeList(*collect))

	// Set the 'Collect' field of the Analyze instance to the sorted collect slice.
	an.Collect = *collect

	// Mark the analysis as complete.
	an.mark()
}

// NewSubAnalyze iterates through JSON data, gathering keys with type info recursively.
func NewSubAnalyze(data map[string]interface{}, prefix string, collect *[]SubAnalyze) {
	// Iterate through each key-value pair in the data map
	for key, value := range data {
		// Create a new prefix by combining the current prefix and key
		newPrefix := fmt.Sprintf("%s.%s", prefix, key)

		// Separate the path and the key value, and it will be easier to handle later on.
		arr := strings.Split(newPrefix, ".")

		// Determine the type of the current value using reflection
		valueType := reflect.TypeOf(value)

		// Create a JSONInfo struct to store information about the current key
		subAnalyze := SubAnalyze{
			Path:    arr[0 : len(arr)-1],
			VarPath: newPrefix,
			Key:     arr[len(arr)-1],
			ShowKey: true,
		}

		// If the value type is not nil, assign its string representation to the info struct
		if valueType != nil {
			subAnalyze.VarType = valueType.String()
		}

		// Append the info struct to the result slice
		*collect = append(*collect, subAnalyze)

		// [recursive function ↩️]
		// Check the type of the value and recursively analyze nested maps or arrays
		switch value.(type) {
		case map[string]interface{}:
			NewSubAnalyze(value.(map[string]interface{}), newPrefix, collect)
		case []interface{}:
			NewSubAnalyze(value.([]interface{})[0].(map[string]interface{}), newPrefix, collect)
		}
	}
}

// markShowKey marks duplicated JSON keys and adds AnalyzeBlocks based on certain conditions.
func (an *Analyze) mark() {
	// Initialize a variable to store the previous JSON key.
	var previousKey string

	// Add an initial AnalyzeBlock named "Enter_NotAnywhere" to the Block slice.
	an.Block = append(an.Block, AnalyzeBlock{BlockName: "Enter_NotAnywhere"})

	// Add an initial AnalyzeBlock named "Enter_NotAnywhere" to the Block slice.
	for i := 0; i < len(an.Collect); i++ {
		// Check if the current key is the same as the previous key.
		if previousKey == an.Collect[i].Key {
			// If it's a duplicate, set ShowKey too false to mark it.
			an.Collect[i].ShowKey = false
		}
		// Update the previousKey with the current key for the next iteration.
		previousKey = an.Collect[i].Key

		// Check if ShowKey is true and the VarType is "map[string]interface {}"
		// or "[]interface {}". If true, add an AnalyzeBlock to the Block slice.
		if an.Collect[i].ShowKey == true &&
			(an.Collect[i].VarType == "map[string]interface {}" || an.Collect[i].VarType == "[]interface {}") {
			an.Block = append(an.Block, AnalyzeBlock{BlockName: "Enter_" + an.Collect[i].Key})
		}
	}
}

// generateUnmarshal generates code to unmarshal JSON data into an Analyze struct using a template.
func generateUnmarshal(an Analyze) string {

	// define a template string
	tmpl := `

import (
	"strconv"
)

const (
	// first, split the JSON string into regions.
	// edit here ! add "= iota"
    {{- range .Block }}
    {{ .BlockName }}
    {{- end }}
)

// UnmarshalByTuning is an automatically generated function that parses JSON data
// and populates the provided TracingData structure.
func UnmarshalByTuning(jsonTracingLog []byte, tData *TracingData) (err error) {
	// Set the position variable first.
	var firstLayerKey, secondLayerKey int
	var countCurlyBrace int
	var countBracket int

	// Using DetectJsonString to extract the key from the JSON trace log.
	var positionNext, keyTail, keyLength int

	// Here, you don't need positionNext++ because 'detect' will perform the movement.
	for positionNext = 0; positionNext < len(jsonTracingLog); {

		// Fetch the key value first.
		positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)

		// Extract the key from the trace log.
		key := string(jsonTracingLog[(keyTail - keyLength):keyTail])
		// fmt.Println(key)

		switch key {
		{{- range .Collect }}
		{{- if eq .ShowKey true }}
		case "{{ .Key }}":
		{{- end }}
			if firstLayerKey == Enter_NotAnywhere {
				// >>>>> Variable Path: {{ .VarPath }}
				// >>>>> Variable Type: {{ .VarType }}

				{{- if and (ne .VarType "float64") (ne .VarType "string") }}
				/*
				// edit here !
				firstLayerKey = Enter_{{ .Key }}
				if firstLayerKey != Enter_ &&
				firstLayerKey != Enter_ &&
				firstLayerKey != Enter_ {
				// If not in the above those areas
				firstLayerKey = Enter_
				}
				*/
				{{- end }}

				{{- if eq .VarType "float64" }}
				if firstLayerKey == Enter_ (editor here ! NotAnywhere ?) {
					positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
					{{ .VarPath }}, err = strconv.Atoi(string(jsonTracingLog[(keyTail - keyLength):keyTail]))
					if err != nil {
						panic("Parsing {{ .Key }} error: " + err.Error())
					}
				}
				fmt.Println("     📎 ----- ----- ----- ----- -----")
				fmt.Println("     📎Key:", key)
				fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
				{{- end }}

				{{- if eq .VarType "string" }}
				if firstLayerKey == Enter_ (editor here ! NotAnywhere ?) {
					positionNext, keyTail, keyLength = DetectJsonElement(positionNext, jsonTracingLog)
					{{ .VarPath }} = string(jsonTracingLog[(keyTail - keyLength):keyTail])
					if err != nil {
						panic("Parsing ChildSpanCount error: " + err.Error())
					}
				}
				fmt.Println("     📎 ----- ----- ----- ----- -----")
				fmt.Println("     📎Key:", key)
				fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
				{{- end }}
			}
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
		{{- end }}
		case "{":
			// edit here ! ... ...
			countCurlyBrace++
			fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
		case "}":
			// edit here ! ... ...
			countCurlyBrace--
			fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
		case "[":
			// edit here ! ... ...
			countBracket++
			fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
		case "]":
			// edit here ! ... ...
			countBracket--
			fmt.Println("     📎 ----- ----- ----- ----- -----")
			fmt.Println("     📎Key:", key)
			fmt.Println("     📎Position:", firstLayerKey, secondLayerKey, countCurlyBrace, countBracket)
		}
	}
	return
}
`
	// create a new template with the name
	t, err := template.New("tracingTemplate").Parse(tmpl)
	if err != nil {
		panic("Error creating template:" + err.Error())
		return ""
	}

	// create a new strings.Builder to store the generated code
	var code strings.Builder
	err = t.Execute(&code, an)
	if err != nil {
		panic("Error executing template:" + err.Error())
		return ""
	}

	// return the generated code as a string
	return code.String()
}
