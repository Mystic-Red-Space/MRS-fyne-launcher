package main

import (
    "encoding/json"
    "io/ioutil"
    "net/http"
)

type mcfile struct {
    Name string `json:"name"`
    Dir  string `json:"dir"`
    Url  string `json:"url"`
    Md5  string `json:"md5"`
}

type modpack struct {
    Name    string `json:"name"`
    Icon    string `json:"icon"`
    Profile string `json:"profile"`
}

func getJson(url string, format interface{}) {
    resp, _ := http.Get(url)
    res, _ := ioutil.ReadAll(resp.Body)
    _ = json.Unmarshal(res, &format)
}

func mcLogin(id string, password string) {

}
