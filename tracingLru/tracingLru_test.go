package tracingLru

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Race_MockRelationshipIDs(t *testing.T) {
	relationshipIDs := MockStandardRelationshipIDs(rawSpanIDs)
	index := make([]int, len(relationshipIDs), len(relationshipIDs))

	for key, value := range relationshipIDs {
		for i := 0; i < len(rawSpanIDs); i++ {
			if value == rawSpanIDs[i] {
				index[key] = i
				break
			}
		}
	}

	for i := 0; i < len(index); i++ {
		if index[i] != 0 && i != 0 {
			require.Less(t, index[i], i)
		}
	}
}
