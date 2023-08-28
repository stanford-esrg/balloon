package balloon

import (
	//"fmt"
	"strings"
	"strconv"
	"encoding/hex"
	"bufio"
)



func HandleCompress( new_ip []byte, str_port string, compare bool, scanner *bufio.Scanner ) ( string, string, []byte, bool, *bufio.Scanner ) {



	//init variables
	str_ip := ""
	var line string


	//MAIN
	scanner.Scan()
    line = scanner.Text()

	//nothing left to scan
	if len(line) == 0 {

		return "", "", nil, false, nil

	}

    line = strings.TrimSuffix(line, "\n")
    s := strings.Split(line,",")

    if len(s) != 2 {
       panic("Error parsing compressed-input list")
    }

    portHex, ipHex := s[0], s[1]

    ipByte, err := hex.DecodeString(ipHex)

    if err != nil {
        panic(err)
    }

    //fmt.Printf("Prev ipByte: %d\n",new_ip)
    //fmt.Printf("Cur ipByte: %d\n",ipByte)

    portByte, emptyPort := strconv.ParseInt(portHex, 16, 16)
    //fmt.Printf("Cur port: %d\n",portByte)
    //fmt.Printf("Prev port: %s\n",str_port)

    //fill in the rest of the IP
    if compare {

        for i, bytte := range ipByte {
            new_ip[i] = bytte
        }

    } else {

        compare = true
        new_ip = ipByte
        str_port =  strconv.Itoa(int(portByte))
    }

    //conver array of bytes to ip address
    for i, bytte := range new_ip {
        str_ip +=strconv.Itoa(int(bytte))
        if i < 3 {
            str_ip += "."
        }
    }
    if emptyPort == nil {
        str_port =  strconv.Itoa(int(portByte))
    }

    //return ip:port
	return str_ip, str_port, new_ip, compare, scanner

}//end

