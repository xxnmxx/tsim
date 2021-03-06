package tsim

import "fmt"

// Corp struct
type Corp struct {
	Accs map[string]Acc
	CIT  CIT
	VAT  VAT
}

func NewCorp() *Corp {
	c := &Corp{}
	c.Accs = make(map[string]Acc)
	c.CIT.Rate = 0.3
	c.VAT.LumpSum = false
	return c
}

// Implements Node Interface
func (c *Corp) TokenLiteral() string { return string(CORP) }
func (c *Corp) expressionNode()      {}

// Implements Object Interface
func (c *Corp) Type() ObjectType { return CORP_OBJ }
func (c *Corp) Inspect() string {
	s := fmt.Sprintf("OP:%v\tCIT:%v\tVAT:%v\t", c.OperatingProfit(), c.Cit(), c.Vat())
	return s
}

// Corp methods
func (c *Corp) OperatingProfit() float64 {
	rev := 0.0
	exp := 0.0
	for _, acc := range c.Accs {
		if acc.Type == Revenue {
			rev += acc.Value
		} else if acc.Type == Expence {
			exp += acc.Value
		}
	}
	return rev - exp
}

func (c *Corp) Cit() float64 { return (c.OperatingProfit() + c.CIT.Adj) * c.CIT.Rate }

func (c *Corp) Vat() float64 {
	c.accToVat()
	if c.VAT.LumpSum {
		return c.VAT.OutputTax() - c.VAT.InputTaxLump()
	}
	return c.VAT.OutputTax() - c.VAT.InputTaxInd()
}

func (c *Corp) accToVat() {
	c.VAT.base = base{}
	for _, acc := range c.Accs {
		switch acc.VAT {
		// Sales VAT
		case V10:
			c.VAT.base.V10 += acc.Value
		case V8:
			c.VAT.base.V8 += acc.Value
		case V8R:
			c.VAT.base.V8R += acc.Value
		case V5:
			c.VAT.base.V5 += acc.Value
		case V0:
			c.VAT.base.V0 += acc.Value
		case VF:
			c.VAT.base.VF += acc.Value
		case VN:
			c.VAT.base.VN += acc.Value
		// Purchase VAT Individual
		case A10t:
			c.VAT.base.A10t += acc.Value
		case A10c:
			c.VAT.base.A10c += acc.Value
		case A10n:
			c.VAT.base.A10n += acc.Value
		case A8t:
			c.VAT.base.A8t += acc.Value
		case A8c:
			c.VAT.base.A8c += acc.Value
		case A8n:
			c.VAT.base.A8n += acc.Value
		case A8Rt:
			c.VAT.base.A8Rt += acc.Value
		case A8Rc:
			c.VAT.base.A8Rc += acc.Value
		case A8Rn:
			c.VAT.base.A8Rn += acc.Value
		case A5t:
			c.VAT.base.A5t += acc.Value
		case A5c:
			c.VAT.base.A5c += acc.Value
		case A5n:
			c.VAT.base.A5n += acc.Value
		// Purchase VAT LumpSum method.
		case A10:
			c.VAT.base.A10 += acc.Value
		case A8:
			c.VAT.base.A8 += acc.Value
		case A5:
			c.VAT.base.A5 += acc.Value
		case AF:
			c.VAT.base.AF += acc.Value
		case AN:
			c.VAT.base.AN += acc.Value
		}
	}
}

// Acc
type AccType string

const (
	Asset     = "Asset"
	Liability = "Liability"
	Revenue   = "Revenue"
	Expence   = "Expence"
)

type VatType string

const (
	V10  = "output10%"
	V8   = "output8%"
	V8R  = "output8%Reduced"
	V5   = "output5%"
	V0   = "outputExempt0%"
	VF   = "outputTaxFree"
	VN   = "outputNonTaxable"
	A10  = "input10%"
	A10t = "input10%CurrTax"
	A10c = "input10%CurrCmn"
	A10n = "input10%CurrNonTax"
	A8   = "input8%"
	A8t  = "input8%CurrTax"
	A8c  = "input8%CurrCmn"
	A8n  = "input8%CurrNonTax"
	A8R  = "input8%Reduced"
	A8Rt = "input8%ReducedCurrTax"
	A8Rc = "input8%ReducedCurrCmn"
	A8Rn = "input8%ReducedCurrNonTax"
	A5   = "input5%"
	A5t  = "input5%CurrTax"
	A5c  = "input5%CurrCmn"
	A5n  = "input5%CurrNonTax"
	AF   = "inputTaxFree"
	AN   = "inputNonTaxable"
)

type Acc struct {
	Type  AccType // Asset etc
	Value float64
	VAT   VatType
}

func (a *Acc) expressionNode() {}
func (a *Acc) TokenLiteral()   {} // dammy

func (c *Corp) CreateAcc(name string, t AccType, v float64, vat VatType) error {
	_, ok := c.Accs[name]
	if ok {
		return fmt.Errorf("%v already exists", name)
	}
	c.Accs[name] = Acc{
		Type:  t,
		Value: v,
		VAT:   vat,
	}
	return nil
}

func (c *Corp) DeleteAcc(name string) error {
	_, ok := c.Accs[name]
	if !ok {
		return fmt.Errorf("%v does not exist", name)
	}
	delete(c.Accs, name)
	// ToDo deduct value from VAT
	return nil
}

// CIT struct
type CIT struct {
	Rate   float64
	Adj    float64
	Credit float64
}

// VAT struct
type VAT struct {
	base    base
	LumpSum bool
}

type base struct {
	// Sales VAT
	V10 float64
	V8  float64
	V8R float64
	V5  float64
	V0  float64
	VF  float64
	VN  float64

	// Purchase Vat Individual method.
	A10t float64
	A10c float64
	A10n float64
	A8t  float64
	A8c  float64
	A8n  float64
	A8Rt float64
	A8Rc float64
	A8Rn float64
	A5t  float64
	A5c  float64
	A5n  float64

	// Purchase VAT LumpSum method.
	A10 float64
	A8  float64
	A5  float64

	// Purchase other.
	AF float64
	AN float64
}

// VAT methods
// Ratio returns TaxRatio.
func (v *VAT) Ratio() float64 {
	deno := v.base.V10 + v.base.V8 + v.base.V8R + v.base.V5 + v.base.V0
	nume := deno + v.base.VF
	return deno / nume
}

// InputTaxInd returns inputtax amount of individual method.
func (v *VAT) InputTaxInd() float64 {
	return v.inputTenInd()*0.1 + v.inputEightInd()*0.08 + v.inputRedEightInd() + v.inputFiveInd()*0.05
}

func (v *VAT) inputTenInd() float64 {
	return v.base.A10t + v.base.A10c*v.Ratio()
}

func (v *VAT) inputEightInd() float64 {
	return v.base.A8t + v.base.A8c*v.Ratio()
}

func (v *VAT) inputRedEightInd() float64 {
	return v.base.A8Rt + v.base.A8Rc*v.Ratio()
}

func (v *VAT) inputFiveInd() float64 {
	return v.base.A5t + v.base.A5c*v.Ratio()
}

// InputTaxLump returns inputtax amount of lump sum method.
func (v *VAT) InputTaxLump() float64 {
	return v.inputTenLump()*0.1 + v.inputEightLump()*0.08 + v.inputFiveLump()*0.05
}

func (v *VAT) inputTenLump() float64 {
	return v.base.A10
}

func (v *VAT) inputEightLump() float64 {
	return v.base.A8
}

func (v *VAT) inputFiveLump() float64 {
	return v.base.A5
}

// OutputTax returns outputtax amount.
func (v *VAT) OutputTax() float64 {
	return v.outputTen()*0.1 + v.outputEight()*0.08 + v.outputFive()*0.05 + v.outputReduced()*0.0
}

func (v *VAT) outputTen() float64 {
	return v.base.V10
}

func (v *VAT) outputEight() float64 {
	return v.base.V8
}

func (v *VAT) outputReduced() float64 {
	return v.base.V8R
}

func (v *VAT) outputFive() float64 {
	return v.base.V5
}
