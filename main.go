package main
import (
	"github.com/vselpublic/gochanger/conf"
	"github.com/golang/glog"
	"flag"
)

func init() {
	flag.Parse()
}

func main() {
	glog.Infoln("DEBUG MODE ON", conf.GetConf().DEBUG)
	glog.Flush()
}