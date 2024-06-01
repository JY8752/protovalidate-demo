package main

import (
	"fmt"
	hellov1 "protovalidate-demo/gen/example/hello/v1"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func validate(v *protovalidate.Validator, msg protoreflect.ProtoMessage) error {
	if err := v.Validate(msg); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	return nil
}

func main() {
	v, err := protovalidate.New()
	if err != nil {
		panic(err)
	}

	var errorList []error

	// if err = validate(v, &hellov1.Hello{}); err != nil {
	// 	errorList = append(errorList, err)
	// }

	// if err = validate(v, &hellov1.DisabledExample{}); err != nil {
	// 	errorList = append(errorList, err)
	// }

	// if err = validate(v, &hellov1.OneofExample{}); err != nil {
	// 	errorList = append(errorList, err)
	// }

	if err = validate(v, &hellov1.StringValidationExample{
		ConstValue:             "const",
		LenValue:               "Hello",
		MinLenValue:            "Hello!!",
		MaxLenValue:            "Hi",
		LenBytesValue:          "ab",
		MinBytesValue:          "abc",
		MaxBytesValue:          "ab",
		PatternValue:           "hello, world!!!",
		PrefixValue:            "Hello, World!!",
		SuffixValue:            "Hello, World",
		ContainsValue:          "apple, banana, orange",
		NotContainsValue:       "apple, orange",
		InValue:                "Go",
		NotInValue:             "Rust",
		EmailValue:             "protovalidate@example.com",
		HostnameValue:          "example.com",
		IpValue:                "255.255.255.255",
		Ipv4Value:              "127.0.0.1",
		Ipv6Value:              "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
		UriValue:               "https://example.com?id=xxxx",
		UriRefValue:            "./example",
		AddressValue:           "example.com",
		UuidValue:              "550e8400-e29b-41d4-a716-446655440000",
		TuuidValue:             "550e8400e29b41d4a716446655440000",
		IpWithPreifxlenValue:   "255.255.255.0/24",
		Ipv4WithPreifxlenValue: "255.255.255.0/24",
		Ipv6WithPreifxlenValue: "2001:0db8:85a3:0000:0000:8a2e:0370:7334/24",
		IpPrefixValue:          "127.0.0.0/16",
		Ip4PrefixValue:         "127.0.0.0/16",
		Ip6PrefixValue:         "2001:db8::/48",
		HostAndPortValue:       "127.0.0.1:3000",
		WellKownRegexValue:     "Contetnt Type",
	}); err != nil {
		errorList = append(errorList, err)
	}

	if len(errorList) != 0 {
		for _, e := range errorList {
			fmt.Println(e.Error())
		}
	} else {
		fmt.Println("Succeded validation!!")
	}
}
