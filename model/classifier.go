package model

func (ftr chunkFeature) Score() float32 {
	score := logitIntercept
	for i := range ftr {
		score += (ftr[i] * logitCoefficients[i])
	}
	return score
}

func (ftr chunkFeature) Predict() bool {
	return ftr.Score() > 0.0
}

func (ftr scoreFeature) Score() float32 {
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

func (ftr scoreFeature) Predict() bool {
	return ftr.Score() > 0.5
}
