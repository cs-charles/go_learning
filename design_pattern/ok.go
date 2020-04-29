package design_pattern

/**
一个表达式返回 2 个参数时使用这种模式：，ok，第一个参数是一个值或者 nil，第二个参数是 true/false
或者一个错误 error。在一个需要赋值的 if 条件语句中，使用这种模式去检测第二个参数值会让代码显得优雅简洁。

*/

//检测map中是否含有某个key值
/**
if value, isPresent = map1[key1]; isPresent {
        Process(value)
}
*/

//检测一个接口是否含有某个类型I
/**
if value, ok := varI.(T); ok {
Process(value)
}
*/

//检测一个通道是否关闭
/**
for input := range ch {
        Process(input)
    }
}

或
for {
        if input, open := <-ch; !open {
            break // 通道是关闭的
        }
        Process(input)
    }

*/
