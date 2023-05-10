package mongodb

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// This one was unfinished.
func Test_Race_mongoClient(t *testing.T) {
	client, err := New("mongodb://localhost:27017")
	require.NoError(t, err)

	defer func() {
		_ = client.Close()
	}()

	collection, err := client.PickUpDocument("openTelemetry2mongodb", "openTelemetry2mongodb")
	require.NoError(t, err)

	tracingData := collection.Search(time.Now().Add(-240*time.Hour), time.Now(), 5, TracingFilter{"instrumentationlibrary.name", "mainTracer"})

	fmt.Println(tracingData)
}
