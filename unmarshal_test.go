package gocsv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Generic1Struct struct {
	field1 string `csv:"1"`
}

func TestUnmarshalGeneric1(t *testing.T) {

	file := "a;b;;c;15\nx;;y;z;20"

	var data []Generic1Struct
	err := Unmarshal([]byte(file), ";", &data, EncodingUTF8, TimezoneUTC)

	assert.Nil(t, err)
}
