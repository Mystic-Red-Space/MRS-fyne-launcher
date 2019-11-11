package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type mcfile struct {
    Name string `json:"name"`
    Dir  string `json:"dir"`
    Url  string `json:"url"`
    Md5  string `json:"md5"`
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

type Agent struct {
    Name    string `json:"name"`
    Version int    `json:"version"`
}

func NewAgent() Agent {
    newagent := Agent{}
    newagent.Name = "Minecraft"
    newagent.Version = 1
    return newagent
}

type LoginPayload struct {
    Agent       Agent  `json:"agent"`
    Username    string `json:"username"`
    Password    string `json:"password"`
    RequestUser bool   `json:"requestUser"`
}

func NewLoginPayload(id string, pass string) LoginPayload {
    newloginpayload := LoginPayload{}
    newloginpayload.Agent = NewAgent()
    newloginpayload.Username = id
    newloginpayload.Password = pass
    newloginpayload.RequestUser = true
    return newloginpayload
}

type modpack struct {
    Name    string `json:"name"`
    Icon    string `json:"icon"`
    Profile string `json:"profile"`
}

func GetJson(url string, format interface{}) {
    resp, _ := http.Get(url)
    res, _ := ioutil.ReadAll(resp.Body)
    _ = json.Unmarshal(res, &format)
}

func McLogin(id string, password string) {
    payload := NewLoginPayload(id, password)
    request, _ := json.Marshal(payload)
    buff := bytes.NewBuffer(request)
    resp, err := http.Post("https://authserver.mojang.com/authenticate", "application/json", buff)
    check(err)
    responce, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(responce))
}

func main() {
    McLogin("doohee006@gmail.com", "dhdh4321@@!!")
}
