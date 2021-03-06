// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/als/als.proto

package als

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *AccessLoggingService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLoggingService)
	if !ok {
		that2, ok := that.(AccessLoggingService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetAccessLog()) != len(target.GetAccessLog()) {
		return false
	}
	for idx, v := range m.GetAccessLog() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetAccessLog()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetAccessLog()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *AccessLog) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AccessLog)
	if !ok {
		that2, ok := that.(AccessLog)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.OutputDestination.(type) {

	case *AccessLog_FileSink:

		if h, ok := interface{}(m.GetFileSink()).(equality.Equalizer); ok {
			if !h.Equal(target.GetFileSink()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetFileSink(), target.GetFileSink()) {
				return false
			}
		}

	case *AccessLog_GrpcService:

		if h, ok := interface{}(m.GetGrpcService()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGrpcService()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGrpcService(), target.GetGrpcService()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *FileSink) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*FileSink)
	if !ok {
		that2, ok := that.(FileSink)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetPath(), target.GetPath()) != 0 {
		return false
	}

	switch m.OutputFormat.(type) {

	case *FileSink_StringFormat:

		if strings.Compare(m.GetStringFormat(), target.GetStringFormat()) != 0 {
			return false
		}

	case *FileSink_JsonFormat:

		if h, ok := interface{}(m.GetJsonFormat()).(equality.Equalizer); ok {
			if !h.Equal(target.GetJsonFormat()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetJsonFormat(), target.GetJsonFormat()) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *GrpcService) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*GrpcService)
	if !ok {
		that2, ok := that.(GrpcService)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetLogName(), target.GetLogName()) != 0 {
		return false
	}

	if len(m.GetAdditionalRequestHeadersToLog()) != len(target.GetAdditionalRequestHeadersToLog()) {
		return false
	}
	for idx, v := range m.GetAdditionalRequestHeadersToLog() {

		if strings.Compare(v, target.GetAdditionalRequestHeadersToLog()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAdditionalResponseHeadersToLog()) != len(target.GetAdditionalResponseHeadersToLog()) {
		return false
	}
	for idx, v := range m.GetAdditionalResponseHeadersToLog() {

		if strings.Compare(v, target.GetAdditionalResponseHeadersToLog()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetAdditionalResponseTrailersToLog()) != len(target.GetAdditionalResponseTrailersToLog()) {
		return false
	}
	for idx, v := range m.GetAdditionalResponseTrailersToLog() {

		if strings.Compare(v, target.GetAdditionalResponseTrailersToLog()[idx]) != 0 {
			return false
		}

	}

	switch m.ServiceRef.(type) {

	case *GrpcService_StaticClusterName:

		if strings.Compare(m.GetStaticClusterName(), target.GetStaticClusterName()) != 0 {
			return false
		}

	}

	return true
}
