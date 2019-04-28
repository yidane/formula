package opt

import (
	"fmt"
	"sync"
)

var defaultFunctionManager = make(map[string]*Function, 256)
var rwLock sync.RWMutex

func Register(f *Function) error {
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
