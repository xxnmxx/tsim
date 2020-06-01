package corp

import "testing"

func TestOperatingProfit(t *testing.T) {
	rev := Acc{
		Type: Revenue,
		VAT:  V10,
	}
	exp := Acc{
		Type: Expence,
		VAT:  A10,
	}
	c := New()
	c.Accs = append(c.Accs, &rev, &exp)
	tt := []struct {
		rev float64
		exp float64
		e   float64
	}{
		{rev: 1000, exp: 500, e: 500},
		{rev: 200, exp: 400, e: -200},
	}
	for i, test := range tt {
		rev.Value = test.rev
		exp.Value = test.exp
		if test.e != c.OperatingProfit() {
			t.Errorf("i: %v,e: %v, a: %v", i, test.e, c.OperatingProfit())
		}
	}
}

func TestVat(t *testing.T) {
	revTax := Acc{Type: Revenue, VAT: V10}
	revEx := Acc{Type: Revenue, VAT: V0}
	revFr := Acc{Type: Revenue, VAT: VF}
	expCurrTax := Acc{Type: Expence, VAT: A10t}
	expCurrCmn := Acc{Type: Expence, VAT: A10c}
	tt := []struct {
		revTax     float64
		revEx      float64
		revFr      float64
		expCurrTax float64
		expCurrCmn float64
		e          float64
	}{
		{revTax: 2000, revEx: 2000., revFr: 1000., expCurrTax: 100., expCurrCmn: 200., e: 174},
		{revTax: 1000., revEx: 2000., revFr: 0., expCurrTax: 100., expCurrCmn: 200., e: 70},
		{revTax: 0., revEx: 2000., revFr: 0., expCurrTax: 100., expCurrCmn: 200., e: -30},
	}
	for i, test := range tt {
		c := New()
		c.Accs = []*Acc{&revTax, &revEx, &revFr, &expCurrTax, &expCurrCmn}
		revTax.Value = test.revTax
		revEx.Value = test.revEx
		revFr.Value = test.revFr
		expCurrTax.Value = test.expCurrTax
		expCurrCmn.Value = test.expCurrCmn
		vat := c.Vat()
		if vat != test.e {
			t.Errorf("\ni: %v\ne: %v\na: %v\nop: %v\nall: %+v\n", i, test.e, vat, c.OperatingProfit(), *c)
		}
	}
}
