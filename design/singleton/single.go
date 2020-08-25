package single

import(
    "sync"
)

type singleton struct {
}

var single *singleton
var once sync.once

func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	returninstance
}
