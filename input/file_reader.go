package input

import "os"

func GetInputData(filename string) string {
	dir := os.Getenv("INPUTDATA")
	bs, err := os.ReadFile(dir + "/" + filename)
	if err != nil {
		panic("couldn't read in file: " + dir + "/" + filename)
	}
	return string(bs)
}
