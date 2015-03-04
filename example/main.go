package main

import (
	"fmt"
	bp "github.com/mapix/bpolicy-go/bpolicy"
	"time"
)

func main() {
	rated_factory := bp.RatedPolicyFactory{Quota: 1, Interval: 10 * time.Second}
	timed_factory := bp.TimedPolicyFactory{Start: &bp.Clock{1, 0, 0}, End: &bp.Clock{24, 0, 0}, Discount: 0.1}
	p1 := rated_factory.Instance(nil)
	p2 := timed_factory.Instance(p1)
	p2.DiscountQuota(0.1)
	p2.CheckPolicy("xxxx")
	fmt.Println(p1)
}
