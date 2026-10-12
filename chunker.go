package chunky

type Chunker interface {
	Split(string) []string
}
