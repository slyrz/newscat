package model

import (
	_ "fmt"
	"github.com/slyrz/newscat/html"
)

// Extractor utilizes the trained model to extract relevant html.Chunks from
// an html.Document.
type Extractor struct {
	ChunkFeatures []ChunkFeature
	ScoreFeatures []ScoreFeature
}

// NewExtractor creates and initalizes a new Extractor.
func NewExtractor() *Extractor {
	return new(Extractor)
}

func (ext *Extractor) Extract(doc *html.Document) []*html.Chunk {
	ext.ChunkFeatures = nil
	ext.ScoreFeatures = nil

	// No chunks? No features.
	if len(doc.Chunks) == 0 {
		return nil
	}

	// We create one feature for each chunk.
	chunkFeatures := make([]ChunkFeature, len(doc.Chunks))
	scoreFeatures := make([]ScoreFeature, len(doc.Chunks))

	// Count the number of words and sentences we encountered for each
	// class. This helps us to detect elements that contain the article text.
	classStats := doc.GetClassStats()
	clusterStats := doc.GetClusterStats()

	chunkFeatureWriter := new(ChunkFeatureWriter)
	for i, chunk := range doc.Chunks {
		// Fill the i-th feature based on the current chunk.
		chunkFeatureWriter.Assign(chunkFeatures[i][:])

		// Write the observations to the feature vector.
		chunkFeatureWriter.WriteElementType(chunk)
		chunkFeatureWriter.WriteParentType(chunk)
		chunkFeatureWriter.WriteSiblingTypes(chunk)
		chunkFeatureWriter.WriteAncestors(chunk)
		chunkFeatureWriter.WriteTextStat(chunk)
		chunkFeatureWriter.WriteTextStatSiblings(chunk)
		chunkFeatureWriter.WriteClassStat(chunk, classStats)
		chunkFeatureWriter.WriteClusterStat(chunk, clusterStats)
	}

	// Detect min and max for each feature component, i.e. column-wise.
	empMin := ChunkFeature{}
	empMax := ChunkFeature{}
	for i := 0; i < len(chunkFeatures); i++ {
		for j, comp := range chunkFeatures[i] {
			switch {
			case comp < empMin[j]:
				empMin[j] = comp
			case comp > empMax[j]:
				empMax[j] = comp
			}
		}
	}

	// Perform MinMax normalization for each feature component.
	for i := 0; i < len(chunkFeatures); i++ {
		feature := &chunkFeatures[i]
		for j := 0; j < len(feature); j++ {
			// If the maximum value isn't greater than one, we assume that
			// the feature is already normalized.
			if empMax[j] > 1.0 {
				feature[j] = (feature[j] - empMin[j]) / (empMax[j] - empMin[j])
			}
		}
	}

	// Now we cluster Chunks by Containers to calculate average score per
	// container.
	clusterContainer := NewClusterMap()
	for i, chunk := range doc.Chunks {
		clusterContainer.Add(chunk.Container, chunk, chunkFeatures[i].Score())
	}

	scoreFeatureWriter := new(ScoreFeatureWriter)
	for i, chunk := range doc.Chunks {
		scoreFeatureWriter.Assign(scoreFeatures[i][:])
		scoreFeatureWriter.WriteChunk(chunk)
		scoreFeatureWriter.WriteCluster(chunk, clusterContainer[chunk.Container])
		scoreFeatureWriter.WriteTitleSimilarity(chunk, doc.Title)
	}

	clusterBlock := NewClusterMap()
	for i, chunk := range doc.Chunks {
		clusterBlock.Add(chunk.Block, chunk, scoreFeatures[i].Score(), float32(chunk.Text.Len()))
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
	ext.ScoreFeatures = scoreFeatures
	return result
}
