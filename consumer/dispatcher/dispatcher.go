package dispatcher

import(
  "io/ioutil"
  "net/http"
  "fmt"
  "bytes"
)

const Url = "http://temptest.proxy.beeceptor.com/my/api/test"

func SendNotification(content []byte) {
  fmt.Println("Sending notification to: ", Url)

  req, err := http.NewRequest("POST", Url, bytes.NewBuffer(content))
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp, err := client.Do(req)
  if err != nil {
      panic(err)
  }
  defer resp.Body.Close()

  fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}
