package di

import "main/bins"

type Di interface {
	Read(string) (*bins.BinList, error)
	Write(*bins.BinList, string) error
}
