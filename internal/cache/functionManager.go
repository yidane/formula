package cache

import (
	"fmt"
	"github.com/yidane/formula/opt"
	"strings"
	"sync"
)

var defaultFunctionManager = make(map[string]*opt.Function, 256)
var rwLock sync.RWMutex

func Register(f *opt.Function) error {
	if f == nil {
		return fmt.Errorf("argument Function can not be nil")
	}

	lowerName := (*f).Name()

	rwLock.Lock()
	defer rwLock.Unlock()

	if _, ok := defaultFunctionManager[lowerName]; ok {
		return fmt.Errorf("function %s duplicate", lowerName)
	}

	defaultFunctionManager[lowerName] = f

	return nil
}

func FindFunction(name string) (*opt.Function, error) {
	name = strings.ToLower(strings.TrimSpace(name))
	if f, ok := defaultFunctionManager[name]; ok {
		return f, nil
	}

	return nil, fmt.Errorf("function %s dot not exists", name)
}
