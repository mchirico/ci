package pkg

import "testing"

func TestPkg(t *testing.T) {

	Sample()
}

func TestSample2(t *testing.T) {
	Sample2()
}

func TestMkdir(t *testing.T) {

	top := "/tmp/spock/"
	full := top + "junk/stuff/dir"
	file := full + "/file"

	Rmdir(top)
	err := WriteString(top+"sample", "file", 0600)
	if err == nil {
		t.FailNow()
	}

	Mkdir(full)
	err = WriteString(file, "This is sample data", 0600)
	if err != nil {
		t.FailNow()
	}

	Rmdir(top)
}
