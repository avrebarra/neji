package neji_test

import (
	"github.com/shrotavre/neji"
	"io/ioutil"
	"testing"
)

func TestRecognize(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		// arrange
		n, err := neji.NewRecognizer("./models", neji.RecognitionModeCNN)
		imgData, err := ioutil.ReadFile("./fixtures/hinata.jpeg")
		if err != nil {
			t.Error("failed loading image test file")
			return
		}

		// act
		faces, err := n.RecognizeFaces(imgData, 10)

		// assert
		if err != nil {
			t.Error("error recognizing faces")
		}
		if len(faces) != 1 {
			t.Error("expected face count mismatch")
		}
	})
}
