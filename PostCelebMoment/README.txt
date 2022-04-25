-Concept: build a portal to photographer can upload their picture, the total picture they can upload daily will depend on their level.
In here, I am setting default the limitation is 5 photos per day.
-Solution: 
	+Create one sample API to upload image to AWS S3
	+Validate the limitation by userId
	+Apply Spring boot framework
	+Apply Junit to write unit test
	+Apply post man application to do integration test
	
-How to run code:
	+Import like a maven project into Eclipse or Inteliji >> Right-click on project >> Run as Spring Boot App
	+Run application like normal java maven project
	+Use post man to send a post request or send request by Go script like below.
	
-How to run unit test:
	+Right click on class PostCelebMomentApplicationTests >> Run as JUnit test
-TO DO:
	+Implement logic to viewer can vote what photo they love and base on that info to upgrade photographer's level
	+User AWS database to store User Tracking information instead of csv file
	+Write more unit test



============GO SCRIPT TO send POST===============================
package main
import (
  "fmt"
  "bytes"
  "mime/multipart"
  "os"
  "path/filepath"
  "io"
  "net/http"
  "io/ioutil"
)

func main() {

  url := "http://localhost:8080/api/storage/celebMoment"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, errFile1 := os.Open("/C:/Users/thanhpn6/Desktop/New folder/VF4_Cell_Icon.jpg")
  defer file.Close()
  part1,
         errFile1 := writer.CreateFormFile("file",filepath.Base("/C:/Users/thanhpn6/Desktop/New folder/VF4_Cell_Icon.jpg"))
  _, errFile1 = io.Copy(part1, file)
  if errFile1 != nil {
    fmt.Println(errFile1)
    return
  }
  _ = writer.WriteField("userId", "thanhpn")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }


  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Set("Content-Type", writer.FormDataContentType())
  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}
============GO SCRIPT TO send POST===============================