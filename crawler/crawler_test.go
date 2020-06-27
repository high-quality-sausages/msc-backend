package crawler

import "testing"

func TestCrawler(t *testing.T) {
	info, err := RequestSongsByName("hello")
	if err != nil {
		t.Error(err)
	}
	t.Log(info)
}
