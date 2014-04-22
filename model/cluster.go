package model

import (
	gonet "code.google.com/p/go.net/html"
	"github.com/slyrz/newscat/html"
)

type Clusters map[*gonet.Node]*Cluster

type Cluster struct {
	Chunks  []*html.Chunk
	Scores  []float32
	Weights []float32
	// Unexported fields.
	average float32
	changed bool
}

func NewCluster() *Cluster {
	result := new(Cluster)
	result.Chunks = make([]*html.Chunk, 0)
	result.Scores = make([]float32, 0)
	result.Weights = make([]float32, 0)
	return result
}

func (cl *Cluster) Add(chunk *html.Chunk, args ...float32) {
	var score float32 = 0.0
	var weight float32 = 0.0
	switch len(args) {
	case 2:
		score, weight = args[0], args[1]
	case 1:
		score, weight = args[0], 1.0
	default:
		panic("Call Add(chunk, score) or Add(chunk, score, weight)")
	}
	cl.Chunks = append(cl.Chunks, chunk)
	cl.Scores = append(cl.Scores, score)
	cl.Weights = append(cl.Weights, weight)
	cl.changed = true
}

func (cl *Cluster) Score() float32 {
	if cl.changed {
		var s float32 = 0.0
		var w float32 = 0.0
		for i := range cl.Chunks {
			s += cl.Weights[i] * cl.Scores[i]
			w += cl.Weights[i]
		}
		cl.average = s / w
		cl.changed = true
	}
	return cl.average
}

func NewClusters() Clusters {
	return make(Clusters)
}

func (cl Clusters) Add(base *gonet.Node, chunk *html.Chunk, args ...float32) {
	cluster, ok := cl[base]
	if !ok {
		cluster = NewCluster()
		cl[base] = cluster
	}
	cluster.Add(chunk, args...)
}
