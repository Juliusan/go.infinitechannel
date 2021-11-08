package util

func PriorityFunMod2(i interface{}) bool {
	return PriorityFunMod(i, 2)
}

func PriorityFunMod3(i interface{}) bool {
	return PriorityFunMod(i, 3)
}

func PriorityFunMod10(i interface{}) bool {
	return PriorityFunMod(i, 10)
}

func PriorityFunMod(i interface{}, mod SimpleHashable) bool {
	return i.(SimpleHashable)%mod == 0
}
