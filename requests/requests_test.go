package requests

import (
	"testing"
	"path/filepath"
	"os"
)

func TestDownloadString(t *testing.T) {
	u := "https://www.oschina.net/"

	b,err := DownloadString(u,nil)

	if err != nil {
		t.Fatal(err)
	}else{
		t.Log(string(b))
	}

}

func TestDownloadAndSaveFile(t *testing.T) {
	u := "http://imgskype.gmw.cn/software/android/chinaGmf-7.37.99.40.apk"

	dstFile := filepath.Join(os.TempDir(),"chinaGmf-7.37.99.40.apk")

	err := DownloadAndSaveFile(u,dstFile);

	if err != nil {
		t.Fatal(err)
	}else{
		t.Log(dstFile)
	}
}