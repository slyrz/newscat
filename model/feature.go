package model

import (
	"github.com/slyrz/newscat/html"
)

// Feature represents a feature vector.
type Feature [39]float32

// FeatureGenerator fills Features with values.
type FeatureGenerator struct {
	Feature *Feature
	Pos     int

	// Unexported fields.
	// These are internal state variables used by some features.
	lastSim float32
	bestSim float32
}

func (fg *FeatureGenerator) SetFeature(f *Feature) {
	fg.Feature = f
	fg.Pos = 0
}

// Write a value of type int, float32 or bool at given offset and skip the
// requested amount of components afterwards.
func (fg *FeatureGenerator) write(val interface{}, off int, skip int) {
	comp := &fg.Feature[fg.Pos+off]
	switch val := val.(type) {
	case int:
		*comp = float32(val)
	case float32:
		*comp = val
	case bool:
		if val {
			*comp = 1.0
		} else {
			*comp = 0.0
		}
	}
	if skip > 0 {
		fg.Skip(skip)
	}
}

// Write value at current position and move to the next.
func (fg *FeatureGenerator) Write(val interface{}) {
	fg.write(val, 0, 1)
}

// Write value at offset, but don't move.
func (fg *FeatureGenerator) WriteAt(val interface{}, off int) {
	fg.write(val, off, 0)
}

// Skip components.
func (fg *FeatureGenerator) Skip(n int) {
	fg.Pos += n
}

func Features(doc *html.Document) []Feature {
	// No chunks? No features.
	if len(doc.Chunks) == 0 {
		return nil
	}

	// We create one feature for each chunk.
	features := make([]Feature, len(doc.Chunks))

	// Count the number of words and sentences we encountered for each
	// class. This helps us to detect elements that contain the article text.
	classStats := doc.GetClassStats()
	clusterStats := doc.GetClusterStats()

	gen := new(FeatureGenerator)
	for i, chunk := range doc.Chunks {
		// Fill the i-th feature based on the current chunk.
		gen.SetFeature(&features[i])

		// Let the generator fill the current feature vector with our
		// observations.
		gen.SetTitleSimilarity(chunk, doc.Title)
		gen.SetElementType(chunk)
		gen.SetParentType(chunk)
		gen.SetSiblingTypes(chunk)
		gen.SetAncestors(chunk)
		gen.SetTextStat(chunk)
		gen.SetTextStatSiblings(chunk)
		gen.SetClassStat(chunk, classStats)
		gen.SetClusterStat(chunk, clusterStats)
	}

	// Detect min and max for each feature component, i.e. column-wise.
	// TODO: Calculate on the fly in generator?
	empMin := Feature{}
	empMax := Feature{}
	for i := 0; i < len(features); i++ {
		for j, comp := range features[i] {
			switch {
			case comp < empMin[j]:
				empMin[j] = comp
			case comp > empMax[j]:
				empMax[j] = comp
			}
		}
	}
	// Perform MinMax normalization for each feature component.
	for i := 0; i < len(features); i++ {
		feature := &features[i]
		for j := 0; j < len(feature); j++ {
			// If the maximum encountered value isn't greater than one, we assume
			// that the feature is already normalized.
			if empMax[j] > 1.0 {
				feature[j] = (feature[j] - empMin[j]) / (empMax[j] - empMin[j])
			}
		}
	}
	return features
}

func (fg *FeatureGenerator) SetTitleSimilarity(chunk *html.Chunk, title *html.Chunk) {
	if chunk.IsHeading() {
		fg.lastSim = chunk.Text.Similarity(title.Text)
		if fg.lastSim > fg.bestSim {
			fg.bestSim = fg.lastSim
		}
	}
	fg.Write(fg.lastSim)
	fg.Write(fg.bestSim)
}

// Entries with a "plus comment" indicate that the next N elements share
// the same offset intentionally.
var elementTypes = map[string]int{
	"p":   0,
	"a":   1,
	"div": 2,
	"h1":  3, // +5
	"h2":  3,
	"h3":  3,
	"h4":  3,
	"h5":  3,
	"h6":  3,
}

func (fg *FeatureGenerator) SetElementType(chunk *html.Chunk) {
	// One hot encoding of the element type.
	fg.WriteAt(true, elementTypes[chunk.Base.Data])
	fg.Skip(4)
}

var parentTypes = map[string]int{
	"p":    0,
	"span": 1,
	"div":  2,
	"li":   3,
}

func (fg *FeatureGenerator) SetParentType(chunk *html.Chunk) {
	// One hot encoding of the chunk's parent's element type.
	if chunk.Base.Parent != nil {
		fg.WriteAt(true, parentTypes[chunk.Base.Parent.Data])
	}
	fg.Skip(4)
}

func (fg *FeatureGenerator) SetSiblingTypes(chunk *html.Chunk) {
	count := 0
	types := map[string]int{"a": 0, "p": 0, "img": 0}
	for _, siblingType := range chunk.GetSiblingTypes() {
		count += 1
		if val, ok := types[siblingType]; ok {
			types[siblingType] = val + 1
		}
	}
	fg.Write(count)
	fg.Write(types["a"])
	fg.Write(types["p"])
	fg.Write(types["img"])
	if count > 0 {
		fg.Write(float32(types["a"]) / float32(count))
		fg.Write(float32(types["p"]) / float32(count))
		fg.Write(float32(types["img"]) / float32(count))
	} else {
		fg.Skip(3)
	}
}

func (fg *FeatureGenerator) SetAncestors(chunk *html.Chunk) {
	fg.Write((chunk.Ancestors & html.AncestorArticle) != 0)
	fg.Write((chunk.Ancestors & html.AncestorAside) != 0)
	fg.Write((chunk.Ancestors & html.AncestorBlockquote) != 0)
	fg.Write((chunk.Ancestors & html.AncestorList) != 0)
}

func (fg *FeatureGenerator) SetTextStat(chunk *html.Chunk) {
	fg.Write(chunk.Text.Words)
	fg.Write(chunk.Text.Sentences)
	fg.Write(chunk.LinkText)
}

func (fg *FeatureGenerator) SetTextStatSiblings(chunk *html.Chunk) {
	if chunk.Prev != nil {
		fg.Write(chunk.Prev.Block == chunk.Block)
		fg.Write(chunk.Prev.Text.Words)
		fg.Write(chunk.Prev.Text.Sentences)
	} else {
		fg.Skip(3)
	}
	if chunk.Next != nil {
		fg.Write(chunk.Next.Block == chunk.Block)
		fg.Write(chunk.Next.Text.Words)
		fg.Write(chunk.Next.Text.Sentences)
	} else {
		fg.Skip(3)
	}
}

var classQuality = map[string]int{
	"article":      +1,
	"content":      +1,
	"headline":     +1,
	"main-content": +1,
	"post":         +1,
	"story":        +1,
	"story-body":   +1,
	"blog":         -1,
	"caption":      -1,
	"col":          -1,
	"comment":      -1,
	"info":         -1,
	"menu":         -1,
	"metadata":     -1,
	"nav":          -1,
	"navigation":   -1,
	"widget":       -1,
}

func (fg *FeatureGenerator) SetClassStat(chunk *html.Chunk, classes map[string]*html.TextStat) {
	var best *html.TextStat = nil
	var qual int = 0
	for _, class := range chunk.Classes {
		if stat, ok := classes[class]; ok {
			if best == nil || (stat.Words/stat.Count) > (best.Words/best.Count) {
				best = stat
			}
		}
		qual += classQuality[class]
	}
	if best != nil {
		fg.Write(true)
		fg.Write(float32(best.Words) / float32(best.Count))
		fg.Write(float32(best.Sentences) / float32(best.Count))
	} else {
		fg.Write(false)
		fg.Skip(2)
	}
	fg.Write(qual > 0)
}

func (fg *FeatureGenerator) SetClusterStat(chunk *html.Chunk, clusters map[*html.Chunk]*html.TextStat) {
	if stat, ok := clusters[chunk]; ok {
		fg.Write(stat.Words)
		fg.Write(stat.Sentences)
		fg.Write(stat.Count)
		fg.Write(float32(stat.Words) / float32(stat.Count))
		fg.Write(float32(stat.Sentences) / float32(stat.Count))
	} else {
		fg.Skip(5)
	}
}
