package utils

import (
	"testing"
)

func TestDicFileToSlice(t *testing.T) {
	v, _ := DicFileToSlice("../dic.txt", "http://192.168.61.215")
	for i:=0;i<50;i++ {
		t.Log(v[i])
	}

}

func TestGenerateByUrl(t *testing.T) {


}
