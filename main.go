package main

import (
	"fmt"
	hellov1 "protovalidate-demo/gen/example/hello/v1"
	"time"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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
		WellKownRegexValue:     "Contetnt-Type",
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.BoolValidationExample{
		TrueValue:  true,
		FalseValue: false,
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.BytesValidationExample{
		ConstValue:    []byte{1, 2, 3, 4},
		LenValue:      []byte{1, 2, 3, 4},
		MinLenValue:   []byte{1, 2, 3},
		MaxLenValue:   []byte{1, 2},
		PatternValue:  []byte{'\x61'},
		PrefixValue:   []byte{'\x01', '\x02', '\x03'},
		SuffixValue:   []byte{'\x01', '\x02', '\x03'},
		ContainsValue: []byte{'\x01', '\x02', '\x03'},
		InValue:       []byte{'\x02', '\x03'},
		NotInValue:    []byte{'\x01', '\x02', '\x03'},
		IpValue:       []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\xFF', '\xFF', '\xFF', '\xFF', '\xFF', '\x00'},
		Ipv4Value:     []byte{'\xFF', '\xFF', '\xFF', '\x00'},
		Ipv6Value:     []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\xFF', '\xFF', '\xFF', '\xFF', '\xFF', '\x00'},
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.DoubleValidationExample{
		ConstValue:  42.0,
		LtValue:     9.0,
		LteValue:    10.0,
		GtValue:     11.0,
		GteValue:    10.0,
		InValue:     11.0,
		NotInValue:  13.0,
		FiniteValue: 1.0,
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.EnumValidationExample{
		ConstValue:       hellov1.EnumValidationExample_MY_ENUM_VALUE1,
		DefinedOnlyValue: hellov1.EnumValidationExample_MY_ENUM_VALUE1,
		InValue:          hellov1.EnumValidationExample_MY_ENUM_VALUE1,
		NotInValue:       hellov1.EnumValidationExample_MY_ENUM_VALUE3,
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.MapValidationExample{
		MinPairsValue: map[string]string{"key1": "value1", "key2": "value2"},
		MaxPairsValue: map[string]string{"key1": "value1", "key2": "value2"},
		KeysValue:     map[string]string{"key1": "value1"},
		ValuesValue:   map[string]string{"key1": "value1"},
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.RepeatedValidationExample{
		MinItemsValue: []string{"elm1", "elm2"},
		MaxItemsValue: []string{"elm1", "elm2"},
		UniqueValue:   []string{"elm1", "elm2", "elm3"},
		ItemsValue:    []string{"abcdefghi"},
	}); err != nil {
		errorList = append(errorList, err)
	}

	intVal, err := anypb.New(wrapperspb.Int32(123))
	if err != nil {
		panic(err)
	}

	boolVal, err := anypb.New(wrapperspb.Bool(true))
	if err != nil {
		panic(err)
	}

	if err = validate(v, &hellov1.AnyValidationExample{
		InValue:    intVal,
		NotInValue: boolVal,
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.DurationValidationExample{
		ConstValue: durationpb.New(5 * time.Second),
		LtValue:    durationpb.New(4 * time.Second),
		LteValue:   durationpb.New(5 * time.Second),
		GtValue:    durationpb.New(6 * time.Second),
		GteValue:   durationpb.New(5 * time.Second),
		InValue:    durationpb.New(5 * time.Second),
		NotInValue: durationpb.New(8 * time.Second),
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.TimestampValidationExample{
		ConstValue:  timestamppb.New(time.Date(2024, 6, 3, 12, 0, 0, 0, time.UTC)),
		LtValue:     timestamppb.New(time.Date(2024, 6, 3, 11, 0, 0, 0, time.UTC)),
		LteValue:    timestamppb.New(time.Date(2024, 6, 3, 12, 0, 0, 0, time.UTC)),
		LtNowValue:  timestamppb.New(time.Now().UTC()),
		GtValue:     timestamppb.New(time.Date(2024, 6, 3, 13, 0, 0, 0, time.UTC)),
		GteValue:    timestamppb.New(time.Date(2024, 6, 3, 12, 0, 0, 0, time.UTC)),
		GtNowValue:  timestamppb.New(time.Now().UTC().Add(time.Hour)),
		WithinValue: timestamppb.New(time.Now().UTC()),
	}); err != nil {
		errorList = append(errorList, err)
	}

	if err = validate(v, &hellov1.FieldConstraintsExample{
		EvenValue:             2,
		RequiredMessageValue:  &hellov1.FieldConstraintsExample_MyValue{},
		RequiredStringValue:   "a",
		RequiredInt32Value:    1,
		RequiredEnumValue:     hellov1.FieldConstraintsExample_STATUS_OK,
		RequiredRepeatedValue: []string{"a"},
		RequiredMapValue:      map[string]string{"a": "a"},
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
