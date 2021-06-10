package singleton

import "sync"

//Singleton 是单例模式类
type Singleton struct{}

var (
	singleton     *Singleton
	singletonOnce sync.Once
)

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	singletonOnce.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
