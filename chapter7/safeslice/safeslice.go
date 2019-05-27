package safeslice

type safeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{} // 关闭通道并返回切片
	Delete(int)
	Len() int
	Update(int, UpdateFunc) // 使用自定义的更新函数更新指定位置的元素
}

type UpdateFunc func(interface{}) interface{}

type commandData struct {
	action commandAction
	value interface{}
	index int
	result chan<- interface{} // 存放带有返回值操作的返回值
	data chan<- []interface{} // 保留底层切片
	updateFunc UpdateFunc     // 切片元素更新逻辑
}

type commandAction int

// slice操作类型
const (
	Append commandAction = iota
	At
	Delete
	Len
	Update
	Close
)

type SafeSlice chan commandData

func (SafeSlice SafeSlice) Append(value interface{}) {
	SafeSlice <- commandData{action:Append, value:value}
}

func (SafeSlice SafeSlice) At(index int) interface{} {
	resultChan := make(chan interface{})
	SafeSlice <- commandData{action:At, index:index, result:resultChan}
	result := <-resultChan
	return result
}

func (SafeSlice SafeSlice) Delete(index int) {
	SafeSlice <- commandData{action:Delete, index:index}
}

func (SafeSlice SafeSlice) Len() int {
	resultChan := make(chan interface{})
	SafeSlice <- commandData{action:Len, result:resultChan}
	return (<-resultChan).(int)
}

func (SafeSlice SafeSlice) Update(index int, updateFunc UpdateFunc) {
	SafeSlice <- commandData{action:Update, updateFunc:updateFunc}
}

func (SafeSlice SafeSlice) Close() []interface{} {
	dataChan := make(chan []interface{})
	SafeSlice <- commandData{action:Close, data:dataChan}
	return <-dataChan
}

func (SafeSlice SafeSlice) run() {
	// 真正的底层slice
	store := make([]interface{}, 0)

	for command := range SafeSlice {
		switch command.action {
		case Append:
			store = append(store, command.value)
		case At:
			command.result <- store[command.index]
		case Delete:
			store = append(store[:command.index], store[command.index+1:]...)
		case Len:
			command.result <- len(store)
		case Update:
			store[command.index] = command.updateFunc(store[command.index])
		case Close:
			command.data <- store
			close(SafeSlice)
		}
	}
}

func New() SafeSlice {
	safeSlice := make(SafeSlice)
	go safeSlice.run()
	return safeSlice
}