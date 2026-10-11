package chunky

type naiveChunker struct {
}

// Split implements [Chunker].
func (n *naiveChunker) Split() {
	panic("unimplemented")
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
	generator := &naiveChunker{}
	return generator
}
