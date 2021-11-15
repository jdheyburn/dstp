package conn

import (
	"context"
	"fmt"
	"net"

	"github.com/ycd/dstp/pkg/common"
)

func RunTest(ctx context.Context, addr common.Address, port common.Port) (common.Output, error) {
	longAddr := fmt.Sprintf("%s:%s", string(addr), string(port))

	output := "connecting over tcp to " + longAddr

	c, err := net.Dial("tcp", longAddr)

	if err != nil {
		return "", err
	}

	c.Close()

	return common.Output(output), nil
}
