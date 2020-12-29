package cs2

// Impl ...
type Impl struct {
	level int
}

// NewImpl ...
func NewImpl(
	level int,
) *Impl {
	return &Impl{
		level: level,
	}
}
