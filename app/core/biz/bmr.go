package biz

func GetBasalMetabolismMifflinStJeorMale(age int, height, weight float32) float32 {
	// https://diagon.arthursonzogni.com/#Math
	//
	//  TMBm = 10 ⋅ weight + 6.25 ⋅ height - 5 ⋅ age + 5
	//                   kg              cm        years

	return 10*weight + 6.25*height - 5*float32(age) + 5
}

func GetBasalMetabolismMifflinStJeorFemale(age int, height, weight float32) float32 {
	// https://diagon.arthursonzogni.com/#Math
	//
	//  TMBf = 10 ⋅ weight + 6.25 ⋅ height - 5 ⋅ age - 161
	//                   kg              cm        years

	return 10*weight + 6.25*height - 5*float32(age) - 161
}

func GetBasalMetabolismRevisedHarrisBenedictMale(age int, height, weight float32) float32 {
	// https://diagon.arthursonzogni.com/#Math
	//
	//  TMBm = 13.397 ⋅ weight + 4.799 ⋅ height - 5.677 ⋅ age + 88.362
	//                       kg              cm        years

	return 13.397*weight + 4.799*height - 5.677*float32(age) + 88.362
}

func GetBasalMetabolismRevisedHarrisBenedictFemale(age int, height, weight float32) float32 {
	// https://diagon.arthursonzogni.com/#Math
	//
	//  TMBm = 9.247 ⋅ weight + 3.098 ⋅ height - 4.330 ⋅ age + 447.593
	//                       kg              cm        years

	return 9.247*weight + 3.098*height - 4.330*float32(age) + 447.593
}
