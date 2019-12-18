package lib

import (
	"fmt"
	"time"

	"github.com/kpango/gache"
)

func Gache(key, value string){
	g := gache.GetGache()
	g.SetWithExpire(key, value, time.Minute * 5)

	_, found := g.Get(key)
	if found == false {
		error.Error(fmt.Errorf("error"))
	}
	g.Delete(key)
	g.Clear()
}
