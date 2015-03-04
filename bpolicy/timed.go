package bpolicy

import (
	"fmt"
)

type TimedPolicy struct {
	factory  *TimedPolicyFactory
	next     IPolicy
	discount float64
}

func (policy *TimedPolicy) DiscountQuota(discount float64) {
	fmt.Println("DiscountQuota in TimedPolicy by %s", discount)
	policy.discount = policy.discount * discount
	if policy.next != nil {
		policy.next.DiscountQuota(discount)
	}
}

func (policy *TimedPolicy) CheckPolicy(identity string) error {
	var check error
	fmt.Println("CheckPolicy in TimedPolicy by %s", identity)
	if policy.next != nil {
		check = policy.next.CheckPolicy(identity)
	}
	return check
}

type TimedPolicyFactory struct {
	Start    *Clock
	End      *Clock
	Discount float64
}

func (factory *TimedPolicyFactory) Instance(next IPolicy) *TimedPolicy {
	return &TimedPolicy{factory, next, factory.Discount}
}
