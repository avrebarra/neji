package neji

import recognizer "github.com/shrotavre/neji/dlib-recognizer"

const (
	// RecognitionModeHOG uses HOG method for recognition
	RecognitionModeHOG int = recognizer.RecognitionModeHOG

	// RecognitionModeCNN uses CNN method for recognition
	RecognitionModeCNN int = recognizer.RecognitionModeCNN
)
