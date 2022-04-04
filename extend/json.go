package extend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type JsonExamples struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	School string `json:"school"`
}

func textJson() {
	var jex JsonExamples
	var unmarshalData JsonExamples
	var decodeData JsonExamples

	jex = JsonExamples{
		Name:   "go",
		Age:    10,
		School: "goole",
	}
	// 创建Json 数据
	by, _ := json.Marshal(jex)
	fmt.Printf("examples %v type %T", string(by), by)
	// 解压数据
	json.Unmarshal(by, &unmarshalData)
	fmt.Printf("unmarshalData %v", unmarshalData)
	fmt.Printf("unmarshalData %v", unmarshalData.Age)
	// 编码JSON数据并以缩进方式返回
	indent_bt_data, _ := json.MarshalIndent(jex, "", " ")
	fmt.Println("indent_bt_data = ", string(indent_bt_data))

	// 将编码完的JSON数据保存到文件
	json_w_fd, _ := os.OpenFile("./logs/out.json", os.O_CREATE, 0666)
	json_encoder := json.NewEncoder(json_w_fd)
	json_encoder.Encode(jex)
	json_w_fd.Close()
	fmt.Println("======================>>>>")
	// 将编码完的JSON数据合并到其他数据中
	dst := bytes.NewBuffer(indent_bt_data)
	json.Indent(dst, by, "", " ")
	fmt.Println(dst.String())
	// 从文件中读取Json 数据
	json_rd_fd, _ := os.Open("./logs/out.json")
	json_decode := json.NewDecoder(json_rd_fd)
	json_decode.Decode(&decodeData)
	fmt.Println("======================>>>>")
	fmt.Println(decodeData)
	json_rd_fd.Close()
}
