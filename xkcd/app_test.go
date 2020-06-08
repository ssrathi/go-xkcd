package xkcd

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

// assertEqual checks if two given values are equal and fatals if not.
func assertEqual(t *testing.T, got interface{}, want interface{}) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got '%+[1]v' (%[1]T), want '%+[2]v' (%[2]T)", got, want)
	}
}

func TestComicSpecificNums(t *testing.T) {
	imgDir, err := ioutil.TempDir(os.TempDir(), "testGoXkcd")
	assertEqual(t, err, nil)
	defer os.RemoveAll(imgDir)

	t.Logf("Image directory: %s", imgDir)
	client := NewClient()

	var tests = []struct {
		comicNum int
		title    string
		dateStr  string
	}{
		{111, "Firefox and Witchcraft - The Connection?", "05-Jun-2006"},
		{222, "Small Talk", "12-Feb-2007"},
		{333, "Getting Out of Hand", "24-Oct-2007"},
	}
	for _, tc := range tests {
		comic, err := client.GetComicMetadata(tc.comicNum)
		assertEqual(t, err, nil)
		assertEqual(t, comic.Title, tc.title)

		dateStr, err := comic.Date()
		assertEqual(t, err, nil)
		assertEqual(t, dateStr, tc.dateStr)

		// Download and check image.
		imgFullPath, err := client.GetComicImage(comic.Img, imgDir)
		assertEqual(t, err, nil)

		_, err = os.Stat(imgFullPath)
		assertEqual(t, err, nil)
	}
}
