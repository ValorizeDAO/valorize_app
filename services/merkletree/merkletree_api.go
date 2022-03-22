package merkletree

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetMerkleRoot(leaves string) (string, error) {
  url := os.Getenv("AWS_SERVERLESS_HOST") + "/get-merkle-root"
  method := "POST"

  payload := strings.NewReader(`{"leaves":` + leaves + `}`)

  client := &http.Client{}

  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return "", err 
  }
  req.Header.Add("Content-Type", "text/plain")

  res, err := client.Do(req)
  if err != nil {
    return "", err 
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return "", err 
  }
  return string(body), nil
}