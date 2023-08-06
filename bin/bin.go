package bin


import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"encoding/hex"
	//"github.com/stanford-esrg/balloon"
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
	NUM_READ := 0
	if len(os.Args) == 3 {
		NUM_READ, error =  strconv.Atoi(os.Args[2])
		if error != nil {
			panic(error)
		}
    }



	defer fd.Close()


	// read compressed file
	scanner := bufio.NewScanner(fd)

	//init variables
	compare := false
	var line string
	var str_ip string
	var str_port string
	new_ip := []byte{}
	counter := 0

	//skip header
	scanner.Scan()

	//MAIN
	for scanner.Scan() {
		counter += 1
        line = scanner.Text()
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

        str_ip = ""
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
        fmt.Printf("%s:%s\n",str_ip,str_port)

		if NUM_READ != 0 && counter > NUM_READ {
			break
		}
    }

}//end of main

