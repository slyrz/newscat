package model

func ChunkFeatureScore(ftr ChunkFeature) float32 {
	score := logitIntercept
	for i := range ftr {
		score += (ftr[i] * logitCoefficients[i])
	}
	return score
}

func ChunkFeaturePredict(ftr ChunkFeature) bool {
	return ChunkFeatureScore(ftr) > 0.0
}

func ScoreFeatureScore(ftr ScoreFeature) float32 {
	score := float32(0.0)
	score += decisionTreeA(ftr)
	score += decisionTreeB(ftr)
	score += decisionTreeC(ftr)
	score += decisionTreeD(ftr)
	score += decisionTreeE(ftr)
	score += decisionTreeF(ftr)
	score += decisionTreeG(ftr)
	score += decisionTreeH(ftr)
	score += decisionTreeI(ftr)
	score += decisionTreeJ(ftr)
	return score / 10.0
}

func ScoreFeaturePredict(ftr ScoreFeature) bool {
	return ScoreFeatureScore(ftr) > 0.5
}
