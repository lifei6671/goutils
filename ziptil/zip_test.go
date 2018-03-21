package ziptil

import (
	"testing"
	"os"
	"path/filepath"
)

func TestZip(t *testing.T) {
	err := Zip("./../",filepath.Join(os.TempDir() , "goutils.zip"))

	if err != nil {
		t.Fatal(err)
	}else{
		t.Log(os.TempDir())
	}
}

func TestUnZip(t *testing.T) {
	dst := filepath.Join(os.TempDir() , "goutils.zip")
	source := filepath.Join(os.TempDir() , "goutils")
	err := Zip("./",dst)

	if err != nil {
		t.Fatal(err)
	}else{
		err = UnZip(dst,source)
		if err != nil {
			t.Fatal(err)
		}else{
			t.Log(os.TempDir())
		}
	}
}