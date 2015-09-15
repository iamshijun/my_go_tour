package main

//import "golang.org/x/tour/reader"
import "fmt"

//实现一个 Reader 类型，它不断生成 ASCII 字符 'A' 的流。

type MyReader struct {
	Letter byte
}

type CharArray []int

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader *MyReader) Read(buffer []byte) (int, error) {
	for i, _ := range buffer {
		buffer[i] = reader.Letter
	}
	return len(buffer), nil
}

func main() {
	//c := 'A'
	//fmt.Printf("%T", c)
	mr := &MyReader{Letter: 'A'}

	buffer := make([]byte, 3, 5)
	num, _ := mr.Read(buffer)

	fmt.Printf("Read %d letter , %v", num, buffer)

}
