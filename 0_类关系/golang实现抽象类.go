package main

import "fmt"

// 多用组合少用继承 / 基于接口编程而非实现
// 没有语法层面的抽象类支持：利用接口和组合继承，可以实现抽象类。
// 继承：可以通过匿名组合实现，多态实现

func main() {
	parent := IPeopleImpl{
		name: "IPeopleImpl",
		age:  10,
	}
	UsePeople(&parent)
}

type IPeople interface {
	GetName() string
	SetName(string)
	GetAge() int
	SetAge(int)
	Run()
}

type IPeopleImpl struct {
	name string
	age int
}
// 方法接收者指针和值的区别，简单示例: https://blog.csdn.net/K346K346/article/details/91150296
func (ip *IPeopleImpl) GetName() string { return "IPeopleImpl" }
func (ip *IPeopleImpl) SetName(newName string) { ip.name = newName }
func (ip *IPeopleImpl) GetAge() int { return ip.age }
func (ip *IPeopleImpl) SetAge(newAge int) { ip.age = newAge }
func (ip *IPeopleImpl) Run() { fmt.Printf("%s:%d\n", ip.name, ip.age)}

type AbstractPeople struct {
	IPeople
	name string
	age int
}

func (a AbstractPeople) GetName() string {
	return a.name
}

func (a *AbstractPeople) SetName(newName string) {
	a.name = newName
}

func (a AbstractPeople) GetAge() int {
	return a.age
}

func (a *AbstractPeople) SetAge(newAge int) {
	a.age = newAge
}

// user.
func UsePeople(p IPeople) {
	fmt.Println(p.GetName())
	fmt.Println(p.GetAge())
	p.SetName("Alice")
	p.SetAge(p.GetAge() + 1)
	p.Run()
}

// 方法重写(匿名组合、继承)和抽象类(接口、函数类型)
// https://blog.csdn.net/zhbinary/article/details/89418195?utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-2.control&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7EBlogCommendFromMachineLearnPai2%7Edefault-2.control

/*
type Daemon interface {
    // start(time.Duration)
    doWork()
}

func start(daemon Daemon, duration time.Duration) { ... }

func main() {
    ...
    start(dA, 1 * time.Second)
    start(dB, 5 * time.Second)
    ...
}
*/

/*
type Task interface {
    doWork()
}

type Daemon struct {
    task Task
}

func (d *Daemon) start(t time.Duration) {
    ticker := time.NewTicker(t)
    // this will call task.doWork() periodically
    go func() {
        for {
            <-ticker.C
            d.task.doWork()
        }
    }()
}

type MyTask struct{}

func (m MyTask) doWork() {
    fmt.Println("Doing my work")
}
 */

/*
type Daemon struct {
    task func()
}

func (d *Daemon) start(t time.Duration) {
    ticker := time.NewTicker(t)
    // this will call task() periodically
    go func() {
        for {
            <-ticker.C
            d.task()
        }
    }()
}

func main() {
    d := Daemon{task: func() {
        fmt.Println("Doing my work")
    }}
    d.start(time.Millisecond * 300)

    time.Sleep(time.Second * 2)
}
*/

/*
type Daemon interface {
    start(time.Duration)
    doWork()
}

// 抽象类
type AbstractDaemon struct {
    Daemon
}

func (a *AbstractDaemon) start(duration time.Duration) {
    ticker := time.NewTicker(duration)

    // this will call daemon.doWork() periodically
    go func() {
        for {
            <- ticker.C
            a.doWork()
        }
    }()
}

// 继承
type ConcreteDaemonA struct {
*AbstractDaemon
foo int
}

func newConcreteDaemonA() *ConcreteDaemonA {
  a:=&AbstractDaemon{}
  r:=&ConcreteDaemonA{a, 0}
  a.Daemon = r
  return r
}

type ConcreteDaemonB struct {
*AbstractDaemon
bar int
}

func newConcreteDaemonB() *ConcreteDaemonB {
  a:=&AbstractDaemon{}
  r:=&ConcreteDaemonB{a, 0}
  a.Daemon = r
  return r
}

func (a *ConcreteDaemonA) doWork() {
    a.foo++
    fmt.Println("A: ", a.foo)
}

func (b *ConcreteDaemonB) doWork() {
    b.bar--
    fmt.Println("B: ", b.bar)
}
*/
