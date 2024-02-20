package auth

import (
	"errors"
	"net/http"
	"strings"
)

// import "net/http"

func GetAPIKey(headers http.Header) (string, error) {
  val := headers.Get("Authorization")

  if val == ""{
    return "", errors.New("No auth info found")
  }

  vals := strings.Split(val, " ")
  if len(vals) != 2 {
    return "", errors.New("Malformed auth header")
  }
  if vals[0] != "ApiKey" {
    return "", errors.New("Mailformed first part of auth header")
  }
  return vals[1], nil
}
