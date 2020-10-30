func singleNumber(nums []int) int {
    num := 0
    dict := make( map[int]int )
​
    for i := 0 ; i < len(nums) ; i++ {
        if _, ok := dict[nums[i]]; ok {
            dict[nums[i]]++
        }else{
            dict[nums[i]] = 1
        }
        
    }
    for key, value := range dict {
        if( value == 1) {
            num = key
        }
    }
    
    return num;
}
