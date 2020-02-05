package neji

import recognizer "github.com/shrotavre/neji/dlib-recognizer"

// Recognizer is compact face recognition instance working over dlib recognition system
type Recognizer struct {
	mode int
	rec  *recognizer.Recognizer
}

// Face is recognized faces object
type Face struct {
	FV []float32
}

// RecognizeFaces will scan and recognize faces found in a image data
func (n *Recognizer) RecognizeFaces(imgData []byte, maxFaces int) (rfs []Face, err error) {
	faces, err := n.rec.Recognize(imgData, maxFaces, n.mode)

	rfs = make([]Face, len(faces))
	for i, f := range faces {
		rfs[i] = Face{f.Descriptor[:]}
	}

	return
}

// NewRecognizer creates new NewRecognizer instance
func NewRecognizer(modelsDir string) (*Recognizer, error) {
	recognitionMode := 0
	c := &recognizer.Config{
		ModelsPath: modelsDir,
	}

	// * try to init recognizer
	reco, err := recognizer.NewRecognizer(c)
	if err != nil {
		return nil, err
	}

	return &Recognizer{
		rec:  reco,
		mode: recognitionMode,
	}, nil
}
