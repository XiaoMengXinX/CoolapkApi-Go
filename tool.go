package coolapk

import (
	"fmt"
	"math/rand"
)

const userAgentTmpl = "Mozilla/5.0 (iPhone; CPU iPhone OS %s like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/%s NetType/%s Language/zh_CN"

func randomUA() string {
	iosVersions := []string{
		"16_0", "16_1", "16_2", "16_3", "16_4", "16_5", "16_6", "16_7",
		"17_0", "17_1", "17_2", "17_3", "17_4", "17_5", "17_6",
		"18_0", "18_1", "18_2", "18_3", "18_4", "18_5", "18_6", "18_7",
	}
	wxVersions := []string{
		"8.0.60", "8.0.61", "8.0.62", "8.0.63", "8.0.64", "8.0.65",
	}
	netTypes := []string{"4G", "5G"}

	iosVer := iosVersions[rand.Intn(len(iosVersions))]
	wxVer := wxVersions[rand.Intn(len(wxVersions))]
	netType := netTypes[rand.Intn(len(netTypes))]

	return fmt.Sprintf(
		userAgentTmpl,
		iosVer,
		wxVer,
		netType,
	)
}
