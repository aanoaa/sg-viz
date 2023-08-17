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
	SrcZone  string `boil:"src_zone"`
	Dst      string `boil:"dst"`
	DstZone  string `boil:"dst_zone"`
	Port     int    `boil:"port"`
	Protocol string `boil:"protocol"`
}

func (p PolicyByHost) ToStrings() []string {
	return []string{p.Src, p.SrcAddr, p.Dst, p.DstAddr, strconv.Itoa(p.Port), p.Protocol}
}

func (p PolicyByGroup) ToStrings() []string {
	return []string{p.Src, p.SrcZone, p.Dst, p.DstZone, strconv.Itoa(p.Port), p.Protocol}
}
