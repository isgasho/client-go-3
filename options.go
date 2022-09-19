package client_go

import "time"

type OptFn func(options *options)

type options struct {
	connTimeOut time.Duration
	// add more options
}

func WithConnTimeout(seconds time.Duration) OptFn {
	return func(options *options) {
		if seconds > 0 {
			options.connTimeOut = seconds
		}
	}
}
