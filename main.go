package main

func main() {

}

func withIf(a []int) int {
	if len(a) > 0 {
		return a[0]
	}
	return 0
}

func withForElement(a []int) int {
	for _, i := range a {
		return i
	}
	return 0
}

func withForIndex(a []int) int {
	for i := range a {
		return a[i]
	}
	return 0
}
