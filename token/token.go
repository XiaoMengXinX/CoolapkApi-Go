package token

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	appID     = "wx7c6be4984041fa23"
	appSecret = "a717e41f8e9254c52da78d70003f24a0"
)

type SystemInfo struct {
	Brand    string
	Model    string
	Platform string
	System   string
}

var mockSystemInfo = SystemInfo{
	Brand:    "iPhone",
	Model:    "unknown<iPhone18,3>",
	Platform: "",
	System:   "",
}

func md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func generateRawDeviceID(sysInfo SystemInfo) string {
	timestampMilli := time.Now().UnixMilli()
	randomNumber := rand.Intn(10000) + 1

	var builder strings.Builder
	builder.WriteString(sysInfo.Brand + ";")
	builder.WriteString(sysInfo.Model + ";")
	builder.WriteString(sysInfo.Platform + ";")
	builder.WriteString(sysInfo.System + ";")
	builder.WriteString(strconv.FormatInt(timestampMilli, 10))
	builder.WriteString(strconv.Itoa(randomNumber) + ";")

	return builder.String()
}

func GenerateHeaders() map[string]string {
	deviceIDMd5 := md5Hash(generateRawDeviceID(mockSystemInfo))
	deviceHeaderStr := fmt.Sprintf("%s; ; ; ; ; %s; %s; ", deviceIDMd5, mockSystemInfo.Brand, mockSystemInfo.Model)
	xAppDevice := base64.StdEncoding.EncodeToString([]byte(deviceHeaderStr))

	timestampSec := time.Now().Unix()
	appSecretMd5 := md5Hash(appSecret)
	signatureStr := fmt.Sprintf("%s/%s%s%d", appID, appSecretMd5, deviceIDMd5, timestampSec)
	signature := md5Hash(signatureStr)
	xAppToken := fmt.Sprintf("%s%s0x%s", signature, deviceIDMd5, fmt.Sprintf("%x", timestampSec))

	headers := map[string]string{
		"X-Requested-With": "XMLHttpRequest",
		"X-Sdk-Int":        "260",
		"X-Sdk-Locale":     "zh-CN",
		"X-App-Id":         appID,
		"X-App-Version":    "1.0",
		"X-App-Code":       "1902250",
		"X-Api-Version":    "9",
		"X-App-Device":     xAppDevice,
		"X-App-Token":      xAppToken,
		"Cookie":           "sec_tc=AQAAANI9pQ7PKw0AcvQGP+h9lgm+XWf9; Path=/; Expires=Mon, 15-Apr-19 08:55:36 GMT; HttpOnly",
	}
	return headers
}
