package simplefactory

import "fmt"

// 简单工厂: 通过不同的类型, 初始化不同的接口实现
// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

/*
抽象产品（Product）	具体产品的父类	描述产品的公共接口
具体产品（Concrete Product）	抽象产品的子类；工厂类创建的目标类	描述生产的具体产品
工厂（Creator）	被外界调用	根据传入不同参数从而创建不同具体产品类的实例
*/

//API is interface
type API interface {
	Say(name string) string
}

//hiAPI is one of API implement
type hiAPI struct{}

//Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

//HelloAPI is another API implement
type helloAPI struct{}

//Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
