package browser

import "github.com/skratchdot/open-golang/open"

type Browser struct{}

func New() *Browser {
	return &Browser{}
}

func (b *Browser) Open(url string) error {
	return open.Start(url)
}
