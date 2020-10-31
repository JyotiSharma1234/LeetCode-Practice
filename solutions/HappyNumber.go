// Example: 
// Input: 19
// Output: true
// Explanation: 
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1    Here we got 1, So 19 is ahappy number

func isHappy(n int) bool {
    flag := false
    var arr [100]int
    num := n
    x := 0
    for  ; num!=0 ;  {
        num = digit_square_sum(num)
        if(num==1){
            flag = true
            break;
        }else if(num == n || include_ele( arr, num )){
            flag=false
            break;
        }
        arr[x] = num
        x = x+1
    }
    
    return flag
}

func digit_square_sum(num int) int {
    rem := 0
    sum := 0
    for ; num!=0 ; {
        rem = num % 10
        sum = sum + rem * rem
        num /= 10
    }
    
    return sum
}

func include_ele(arr [100]int, ele int) bool{
    flag := false
    for i :=0 ; i<len(arr); i++{
        if(arr[i] == ele){
            flag = true
            break
        }
    }
    
    return flag
}
