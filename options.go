package chunky

type naiveGeneratorOptions struct {
}

type naiveGeneratorOptFunc func(opts *naiveGeneratorOptions)

func newDefaultNaiveGeneratorOptions() *naiveGeneratorOptions {
	return &naiveGeneratorOptions{}
}
