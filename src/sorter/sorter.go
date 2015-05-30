package main 

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"

var infile2 *string = flag.String("i", "unsorted.dat","File  contains values for sorting")
var outfile2 *string = flag.String("o","sorted.dat","File to receive sorted valuew")  
var algorithm2 *string = flag.String("a","qsort","Sort algorithm")

func readValues(infile string)(value []int, err error){
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file",infile)
		return
	}
	
	defer file.Close()
	
	br := bufio.NewReader(file)	
	
//    values = make([]int, 0)
	var values []int
	
	for {
		line, isPrefix, err1 := br.ReadLine()
		
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
//			return
		}
			
		if isPrefix {
			fmt.Println("A too long line, seems unexpected")
			return
		}		
		
		str :=  string(line)
		
		value, err1 := strconv.Atoi(str) 
		
		if err1 != nil {
			err = err1
//			return
			break
			
		}

		values = append(values,value)
	}
	return
}

func writeValue(values []int, outfile string) error {
	file, err := os.Create(outfile)	
	if err !=  nil {
		fmt.Println("failed to create the output file", outfile)
		return err
	}
	defer file.Close()
	
	for _, value := range values {
		
	}	
	return nil
}

func main(){
	flag.Parse()
	
	if infile2 != nil {
		fmt.Println("infile =", *infile2, "outfile =", *outfile2, "algorithm =", *algorithm2)		
	}	
}	

