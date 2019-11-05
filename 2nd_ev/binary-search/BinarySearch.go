/*package main
import "fmt"

func pesquisaBinaria(slice int, key []int) int {

    low := 0
    high := len(key) - 1

    for low <= high{
        median := (low + high) / 2

        if key[median] < slice {
            low = median + 1
        }
        if key[median] == slice {
            return key[median]

        }else{
        high = median - 1
        }
    }
    return -1

func main(){
    fmt.Printf("%d/n",pesquisaBinaria(63,{1,2, 9, 20, 31, 45, 63, 70, 100} ))
}
*/
package main
import (
    "fmt"
)
func BinarySearch(slice int, key []int) int {

    low := 0
    high := len(key) - 1

    for low <= high{
        median := (low + high) / 2

        if key[median] < slice {
            low = median + 1
        }else if key[median] == slice {
            return median

        }else{
        high = median
        }
    }
    return -1
}
//tests
func main(){

    var items = []int{1,2, 9, 20, 31, 45, 63, 70, 100,120}
    fmt.Printf("%d\n",BinarySearch(45,items))
}