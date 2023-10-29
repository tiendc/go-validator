package stringvalidation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_IsEmail(t *testing.T) {
	assert.True(t, gofn.Head(IsEmail("test@example.com")))
	assert.True(t, gofn.Head(IsEmail("Test@eXample.com")))
	assert.False(t, gofn.Head(IsEmail("Test@eXample")))
}

func Test_IsExistingEmail(t *testing.T) {
	assert.True(t, gofn.Head(IsExistingEmail("Test@eXample.com")))
	assert.False(t, gofn.Head(IsExistingEmail("Test@eXample")))
}

func Test_IsURL(t *testing.T) {
	assert.True(t, gofn.Head(IsURL("http://host.com")))
	assert.True(t, gofn.Head(IsURL("localhost:10000")))
	assert.False(t, gofn.Head(IsURL("host")))
}

func Test_IsRequestURL(t *testing.T) {
	assert.True(t, gofn.Head(IsRequestURL("http://host.com:3000")))
	assert.False(t, gofn.Head(IsRequestURL("abc3000")))
}

func Test_IsRequestURI(t *testing.T) {
	assert.True(t, gofn.Head(IsRequestURI("http://host.com:3000")))
	assert.False(t, gofn.Head(IsRequestURI("abc3000")))
}

func Test_IsAlpha(t *testing.T) {
	assert.True(t, gofn.Head(IsAlpha("abc")))
	assert.False(t, gofn.Head(IsAlpha("abc123")))
}

func Test_IsUTFLetter(t *testing.T) {
	assert.True(t, gofn.Head(IsUTFLetter("Tiến")))
	assert.False(t, gofn.Head(IsUTFLetter("abc123")))
}

func Test_IsAlphanumeric(t *testing.T) {
	assert.True(t, gofn.Head(IsAlphanumeric("abc012")))
	assert.False(t, gofn.Head(IsAlphanumeric("abc-123")))
}

func Test_IsUTFLetterNumeric(t *testing.T) {
	assert.True(t, gofn.Head(IsUTFLetterNumeric("abc012")))
	assert.False(t, gofn.Head(IsUTFLetterNumeric("abc+123")))
}

func Test_IsNumeric(t *testing.T) {
	assert.True(t, gofn.Head(IsNumeric("012")))
	assert.False(t, gofn.Head(IsNumeric("-123")))
}

func Test_IsUTFNumeric(t *testing.T) {
	assert.True(t, gofn.Head(IsUTFNumeric("012")))
	assert.False(t, gofn.Head(IsUTFNumeric("abc123")))
}

func Test_IsUTFDigit(t *testing.T) {
	assert.True(t, gofn.Head(IsUTFDigit("123456")))
	assert.False(t, gofn.Head(IsUTFDigit("abc123")))
}

func Test_IsHexadecimal(t *testing.T) {
	assert.True(t, gofn.Head(IsHexadecimal("abdcef")))
	assert.False(t, gofn.Head(IsHexadecimal("abdcefg")))
}

func Test_IsHexcolor(t *testing.T) {
	assert.True(t, gofn.Head(IsHexcolor("#ffffff")))
	assert.False(t, gofn.Head(IsHexcolor("#fffffg")))
}

func Test_IsRGBcolor(t *testing.T) {
	assert.True(t, gofn.Head(IsRGBcolor("rgb(1,1,1)")))
	assert.False(t, gofn.Head(IsRGBcolor("rgb(1,1,a)")))
}

func Test_IsLowerCase(t *testing.T) {
	assert.True(t, gofn.Head(IsLowerCase("abc123")))
	assert.False(t, gofn.Head(IsLowerCase("aBc123")))
}

func Test_IsUpperCase(t *testing.T) {
	assert.True(t, gofn.Head(IsUpperCase("ABC123")))
	assert.False(t, gofn.Head(IsUpperCase("aBc123")))
}

func Test_HasLowerCase(t *testing.T) {
	assert.True(t, gofn.Head(HasLowerCase("aBC123")))
	assert.False(t, gofn.Head(HasLowerCase("ABC123")))
}

func Test_HasUpperCase(t *testing.T) {
	assert.True(t, gofn.Head(HasUpperCase("aBc123")))
	assert.False(t, gofn.Head(HasUpperCase("abc123")))
}

func Test_IsInt(t *testing.T) {
	assert.True(t, gofn.Head(IsInt("-123")))
	assert.False(t, gofn.Head(IsInt("123a")))
}

func Test_IsFloat(t *testing.T) {
	assert.True(t, gofn.Head(IsFloat("-123.123")))
	assert.False(t, gofn.Head(IsFloat("123.123f")))
}

func Test_HasWhitespaceOnly(t *testing.T) {
	assert.True(t, gofn.Head(HasWhitespaceOnly("   ")))
	assert.False(t, gofn.Head(HasWhitespaceOnly("ab c")))
}

func Test_HasWhitespace(t *testing.T) {
	assert.True(t, gofn.Head(HasWhitespace(" b  c")))
	assert.False(t, gofn.Head(HasWhitespace("abc")))
}

func Test_IsUUIDv3(t *testing.T) {
	assert.True(t, gofn.Head(IsUUIDv3("2037da5d-a759-3ba8-bfeb-84519bb669c6")))
	assert.False(t, gofn.Head(IsUUIDv3("2037da5d-a759-3ba8-bfeb-84519bb669g6")))
}

func Test_IsUUIDv4(t *testing.T) {
	assert.True(t, gofn.Head(IsUUIDv4("fb3e2e7c-e478-4d76-aa84-9880d6eb67f4")))
	assert.False(t, gofn.Head(IsUUIDv4("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")))
}

func Test_IsUUIDv5(t *testing.T) {
	assert.True(t, gofn.Head(IsUUIDv5("aa43954a-ced3-5b84-9931-d3516b2e1867")))
	assert.False(t, gofn.Head(IsUUIDv5("aa43954a-ced3-5b84-9931-d3516b2e186g")))
}

func Test_IsUUID(t *testing.T) {
	assert.True(t, gofn.Head(IsUUID("2037da5d-a759-3ba8-bfeb-84519bb669c6")))
	assert.True(t, gofn.Head(IsUUID("fb3e2e7c-e478-4d76-aa84-9880d6eb67f4")))
	assert.True(t, gofn.Head(IsUUID("aa43954a-ced3-5b84-9931-d3516b2e1867")))
	assert.False(t, gofn.Head(IsUUID("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")))
}

func Test_IsULID(t *testing.T) {
	assert.True(t, gofn.Head(IsULID("01G65Z755AFWAKHE12NY0CQ9FH")))
	assert.False(t, gofn.Head(IsULID("01G65Z755AFWAKHE12NY0CQ9FHx")))
}

func Test_IsCreditCard(t *testing.T) {
	assert.True(t, gofn.Head(IsCreditCard("4242-8888-8888-8888")))
	assert.False(t, gofn.Head(IsCreditCard("5242-8888-8888-8888")))
}

func Test_IsISBN13(t *testing.T) {
	assert.True(t, gofn.Head(IsISBN13("978-3-16-148410-0")))
	assert.False(t, gofn.Head(IsISBN13("0-306-40615-2")))
}

func Test_IsISBN10(t *testing.T) {
	assert.True(t, gofn.Head(IsISBN10("0-306-40615-2")))
	assert.False(t, gofn.Head(IsISBN10("978-3-16-148410-0")))
}

func Test_IsISBN(t *testing.T) {
	assert.True(t, gofn.Head(IsISBN("978-3-16-148410-0")))
	assert.True(t, gofn.Head(IsISBN("0-306-40615-2")))
}

func Test_IsJSON(t *testing.T) {
	assert.True(t, gofn.Head(IsJSON("123")))
	assert.True(t, gofn.Head(IsJSON(`{"k":123}`)))
	assert.False(t, gofn.Head(IsJSON(`{"k":v}`)))
}

func Test_IsMultibyte(t *testing.T) {

}

func Test_IsASCII(t *testing.T) {
	assert.True(t, gofn.Head(IsASCII("hello")))
	assert.False(t, gofn.Head(IsASCII("hello Tiến")))
}

func Test_IsPrintableASCII(t *testing.T) {
	assert.True(t, gofn.Head(IsPrintableASCII("hello")))
	assert.False(t, gofn.Head(IsPrintableASCII("hello\t")))
}

func Test_IsFullWidth(t *testing.T) {
	assert.True(t, gofn.Head(IsFullWidth("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")))
	assert.False(t, gofn.Head(IsFullWidth("abc123")))
}

func Test_IsHalfWidth(t *testing.T) {
	assert.True(t, gofn.Head(IsHalfWidth("abc123")))
	assert.False(t, gofn.Head(IsHalfWidth("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")))
}

func Test_IsVariableWidth(t *testing.T) {
	assert.True(t, gofn.Head(IsVariableWidth("Ｔｙｐｅ　ｏｒ　Paste")))
	assert.False(t, gofn.Head(IsVariableWidth("Ｔｙｐｅ")))
	assert.False(t, gofn.Head(IsVariableWidth("Type")))
}

func Test_IsBase64(t *testing.T) {
	assert.True(t, gofn.Head(IsBase64("aGVsbG8gd29ybGQ=")))
	assert.False(t, gofn.Head(IsBase64("hello")))
}

func Test_IsFilePath(t *testing.T) {
	assert.True(t, gofn.Head(IsFilePath("C:\\abc")))
	assert.True(t, gofn.Head(IsFilePath("/root")))
}

func Test_IsWinFilePath(t *testing.T) {
	assert.True(t, gofn.Head(IsWinFilePath("C:\\abc")))
	assert.False(t, gofn.Head(IsWinFilePath("/root")))
}

func Test_IsUnixFilePath(t *testing.T) {
	assert.True(t, gofn.Head(IsUnixFilePath("/root/path")))
}

func Test_IsDataURI(t *testing.T) {
	assert.True(t, gofn.Head(IsDataURI("data:image/png;base64,iVBORw0OHwAAAABJRU5ErkJggg==")))
	assert.False(t, gofn.Head(IsDataURI("example.com")))
}

func Test_IsMagnetURI(t *testing.T) {
	assert.True(t, gofn.Head(IsMagnetURI("magnet:?xt=urn:xxxxxx:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&dn=bbb&tr=ccc")))
	assert.False(t, gofn.Head(IsMagnetURI("example.com")))
}

func Test_IsISO3166Alpha2(t *testing.T) {
	assert.True(t, gofn.Head(IsISO3166Alpha2("VN")))
	assert.False(t, gofn.Head(IsISO3166Alpha2("VND")))
}

func Test_IsISO3166Alpha3(t *testing.T) {
	assert.True(t, gofn.Head(IsISO3166Alpha3("USA")))
	assert.False(t, gofn.Head(IsISO3166Alpha3("USD")))
}

func Test_IsISO639Alpha2(t *testing.T) {
	assert.True(t, gofn.Head(IsISO639Alpha2("vi")))
	assert.False(t, gofn.Head(IsISO639Alpha2("vie")))
}

func Test_IsISO639Alpha3b(t *testing.T) {
	assert.True(t, gofn.Head(IsISO639Alpha3b("vie")))
	assert.False(t, gofn.Head(IsISO639Alpha3b("vi")))
}

func Test_IsDNSName(t *testing.T) {
	assert.True(t, gofn.Head(IsDNSName("MX")))
	assert.False(t, gofn.Head(IsDNSName("")))
}

func Test_IsSHA3224(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA3224("d9cf77bff9d00e47ad2d4841539bf72b0cfeff5e106819625e4f99f4")))
	assert.False(t, gofn.Head(IsSHA3224("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606e"+
		"fb61c8cb4f173e3a526761dd")))
}

func Test_IsSHA3256(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA3256("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")))
	assert.False(t, gofn.Head(IsSHA3256("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606e"+
		"fb61c8cb4f173e3a526761dd")))
}

func Test_IsSHA3384(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA3384("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606ef"+
		"b61c8cb4f173e3a526761dd")))
	assert.False(t, gofn.Head(IsSHA3384("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")))
}

func Test_IsSHA3512(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA3512("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b"+
		"37ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")))
	assert.False(t, gofn.Head(IsSHA3512("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606ef"+
		"b61c8cb4f173e3a526761dd")))
}

func Test_IsSHA512(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA512("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b3"+
		"7ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")))
	assert.False(t, gofn.Head(IsSHA512("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb"+
		"61c8cb4f173e3a526761dd")))
}

func Test_IsSHA384(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA384("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb"+
		"61c8cb4f173e3a526761dd")))
	assert.False(t, gofn.Head(IsSHA384("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")))
}

func Test_IsSHA256(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA256("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")))
	assert.False(t, gofn.Head(IsSHA256("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606ef"+
		"b61c8cb4f173e3a526761dd")))
}

func Test_IsSHA1(t *testing.T) {
	assert.True(t, gofn.Head(IsSHA1("4b2b79b6f371ca18f1216461cffeaddf6848a50e")))
	assert.False(t, gofn.Head(IsSHA1("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")))
}

func Test_IsTiger192(t *testing.T) {
	assert.True(t, gofn.Head(IsTiger192("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")))
	assert.False(t, gofn.Head(IsTiger192("8d0484881b88804442e390b0784003a1981db0b3")))
}

func Test_IsTiger160(t *testing.T) {
	assert.True(t, gofn.Head(IsTiger160("8d0484881b88804442e390b0784003a1981db0b3")))
	assert.False(t, gofn.Head(IsTiger160("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")))
}

func Test_IsTiger128(t *testing.T) {
	assert.True(t, gofn.Head(IsTiger128("b8d02ffdd4d054b5327eed23b20e8716")))
	assert.False(t, gofn.Head(IsTiger128("8d0484881b88804442e390b0784003a1981db0b3")))
}

func Test_IsRipeMD160(t *testing.T) {
	assert.True(t, gofn.Head(IsRipeMD160("2330032576b1cc1df37b92abec140368c7396745")))
	assert.False(t, gofn.Head(IsRipeMD160("afbed790a7824ade4a3ae531285f8fe8")))
}

func Test_IsRipeMD128(t *testing.T) {
	assert.True(t, gofn.Head(IsRipeMD128("afbed790a7824ade4a3ae531285f8fe8")))
	assert.False(t, gofn.Head(IsRipeMD128("2330032576b1cc1df37b92abec140368c7396745")))
}

func Test_IsCRC32(t *testing.T) {
	assert.True(t, gofn.Head(IsCRC32("cd71f992")))
	assert.False(t, gofn.Head(IsCRC32("abc123")))
}

func Test_IsCRC32b(t *testing.T) {
	assert.True(t, gofn.Head(IsCRC32b("94d6f306")))
	assert.False(t, gofn.Head(IsCRC32b("abc123")))
}

func Test_IsMD5(t *testing.T) {
	assert.True(t, gofn.Head(IsMD5("ec02c59dee6faaca3189bace969c22d3")))
	assert.False(t, gofn.Head(IsMD5("abc")))
}

func Test_IsMD4(t *testing.T) {
	assert.True(t, gofn.Head(IsMD4("538a2a786ca1bc47694c0bc3f3ac3228")))
	assert.False(t, gofn.Head(IsMD4("abc")))
}

func Test_IsDialString(t *testing.T) {
	assert.True(t, gofn.Head(IsDialString("golang.org:3000")))
	assert.False(t, gofn.Head(IsDialString("abc")))
}

func Test_IsIP(t *testing.T) {
	assert.True(t, gofn.Head(IsIP("1.1.1.1")))
	assert.True(t, gofn.Head(IsIP("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")))
}

func Test_IsPort(t *testing.T) {
	assert.True(t, gofn.Head(IsPort("1234")))
	assert.False(t, gofn.Head(IsPort("70000")))
}

func Test_IsIPv4(t *testing.T) {
	assert.True(t, gofn.Head(IsIPv4("1.1.1.1")))
	assert.False(t, gofn.Head(IsIPv4("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")))
}

func Test_IsIPv6(t *testing.T) {
	assert.True(t, gofn.Head(IsIPv6("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")))
	assert.False(t, gofn.Head(IsIPv6("1.1.1.1")))
}

func Test_IsCIDR(t *testing.T) {
	assert.True(t, gofn.Head(IsCIDR("255.255.255.0/32")))
	assert.False(t, gofn.Head(IsCIDR("255.255.255.255")))
}

func Test_IsMAC(t *testing.T) {
	assert.True(t, gofn.Head(IsMAC("01-80-C2-FF-FF-FF")))
	assert.False(t, gofn.Head(IsMAC("1.1.1.1")))
}

func Test_IsHost(t *testing.T) {
	assert.True(t, gofn.Head(IsHost("example.com")))
	assert.True(t, gofn.Head(IsHost("11.22.33.33")))
	assert.False(t, gofn.Head(IsHost("abc/123")))
}

func Test_IsMongoID(t *testing.T) {
	assert.True(t, gofn.Head(IsMongoID("507f1f77bcf86cd799439011")))
	assert.False(t, gofn.Head(IsMongoID("abc123")))
}

func Test_IsLatitude(t *testing.T) {
	assert.True(t, gofn.Head(IsLatitude("38.8951")))
	assert.False(t, gofn.Head(IsLatitude("abc123")))
}

func Test_IsLongitude(t *testing.T) {
	assert.True(t, gofn.Head(IsLongitude("-77.0364")))
	assert.False(t, gofn.Head(IsLongitude("abc123")))
}

func Test_IsIMEI(t *testing.T) {
	assert.True(t, gofn.Head(IsIMEI("11222222333333")))
	assert.False(t, gofn.Head(IsIMEI("abc123")))
}

func Test_IsIMSI(t *testing.T) {
	assert.True(t, gofn.Head(IsIMSI("313460000000001")))
	assert.False(t, gofn.Head(IsIMSI("11222222333333")))
}

func Test_IsRsaPublicKey(t *testing.T) {
	assert.True(t, gofn.Head(IsRsaPublicKey("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsF"+
		"Cmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ==", 512)))
	assert.False(t, gofn.Head(IsRsaPublicKey("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsF"+
		"Cmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ==", 1024)))
}

func Test_IsRegex(t *testing.T) {
	assert.True(t, gofn.Head(IsRegex("[0-9a-z]+")))
}

func Test_IsSSN(t *testing.T) {
	assert.True(t, gofn.Head(IsSSN("111-22-3333")))
	assert.False(t, gofn.Head(IsSSN("111.22.3333")))
}

func Test_IsSemver(t *testing.T) {
	assert.True(t, gofn.Head(IsSemver("1.2.3")))
	assert.False(t, gofn.Head(IsSemver("1.2")))
}

func Test_IsTime(t *testing.T) {
	assert.True(t, gofn.Head(IsTime("2023-10-01", "2006-02-01")))
	assert.False(t, gofn.Head(IsTime("2023/10/01", "2006-02-01")))
}

func Test_IsUnixTime(t *testing.T) {
	assert.True(t, gofn.Head(IsUnixTime("123456789")))
	assert.False(t, gofn.Head(IsUnixTime("12345678912345678912345")))
}

func Test_IsRFC3339(t *testing.T) {
	assert.True(t, gofn.Head(IsRFC3339("2020-12-09T16:09:53-00:00")))
	assert.False(t, gofn.Head(IsRFC3339("2020-12-09 16:09:53+00:00")))
}

func Test_IsRFC3339WithoutZone(t *testing.T) {
	assert.True(t, gofn.Head(IsRFC3339WithoutZone("2020-12-09T16:09:53")))
	assert.False(t, gofn.Head(IsRFC3339WithoutZone("2020-12-09T16:09:53-00:00")))
}

func Test_IsISO4217(t *testing.T) {
	assert.True(t, gofn.Head(IsISO4217("USD")))
	assert.False(t, gofn.Head(IsISO4217("usd")))
}

func Test_IsE164(t *testing.T) {
	assert.True(t, gofn.Head(IsE164("+442071838750")))
	assert.False(t, gofn.Head(IsE164("abc123")))
}
