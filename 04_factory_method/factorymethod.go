package factorymethod

//OperatorFactory 是工厂接口
type OperatorFactory interface {
	Create() Operator
}

//PlusOperatorFactory 是 PlusOperator 的工厂类
type PlusOperatorFactory struct{}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		ob: &OperatorBase{},
	}
}

//MinusOperatorFactory 是 MinusOperator 的工厂类
type MinusOperatorFactory struct{}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

type OwerOperatorFactory struct{}

func (OwerOperatorFactory) Crete() Operator {
	return &QwerOperator {
		OperatorBase: OperatorBase{},
	}
}

// 具体的产品实现 接口Operator
// 每种产品都有具体的 工厂方法: 通过create方法延迟 具体对象的创建 Create() Operator
// 抽象工厂接口

/*
抽象产品（Product）	具体产品的父类	描述具体产品的公共接口
具体产品（Concrete Product）	抽象产品的子类；工厂类创建的目标类	描述生产的具体产品

抽象工厂（Creator）	具体工厂的父类	描述具体工厂的公共接口
具体工厂（Concrete Creator）	抽象工厂的子类；被外界调用	描述具体工厂；实现FactoryMethod工厂方法创建产品的实例
*/

//Operator 是被封装的实际类接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

//OperatorBase 是Operator 接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

//SetA 设置 A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

//SetB 设置 B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// 聚合/组合/关联
// PlusOperator Operator 的实际加法实现
type PlusOperator struct {
	ob *OperatorBase
}

func (o PlusOperator) SetA(a int) {
	o.ob.a = a
}

func (o PlusOperator) SetB(b int) {
	o.ob.b = b
}

//Result 获取结果
func (o PlusOperator) Result() int {
	return o.ob.a + o.ob.b
}

// 匿名聚合(继承)
//MinusOperator Operator 的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

//Result 获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}

// 匿名组合(继承)
type QwerOperator struct {
	OperatorBase
}

//Result 获取结果
func (o QwerOperator) Result() int {
	return o.a - o.b
}

