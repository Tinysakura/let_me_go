package my_oslice

import (
	"bytes"
	"errors"
)

type Slice []interface{}

var Less func(interface{}, interface{}) bool

func New(less func(interface{}, interface{}) bool) *Slice {
	Less = less

	return &Slice{}
}

func NewStringSlice() *Slice {
	Less = func(e1 interface{}, e2 interface{}) bool {
		return e1.(string) < e2.(string)
	}

	return &Slice{}
}

func NewIntSlice() *Slice {
	Less = func(e1 interface{}, e2 interface{}) bool {
		return e1.(int) < e2.(int)
	}

	return &Slice{}
}

func (Slice *Slice) Clear() {
	*Slice = []interface{}{}
}

func (Slice *Slice) Add(index int, value interface{}) {
	result := make([]interface{}, len(*Slice) + 1)
	at := copy(result, (*Slice)[:index])
	at += copy(result[at:], []interface{}{value})
	copy(result[at:], (*Slice)[index:])

	(*Slice).Clear()
	*Slice = append(*Slice, result...)
}

func (Slice *Slice) Remove(value interface{}) bool {
	index := (*Slice).Index(value)

	if index == -1 {
		return false
	}

	*Slice = append((*Slice)[:index], (*Slice)[index+1:]...)

	return true
}

func (Slice *Slice) Index(value interface{}) int {
	for index, element := range *Slice {
		if element == value {
			return index
		}
	}

	return -1
}

func (Slice *Slice) At(index int) (interface{}, error) {
	if len(*Slice) - 1 < index {
		return nil, errors.New("index out of bounds")
	}

	return (*Slice)[index], nil
}

func (Slice *Slice) Len() int {
	return len(*Slice)
}

func (Slice *Slice) String() string {
	buffer := bytes.NewBufferString("[")

	for index, element := range *Slice {
		s, e := element.(string)
		if e {
			if index != len(*Slice) - 1 {
				buffer.WriteString(s)
				buffer.WriteString(",")
			} else {
				buffer.WriteString(s)
			}
		}
	}

	buffer.WriteString("]")
	return buffer.String()
}


