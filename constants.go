package neji

// RecognitionMode implies what mode should be used in the recognition
type RecognitionMode int

const (
	// RecognitionModeHOG uses HOG method for recognition
	RecognitionModeHOG RecognitionMode = 0

	// RecognitionModeCNN uses CNN method for recognition
	RecognitionModeCNN RecognitionMode = 1
)
