package model

import (
	gonet "code.google.com/p/go.net/html"
	"github.com/slyrz/newscat/html"
)

// A clusterMap groups Clusters by HTML nodes.
type clusterMap map[*gonet.Node]*cluster

// A cluster stores a group of html.Chunks and their scores.
type cluster struct {
	Chunks  []*html.Chunk
	Scores  []float32
	Weights []float32
	// Unexported fields.
	average float32 // weighted average score
	changed bool    // denotes struct changes after average calculation
}

// newCluster creates and initalizes a new cluster.
func newCluster() *cluster {
	result := new(cluster)
	result.Chunks = make([]*html.Chunk, 0)
	result.Scores = make([]float32, 0)
	result.Weights = make([]float32, 0)
	return result
}

// Add adds the html.Chunk chunk to the cluster. The variadic float32 parameter
// args must either be (score,) or (score, weight).
func (cl *cluster) Add(chunk *html.Chunk, args ...float32) {
	var score float32 = 0.0
	var weight float32 = 0.0
	switch len(args) {
	case 2:
		score, weight = args[0], args[1]
	case 1:
		score, weight = args[0], 1.0
	default:
		panic("call Add(chunk, score) or Add(chunk, score, weight)")
	}
	cl.Chunks = append(cl.Chunks, chunk)
	cl.Scores = append(cl.Scores, score)
	cl.Weights = append(cl.Weights, weight)
	cl.changed = true
}

// Score calculates the weighted average of all chunk scores in cluster.
func (cl *cluster) Score() float32 {
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

// newClusterMap creates and initalizes a new clusterMap.
func newClusterMap() clusterMap {
	return make(clusterMap)
}

// Add adds the html.Chunk chunk to the cluster indexed by key.
func (cl clusterMap) Add(key *gonet.Node, chunk *html.Chunk, args ...float32) {
	cluster, ok := cl[key]
	if !ok {
		cluster = newCluster()
		cl[key] = cluster
	}
	cluster.Add(chunk, args...)
}
