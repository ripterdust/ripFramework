package responser

import "ripframework/ripFramework/constants"

type Responser struct {
	status int
}

func New() *Responser {
	responser := &Responser{status: constants.HTTP_SUCCESS}

	return responser
}
