package model

import (
	gonet "code.google.com/p/go.net/html"
	"github.com/slyrz/newscat/html"
)

type Clusters map[*gonet.Node]*Cluster

type Cluster struct {
	Chunks   []*html.Chunk
	Scores   []float32
	ScoreAvg float32
	ScoreSum float32
}

func NewCluster() *Cluster {
	result := new(Cluster)
	result.Chunks = make([]*html.Chunk, 0)
	result.Scores = make([]float32, 0)
	return result
}

func (cl *Cluster) Add(chunk *html.Chunk, score float32) {
	cl.Chunks = append(cl.Chunks, chunk)
	cl.Scores = append(cl.Scores, score)
	cl.ScoreSum += score
	cl.ScoreAvg = cl.ScoreSum / float32(len(cl.Chunks))
}

func NewClusters() Clusters {
	return make(Clusters)
}

func (cl Clusters) Add(base *gonet.Node, chunk *html.Chunk, score float32) {
	cluster, ok := cl[base]
	if !ok {
		cluster = NewCluster()
		cl[base] = cluster
	}
	cluster.Add(chunk, score)
}
