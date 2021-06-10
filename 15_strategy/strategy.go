package strategy

import "fmt"

// 一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。
// 创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。策略对象改变 context 对象的执行算法。
// 支付对象
type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy // 支付策略(现金支付、银行卡支付)
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

func NewPayment(name, cardid string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardid,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

// 支付策略接口
type PaymentStrategy interface {
	Pay(*PaymentContext)
}

// 现金支付
type Cash struct{}

func (*Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by cash", ctx.Money, ctx.Name)
}

// 银行卡支付
type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s by bank account %s", ctx.Money, ctx.Name, ctx.CardID)

}


// operation 策略接口
type OperationStrategy interface {
	DoOperation(num1, num2 int) int
}

// 实现接口的实体类
type OperationAdd struct{}
func (oa *OperationAdd) DoOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSubtract struct{}
func (os *OperationSubtract) DoOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationMultiply struct{}
func (om *OperationMultiply) DoOperation(num1, num2 int) int {
	return num1 * num2
}

type Calculation struct {
	strategy OperationStrategy
}

func (c *Calculation) SetStrategy(s OperationStrategy) {
	c.strategy = s
}

func (c *Calculation) ExecuteStrategy(num1, num2 int) {
	c.strategy.DoOperation(num1, num2)
}

// 构造函数
func NewCalculation(strategy OperationStrategy) *Calculation {
	return &Calculation{strategy: strategy}
}