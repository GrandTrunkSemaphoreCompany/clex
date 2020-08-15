package encoding

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var update = flag.Bool("update", false, "update golden file")

func TestMakeClacksFromByteUsingInvalidDirectory(t *testing.T) {
	dir := "/tmp/clex/unit/invalid"

	s := "a A"
	im := new(Image)
	im.BasePath = dir

	_, err := im.Write([]byte(s))
	if err == nil {
		t.Fatalf("Test should fail on not created directory")
	}

}

func TestMakeClacksFromByte(t *testing.T) {
	dir := fmt.Sprintf("%s/clacks-unit/writing/%d", os.TempDir(), time.Now().UTC().UnixNano())

	err := os.MkdirAll(dir, 0700)
	if err != nil {
		t.Fatal(err)
	}

	s := "a A"
	im := new(Image)
	im.BasePath = dir + "/"

	_, err = im.Write([]byte(s))
	if err != nil {
		t.Fatal(err)
	}

	// check directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}

	if len(files) != 3 {
		t.Errorf("%d files found instead of 3", len(files))
	}

	// check files
	for _, filename := range files {
		compare, err := compareFiles(dir, "clacks/testdata", filename.Name())

		if err != nil {
			t.Fatal(err)
		}

		if !compare {
			t.Errorf("File %s is different", filename.Name())
		}

	}

	//
}

func compareFiles(source string, generated string, filename string) (result bool, error error) {
	sourceFilepath := source + "/" + filename
	generatedFilepath := generated + "/" + filename

	fmt.Println(sourceFilepath)
	fmt.Println(generatedFilepath)

	file, err := os.Open(sourceFilepath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

	return true, nil
}
