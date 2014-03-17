package model

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"os"
	"path"
)

var (
	// Default path to the model file, it's relative to our binary.
	Path = path.Join(path.Dir(os.Args[0]), "data/model")
)

// Atoms is the fixed size SVM part. Combining these fields in a separate
// struct allows us to read them altogether. We need to know these fields
// before we can read the variable length arrays.
type Atoms struct {
	N int32   // number of classes
	L int32   // number of support vectors
	G float32 // parameter gamma of RBF
}

// Support Vector Machine used to classify Features.
type SVM struct {
	Atoms
	Nsv []int32   // number of support vectors for each class
	Cof []float32 // coefficients used in decision functions
	Rho []float32 // constants in decision function
	Vec []Feature // support vectors
}

// Open creates a new SVM struct and loads the SVM fields from a given path.
func Open(path string) (*SVM, error) {
	instance, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer instance.Close()

	svm := new(SVM)
	return svm, svm.Load(instance)
}

// ReadFields reads binary data from an io.Reader object and panics if the
// read operation failed.
func (svm *SVM) ReadField(r io.Reader, data interface{}) {
	if err := binary.Read(r, binary.LittleEndian, data); err != nil {
		panic(err)
	}
}

// Load tries to read the SVM struct fields from an object providing the
// io.Reader interface.
func (svm *SVM) Load(r io.Reader) (err error) {
	// Catch the panics and return errors instead.
	defer func() {
		if recover() != nil {
			err = errors.New("loading format failed")
		} else {
			err = nil
		}
	}()

	svm.ReadField(r, &svm.Atoms)

	svm.Nsv = make([]int32, svm.N)
	svm.Cof = make([]float32, svm.L)
	svm.Rho = make([]float32, svm.N*(svm.N-1)/2)
	svm.Vec = make([]Feature, svm.L)

	svm.ReadField(r, &svm.Nsv)
	svm.ReadField(r, &svm.Cof)
	svm.ReadField(r, &svm.Rho)
	svm.ReadField(r, &svm.Vec)
	return
}

// Radial Basis Function.
func (svm *SVM) RBF(a, b *Feature) float32 {
	var s float32 = 0.0
	var v float32 = 0.0
	for i := range a {
		v = (a[i] - b[i])
		s += v * v
	}
	return float32(math.Exp(float64(-svm.G * s)))
}

// Predict returns true if a feature vector presumably belongs to the
// article text.
func (svm *SVM) Predict(x *Feature) bool {
	var s float32 = 0.0
	var i int32 = 0
	var j int32 = 0
	var k int32 = 0

	kern := make([]float32, svm.L)
	for i = 0; i < svm.L; i++ {
		kern[i] = svm.RBF(x, &svm.Vec[i])
	}

	i = 0
	j = svm.Nsv[0]

	for k = 0; k < svm.Nsv[0]; k++ {
		s += svm.Cof[i+k] * kern[i+k]
	}

	for k = 0; k < svm.Nsv[1]; k++ {
		s += svm.Cof[j+k] * kern[j+k]
	}

	return (s - svm.Rho[0]) <= 0
}
