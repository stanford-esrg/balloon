package balloon

import (
	"fmt"
	//"strings"
	//"strconv"
	//"encoding/hex"
	"bufio"
	"net"
	"math"
	"math/big"
	"math/rand"
	//"os"
)


func IP4toInt(ip string) uint32 {
	IPv4Int := big.NewInt(0)
	IPv4Int.SetBytes(net.ParseIP(ip).To4())
	return uint32(IPv4Int.Int64())
}


func IntToIP(ip uint32) string {
	result := make(net.IP, 4)
	result[3] = byte(ip)
	result[2] = byte(ip >> 8)
	result[1] = byte(ip >> 16)
	result[0] = byte(ip >> 24)
	return result.String()
}

func subs2Int( subnets []string ) []uint32 {

	var subnetsint []uint32
	for _, sub := range subnets {
		subnetsint = append( subnetsint, IP4toInt( sub ) )

	}

	return subnetsint

}

func decompressSubs( bit_permutations []int, subnets []string, str_port string) {

	subnetsint :=  subs2Int( subnets )

	var b_ip uint32
	var str_ip string
	for _, bit := range bit_permutations {


		for _, sub := range subnetsint {
			//fmt.Printf("%d\n",sub)
			//fmt.Printf("%d\n",bit)
			b_ip = sub + uint32(bit)
			str_ip = IntToIP(b_ip)
			fmt.Printf("%s:%s\n",str_ip,str_port)
		}

		//fmt.Printf("%d\n",bit)
	}
	return

}


func HandleSubnets( NUM_READ int64 , scanner *bufio.Scanner, SUB_SIZE int64 ) {
	//init variables
	compare := false
	var counter int64 = 0
	var new_ip []byte
	var str_ip string
	var prev_port string = "init"
	var str_port string

	//create a random walk around a sequence of random numbers
	bit := int( math.Pow(2, 32-float64(SUB_SIZE)))
	bit_permutations := rand.Perm(bit)

	var subnets []string

	//Not sure what role NUM_READ plays
	//# of subnets to read?

    //return ip:port
	for counter < NUM_READ {
		str_ip, str_port, new_ip, compare, scanner =
			 HandleCompress( new_ip, str_port, compare, scanner)

		//fmt.Printf("file: %s:%s\n",str_ip,str_port)


		if prev_port == "init" {
			prev_port = str_port
		//the group of subnets per port has changed
		} else if str_port != prev_port {
			decompressSubs( bit_permutations, subnets, prev_port )
			prev_port = str_port
			//empty the subnets
			subnets = subnets[:0]
		}

		// add subnet to array
		subnets = append( subnets, str_ip )

		//run out of lines to read in file
		if str_ip == "" {
			break
		}

		counter = counter + 1
	}

}//end

