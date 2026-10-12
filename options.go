package chunky

const (
	defaultChunkSize   = 512
	defautlOverlapSize = 64
)

type naiveGeneratorOptions struct {
	ChunkSize   int
	OverlapSize int
}

type naiveGeneratorOptFunc func(opts *naiveGeneratorOptions)

func newDefaultNaiveGeneratorOptions() *naiveGeneratorOptions {
	return &naiveGeneratorOptions{
		ChunkSize:   defaultChunkSize,
		OverlapSize: defautlOverlapSize,
	}
}

func WithChunkSize(chunkSize int) naiveGeneratorOptFunc {
	return func(opts *naiveGeneratorOptions) {
		opts.ChunkSize = chunkSize
	}
}

func WithOverlapSize(overlapSize int) naiveGeneratorOptFunc {
	return func(opts *naiveGeneratorOptions) {
		opts.OverlapSize = overlapSize
	}
}
