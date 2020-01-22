package lstags

type Tag interface {
	GetFilename() string
	GetLine() int
	GetColumn() int
	GetSignature() string
}
