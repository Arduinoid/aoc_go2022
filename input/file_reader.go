package input

import (
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func GetInputData(filename string) []string {
	err := godotenv.Load()
	if err != nil {
		panic("could not load environment")
	}
	dir := os.Getenv("INPUTDATA")
	bs, err := os.ReadFile(dir + "/" + filename)
	if err != nil {
		panic("couldn't read in file: " + dir + "/" + filename)
	}
	return strings.Split(strings.ReplaceAll(string(bs), "\r", ""), "\n")
}
