package ilog

const (
	DefaultFile = "itea.log"
)

type IOption interface {
	do(ILog)
}

type OptionFunc func(ILog)

func (o OptionFunc) do(l ILog)  {
	o(l)
}

func EnableRotate() IOption {
	return OptionFunc(func(l ILog) {
		if f, ok := l.(*File); ok {
			f.rotate = true
		}
	})
}

func WithFile(file string) IOption {
	return OptionFunc(func(l ILog) {
		if f, ok := l.(*File); ok {
			f.logfile = file
		}
	})
}