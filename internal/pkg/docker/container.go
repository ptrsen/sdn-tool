package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/ptrsen/sdn-tool/internal/pkg/system"
	"io"
	"io/ioutil"
	"os"
)


/*********************************************************************************
*	Pullimage :
*			Function to pull Docker Image from Dockerfile
*
**********************************************************************************/

func Pullimage(dockerfilePath string, dockerfileName string) {

	ctx := context.Background()  		// Create context
	cli, err :=  client.NewEnvClient()  // Create Docker client
	system.CheckError("Fail to create Docker client, check docker installation",err)

	// Create tarfile from Dockerfile
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	defer func() {
		er := tw.Close()
		system.CheckError("Error closing tarfile",er)
	}()

	// Open Dockerfile
	dockerFileReader, err := os.Open(dockerfilePath + "/" + dockerfileName)
	system.CheckError("Error opening Dockefile", err)

	// Read Dockerfile
	readDockerFile, err := ioutil.ReadAll(dockerFileReader)
	system.CheckError("Error reading Dockefile", err)

	// Create Tarfile
	tarHeader := &tar.Header{
		Name: dockerfileName,
		Size: int64(len(readDockerFile)),
	}
	err = tw.WriteHeader(tarHeader)
	system.CheckError("Error writing tar header", err)
	_, err = tw.Write(readDockerFile)
	system.CheckError("Error writing tar body", err)

	dockerFileTarReader := bytes.NewReader(buf.Bytes())


	// Docker image options
	imageOptions := types.ImageBuildOptions{
		Tags: []string{},
		Context:    dockerFileTarReader,
		Dockerfile: dockerfileName,
		/* Other
		CPUSetCPUs:   "2",
		CPUSetMems:   "12",
		CPUShares:    20,
		CPUQuota:     10,
		CPUPeriod:    30,
		Memory:       256,
		MemorySwap:   512,
		ShmSize:      10,*/
		Remove:     true}

	// Build Docker image
	imageBuildResponse, err := cli.ImageBuild(ctx, dockerFileTarReader,imageOptions)
	system.CheckError("Error building docker image",err)

	defer func() {
		er := imageBuildResponse.Body.Close()
		system.CheckError("Error getting docker image build response",er)
	}()

	 //Print just to see the response in Stdout
	_, err = io.Copy(os.Stdout, imageBuildResponse.Body)
	system.CheckError("Error reading build response",err)

	ctx.Done()
	err = cli.Close()
	system.CheckError("Error closing docker client",err)

}
/**********************************************************************************/





/*********************************************************************************
*	CreateContainer :
*			Function to create Docker Container
*			return error, output string
**********************************************************************************/

func CreateContainer( containerName string, imageName string, logVolumeDirectory [2]string, confVolumeDirectory [2]string, networkName string, ipv4 string, ipv6 string) {

	ctx := context.Background() // Create context
	cli, err :=  client.NewEnvClient() // Create Docker client
	system.CheckError("Fail to create Docker client, Check docker installation",err)

	containerConfig := &container.Config{  		// Container configuration
		Image:        	imageName,	 			// Image Name
		Tty:          	true,     	 			// Attach standard streams to a tty.
		AttachStdin:  	true,     	 			// Attach the standard input, makes possible user interaction
		AttachStderr: 	true,    	 			// Attach the standard error
		AttachStdout: 	true,    	 			// Attach the standard output
	}

	// Host configuration
	hostConfig := &container.HostConfig{
		Privileged: true,
		Sysctls:  map[string]string{},

		NetworkMode:  container.NetworkMode(networkName), // networkMode -> docker network name to attach container

	   /*	PortBindings: nat.PortMap{
			nat.Port("6640/tcp"): []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "6640"}},
			nat.Port("6653/tcp"): []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "6653"}},
			nat.Port("8101/tcp"): []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "8101"}},
			nat.Port("8181/tcp"): []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "8181"}},
			nat.Port("9876/tcp"): []nat.PortBinding{{HostIP: "127.0.0.1", HostPort: "9876"}},
		}, */

		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: logVolumeDirectory[0] + "/" + containerName,  //  Container logs -> host local path
				Target: logVolumeDirectory[1],    // Container logs -> path inside Container
			},
			{
				Type:   mount.TypeBind,
				Source: confVolumeDirectory[0] + "/" + containerName,  ///  Container configuration files -> host local path
				Target: confVolumeDirectory[1],    // Container configuration files  -> path inside Container
			},
		},

	}


	// Network configuration
	netConfig := &network.NetworkingConfig{ //}
		EndpointsConfig: map[string]*network.EndpointSettings{
			networkName : {  // },
				IPAMConfig: &network.EndpointIPAMConfig{
					IPv4Address:  ipv4,
					IPv6Address:  ipv6,
				},
			},
		},
	}

	//  Run configuration
	runOptions :=  types.ContainerStartOptions{}  // Default

	// Create Container
	respContainer, err := cli.ContainerCreate(ctx, containerConfig, hostConfig, netConfig, containerName)
	system.CheckError("Error creating container", err)

	// Start Container
	err = cli.ContainerStart(ctx, respContainer.ID, runOptions)
	system.CheckError("Error starting container", err)

	//return err, "Container ID: " + respContainer.ID   // Return container ID if everything is good

	ctx.Done()
	err = cli.Close()
	system.CheckError("Error closing docker client",err)

}

/*********************************************************************************/


/*********************************************************************************
*	CreateDockerNetwork :
*			Function to create Docker Network
*			return error, output string
**********************************************************************************/

func CreateDockerNetwork (networkName string, subnetipv4 string, subnetipv6 string ) {

	ctx := context.Background() // Create context
	cli, err :=  client.NewEnvClient() // Create Docker client
	system.CheckError("Fail to create Docker client, Check docker installation",err)

	// IPAM Driver Configuration
	ipamConf := network.IPAM{
		Driver: "default",
		Config: []network.IPAMConfig{
			{
				Subnet: subnetipv4 ,    // ipv4     <- check Int2ipv4 funcion
			},
			{
				Subnet: subnetipv6 ,    // global ipv6 - link ipv6 fe80::/64  <-  check Int2ipv6 function
			},
		},
		Options: make(map[string]string, 0),
	}

	// Network options
	networkCreateOptions := types.NetworkCreate{
		Driver:         	"bridge",
		EnableIPv6:     	true,
		IPAM: 				&ipamConf,
		Internal:   		false,
		Attachable:     	true,
		CheckDuplicate :	true,
		Options: map[string]string{
			"com.docker.network.bridge.name":   networkName,
		},
	}

	respNetwork, err := cli.NetworkCreate(ctx, networkName , networkCreateOptions)
	system.CheckError("Unable to create Docker Network",err)

	fmt.Println(respNetwork)

	ctx.Done()
	err = cli.Close()
	system.CheckError("Error closing docker client",err)
}


