package connectors

import (
	"fmt"
	ignite "github.com/amsokol/ignite-go-client/binary/v1"
)

func ConnectToIgniteCluster(connectionInfo ignite.ConnInfo) ignite.Client {

	c, err := ignite.Connect(connectionInfo)
	if err != nil {
		fmt.Printf("failed connect to server: %v", err)
	}

	return c
}
