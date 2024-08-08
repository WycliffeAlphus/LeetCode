func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var median float64
    for _, ch := range nums2 {
        nums1 = append(nums1, ch)
    }
    
    for i := 0; i < len(nums1); i++ {
        for j := 0; j < len(nums1); j++ {
            if j+1 < len(nums1)&&nums1[j] > nums1[j+1] {
                nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
            } 
        }
    }
    
    if len(nums1)%2 == 0 {
        ind:= len(nums1)/2 
    median = (float64(nums1[ind-1]) + float64(nums1[ind])) / float64(2)
    } else{
        ind := len(nums1)/2
        median = float64(nums1[ind])
    }
    return median
}