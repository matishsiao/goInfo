package goInfo

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	info, err := GetInfo()
	assert.NoError(t, err)
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.GoOS))
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.Kernel))
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.Core))
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.Platform))
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.OS))
	assert.Equal(t, reflect.TypeOf(""), reflect.TypeOf(info.Hostname))
	assert.Equal(t, reflect.TypeOf(0), reflect.TypeOf(info.CPUs))
}
