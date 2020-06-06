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
	// defer os.RemoveAll(imgDir)

	t.Logf("Image directory: %s", imgDir)
	client := NewClient()

	var tests = []struct {
		comicNum int
		title    string
	}{
		{111, "Firefox and Witchcraft - The Connection?"},
		{222, "Small Talk"},
		{333, "Getting Out of Hand"},
	}
	for _, tc := range tests {
		comic, err := client.GetComicMetadata(tc.comicNum)
		assertEqual(t, err, nil)
		got := comic.Title
		want := tc.title
		assertEqual(t, got, want)

		// Download and check image.
		imgFullPath, err := client.GetComicImage(comic.Img, imgDir)
		assertEqual(t, err, nil)

		_, err = os.Stat(imgFullPath)
		assertEqual(t, err, nil)
	}
}
