package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L.
#include <stdlib.h>
*/
import "C"

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
	"strings"
)

//export calculateSessionKey
func calculateSessionKey() *C.char {
    currentTime := time.Now()
    date := currentTime.Format("2006-01-02")
    timeComponent := currentTime.Format("15:04")

    noisePattern := "some-fixed-noise-pattern"

    combinedString := date + timeComponent + noisePattern

    hash := sha256.New()
    hash.Write([]byte(combinedString))
    sessionKey := hex.EncodeToString(hash.Sum(nil))

    jsonData := fmt.Sprintf(`{"Session": "%s", "Created": "%s"}`, sessionKey, timeComponent)

    // Return the JSON object as a C string
    return C.CString(jsonData)
}

// SessionInfo represents the JSON input for validation
type SessionInfo struct {
    Session string `json:"Session"`
    Created string `json:"Created"`
}

//export validateSessionKey
func validateSessionKey(jsonStr *C.char) bool {
    jsonBytes := C.GoString(jsonStr)
    jsonBytes = jsonBytes[1 : len(jsonBytes)-1] // remove outer quotes
    jsonBytes = strings.Replace(jsonBytes, "\\\"", "\"", -1) // remove backslashes
    var sessionInfo SessionInfo
    err := json.Unmarshal([]byte(jsonBytes), &sessionInfo)
    if err != nil {
        fmt.Println("Error parsing JSON:", err)
        return false
    }

    if sessionInfo.Session == "" || sessionInfo.Created == "" {
        fmt.Println("Missing session key or created time in JSON")
        return false
    }

    createdTime, err := time.Parse("15:04", sessionInfo.Created)
    if err != nil {
        fmt.Println("Error parsing time:", err)
        return false
    }

    currentDate := time.Now().Format("2006-01-02")
    fmt.Println(currentDate)
    expectedKey := calculateSessionKeyFromDateTime(currentDate, createdTime)

    return sessionInfo.Session == expectedKey
}

func calculateSessionKeyFromDateTime(date string, createdTime time.Time) string {
	timeComponent := createdTime.Format("15:04")
	noisePattern := "some-fixed-noise-pattern"
	combinedString := date + timeComponent + noisePattern

	hash := sha256.New()
	hash.Write([]byte(combinedString))
	sessionKey := hex.EncodeToString(hash.Sum(nil))

	return sessionKey
}

func main() {}
