package types

import "strconv"

type PolicyByHost struct {
	Src      string `boil:"src"`
	SrcAddr  string `boil:"src_addr"`
	Dst      string `boil:"dst"`
	DstAddr  string `boil:"dst_addr"`
	Port     int    `boil:"port"`
	Protocol string `boil:"protocol"`
}

type PolicyByGroup struct {
	Src      string `boil:"src"`
	Dst      string `boil:"dst"`
	Port     int    `boil:"port"`
	Protocol string `boil:"protocol"`
}

func (p PolicyByHost) ToStrings() []string {
	return []string{p.Src, p.SrcAddr, p.Dst, p.DstAddr, strconv.Itoa(p.Port), p.Protocol}
}

func (p PolicyByGroup) ToStrings() []string {
	return []string{p.Src, p.Dst, strconv.Itoa(p.Port), p.Protocol}
}
