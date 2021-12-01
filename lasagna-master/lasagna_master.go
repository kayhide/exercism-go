package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer

}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauces := 0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles++
		case "sauce":
			sauces++
		}
	}
	return noodles * 50, float64(sauces) * 0.2
}

func AddSecretIngredient(xs, ys []string) []string {
	zs := make([]string, len(ys)+1)
	for i, y := range ys {
		zs[i] = y
	}
	zs[len(ys)] = xs[len(xs)-1]
	return zs
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	scale := float64(portions) / 2.0
	res := make([]float64, len(amounts))
	for i, q := range amounts {
		res[i] = q * scale
	}
	return res
}
