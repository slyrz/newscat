package model

import (
	"github.com/slyrz/newscat/html"
	"math"
	"strings"
)

// ChunkExtractor utilizes the trained model to extract relevant html.Chunks from
// an html.Article.
type ChunkExtractor struct {
	ChunkFeatures []chunkFeature
	BoostFeatures []boostFeature
}

// LinkExtractor extracts possible links to news articles from a html.Website.
type LinkExtractor struct {
}

// NewChunkExtractor creates and initalizes a new ChunkExtractor.
func NewChunkExtractor() *ChunkExtractor {
	return new(ChunkExtractor)
}

// Extract returns a list of relevant text chunks found in article.
//
// How it works
//
// This function creates a feature vector for each chunk found in article.
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
func (ext *ChunkExtractor) Extract(article *html.Article) []*html.Chunk {
	ext.ChunkFeatures = nil
	ext.BoostFeatures = nil

	// No chunks? No features.
	if len(article.Chunks) == 0 {
		return nil
	}

	// We create one feature for each chunk.
	chunkFeatures := make([]chunkFeature, len(article.Chunks))
	boostFeatures := make([]boostFeature, len(article.Chunks))

	// Count the number of words and sentences we encountered for each
	// class. This helps us to detect elements that contain the article text.
	classStats := article.GetClassStats()
	clusterStats := article.GetClusterStats()

	chunkFeatureWriter := new(chunkFeatureWriter)
	for i, chunk := range article.Chunks {
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
	for i, chunk := range article.Chunks {
		clusterContainer.Add(chunk.Container, chunk, chunkFeatures[i].Score())
	}

	boostFeatureWriter := new(boostFeatureWriter)
	for i, chunk := range article.Chunks {
		boostFeatureWriter.Assign(boostFeatures[i][:])
		boostFeatureWriter.WriteChunk(chunk)
		boostFeatureWriter.WriteCluster(chunk, clusterContainer[chunk.Container])
		boostFeatureWriter.WriteTitleSimilarity(chunk, article.Title)
	}

	// Cluster chunks by block and add those blocks to the result whose average
	// score is above prediction level. This makes sure that we don't split large
	// blocks.
	clusterBlock := newClusterMap()
	for i, chunk := range article.Chunks {
		clusterBlock.Add(chunk.Block, chunk, boostFeatures[i].Score(), float32(chunk.Text.Len()))
	}

	// Keep blocks together.
	result := make([]*html.Chunk, 0, 8)
	for _, chunk := range article.Chunks {
		if clusterBlock[chunk.Block].Score() > 0.5 {
			result = append(result, chunk)
		}
	}

	// Make them accessible.
	ext.ChunkFeatures = chunkFeatures
	ext.BoostFeatures = boostFeatures
	return result
}

// NewLinkExtractor creates and initalizes a new LinkExtractor.
func NewLinkExtractor() *LinkExtractor {
	return new(LinkExtractor)
}

// Extract returns a list of links to possible news articles found on website.
//
// TODO
func (ext *LinkExtractor) Extract(website *html.Website) []*html.Link {
	if len(website.Links) == 0 {
		return nil
	}

	// Calculate the host name frequency. We assume that the more often
	// we encounter a host name, the more important are its links.
	hostFreq := make(map[string]float64)
	unitFrac := 1.0 / float64(len(website.Links))
	for _, link := range website.Links {
		hostFreq[link.URL.Host] = hostFreq[link.URL.Host] + unitFrac
	}

	// TODO
	result := make([]*html.Link, 0, 8)
	for _, link := range website.Links {
		score := calcEntropy(link.URL.Path) * hostFreq[link.URL.Host]
		if score >= 3.0 {
			result = append(result, link)
		}
	}
	return result
}

var (
	entropy = map[rune]float64{
		'a': -0.10634380971489953,
		'b': -0.1530425750819142,
		'c': -0.054512771770650156,
		'd': -0.2277081908035247,
		'e': -0.019143424124981535,
		'f': -0.10631901817532326,
		'g': -0.08950958783746302,
		'h': -0.02138890867431747,
		'i': -0.2040842535452831,
		'j': -0.08066894356623576,
		'k': -0.14676065483416048,
		'l': -0.054715501033775864,
		'm': -0.22756818123774952,
		'n': -0.2036576755124059,
		'o': -0.011796162629390488,
		'p': -0.19352586509733832,
		'q': -0.061206381465562155,
		'r': -0.12701158290463668,
		's': -0.10100191941724014,
		't': -0.022872200444912123,
		'u': -0.05095402954111724,
		'v': -0.09120240867794169,
		'w': -0.20676837174037171,
		'x': -0.0856356305628621,
		'y': -0.052914228421479595,
		'z': -0.08290376493347784,
		'0': -0.20509019606654505,
		'1': -0.04936506198036805,
		'2': -0.052649170776553546,
		'3': -0.28439304075529853,
		'4': -0.11790007547597285,
		'5': -0.09902560480951333,
		'6': -0.0624913286633471,
		'7': -0.1247640940220319,
		'8': -0.17807011198199005,
		'9': -0.2233673033869976,
	}
)

// calcEntropy returns the entropy of a string.
func calcEntropy(s string) float64 {
	result := 0.0
	for _, r := range strings.ToLower(s) {
		result += entropy[r]
	}
	return math.Abs(result)
}
