package read_write_package

import (
	"encoding/json"
	"log"
	"os"
)

//JSON 与 Go 类型对应如下：
//
//bool 对应 JSON 的 booleans
//float64 对应 JSON 的 numbers
//string 对应 JSON 的 strings
//nil 对应 JSON 的 null
//不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：
//
//JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map [string] T（T 是 json 包中支持的任何类型）
//Channel，复杂类型和函数类型不能被编码
//不支持循环数据结构；它将引起序列化进入一个无限循环
//指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

//序列化
func Marshal() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	//出于安全考虑，在 web 应用中最好使用 json.MarshalforHTML() 函数，
	//其对数据执行 HTML 转码，所以文本可以被安全地嵌在 HTML <script> 标签中。
	//序列化后的byte[]数组
	byte, err := json.Marshal(vc)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(byte)

	//序列化到文件中
	file, _ := os.OpenFile("./files/marshal_json.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	encode := json.NewEncoder(file)
	err = encode.Encode(vc)
	if err != nil {
		log.Println(err)
	}
}

//反序列化
func UnMarshal() {
	var vc VCard
	file, _ := os.Open("./files/marshal_json.txt")
	defer file.Close()
	decode := json.NewDecoder(file)
	err := decode.Decode(&vc)
	if err != nil {
		log.Println(err)
	}
	log.Printf("unmarshal value is %s", vc)

	var buf []byte
	_, err1 := file.Read(buf)
	if err1 != nil {
		log.Println(err1)
	}
	json.Unmarshal(buf, &vc)
	log.Printf("unmarshal value is %s", vc)
}
