package bin


import (
	"os"
	"fmt"
	"bufio"
	//"strings"
	"strconv"
	"math"
	"github.com/stanford-esrg/balloon"
)



func BalloonMain() {

    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

	fd, error := os.Open(os.Args[1])
	if error != nil {
		panic(error)
	}

	//optional second parameter to provide the num services to scan
	var NUM_READ int64 = 0
	if len(os.Args) > 2 {
		if os.Args[2] != "-" {

			NUM_READ, error =  strconv.ParseInt(os.Args[2], 10, 64)
			fmt.Fprintf(os.Stderr,"Writing up to %d services\n", NUM_READ)
			if error != nil {
				panic(error)
			}
		} else {

			NUM_READ = math.MaxInt64

		}
    } else {

		NUM_READ = math.MaxInt64

	}



	defer fd.Close()


	// read compressed file
	scanner := bufio.NewScanner(fd)

	//skip header
	scanner.Scan()

	if len(os.Args) > 3 {

		SUB_SIZE, error :=  strconv.ParseInt(os.Args[3], 10, 64)
        fmt.Fprintf(os.Stderr,"Scanning /%d subnetworks\n", SUB_SIZE)
        if error != nil {
            panic(error)
        }

		balloon.HandleSubnets( NUM_READ, scanner, SUB_SIZE )

	} else {


		balloon.HandleServices( NUM_READ, scanner )

    }

}//end of main

