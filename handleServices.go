package balloon

import (
	"fmt"
	//"strings"
	//"strconv"
	//"encoding/hex"
	"bufio"
)



func HandleServices( NUM_READ int64 , scanner *bufio.Scanner ) {

	//init variables
	compare := false
	var counter int64 = 0
	var new_ip []byte
	var str_ip string
	var str_port string


    //return ip:port
	for counter < NUM_READ {

		str_ip, str_port, new_ip, compare, scanner =
			 HandleCompress( new_ip, str_port, compare, scanner)

		//less lines to read than num_lines
		if str_ip == "" {
			break
		}

		fmt.Printf("%s:%s\n",str_ip,str_port)

		counter += 1

	}

}//end

