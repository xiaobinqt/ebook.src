package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	var (
		set = make(map[int]struct{}, 0)
		ret = make([]int, 0)
	)

	for _, each := range nums1 {
		set[each] = struct{}{}
	}

	for _, each := range nums2 {
		if _, ok := set[each]; ok {
			ret = append(ret, each)
			delete(set, each)
		}
	}

	return ret
}
