package system

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"time"
)


/**********************************************************************************
*	CheckError :
				Check error, exits if something wrong
***********************************************************************************/

func CheckError(output string, err error) {
	if err  != nil {
		fmt.Print(output)
		fmt.Println(err)
		os.Exit(1)
	}
}

/**********************************************************************************/



/*********************************************************************************
*	ShellExec :
*			Execute shell command from Go, Standard Output
*			returns error
**********************************************************************************/

func ShellExec(dir string, cmd string, args ...string) error {

	ctx := context.Background()

	shell := exec.CommandContext(ctx,cmd,args...)
	shell.Dir = dir
	shell.Stdout = os.Stdout
	shell.Stderr = os.Stderr

	err := shell.Run()
	ctx.Done()

	return err
}
/**********************************************************************************/


/*********************************************************************************
*	ShellExecOutput :
*			Execute command from Go
*			return error, output string
**********************************************************************************/

func ShellExecOutput(dir string, cmd string, args ...string) (output string, er error) {

	ctx := context.Background()

	shell := exec.CommandContext(ctx,cmd,args...)

	shell.Dir = dir
	outStr, err := shell.Output()
	str := string(outStr)

	if err == nil {
		str = str[ : len(str)-1]
	}

	ctx.Done()

	return str, err
}

/**********************************************************************************/



/*********************************************************************************
*	GenerateMac:
*			Generate Random MAC Address
*			returns MAC
**********************************************************************************/

func GenerateMac() net.HardwareAddr {
	buf := make([]byte, 6)
	var mac net.HardwareAddr
	rand.Seed(time.Now().Unix())
	_, err  := rand.Read(buf)
	if err != nil {}

	// Set the local bit
	buf[0] = (buf[0] | 2) & 0xfe // Set local bit, ensure unicast address
	mac = append(mac, buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])

	return mac
}

/**********************************************************************************/




/*


func Int2ipv4(lo uint16) net.IP {
	ip := make(net.IP, net.IPv4len)  	 //  ipv4 -> 4 bytes (32 bits)  10.12.0.0/16
	ip[0] = 10                           //  10 -> 10.0.0.0
	ip[1] = 12							 //	 12 -> 10.12.0.0
	binary.BigEndian.PutUint16(ip[2:], lo)   //  Max hosts [2^16 -2]  -> broadcast addr 10.12.255.255 , Default Gateway  10.12.255.254
	return ip
}

func Int2ipv6(lo uint64) net.IP {
	ip := make(net.IP, net.IPv6len)  	// ipv4 -> 16 bytes (128 bits)  2001:db8::0/64
	ip[0]= 32							//  0x20 -> 2000::0
	ip[1]= 1                            //  0x01 -> 2001::0
	ip[2]= 13							//  0x0d -> 2001:0d::0
	ip[3]= 184 							//  0xb8 -> 2001:db8::0
	ip[4]= 0
	ip[5]= 0
	ip[6]= 0
	ip[7]= 0
	binary.BigEndian.PutUint64(ip[8:],lo)  //  Max hosts [2^64]
	return ip
}


 */