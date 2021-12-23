package tool

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeAdd(t *testing.T) {
	fmt.Println(TimeAdd(time.Now().Format(Date), "day", 3))
}
