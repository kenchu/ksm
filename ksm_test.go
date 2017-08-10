package ksm

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"github.com/easonlin404/ksm/rsa"
	"encoding/hex"
)

type spcContainerTest struct {
	filePath string
	size     int

	iv string
	spec string
}

var spcContainerTests = []spcContainerTest{
	{"testdata/spc1.bin", 2688,
	"5d1644eaec11f983147541e46eeb2774",
		"926648b9861ec0471ba21758851c3dda31c93b1dd601aa4ead4415a20759aab9a6d89f551385856e73571729df2f1d46d25c13da2ad75d00fd3413ebd96ca47d02955c569f7fab40f1a7fb23414167a653eabdf1ad283df5e07e7cf4aa2fbac64f1d460fdf9a21eeb27a7f60727853a414c1c450c525e8dab6a3f13cfa57171a",
	},
}

func TestGenCKC(t *testing.T) {
	for _, test := range spcContainerTests {
		f, err := os.Open(test.filePath)
		assert.NoError(t, err)
		defer f.Close()

		//spcMessage, err := ioutil.ReadAll(f)
		//assert.NoError(t, err)

	}

}

func TestParseSPCContainer(t *testing.T) {
	FileReader:=rsa.FileReader{FileName:"testdata/Development Credentials/dev_private_key.pem"}

	for _, test := range spcContainerTests {
		f, err := os.Open(test.filePath)
		defer f.Close()
		assert.NoError(t, err)


		spcMessage, err := ioutil.ReadAll(f)
		assert.NoError(t, err)

		spcContainer, err := ParseSPCContainer(spcMessage,FileReader)
		assert.NoError(t, err)

		assert.Equal(t, test.iv, hex.EncodeToString(spcContainer.AesKeyIV))
		assert.Equal(t, test.spec, hex.EncodeToString(spcContainer.EncryptedAesKey))


	}
}
