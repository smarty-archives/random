package random

import (
	"testing"

	"github.com/smartystreets/assertions/should"
	"github.com/smartystreets/gunit"
)

func TestFixture(t *testing.T) {
	gunit.Run(new(Fixture), t)
}

type Fixture struct {
	*gunit.Fixture

	random *Random // nil (production)
	faked  *Random // not nil (testing)
}

func (this *Fixture) Setup() {
	this.faked = new(Random)
}

func (this *Fixture) TestDiagnosticPrinting() {
	this.Println("These values should be random--different each time you run these tests.")
	this.Println("-----------------------------------------------------------------------")
	this.Println("this.random.Bytes(16)  |", this.random.Bytes(16))
	this.Println("this.random.Base64(16) |", this.random.Base64(16))
	this.Println("this.random.Base62(16) |", this.random.Base62(16))
	this.Println("this.random.Uint32     |", this.random.Uint32(10000, 1))
	this.Println("this.random.Hex(16)    |", this.random.Hex(16))
	this.Println("this.random.GUID       |", this.random.GUID())
	this.Println("this.random.GUIDString |", this.random.GUIDString())
	this.Println("-----------------------------------------------------------------------")
	this.Println()
	this.Println("These values should be consistent--unchanging each time you run these tests.")
	this.Println("----------------------------------------------------------------------------")
	this.Println("this.faked.Bytes(16)  |", this.faked.Bytes(16))
	this.Println("this.faked.Base64(16) |", this.faked.Base64(16))
	this.Println("this.faked.Base62(16) |", this.faked.Base62(16))
	this.Println("this.faked.Uint32     |", this.faked.Uint32(10000, 1))
	this.Println("this.faked.Hex(16)    |", this.faked.Hex(16))
	this.Println("this.faked.GUID       |", this.faked.GUID())
	this.Println("this.faked.GUIDString |", this.faked.GUIDString())
	this.Println("-----------------------------------------------------------------------------")
}

func (this *Fixture) TestProductionGradeRandomizedFunctions() {
	this.So(this.random.Bytes(16), should.NotResemble, this.random.Bytes(16))
	this.So(this.random.Uint32(1, 10000), should.NotEqual, this.random.Uint32(1, 10000))
	this.So(this.random.Hex(16), should.NotResemble, this.random.Hex(16))
	this.So(this.random.Base64(16), should.NotResemble, this.random.Base64(16))
	this.So(this.random.Base62(16), should.NotResemble, this.random.Base62(16))
	this.So(this.random.GUID(), should.NotResemble, this.random.GUID())
	this.So(this.random.GUIDString(), should.NotResemble, this.random.GUIDString())
}

func (this *Fixture) TestNonRandomValuesUsefulForTestEnvironments() {
	this.So(this.faked.Bytes(1), should.Resemble, []byte{1})
	this.So(this.faked.Bytes(2), should.Resemble, []byte{2, 2})
	this.So(this.faked.Bytes(3), should.Resemble, []byte{3, 3, 3}) // you get the idea...

	this.So(this.faked.Uint32(1, 3215), should.Equal, 3214)
	this.So(this.faked.Hex(4), should.Resemble, "04040404")
	this.So(this.faked.Base64(5), should.Resemble, "BQUFBQU=")
	this.So(this.faked.Base62(5), should.Resemble, "BQUFBQU1")
	this.So(this.faked.GUID(), should.Resemble, []byte{16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16})
	this.So(this.faked.GUIDString(), should.Resemble, "10101010-1010-1010-1010-101010101010")
}
