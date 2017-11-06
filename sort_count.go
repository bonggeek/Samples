package main
import "fmt"

func countsort(array []byte) {
    count := make(map[byte]int)
    count['a'] = 0
    count['b'] = 0
    count['c'] = 0

    for _,a := range array {
        count[a] = count[a] + 1
    }

    j := 0
    for k, v := range count {
        for m := 0; m < v; m++ {
            array[j] = k;
            j++
        }
    }

}

func main() {
    array := []byte {'a', 'b', 'c', 'c', 'b', 'a', 'c', 'b', 'a', 'c'}
    fmt.Println(string(array))
    countsort(array)
    fmt.Println(string(array))
}
