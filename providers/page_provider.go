package providers

import "github.com/ccutch/webapp/pages"

// PageProvider is a generalized implementation of a provider which renders pages.
type PageProvider struct {
	Pages []pages.Page
}

func NewPageProvider(...pages Page) *PageProvider {
  p := &PageProvider{
    Pages: pages,
  }
  return p
}

func (p *PageProvider) Mount() handler {
  r := mux.NewRouter()
  for _, pg := range p.Pages {
    r.HandleFunc(pg.Page(), pg.Serve)    
  }
  return 
}