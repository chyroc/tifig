package tifig_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Chyroc/tifig"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	as := assert.New(t)

	pwd, err := os.Getwd()
	as.Nil(err)

	input := fmt.Sprintf("%s/test.heic", pwd)

	{
		f, err := ioutil.TempFile("", "tifig-test-*.png")
		as.Nil(err)

		output, err := tifig.Convert(input, f.Name())
		as.Nil(err)
		as.Equal(f.Name(), output)
	}

	{
		bs, err := ioutil.ReadFile(input)
		as.Nil(err)

		f, err := ioutil.TempFile("", "tifig-test-*.png")
		as.Nil(err)

		output, err := tifig.ConvertReader(bytes.NewReader(bs), f.Name())
		as.Nil(err)
		as.Equal(f.Name(), output)
	}
}
