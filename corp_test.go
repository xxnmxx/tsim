package corp

import (
	"log"
	"testing"
)

func TestOperatingProfit(t *testing.T) {
	c := New()
	tt := []struct {
		rt  AccType
		rv  VatType
		rvl float64
		et  AccType
		ev  VatType
		evl float64
		e   float64
	}{
		{rt: Revenue, rv: V10, rvl: 1000, et: Expence, ev: A10, evl: 500, e: 500},
		{rt: Revenue, rv: V10, rvl: 2000, et: Expence, ev: A10, evl: 2500, e: -500},
	}
	for i, test := range tt {
		if err := c.CreateAcc("rev", test.rt, test.rvl, test.rv); err != nil {
			log.Fatal(err)
		}
		if err := c.CreateAcc("exp", test.et, test.evl, test.ev); err != nil {
			log.Fatal(err)
		}
		if test.e != c.OperatingProfit() {
			t.Errorf("i: %v,e: %v, a: %v", i, test.e, c.OperatingProfit())
		}
		if err := c.DeleteAcc("rev"); err != nil {
			log.Fatal(err)
		}
		if err := c.DeleteAcc("exp"); err != nil {
			log.Fatal(err)
		}
	}
}

func TestVat(t *testing.T) {
	c := New()
	tt := []struct {
		revTax     float64
		revEx      float64
		revFr      float64
		expCurrTax float64
		expCurrCmn float64
		ep          float64
		et          float64
	}{
		{revTax: 2000, revEx: 2000., revFr: 1000., expCurrTax: 100., expCurrCmn: 200., ep:4700,et: 174},
		{revTax: 1000., revEx: 2000., revFr: 0., expCurrTax: 100., expCurrCmn: 200.,ep:2700, et: 70},
		{revTax: 0., revEx: 2000., revFr: 0., expCurrTax: 100., expCurrCmn: 200.,ep:1700, et: -30},
	}
	for i, test := range tt {
		c.CreateAcc("rt",Revenue,test.revTax,V10)
		c.CreateAcc("re",Revenue,test.revEx,V0)
		c.CreateAcc("rf",Revenue,test.revFr,VF)
		c.CreateAcc("et",Expence,test.expCurrTax,A10t)
		c.CreateAcc("ec",Expence,test.expCurrCmn,A10c)
		if c.OperatingProfit() != test.ep || c.Vat() != test.et {
			t.Errorf("\ni: %v\nep: %v\nap: %v\net: %v\nat: %v\nall: %+v\n", i, test.ep,c.OperatingProfit(),test.et, c.Vat(),  *c)
		}
		c.DeleteAcc("rt")
		c.DeleteAcc("re")
		c.DeleteAcc("rf")
		c.DeleteAcc("et")
		c.DeleteAcc("ec")
	}
}
