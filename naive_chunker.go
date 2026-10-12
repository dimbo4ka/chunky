package chunky

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

var _ Chunker = (*naiveChunker)(nil)

func NewNaiveChunker(optFuncs ...naiveGeneratorOptFunc) Chunker {
	opts := newDefaultNaiveGeneratorOptions()
	for _, f := range optFuncs {
		f(opts)
	}

	return newNaiveChunkerWithOpts(opts)
}

func newNaiveChunkerWithOpts(opts *naiveGeneratorOptions) *naiveChunker {
	generator := &naiveChunker{
		chunkSize:   opts.ChunkSize,
		overlapSize: opts.OverlapSize,
	}
	return generator
}
