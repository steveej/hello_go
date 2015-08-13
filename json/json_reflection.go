package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type T struct {
	A string
	B string
}

var j1 string = `name:lan,mode:l3,ipam:{routes:[{dst:10.0.99.0/16,gw:192.168.0.99}]}`
var j2 string = `{name:lan,mode:l3,ipam:{routes:[{dst:fe81::/16,gw:fe80:1}]}}`

func main() {
	fmt.Println(T{})
	var blub interface{}
	re := regexp.MustCompile(`([/\w\d\.]+)([:,}])`)
	jPrepped := re.ReplaceAllString("{"+j1+"}", `"${1}"${2}`)
	fmt.Println(jPrepped)
	err := json.Unmarshal([]byte(jPrepped), &blub)
	fmt.Println(err, blub)
}
