package mtx
import (
	"testing"
)
func TestMatrix(t *testing.T) {
	m := IdMatrix(3)
	n := CloneMtx(m)
	t.Logf("%v",m)
	if !Equal(m,Transpose(m)) {
		t.Fail()
	}
	Scale(m,2)
	t.Logf("%v",m)
	v := Inverse(m)
	t.Logf("%v",v)
	Mul(m,v)
	t.Logf("%v",m)
	if !Equal(m, n) {
		t.Fail()
	}
	m = Matrix{{5,6,6,8},{2,2,2,8},{6,6,2,8},{2,3,6,7}}
	mi := Inverse(m)
	t.Logf("%v",mi)
	if !Equal(Mul(m,mi),IdMatrix(4)) {
		t.Fail()
	}
	y := BuildVec(5,func(i int) float64 { return 1.0/float64(i+1)})
	p := PolyFit(RangeVec(1,5),y)
	t.Logf("%v",p)
	pb := PolyFit(RangeVec(1,5),y)
	t.Logf("%v",pb)
	f := MakePoly(p,false)
	fb := MakePoly(pb,false)
	t.Logf("%v",BuildVec(10,func(i int) float64 { return f(float64(i-2)) }))
	t.Logf("%v",BuildVec(10,func(i int) float64 { return fb(float64(i-2)) }))
	t.Logf("%v",BuildVec(10,func(i int) float64 { return 1.0/float64(i-2) }))
	t.Logf("%v",RangeMtx(0,10,10))
	t.Logf("%v",Transpose(RangeMtx(0,10,10)))
}
