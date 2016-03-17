package regx_test

import (
	"fmt"

	"github.com/britannic/blacklist/regx"
)

func ExampleCompile() {
	rx := regx.Regex()
	fmt.Println(rx)
	// Output: &{^(?:[\/*]+)(.*?)(?:[*\/]+)$ ^(?:description)+\s"?([^"]+)?"?$ ^(?:disabled)+\s([\S]+)$ ^(?:address=[/][.]{0,1}.*[/])(.*)$ \b((?:(?:[^.-/]{0,1})[a-zA-Z0-9-_]{1,63}[-]{0,1}[.]{1})+(?:[a-zA-Z]{2,63}))\b ^(?:address=[/][.]{0,1})(.*)(?:[/].*)$ (?:^(?:http|https){1}:)(?:\/|%2f){1,2}(.*) ^(source)+\s([\S]+)\s[{]{1}$ [{] ^([\w-]+)$ ^((?:include|exclude)+)\s([\S]+)$ ^$ ^([\w-]+)\s["']{0,1}(.*?)["']{0,1}$ ^([\w-]+)\s[{]{1}$ [}] (?:#.*|\{.*|[/[].*)\z}
}
