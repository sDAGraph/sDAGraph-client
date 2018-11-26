package params

import (
	"fmt"
        "testing"
)

type newput struct{
        tmp string
}

func TestParams(t *testing.T) {
	c := Chain() 
	fmt.Println(c)
	fmt.Println(c.Version)
	fmt.Println(c.Version.Sue["dev"])
	fmt.Println(c.Version.Sue["dev"].MongoIp)
}
