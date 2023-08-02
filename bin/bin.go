package bin


import (
	os

)



func BalloonMain() {

    if len(os.Args) < 2 {
        fmt.Println("Missing parameter, provide file name!")
        return
    }

	fd, error := os.Open(os.Args[1])

	if error != nil {
		fmt.Println(error)
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

	//skip header
	scanner.Scan()

	//MAIN
	for scanner.Scan() {
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

    }

}//end of main

