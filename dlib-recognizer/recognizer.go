package recognizer

// #cgo pkg-config: dlib-1
// #cgo CXXFLAGS: -std=c++1z -Wall -O3 -DNDEBUG -march=native
// #cgo LDFLAGS: -ljpeg
// #include <stdlib.h>
// #include <stdint.h>
// #include "facerec.h"
import "C"
import (
	"fmt"
	"image"
	"unsafe"
)

const (
	RecognitionModeHOG int = 0
	RecognitionModeCNN int = 1

	rectLen  = 4
	descrLen = 128
	shapeLen = 2
)

type Recognizer struct {
	facerec *C.facerec
}

type Config struct {
	ModelsPath string

	Size      int
	Padding   int
	Jittering int
}

type Face struct {
	Rectangle  image.Rectangle
	Descriptor Descriptor
	Shapes     []image.Point
}

type Descriptor [128]float32

func NewRecognizer(cfg *Config) (rec *Recognizer, err error) {
	cModelDir := C.CString(cfg.ModelsPath)
	defer func() {
		C.free(unsafe.Pointer(cModelDir))
		}()

		// init facerec
	ptr := C.facerec_init(cModelDir)
	if ptr.err_str != nil {
		err = fmt.Errorf("dlib: %s (code %d)", C.GoString(ptr.err_str), int(ptr.err_code))
		defer C.facerec_free(ptr)
		defer C.free(unsafe.Pointer(ptr.err_str))
		
		return
	}
	
	// set additional configs
	if cfg.Size != 0 || cfg.Padding != 0  || cfg.Jittering != 0 {
		cSize := C.ulong(cfg.Size)
		cPadding := C.double(cfg.Padding)
		cJittering := C.int(cfg.Jittering)
		// TODO: free other params pointers?

		C.facerec_config(ptr, cSize, cPadding, cJittering)
	}

	return &Recognizer{ptr}, nil
}

// Recognize will recognize faces in imgData
func (r *Recognizer) Recognize(imgData []byte, maxFaces int, recotype int) (faces []Face, err error) {
	if len(imgData) == 0 {
		err = fmt.Errorf("dlib: empty image")
		return
	}

	cImgData := (*C.uint8_t)(&imgData[0])
	cLen := C.int(len(imgData))
	cMaxFaces := C.int(maxFaces)
	cRecoType := C.int(recotype)

	ret := C.facerec_recognize(r.facerec, cImgData, cLen, cMaxFaces, cRecoType)
	defer C.free(unsafe.Pointer(ret))

	if ret.err_str != nil {
		defer C.free(unsafe.Pointer(ret.err_str))
		err = fmt.Errorf("dlib: %s (code %d)", C.GoString(ret.err_str), int(ret.err_code))
		return
	}

	numFaces := int(ret.num_faces)
	if numFaces == 0 {
		return
	}
	numShapes := int(ret.num_shapes)

	// Copy faces data to Go structure.
	defer C.free(unsafe.Pointer(ret.shapes))
	defer C.free(unsafe.Pointer(ret.rectangles))
	defer C.free(unsafe.Pointer(ret.descriptors))

	rDataLen := numFaces * rectLen
	rDataPtr := unsafe.Pointer(ret.rectangles)
	rData := (*[1 << 30]C.long)(rDataPtr)[:rDataLen:rDataLen]

	dDataLen := numFaces * descrLen
	dDataPtr := unsafe.Pointer(ret.descriptors)
	dData := (*[1 << 30]float32)(dDataPtr)[:dDataLen:dDataLen]

	sDataLen := numFaces * numShapes * shapeLen
	sDataPtr := unsafe.Pointer(ret.shapes)
	sData := (*[1 << 30]C.long)(sDataPtr)[:sDataLen:sDataLen]

	for i := 0; i < numFaces; i++ {
		face := Face{}
		x0 := int(rData[i*rectLen])
		y0 := int(rData[i*rectLen+1])
		x1 := int(rData[i*rectLen+2])
		y1 := int(rData[i*rectLen+3])
		face.Rectangle = image.Rect(x0, y0, x1, y1)
		copy(face.Descriptor[:], dData[i*descrLen:(i+1)*descrLen])
		for j := 0; j < numShapes; j++ {
			shapeX := int(sData[(i*numShapes+j)*shapeLen])
			shapeY := int(sData[(i*numShapes+j)*shapeLen+1])
			face.Shapes = append(face.Shapes, image.Point{shapeX, shapeY})
		}
		faces = append(faces, face)
	}
	return
}

// RecognizeHOG will recognize image data using HOG mode
func (r *Recognizer) RecognizeHOG(imgData []byte, maxFaces int) (faces []Face, err error) {
	return r.Recognize(imgData, maxFaces, RecognitionModeHOG)
}

// RecognizeCNN will recognize image data using CNN mode
func (r *Recognizer) RecognizeCNN(imgData []byte, maxFaces int) (faces []Face, err error) {
	return r.Recognize(imgData, maxFaces, RecognitionModeCNN)
}

// Close frees up resources taken by recognizer
func (r *Recognizer) Close() {
	C.facerec_free(r.facerec)
	r.facerec = nil
}
