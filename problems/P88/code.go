package P88

func merge(nums1 []int, m int, nums2 []int, n int) {
	l1 := m - 1
	l2 := n - 1
	l := m + n - 1
	for l1 >= 0 && l2 >= 0 {
		if nums1[l1] > nums2[l2] {
			nums1[l] = nums1[l1]
			l1--
			l--
		} else {
			nums1[l] = nums2[l2]
			l2--
			l--
		}
	}

	for l2 >= 0 {
		nums1[l] = nums2[l2]
		l2--
		l--
	}
}
