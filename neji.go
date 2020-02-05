package neji

import recognizer "github.com/shrotavre/neji/dlib-recognizer"

// Recognizer is compact face recognition instance working over dlib recognition system
type Recognizer struct {
	R    *recognizer.Recognizer
	mode int
}

// Face is recognized faces object
type Face struct {
	FV []float32
}

// RecognizeFaces will scan and recognize faces found in a image data
func (n *Recognizer) RecognizeFaces(imgData []byte, maxFaces int) (rfs []Face, err error) {
	faces, err := n.R.Recognize(imgData, maxFaces, n.mode)

	rfs = make([]Face, len(faces))
	for i, f := range faces {
		rfs[i] = Face{f.Descriptor[:]}
	}

	return
}

// Close will close and clear recognizer's occupied resources
func (n *Recognizer) Close() error {
	n.R.Close()
	n.R = nil

	return nil
}

// NewRecognizer creates new NewRecognizer instance
func NewRecognizer(modelsDir string, mode int) (*Recognizer, error) {
	c := &recognizer.Config{
		ModelsPath: modelsDir,
	}

	// * try to init recognizer
	reco, err := recognizer.NewRecognizer(c)
	if err != nil {
		return nil, err
	}

	return &Recognizer{
		R:    reco,
		mode: mode,
	}, nil
}
