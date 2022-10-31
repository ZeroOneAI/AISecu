package debug

import (
	"fmt"
	"github.com/ZeroOneAI/AISecu/pkg/utils/datastruct/stack"
	"runtime"
	"strings"
	"sync"
)

type Debugger struct {
	fnStack *stack.Stack[string]
	mu      sync.Mutex
}

func NewDebugger() *Debugger {
	return &Debugger{
		fnStack: stack.New[string](),
	}
}

func (d *Debugger) FnStart() {
	pc, _, _, ok := runtime.Caller(1)
	fnName := "unknown function"
	if ok {
		details := runtime.FuncForPC(pc)
		fnName = details.Name()
	}
	d.mu.Lock()
	fmt.Println(strings.Repeat("	", d.fnStack.Length()) + fnName + " start")
	d.fnStack.Push(fnName)
	d.mu.Unlock()
}

func (d *Debugger) FnEnd() {
	d.mu.Lock()
	fnName, _ := d.fnStack.Pop()
	fmt.Println(strings.Repeat("	", d.fnStack.Length()) + fnName + " end")
	d.mu.Unlock()
}
