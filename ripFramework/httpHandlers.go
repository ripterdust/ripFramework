package ripframework

import (
	"ripframework/ripFramework/constants"
)

func (f *Framework) setRoute(verb string, uri string, handler handleFunction) {
	f.handlers[verb][uri] = handler
}

func (f *Framework) Get(uri string, handler handleFunction) {
	f.setRoute(constants.GET, uri, handler)
}

func (f *Framework) Post(uri string, handler handleFunction) {
	f.setRoute(constants.POST, uri, handler)
}

func (f *Framework) Put(uri string, handler handleFunction) {
	f.setRoute(constants.PUT, uri, handler)
}

func (f *Framework) Delete(uri string, handler handleFunction) {
	f.setRoute(constants.DELETE, uri, handler)
}
