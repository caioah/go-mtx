package mtx
import("math")
type Matrix = [][]float64
type Vector = []float64
const (
	X = iota
		Y
		Z
)
func NewMatrix(r,c int) Matrix {
	ret := make(Matrix,r)
	for i:=0;i<r;i++ {
		ret[i] = make(Vector,c)
	}
	return ret
}
func IdMatrix(r int) Matrix {
	ret := make(Matrix,r)
	for i:=0;i<r;i++ {
		ret[i] = make(Vector,r)
		ret[i][i] = 1
	}
	return ret
}
func SkewMtx(v Vector) Matrix {
	ret := make(Matrix,len(v))
	for i:=0;i<len(v);i++ {
		ret[i] = make(Vector,len(v))
		ret[i][i] = v[i]
	}
	return ret
}
func MapVec(f func(...float64) float64, v ...Vector) Vector {
	if len(v) == 0 {
		return nil
	}
	ret := make(Vector,len(v))
	for i:=0;i<len(v);i++ {
		ret[i] = f(v[i]...)
	}
	return ret
}
func FoldVec(start float64,f func(float64,...float64) float64,v ...Vector) float64 {
	if len(v) == 0 {
		return start
	}
	sum := start
	for i:=0;i<len(v);i++ {
		sum = f(sum,v[i]...)
	}
	return sum
}
func FilterVec(f func(float64) bool,v Vector) Vector {
	var ret Vector
	for i:=0;i<len(v);i++ {
		if f(v[i]) {
			ret = append(ret,v[i])
		}}
	return ret
}
func ReverseVec(v Vector) Vector {
	for i,j:=0,len(v)-1;i<j;i,j = i+1,j-1 {
		v[i],v[j] = v[j],v[i]
	}
	return v
}
func CopyVec(v Vector,start,end int,u Vector,offs int) Vector {
	for i,j:=start,offs;i<end;i,j = i+1,j+1 {
		u[j] = v[i]
	}
	return u
}
func CloneVec(v Vector) Vector {
	ret := make(Vector,len(v))
	return CopyVec(v,0,len(v),ret,0)
}
func BuildVec(k int,f func(int) float64) Vector {
	ret := make(Vector,k)
	for i:=0;i<k;i++ {
		ret[i]=f(i)
	}
	return ret
}
func BuildMtx(r,c int,f func(int,int) float64) Matrix {
	ret := make(Matrix,r)
	for i:=0;i<r;i++ {
		ret[i] = BuildVec(c,func(j int) float64 {
			return f(i,j)
		})}
	return ret
}
func MakeVec(k int, val float64) Vector {
	ret := make(Vector,k)
	for i:=0;i<k;i++ {
		ret[i] = val
	}
	return ret
}
func MakeMtx(r,c int, val float64) Matrix {
	ret := make(Matrix,r)
	for i:=0;i<r;i++ {
		ret[i] = MakeVec(c,val)
	}
	return ret
}
func RangeVec(start float64,k int) Vector {
	return BuildVec(k,func(i int) float64 {
		return start+float64(i)
	})}
func RangeMtx(start float64,r,c int) Matrix {
	return BuildMtx(r,c,func(x,y int) float64 {
		return (float64(x)*float64(c))+float64(y)
	})}
func ReverseMtx(m Matrix) Matrix {
	for i,j := 0,len(m)-1;i<j;i,j=i+1,j-1 {
		m[i],m[j] = ReverseVec(m[j]),ReverseVec(m[i])
	}
	return m
}
func CopyMtx(m Matrix,r,c,lr,lc int,n Matrix,x,y int) Matrix {
	for i,j:=r,x;i<lr;i,j = i+1,j+1 {
		CopyVec(m[i],c,lc,n[j],y)
	}
	return n
}
func CloneMtx(m Matrix) Matrix {
	ret := NewMatrix(len(m),len(m[0]))
	return CopyMtx(m,0,0,len(m),len(m[0]),ret,0,0)
}
func MapMtx(f func(...float64) float64, m Matrix) Matrix {
	for i:=0;i<len(m);i++ {
		MapVec(f,m[i])
	}
	return m
}
func FoldMtx(start float64,f func(float64,...float64) float64, m Matrix) float64 {
	sum := start
	for i:=0;i<len(m);i++ {
		sum = FoldVec(sum,f,m[i])
	}
	return sum
}
func FilterMtx(f func(float64)bool, m Matrix) Vector {
	var ret Vector
	for i:=0;i<len(m);i++ {
		ret = append(ret,FilterVec(f,m[i])...)
	}
	return ret
}

func EqualVec(v,u Vector) bool {
	if len(v) != len(u) {
		return false
	}
	for i:=0;i<len(v);i++ {
		if v[i] != u[i] {
			return false
		}}
	return true
}
func Equal(m,n Matrix) bool {
	if len(m) != len(n) {
		return false
	}
	for i:=0;i<len(m);i++ {
		if !EqualVec(m[i],n[i]) {
			return false
		}}
	return true
}
func Transpose(m Matrix) Matrix {
	return BuildMtx(len(m[0]),len(m),func(i,j int) float64 {
		return m[j][i]
	})}
func Transpose2(m []Matrix) []Matrix {
	ret := make([]Matrix,len(m[0]))
	for i:=0;i<len(m[0]);i++ {
		ret[i] = make(Matrix,len(m))
		for j:=0;j<len(m);j++ {
			ret[i][j] = CloneVec(m[j][i])
		}} /*
			for k:=0;k<len(m);k++ {
				ret[i][j][k] = m[j][i][j]
			}}}
			//*/
	return ret
}
func Transpose3(m [][]Matrix) [][]Matrix {
	ret := make([][]Matrix,len(m[0]))
	for i:=0;i<len(m[0]);i++ {
		ret[i] = make([]Matrix,len(m))
		for j:=0;j<len(m);j++ {
			ret[i][j] = CloneMtx(m[j][i])
		}}
	return ret
}
func ScaleVec(v Vector,x float64) Vector {
	for i:=0;i<len(v);i++ {
		v[i] *= x
	}
	return v
}
func Scale(m Matrix, x float64) Matrix {
	for i:=0;i<len(m);i++ {
		ScaleVec(m[i],x)
	}
	return m
}
func Dot(v,u Vector) float64 {
	if len(v) != len(u) {
		panic("Dot: expected same length")
	}
	sum := float64(0)
	for i:=0;i<len(v);i++ {
		sum += v[i]*u[i]
	}
	return sum
}
func VMul(m Matrix,v Vector) Vector {
	if len(v) != len(m[0]) {
		panic("VMul: expected same length")
	}
	ret := make([]float64,len(m))
	for i:=0;i<len(m);i++ {
		ret[i] = Dot(m[i],v)
	}
	return ret
}
func VAdd(v,u Vector) Vector {
	if len(v) != len(u) {
		panic("VAdd: expected same length")
	}
	for i:=0;i<len(v);i++ {
		v[i] += u[i]
	}
	return v
}
func Add(m,n Matrix) Matrix {
	if len(m) != len(n) {
		panic("Add: expected same length")
	}
	for i:=0;i<len(m);i++ {
		VAdd(m[i],n[i])
	}
	return m
}
func Mul(m,n Matrix) Matrix {
	if len(m) != len(n[0]) || len(m[0]) != len(n) {
		panic("Mul: expected same length")
	}
	tn := Transpose(n)
	for i:=0;i<len(m);i++ {
		m[i] = VMul(tn,m[i])
	}
	return m
}
func Minor(m Matrix,r,c int) Matrix {
	ret := make(Matrix,len(m)-1)
	var x,y int
	for i:=0;i<len(m)-1;i++ {
		if i >= r {
			x = i+1
		} else {
			x = i
		}
		ret[i] = make(Vector,len(m[0])-1)
		for j:=0;j<len(m[0])-1;j++ {
			if j >= c {
				y = j+1
			} else {
				y = j
			}
			ret[i][j] = m[x][y]
		}}
	return ret
}
func Det(m Matrix) float64 {
	if len(m) != len(m[0]) {
		panic("Det: expected square matrix")
	}
	if sz := len(m); sz < 2 {
		return 0
	} else if sz == 2 {
		return (m[0][0]*m[1][1]) - (m[0][1]*m[1][0])
	} else {
		d := float64(0)
		for i:=0;i<sz;i++ {
			s := float64(1)
			if i%2 != 0 {
				s = -1
			}
			d += s*Det(Minor(m,0,i))*m[0][i]
		}
		return d
	}}
func TCof(m Matrix) Matrix {
	ret := make(Matrix,len(m[0]))
	for i:=0;i<len(m[0]);i++ {
		ret[i] = make(Vector,len(m))
		for j:=0;j<len(m);j++ {
			s := float64(1)
			if (i+j)%2 != 0 {
				s = -1
			}
			ret[i][j] = s*Det(Minor(m,j,i))
		}}
	return ret
}
func Cofactor(m Matrix) Matrix {
	return Transpose(TCof(m))
}
func Inverse(m Matrix) Matrix {
	d := Det(m)
	if d == 0 {
		return nil
	} else {
		return Scale(TCof(m),1.0/d)
	}}
func MakePoly(c Vector, left bool) func(float64) float64 {
	if left {
		return func(x float64) float64 {
			sum := float64(0)
			for i:=0;i<len(c);i++ {
				sum = (sum*x) + c[i]
			}
			return sum
		}
	} else {
		return func(x float64) float64 {
			sum := float64(0)
			for i:=len(c);i>0;i-- {
				sum = (sum*x) + c[i-1]
			}
			return sum
		}}}
func PolyFit(x,y Vector) Vector {
	if len(x) != len(y) {
		panic("PolyFit: expected same length")
	}
	return VMul(Inverse(BuildMtx(len(x),len(y),func(r,c int) float64 {
		return math.Pow(x[r],float64(c))
	})),y)
}
func PolyBest(x,y Vector) Vector {
	if len(x) != len(y) {
		panic("PolyBest: expected same length")
	}
	m := BuildMtx(len(x),len(y),func(r,c int) float64 {
		return math.Pow(x[r],float64(c))
	})
	mt := Transpose(m)
	return VMul(Mul(Inverse(Mul(mt,m)),mt),y)
}
func Convolve(k,m Matrix) float64 {
	if len(k) != len(m) || len(k[0]) != len(m[0]) {
		panic("Convolve: expected same length")
	}
	sum := float64(0)
	for i:=0;i<len(k);i++ {
		for j:=0;j<len(k[0]);j++ {
			sum += k[len(k)-i-1][len(k[0])-j-1]*m[i][j]
		}}
	return sum
}
