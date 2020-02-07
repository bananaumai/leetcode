package main

func minInsertions1(s string) int {
	bs := []byte(s)
	cnt1 := count(bs)
	cnt2 := count(reverse(bs))
	if cnt1 < cnt2 {
		return cnt1
	} else {
		return cnt2
	}
}

func reverse(bs []byte) []byte {
	var res []byte
	for i := len(bs) - 1; i >= 0; i-- {
		res = append(res, bs[i])
	}
	return res
}

func count(bs []byte) int {
	if len(bs) <= 1 {
		return 0
	}

	if len(bs) == 2 && bs[0] == bs[1] {
		return 0
	}

	var lo []byte
	var mid []byte
	var hi []byte
	var found bool
Loop:
	for i, b := range bs {
		for j := len(bs) - 1; j > i; j-- {
			if b == bs[j] {
				lo = bs[:i]
				mid = bs[i+1 : j]
				hi = bs[j+1:]
				found = true
				break Loop
			}
		}
	}

	if found {
		return len(lo) + count(mid) + len(hi)
	} else {
		return len(bs) - 1
	}
}
