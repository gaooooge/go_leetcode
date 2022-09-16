package easy

func firstBadVersion(n int) int {
	left := 1
    right := n

    for left < right {
        mix := (right-left)/2 + left
        if isBadVersion(mix) {
            right = mix
        }else{
            left = mix + 1
        }
        fmt.Println(mix)
    }
    return left
}

func TestFirstBadVersion(t *testing.T) {
	n := 5
	res := firstBadVersion(n)
	fmt.Println(res)
}
