package bpolicy

type IPolicy interface {
	DiscountQuota(float64)
	CheckPolicy(string) error
}
