package utils

import(
	"crypto/rand"
	"strings"
    "context"
    "sync"
    "io/ioutil"
    "net/http"
    "time"    
)

const randomchars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
var Ctx = context.Background()

func RandomString(length int) string{
	bytes := make([]byte, length)
    //There is an entropy bug here with a lot of concurrency
    var mu sync.Mutex
    mu.Lock()
    _, err := rand.Read(bytes)
    mu.Unlock()
    if err != nil {
        //...
    }    
	
    for i, b := range bytes {
        bytes[i] = randomchars[b%byte(len(randomchars))]
    }
	return string(bytes)
}

func TrimLeftChars(s string, n int) string {
    m := 0
    for i := range s {
        if m >= n {
            return s[i:]
        }
        m++
    }
    return s[:0]
}

func RemoveQuotes(s string) string {
    if s == ""{
        return ""
    }
	s = strings.TrimSuffix(s, `"`)
	return strings.TrimPrefix(s, `"`)
}

func StringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func ArrayContains(arr []string, search string) bool{
    for _, v := range arr{
        if v == search{
            return true
        }
    }
    return false
}

/*
func GetValues(m map[string][]string) []string{
    ret := []string{}
    for _,v := range m{
        ret = append(ret, v)
    }
    return ret
}

func GetKeys(m map[string][]string) []string{
    ret := []string{}
    for k,_ := range m{
        ret = append(ret, k)
    }
    return ret
}*/

func ArraySlice(arr []interface{}, index int) []interface{}{
    copy(arr[index:], arr[index+1:])
    arr[len(arr)-1] = ""
    arr = arr[:len(arr)-1]
    return arr
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
    cp := make(map[string]interface{})
    for k, v := range m {
        vm, ok := v.(map[string]interface{})
        if ok {
            cp[k] = CopyMap(vm)
        } else {
            cp[k] = v
        }
    }

    return cp
}

func OpenFile(path string) ([]byte, error){
    var ret []byte
    if strings.HasPrefix(path, "https://"){
        client := &http.Client{
            Timeout: time.Second * 15,
        }
        req, _ := http.NewRequest("GET", path, nil)
        res, err := client.Do(req)
        if err != nil {
            return nil, err
        }
        defer res.Body.Close()
        ret, _ = ioutil.ReadAll(res.Body)
    }else{
        var err error
        ret, err = ioutil.ReadFile(path)
        if err != nil{
            return nil, err
        }
    }
    return ret, nil
}

func StripSpaces(str string) string{
    return strings.Replace(str, " ", "", -1)    
}