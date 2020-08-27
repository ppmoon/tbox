package tbox_test

import "testing"
import "github.com/ppmoon/tbox"

func TestDownloadFileAndStore(t *testing.T) {
	filePath := "logo.png"
	url := "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png"
	err := tbox.DownloadFileAndStore(filePath, url)
	if err != nil {
		t.Error(err)
		return
	}
}
