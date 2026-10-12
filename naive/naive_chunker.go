package naive

import "github.com/dimbo4ka/chunky"

type naiveChunker struct {
	chunkSize   int
	overlapSize int
}

func (n *naiveChunker) Split(text string) []string {
	chunks := []string{}
	parsedText := []rune(text)
	textLen := len(parsedText)
	chunkDiff := n.chunkSize - n.overlapSize

	for i := range textLen / chunkDiff {
		l := i * chunkDiff
		r := min(textLen, l+n.chunkSize)
		chunk := string(parsedText[l:r])
		chunks = append(chunks, chunk)
	}

	return chunks
}

var _ chunky.Chunker = (*naiveChunker)(nil)

func NewChunker(optFuncs ...OptFunc) chunky.Chunker {
	opts := newDefaultOptions()
	for _, f := range optFuncs {
		f(opts)
	}

	return newWithOpts(opts)
}

func newWithOpts(opts *Options) *naiveChunker {
	generator := &naiveChunker{
		chunkSize:   opts.ChunkSize,
		overlapSize: opts.OverlapSize,
	}
	return generator
}
