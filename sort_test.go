package sort

import (
	"reflect"
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	tt := []struct {
		name   string
		list   interface{}
		op     []Option
		result interface{}
	}{
		{
			name: "empty slice",
			list: nil,
			op: []Option{
				{
					Fields:    "",
					OrderType: Asc,
				},
			},
			result: nil,
		},
		{
			name: "empty ops",
			list: []struct {
				name string
			}{{name: "2"}, {name: "1"}},
			op: nil,
			result: []struct {
				name string
			}{{name: "2"}, {name: "1"}},
		},
		{
			name: "sort ptr slice",
			list: []*struct {
				name string
			}{
				{name: "lily"},
				{name: "honey"},
			},
			op: []Option{
				{
					Fields:    "name",
					OrderType: Asc,
				},
			},
			result: []*struct {
				name string
			}{
				{name: "honey"},
				{name: "lily"},
			},
		},
		{
			name: "sort struct slice",
			list: []struct {
				name string
			}{
				{name: "lily"},
				{name: "honey"},
			},
			op: []Option{
				{
					Fields:    "name",
					OrderType: Asc,
				},
			},
			result: []struct {
				name string
			}{
				{name: "honey"},
				{name: "lily"},
			},
		},
		{
			name: "sort by multi fields",
			list: []struct {
				name string
				age  int
			}{
				{name: "lily", age: 10},
				{name: "lily", age: 11},
				{name: "honey", age: 12},
				{name: "honey", age: 11},
			},
			op: []Option{
				{
					Fields:    "name",
					OrderType: Asc,
				},
				{
					Fields:    "age",
					OrderType: Desc,
				},
			},
			result: []struct {
				name string
				age  int
			}{

				{name: "honey", age: 12},
				{name: "honey", age: 11},
				{name: "lily", age: 11},
				{name: "lily", age: 10},
			},
		},
	}
	for _, test := range tt {
		t.Run(test.name, func(t *testing.T) {
			Sort(test.list, test.op...)
			if !reflect.DeepEqual(test.list, test.result) {
				t.Fatalf("%+v != %+v", test.list, test.result)
			}
		})
	}
}

func TestSortByFunc(t *testing.T) {
	list := []struct {
		name string
		age  int
	}{
		{name: "lily", age: 10},
		{name: "lily", age: 11},
		{name: "honey", age: 12},
		{name: "honey", age: 11},
	}
	result := []struct {
		name string
		age  int
	}{
		{name: "honey", age: 12},
		{name: "honey", age: 11},
		{name: "lily", age: 11},
		{name: "lily", age: 10},
	}
	ByFunc(list, func(i, j int) int {
		return strings.Compare(list[i].name, list[j].name) // 正序
	}, func(i, j int) int {
		return 0 - (list[i].age - list[j].age) // 倒序
	})

	if !reflect.DeepEqual(list, result) {
		t.Fatalf("%+v != %+v", list, result)
	}
}
