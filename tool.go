package coolapk

import (
	"fmt"
	"math/rand"
	"time"
)

const userAgentTmpl = `Dalvik/2.1.0 (Linux; U; Android %s; %s Build/%s; %s) (#Build; %s; %s; %s; %s) +CoolMarket/%s-%s-universal`

var coolapkVersions = [][]string{
	{"12.4.1", "2208081"},
	{"12.4", "2207271"},
	{"12.3.2", "2207151"},
	{"12.3.1", "2206171"},
	{"12.3", "2205191"},
	{"12.2.1", "2204291"},
	{"12.2", "2204151"},
}
var androidVersions = []string{"9", "10", "11", "12"}
var models = []string{
	"Pixel 3",
	"Pixel 3 XL",
	"Pixel 3a",
	"Pixel 4",
	"Pixel 4 XL",
	"Pixel 4a",
	"Pixel 5",
	"Pixel 5a",
}
var buildNumbers = []string{
	"SP2A.220505.002",
	"SP2A.220405.003",
	"SP2A.220305.012",
	"SQ1A.220105.002",
	"SQ1A.211205.008",
	"SP1A.211105.004",
	"RQ3A.211001.001",
	"RQ3A.210905.001",
	"RQ3A.210705.001",
	"RP1A.200720.009",
}

func getRandomUA(tmpl string) string {
	rand.Seed(time.Now().UnixNano())
	androidVer := androidVersions[rand.Intn(len(androidVersions))]
	coolapkVer := coolapkVersions[rand.Intn(len(coolapkVersions))]
	model := models[rand.Intn(len(models))]
	buildNumber := buildNumbers[rand.Intn(len(buildNumbers))]
	return fmt.Sprintf(tmpl, androidVer, model, buildNumber, androidVer, "google", model, buildNumber, androidVer, coolapkVer[0], coolapkVer[1])
}
