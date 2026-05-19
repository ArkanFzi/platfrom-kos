package main

import (
"fmt"
"os"
    "context"

"github.com/cloudinary/cloudinary-go/v2"
"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func main() {
url := os.Getenv("CLOUDINARY_URL")
cld, _ := cloudinary.NewFromURL(url)

res, err := cld.Upload.Upload(context.Background(), "https://res.cloudinary.com/demo/image/upload/sample.jpg", uploader.UploadParams{
Folder: "rooms",
})
if err != nil {
fmt.Println("Error:", err)
return
}
fmt.Printf("Result: %+v\n", res)
}
