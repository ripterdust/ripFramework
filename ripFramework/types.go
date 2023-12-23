package ripframework

import (
	"ripframework/ripFramework/requester"
	"ripframework/ripFramework/responser"
)

type handleFunction func(*requester.Requester, *responser.Responser)
