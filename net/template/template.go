package template

import (
	"fmt"
	"html/template"
	"os"
)

/**
通过执行 template 将模板与数据结构合并，在大多数情况下是一个结构体或者一个结构体的切片。
它可以重写一段文本，通过传递给 templ.Execute () 的数据项进行替换生成新的内容。
只有能被导出的数据项才可以用于模板合并。操作可以是数据评估或控制结构，并通过 「{{」和「}}」定义。
*/

type Person struct {
	Name                string
	nonExportedAgeField string // 非导出字段
}

//## 字段替代: {{.FieldName}}
// {{.}} 直接显示两个参数
//|html 部分告诉 template 引擎在输出 FieldName 的值之前要通过 html 格式化它。
//他会转义特殊的 html 字符（ 如：会将 > 替换成 &gt; ）, 这样可以防止用户的数据破坏 HTML 表单。
func TestTemplate() {
	t := template.New("hello")

	t, _ = t.Parse("hello {{.Name}}!")

	p := Person{Name: "Mary", nonExportedAgeField: "31"} // data

	if err := t.Execute(os.Stdout, p); err != nil {

		fmt.Println("There was an error:", err.Error())

	}
}

//检查模板的语法是否定义正确，对 Parse 的结果执行 Must 函数
//如果出现错误时会出现一个运行时的panic
func TestTemplateMust() {
	template.Must(template.New("right").Parse("/*这是一个注释 */ some static text: {{ .Name }}"))
	template.Must(template.New("error").Parse("some static text {{ .Name }"))
}

//If-else
func TestIfElse() {
	tEmpty := template.New("template test")

	// if 是一个空管道时的内容
	tEmpty = template.Must(tEmpty.
		Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))

	tEmpty.Execute(os.Stdout, nil)

	tWithValue := template.New("template test")

	// 如果条件满足，则为非空管道
	tWithValue = template.Must(tWithValue.
		Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n"))

	tWithValue.Execute(os.Stdout, nil)

	tIfElse := template.New("template test")

	// 如果条件满足，则为非空管道

	tIfElse = template.Must(tIfElse.
		Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))

	tIfElse.Execute(os.Stdout, nil)
}

//.与with-end
//with 语句将点的值设置为管道的值。如果管道是空的，就会跳过 with 到 end 之前的任何内容；当嵌套使用时，点会从最近的范围取值。在下面这个程序中会说明：
func TestWithEnd() {
	t := template.New("test")

	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")

	t.Execute(os.Stdout, nil)

	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}} {{end}}!\n")

	t.Execute(os.Stdout, nil)
}

//##模板变量 $
func TestLocalVar() {
	t := template.New("test")

	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))

	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))

	t.Execute(os.Stdout, nil)

	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}} {{end}}!\n"))

	t.Execute(os.Stdout, nil)

}

//##Range-end
//range 在循环的集合中使用： 管道的值必须是一个数组、切片或者 map 。如果管道的值的长度为零，点不会被影响并且 T0 会被执行；否则将点设置成拥有连续元素的数组、切片或者 map， T1 就会被执行。
//{{range pipeline}} T1 {{else}} T0 {{end}}
