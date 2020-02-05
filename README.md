![GitHub Logo](./neji.jpg)

# neji [![GoDoc](https://godoc.org/github.com/shrotavre/neji?status.svg)](http://godoc.org/github.com/shrotavre/neji) [![Go Report Card](https://goreportcard.com/badge/shrotavre/neji)](https://goreportcard.com/report/github.com/shrotavre/neji)

Neji is a face recognition library working over dlib face recognition model. Yes, it needs `dlib` to work.

This library will only help you to recognize face feature vectors inside a given image (currently only support `.JPEG` files format). For determining closeness of faces you can implement another distance calculation function over the extracterd feature vectors here.


## Usage

~~~ go
// main.go
package main

import (
	"fmt"

	"github.com/shrotavre/neji"
)

func main() {
    // this assumes you already downloaded dlib models needed and placed in
    // ./models folder
    n, err := neji.NewNeji("./models")
    imgData, err := ioutil.ReadFile("./fixtures/hinata.jpeg")
    if err != nil {
        t.Error("failed loading image test file")
        return
    }

    faces, err := n.RecognizeFaces(imgData, 10)

    // this will list all feature vectors found in the fixture image.
    // The fixture image only has 1 faces though.  
    fmt.Println(faces)
}
~~~


## Models

The dlib models used within this package is/are: 
- `dlib_face_recognition_resnet_model_v1.dat`
- `mmod_human_face_detector.dat`
- `shape_predictor_5_face_landmarks.dat`