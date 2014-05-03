package model

import (
	"github.com/slyrz/newscat/html"
)

// Extractor utilizes the trained model to extract relevant html.Chunks from
// an html.Document.
type Extractor struct {
	ChunkFeatures []chunkFeature
	BoostFeatures []boostFeature
}

// NewExtractor creates and initalizes a new Extractor.
func NewExtractor() *Extractor {
	return new(Extractor)
}

// Extract returns a list of relevant article content chunks found in
// the document.
//
// How it works
//
// This function creates a feature vector for each chunk found in document.
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
func (ext *Extractor) Extract(doc *html.Document) []*html.Chunk {
	ext.ChunkFeatures = nil
	ext.BoostFeatures = nil

	// No chunks? No features.
	if len(doc.Chunks) == 0 {
		return nil
	}

	// We create one feature for each chunk.
	chunkFeatures := make([]chunkFeature, len(doc.Chunks))
	boostFeatures := make([]boostFeature, len(doc.Chunks))

	// Count the number of words and sentences we encountered for each
	// class. This helps us to detect elements that contain the article text.
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

	// Now we cluster Chunks by Containers to calculate average score per
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

	// Cluster chunks by block and add those blocks to the result whose average
	// score is above prediction level. This makes sure that we don't split large
	// blocks.
	clusterBlock := newClusterMap()
	for i, chunk := range doc.Chunks {
		clusterBlock.Add(chunk.Block, chunk, boostFeatures[i].Score(), float32(chunk.Text.Len()))
	}

	// Keep blocks together.
	result := make([]*html.Chunk, 0, 8)
	for _, chunk := range doc.Chunks {
		if clusterBlock[chunk.Block].Score() > 0.5 {
			result = append(result, chunk)
		}
	}

	// Make them accessible.
	ext.ChunkFeatures = chunkFeatures
	ext.BoostFeatures = boostFeatures
	return result
}
