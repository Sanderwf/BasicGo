package reflect

import (
	"fmt"
	"reflect"
)

func DisplayType(t reflect.Type, tab string) {
	if t == nil {
		fmt.Printf("<nil>")
		return
	}
	switch t.Kind() {

	case reflect.Int, reflect.Float32, reflect.Bool, reflect.String:
		fmt.Printf("%s%s", tab, t.Name())

	case reflect.Array, reflect.Slice:
		fmt.Printf("%s%s", tab, t)

	case reflect.Map:
		fmt.Printf("%smap{\n", tab)
		fmt.Printf("%s\tKey: ", tab)
		fmt.Printf("%s%s: ", tab, t.Key())
		fmt.Println()
		fmt.Printf("%s\tValue: ", tab)
		fmt.Printf("%s%s: ", tab, t.Elem())
		fmt.Println()
		fmt.Printf("%s}", tab)

	case reflect.Func:
		fmt.Printf("%sfunc (", tab)
		// 遍历参数
		for i := 0; i < t.NumIn(); i++ {
			fmt.Printf("%s", t.In(i))
			if i != t.NumIn()-1 {
				fmt.Printf(", ")
			}
		}
		// 可变参数
		if t.IsVariadic() {
			fmt.Printf("...")
		}
		fmt.Printf(") ")

		// 遍历返回值
		if t.NumOut() > 0 {
			fmt.Printf("(")
			for i := 0; i < t.NumOut(); i++ {
				fmt.Printf("%s", t.Out(i))
				if i != t.NumIn()-1 {
					fmt.Printf(", ")
				}
			}
			fmt.Printf(") ")
		}
		fmt.Printf("{}")

	case reflect.Struct:
		fmt.Printf("%stype %s struct {\n", tab, t.Name())

		// 遍历属性
		fmt.Printf("%s\tFields(%d)\n", tab, t.NumField())
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fmt.Printf("%s\t\t%s\t%s\t`%s`", tab, field.Name, field.Tag)
			fmt.Printf(",\n")
		}

		// 遍历方法
		fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			DisplayMethod(t.Method(i), tab+"\t\t")
			fmt.Printf(",\n")
		}
		fmt.Printf("%s}", tab)

	case reflect.Ptr:
		fmt.Printf("%s*{\n", tab)
		DisplayType(t.Elem(), tab+"\t")
		// 打印指针变量的方法
		if t.NumMethod() > 0 {
			fmt.Printf("\n%s\tMethods(%d):\n", tab, t.NumMethod())
			for i := 0; i < t.NumMethod(); i++ {
				DisplayMethod(t.Method(i), tab+"\t\t")
				if i != t.NumMethod()-1 {
					fmt.Printf(",\n")
				}
			}
		}
		fmt.Printf("\n%s", tab)

	default:
		fmt.Printf("-------------------------")
	}
}

func DisplayMethod(method reflect.Method, tab string) {
	t := method.Type
	fmt.Printf("%sfunc %s(", tab, method.Name)

	// 遍历参数
	for i := 0; i < t.NumIn(); i++ {
		fmt.Printf("%s", t.In(i))
		if i != t.NumIn()-1 {
			fmt.Printf(", ")
		}
	}
	if t.IsVariadic() {
		fmt.Printf("...")
	}

	// 遍历返回值
	if t.NumOut() > 0 {
		fmt.Printf("(")
		for i := 0; i < t.NumOut(); i++ {
			fmt.Printf("%s", t.Out(i))
			if i != t.NumIn()-1 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf(") ")
	}
	fmt.Printf("{}")
}
