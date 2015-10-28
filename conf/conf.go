package conf
import (
	"encoding/json"
	"os"
	"sync"
)

type Configuration struct {
	filename               map[string]string
	DEBUG                  bool
}

var (
	conf Configuration
 	once sync.Once
)

func InitConf() {
	path := os.Args[1]
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	conf = Configuration{}
	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		panic(err)
	}
}

func GetConf() *Configuration {
	once.Do(InitConf)
	return &conf
}