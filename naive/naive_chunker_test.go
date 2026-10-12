package naive

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func TestNaiveChunker(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		text        string
		chunkSize   int
		overlapSize int
		chunks      []string
	}{
		{
			name:        "cyrillic",
			text:        "Львы — представители семейства кошачьих",
			chunkSize:   10,
			overlapSize: 2,
			chunks: []string{
				"Львы — пре",
				"редставите",
				"тели семей",
				"ейства кош",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chunker := NewChunker(WithChunkSize(tt.chunkSize), WithOverlapSize(tt.overlapSize))
			chunks := chunker.Split(tt.text)
			require.Equal(t, tt.chunks, chunks)
		})
	}
}
