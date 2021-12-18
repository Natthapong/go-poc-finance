package goalseek

import(
	"math"
)

type Result struct{
	TargetValue float64
	AccuracyLevel float64
	Iterations int
	IsGoalReached bool
	ClosetValue float64
}

const delta float64 = 0.0001

func Run(funcInside func(x float64)float64, targetValue float64, limitIterations int) *Result {
	
	iterations:=0
	var guess float64 = 0
	var accuracyLevel float64 = 0.0000001

	result1:=(funcInside(guess)-targetValue)

	for math.Abs(result1)>accuracyLevel && iterations<limitIterations {

		newGuess:=(guess+delta)
		result2:=(funcInside(newGuess)-targetValue)

		if (result2-result1) !=0{
			guess=guess - result1 * (newGuess - guess) / (result2 - result1)
		}else{
			break
		}

		result1 = (funcInside(guess)-targetValue)

		iterations++
	}

	if iterations>limitIterations {
		iterations=limitIterations 
	}

	isGoalReache:=false
	if math.Abs(result1) <= accuracyLevel{isGoalReache=true}

	return &Result{
		TargetValue:targetValue,
		AccuracyLevel:accuracyLevel,
		Iterations:iterations,
		IsGoalReached:isGoalReache,
		ClosetValue:guess,
	}

}
