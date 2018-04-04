// Code generated by thriftrw v1.11.0. DO NOT EDIT.
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

// ThriftTest_TestOneway_Args represents the arguments for the ThriftTest.testOneway function.
//
// The arguments for testOneway are sent and received over the wire as this struct.
type ThriftTest_TestOneway_Args struct {
	SecondsToSleep *int32 `json:"secondsToSleep,omitempty"`
}

// ToWire translates a ThriftTest_TestOneway_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *ThriftTest_TestOneway_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.SecondsToSleep != nil {
		w, err = wire.NewValueI32(*(v.SecondsToSleep)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ThriftTest_TestOneway_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ThriftTest_TestOneway_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v ThriftTest_TestOneway_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *ThriftTest_TestOneway_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x int32
				x, err = field.Value.GetI32(), error(nil)
				v.SecondsToSleep = &x
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// String returns a readable string representation of a ThriftTest_TestOneway_Args
// struct.
func (v *ThriftTest_TestOneway_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.SecondsToSleep != nil {
		fields[i] = fmt.Sprintf("SecondsToSleep: %v", *(v.SecondsToSleep))
		i++
	}

	return fmt.Sprintf("ThriftTest_TestOneway_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this ThriftTest_TestOneway_Args match the
// provided ThriftTest_TestOneway_Args.
//
// This function performs a deep comparison.
func (v *ThriftTest_TestOneway_Args) Equals(rhs *ThriftTest_TestOneway_Args) bool {
	if !_I32_EqualsPtr(v.SecondsToSleep, rhs.SecondsToSleep) {
		return false
	}

	return true
}

// GetSecondsToSleep returns the value of SecondsToSleep if it is set or its
// zero value if it is unset.
func (v *ThriftTest_TestOneway_Args) GetSecondsToSleep() (o int32) {
	if v.SecondsToSleep != nil {
		return *v.SecondsToSleep
	}

	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "testOneway" for this struct.
func (v *ThriftTest_TestOneway_Args) MethodName() string {
	return "testOneway"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be OneWay for this struct.
func (v *ThriftTest_TestOneway_Args) EnvelopeType() wire.EnvelopeType {
	return wire.OneWay
}

// ThriftTest_TestOneway_Helper provides functions that aid in handling the
// parameters and return values of the ThriftTest.testOneway
// function.
var ThriftTest_TestOneway_Helper = struct {
	// Args accepts the parameters of testOneway in-order and returns
	// the arguments struct for the function.
	Args func(
		secondsToSleep *int32,
	) *ThriftTest_TestOneway_Args
}{}

func init() {
	ThriftTest_TestOneway_Helper.Args = func(
		secondsToSleep *int32,
	) *ThriftTest_TestOneway_Args {
		return &ThriftTest_TestOneway_Args{
			SecondsToSleep: secondsToSleep,
		}
	}

}
