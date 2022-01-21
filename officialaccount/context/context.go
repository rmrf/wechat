package context

import (
	"github.com/rmrf/wechat/v2/credential"
	"github.com/rmrf/wechat/v2/officialaccount/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
