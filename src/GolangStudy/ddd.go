package main
 
import (
    "encoding/json"
    "errors"
    "fmt"
)
 
var parseJsonError = errors.New("json parse error")
var toJsonError = errors.New("to json error")
 
 
//反序列化为map
func parseMap(a string) (map[string]interface{}, error) {
    // fmt.Printf("原始字符串: %s\n", a)
    var m map[string]interface{}
    if err := json.Unmarshal([]byte(a), &m); err != nil {
        fmt.Println("Error =", err)
        return m, parseJsonError
    }
    for k, v := range m {
        fmt.Printf("k=%s,v类型为%T,v=%v\n", k, v, v)
    }
    return m, nil
}
 
func main() {
    j := `{"time":"12:48:21","src_ip":"10.0.0.2","method":"GET","url":"http:\/\/123.58.224.8:30377\/?id=cat%20\/flag","data":"","waf":true,"404":false}`
    if m, e := parseMap(j); e == nil {
        fmt.Printf("转换map为: %v\n", m)
    }
}