package naive

const (
	defaultChunkSize   = 512
	defautlOverlapSize = 64
)

type Options struct {
	ChunkSize   int
	OverlapSize int
}

type OptFunc func(opts *Options)

func newDefaultOptions() *Options {
	return &Options{
		ChunkSize:   defaultChunkSize,
		OverlapSize: defautlOverlapSize,
	}
}

func WithChunkSize(chunkSize int) OptFunc {
	return func(opts *Options) {
		opts.ChunkSize = chunkSize
	}
}

func WithOverlapSize(overlapSize int) OptFunc {
	return func(opts *Options) {
		opts.OverlapSize = overlapSize
	}
}
