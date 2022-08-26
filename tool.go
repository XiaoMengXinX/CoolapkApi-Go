package coolapk

import (
	"fmt"
	"math/rand"
	"time"
)

const userAgentTmpl = `Dalvik/2.1.0 (Linux; U; Android %d) +CoolMarket/%d-%d-universal`

var coolapkVersions = [][]string{
	{"12.4.1", "2208081"},
	{"12.4", "2207271"},
	{"12.3.2", "2207151"},
	{"12.3.1", "2206171"},
	{"12.3", "2205191"},
	{"12.2.1", "2204291"},
	{"12.2", "2204151"},
}
var androidVersions = []string{"6.0", "6.0.1", "7.0", "7.1.1", "7.1.2", "8.0", "8.1", "9", "10", "11", "12", "12"}

func getRandomUA(tmpl string) string {
	rand.Seed(time.Now().UnixNano())
	androidVer := androidVersions[rand.Intn(len(androidVersions))]
	coolapkVer := coolapkVersions[rand.Intn(len(coolapkVersions))]
	return fmt.Sprintf(tmpl, androidVer, coolapkVer[0], coolapkVer[1])
}
