package bpolicy

import (
	"fmt"
	"time"
)

type RatedPolicy struct {
	factory *RatedPolicyFactory
	next    IPolicy
	quota   int64
}

func (policy *RatedPolicy) DiscountQuota(discount float64) {
	fmt.Println("DiscountQuota in RatedPolicy by %s", discount)
	if policy.next != nil {
		policy.next.DiscountQuota(discount)
	}
}

func (policy *RatedPolicy) CheckPolicy(identity string) error {
	var check error
	fmt.Println("CheckPolicy in RatedPolicy by %s", identity)
	if policy.next != nil {
		check = policy.next.CheckPolicy(identity)
	}
	return check
}

type RatedPolicyFactory struct {
	Quota    int64
	Interval time.Duration
}

func (factory *RatedPolicyFactory) Instance(next IPolicy) *RatedPolicy {
	return &RatedPolicy{factory, next, factory.Quota}
}
