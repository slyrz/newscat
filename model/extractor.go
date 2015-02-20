package model

import (
	"errors"
	"github.com/slyrz/newscat/html"
	"github.com/slyrz/newscat/util"
)

var (
	ErrNoChunks    = errors.New("document contains no chunks")
	ErrEmptyResult = errors.New("nothing found")
)

// Extractor utilizes the trained model to extract relevant html.Chunks from
// an html.Document.
type Extractor struct {
	Labels []bool
}

// NewExtractor creates and initializes a new Extractor.
func NewExtractor() *Extractor {
	return new(Extractor)
}

// Extract returns a list of relevant text chunks found in doc.
//
// How it works
//
// This function creates a feature vector for each chunk found in doc.
// A feature vector contains a numerical representation of the chunk's
// properties like HTML element type, parent element type, number of words,
// number of sentences and stuff like this.
//
// A logistic regression model is used to calculate scores based on these
// feature vectors. Then, in some kind of meta / ensemble learning approach,
// a second type of feature vector is created based on these scores.
// This feature vector is fed to our random forest and finally
// the random forest's predictions are used to generate the result.
//
// By now you might have noticed that I'm exceptionally bad at naming and
// describing things properly.
func (ext *Extractor) Extract(doc *html.Document) (*util.Article, error) {
	*ext = Extractor{}
	if len(doc.Chunks) == 0 {
		return nil, ErrNoChunks
	}

	chunkFeatures := make([]chunkFeature, len(doc.Chunks))
	boostFeatures := make([]boostFeature, len(doc.Chunks))

	// Count the number of words and sentences we encountered for each
	// class. This helps us to detect elements that contain the doc text.
	classStats := doc.GetClassStats()
	clusterStats := doc.GetClusterStats()

	chunkFeatureWriter := new(chunkFeatureWriter)
	for i, chunk := range doc.Chunks {
		chunkFeatureWriter.Assign(chunkFeatures[i][:])
		chunkFeatureWriter.WriteElementType(chunk)
		chunkFeatureWriter.WriteParentType(chunk)
		chunkFeatureWriter.WriteSiblingTypes(chunk)
		chunkFeatureWriter.WriteAncestors(chunk)
		chunkFeatureWriter.WriteTextStat(chunk)
		chunkFeatureWriter.WriteTextStatSiblings(chunk)
		chunkFeatureWriter.WriteClassStat(chunk, classStats)
		chunkFeatureWriter.WriteClusterStat(chunk, clusterStats)
	}

	// Detect the minimum and maximum value for each element in the
	// feature vector.
	empMin := chunkFeature{}
	empMax := chunkFeature{}
	for i := range chunkFeatures {
		for j, val := range chunkFeatures[i] {
			switch {
			case val < empMin[j]:
				empMin[j] = val
			case val > empMax[j]:
				empMax[j] = val
			}
		}
	}

	// Perform MinMax normalization.
	for i := range chunkFeatures {
		feature := &chunkFeatures[i]
		for j, val := range chunkFeatures[i] {
			// If the maximum value is not greater than one, we assume that the feature is
			// already normalized and leave it untouched.
			if empMax[j] > 1.0 {
				feature[j] = (val - empMin[j]) / (empMax[j] - empMin[j])
			}
		}
	}

	// Now cluster chunks by containers to calculate average score per
	// container.
	clusterContainer := newClusterMap()
	for i, chunk := range doc.Chunks {
		clusterContainer.Add(chunk.Container, chunk, chunkFeatures[i].Score())
	}

	boostFeatureWriter := new(boostFeatureWriter)
	for i, chunk := range doc.Chunks {
		boostFeatureWriter.Assign(boostFeatures[i][:])
		boostFeatureWriter.WriteChunk(chunk)
		boostFeatureWriter.WriteCluster(chunk, clusterContainer[chunk.Container])
		boostFeatureWriter.WriteTitleSimilarity(chunk, doc.Title)
	}

	// Cluster chunks by block.
	clusterBlock := newClusterMap()
	for i, chunk := range doc.Chunks {
		clusterBlock.Add(chunk.Block, chunk, boostFeatures[i].Score(), float32(chunk.Text.Len()))
	}

	// Label all chunks whose blocks have a score above prediction level.
	// This makes sure that we don't split large blocks.
	ext.Labels = make([]bool, len(doc.Chunks))
	for i, chunk := range doc.Chunks {
		if cluster, ok := clusterBlock[chunk.Block]; ok {
			ext.Labels[i] = cluster.Score() > 0.5
		}
	}

	result := &util.Article{Title: doc.Title.String()}
	for i, chunk := range doc.Chunks {
		if cluster, ok := clusterBlock[chunk.Block]; ok && ext.Labels[i] {
			text := util.NewText()
			for _, chunk := range cluster.Chunks {
				text.WriteText(chunk.Text)
			}
			if chunk.IsHeading() {
				result.Append(util.Heading(text.String()))
			} else {
				result.Append(util.Paragraph(text.String()))
			}
			delete(clusterBlock, chunk.Block)
		}
	}
	if len(result.Text) == 0 {
		return nil, ErrEmptyResult
	}
	return result, nil
}
