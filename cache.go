package formula

import (
	"github.com/yidane/formula/internal/cache"
	"github.com/yidane/formula/opt"
)

func Register(f *opt.Function) error {
	return cache.Register(f)
}

func RegisterGlobalParameter(name string, value interface{}) {

}
