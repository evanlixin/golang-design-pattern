package observer

import (
	"context"
	"errors"
	"fmt"
)

// 被观察者
type Subject struct {
	observers []Observer  // handler
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make([]Observer, 0),
	}
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}


// 观察者
type Observer interface {
	Update(*Subject)
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

type Reader struct {
	name string
}

func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}

// handler 注入联动  数据传输，对象注入、函数注入
type SubjectX struct {
	handler Handler
	context string
}

func (sx *SubjectX) Update(s string) {
	sx.context = s
	sx.handler.Handle(context.Background(), []byte(s))
}

type TestHandler struct { }

func (t *TestHandler) HandleQueueEvent(ctx context.Context, data []byte) {
	fmt.Println(data)
}

// EventHandler discovery event handler interface
type EventHandler func(svcs []string)     // 函数注入传递数据
/*
// RegisterEventHandler register event callback function
func (md *ModuleDiscovery) RegisterEventHandler(eHandler EventHandler) {
	md.handler = eHandler
}
*/

/*

// EventHandler callback function when server changes
type EventHandler func(module string)  注入、交互观察者    以对象交互方式注入

type MicroDiscovery struct {
	ctx      context.Context
	stopFunc context.CancelFunc
	//event callback
	eventHandler  EventHandler
}
*/

// Handler handle event to data and inject consumer
type Handler interface {
	Name() string
	Handle(ctx context.Context, data []byte) error
}

// HandlerWrap function for Handler interface
func HandlerWrap(name string, f func(ctx context.Context, data []byte) error) *HandlerWrapper {
	return &HandlerWrapper{name, f}
}

// HandlerWrapper for Handler
type HandlerWrapper struct {
	NameValue string
	Impl      func(ctx context.Context, data []byte) error
}

// Handle of hw
func (hw *HandlerWrapper) Handle(ctx context.Context, data []byte) error {
	if hw == nil {
		return errors.New("nil handler")
	}
	return hw.Impl(ctx, data)
}

// Name of the handler
func (hw *HandlerWrapper) Name() string {
	if hw == nil {
		return "nil"
	}
	return hw.NameValue
}


