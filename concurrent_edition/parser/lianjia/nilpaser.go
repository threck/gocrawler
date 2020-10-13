package lianjia

import (
	"fmt"
	"gocrawler/concurrent_edition/engine"
)

func NilParser(b []byte, s string) (engine.ParseResult, error) {
	fmt.Printf("i am nilParser...... from parser: %s \n", s)
	return engine.ParseResult{}, nil
}
