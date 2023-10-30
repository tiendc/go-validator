package validation

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_StrLen(t *testing.T) {
	errs := StrLen(gofn.New("chào"), 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrLen(gofn.New("chào"), 1, 3).Exec()
	assert.Equal(t, "len", errs[0].Type())
}

func Test_StrByteLen(t *testing.T) {
	errs := StrByteLen(gofn.New("ab "), 2, 10).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrByteLen(gofn.New("chào"), 1, 4).Exec()
	assert.Equal(t, "byte_len", errs[0].Type())
}

func Test_StrEQ(t *testing.T) {
	errs := StrEQ(gofn.New("abc"), "abc").Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrEQ(gofn.New("abc"), "aBc").Exec()
	assert.Equal(t, "eq", errs[0].Type())
}

func Test_StrIn(t *testing.T) {
	errs := StrIn(gofn.New("chào"), "", "a", "chào").Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrIn(gofn.New("a"), "", "b").Exec()
	assert.Equal(t, "in", errs[0].Type())
}

func Test_StrNotIn(t *testing.T) {
	errs := StrNotIn(gofn.New("chào"), "", "a", "b").Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrNotIn(gofn.New("a"), "", "a", "b").Exec()
	assert.Equal(t, "not_in", errs[0].Type())
}

func Test_StrMatch(t *testing.T) {
	re := regexp.MustCompile("[0-9]+")
	errs := StrMatch(gofn.New("123"), re).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrMatch(gofn.New("abc"), re).Exec()
	assert.Equal(t, "match", errs[0].Type())
}

func Test_StrByteMatch(t *testing.T) {
	re := regexp.MustCompile("[0-9]+")
	errs := StrByteMatch(gofn.New("123"), re).Exec()
	assert.Equal(t, 0, len(errs))

	errs = StrByteMatch(gofn.New("abc"), re).Exec()
	assert.Equal(t, "byte_match", errs[0].Type())
}

// nolint: lll
func Test_StrIs_Functions(t *testing.T) {
	assert.Equal(t, 0, len(StrIsEmail(gofn.New("test@example.com")).Exec()))
	assert.Equal(t, "is_email", StrIsEmail(gofn.New("Test@eXample")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsExistingEmail(gofn.New("test@example.com")).Exec()))
	assert.Equal(t, "is_existing_email", StrIsExistingEmail(gofn.New("Test@eXample")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsURL(gofn.New("http://host.com")).Exec()))
	assert.Equal(t, "is_url", StrIsURL(gofn.New("host")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRequestURL(gofn.New("http://host.com:3000")).Exec()))
	assert.Equal(t, "is_request_url", StrIsRequestURL(gofn.New("abc3000")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRequestURI(gofn.New("http://host.com:3000")).Exec()))
	assert.Equal(t, "is_request_uri", StrIsRequestURI(gofn.New("abc3000")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsAlpha(gofn.New("abc")).Exec()))
	assert.Equal(t, "is_alpha", StrIsAlpha(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUTFLetter(gofn.New("Tiến")).Exec()))
	assert.Equal(t, "is_utf_letter", StrIsUTFLetter(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsAlphanumeric(gofn.New("abc012")).Exec()))
	assert.Equal(t, "is_alpha_numeric", StrIsAlphanumeric(gofn.New("abc-123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUTFLetterNumeric(gofn.New("abc012")).Exec()))
	assert.Equal(t, "is_utf_letter_numeric", StrIsUTFLetterNumeric(gofn.New("abc+123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsNumeric(gofn.New("012")).Exec()))
	assert.Equal(t, "is_numeric", StrIsNumeric(gofn.New("-123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUTFNumeric(gofn.New("012")).Exec()))
	assert.Equal(t, "is_utf_numeric", StrIsUTFNumeric(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUTFDigit(gofn.New("123456")).Exec()))
	assert.Equal(t, "is_utf_digit", StrIsUTFDigit(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsHexadecimal(gofn.New("abdcef")).Exec()))
	assert.Equal(t, "is_hexadecimal", StrIsHexadecimal(gofn.New("abdcefg")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsHexcolor(gofn.New("#ffffff")).Exec()))
	assert.Equal(t, "is_hexcolor", StrIsHexcolor(gofn.New("#fffffg")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRGBcolor(gofn.New("rgb(1,1,1)")).Exec()))
	assert.Equal(t, "is_rgbcolor", StrIsRGBcolor(gofn.New("rgb(1,1,a)")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsLowerCase(gofn.New("abc123")).Exec()))
	assert.Equal(t, "is_lower_case", StrIsLowerCase(gofn.New("aBc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUpperCase(gofn.New("ABC123")).Exec()))
	assert.Equal(t, "is_upper_case", StrIsUpperCase(gofn.New("aBc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrHasLowerCase(gofn.New("aBC123")).Exec()))
	assert.Equal(t, "has_lower_case", StrHasLowerCase(gofn.New("ABC123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrHasUpperCase(gofn.New("aBc123")).Exec()))
	assert.Equal(t, "has_upper_case", StrHasUpperCase(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsInt(gofn.New("-123")).Exec()))
	assert.Equal(t, "is_int", StrIsInt(gofn.New("123a")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsFloat(gofn.New("-123.123")).Exec()))
	assert.Equal(t, "is_float", StrIsFloat(gofn.New("123.123f")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrHasWhitespaceOnly(gofn.New("   ")).Exec()))
	assert.Equal(t, "has_whitespace_only", StrHasWhitespaceOnly(gofn.New("ab c")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrHasWhitespace(gofn.New(" b  c")).Exec()))
	assert.Equal(t, "has_whitespace", StrHasWhitespace(gofn.New("abc")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv3(gofn.New("2037da5d-a759-3ba8-bfeb-84519bb669c6")).Exec()))
	assert.Equal(t, "is_uuid_v3", StrIsUUIDv3(gofn.New("2037da5d-a759-3ba8-bfeb-84519bb669g6")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv4(gofn.New("fb3e2e7c-e478-4d76-aa84-9880d6eb67f4")).Exec()))
	assert.Equal(t, "is_uuid_v4", StrIsUUIDv4(gofn.New("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv5(gofn.New("aa43954a-ced3-5b84-9931-d3516b2e1867")).Exec()))
	assert.Equal(t, "is_uuid_v5", StrIsUUIDv5(gofn.New("aa43954a-ced3-5b84-9931-d3516b2e186g")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUUID(gofn.New("2037da5d-a759-3ba8-bfeb-84519bb669c6")).Exec()))
	assert.Equal(t, "is_uuid", StrIsUUID(gofn.New("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsULID(gofn.New("01G65Z755AFWAKHE12NY0CQ9FH")).Exec()))
	assert.Equal(t, "is_ulid", StrIsULID(gofn.New("01G65Z755AFWAKHE12NY0CQ9FHx")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsCreditCard(gofn.New("4242-8888-8888-8888")).Exec()))
	assert.Equal(t, "is_credit_card", StrIsCreditCard(gofn.New("5242-8888-8888-8888")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISBN13(gofn.New("978-3-16-148410-0")).Exec()))
	assert.Equal(t, "is_isbn13", StrIsISBN13(gofn.New("0-306-40615-2")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISBN10(gofn.New("0-306-40615-2")).Exec()))
	assert.Equal(t, "is_isbn10", StrIsISBN10(gofn.New("978-3-16-148410-0")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISBN(gofn.New("978-3-16-148410-0")).Exec()))

	assert.Equal(t, 0, len(StrIsJSON(gofn.New("123")).Exec()))
	assert.Equal(t, "is_json", StrIsJSON(gofn.New(`{"k":v}`)).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMultibyte(gofn.New("Chào buổi sáng")).Exec()))
	assert.Equal(t, "is_multibyte", StrIsMultibyte(gofn.New("abc 123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsASCII(gofn.New("hello")).Exec()))
	assert.Equal(t, "is_ascii", StrIsASCII(gofn.New("hello Tiến")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsPrintableASCII(gofn.New("hello")).Exec()))
	assert.Equal(t, "is_printable_ascii", StrIsPrintableASCII(gofn.New("hello\t")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsFullWidth(gofn.New("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")).Exec()))
	assert.Equal(t, "is_full_width", StrIsFullWidth(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsHalfWidth(gofn.New("abc123")).Exec()))
	assert.Equal(t, "is_half_width", StrIsHalfWidth(gofn.New("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsVariableWidth(gofn.New("Ｔｙｐｅ　ｏｒ　Paste")).Exec()))
	assert.Equal(t, "is_variable_width", StrIsVariableWidth(gofn.New("Ｔｙｐｅ")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsBase64(gofn.New("aGVsbG8gd29ybGQ=")).Exec()))
	assert.Equal(t, "is_base64", StrIsBase64(gofn.New("hello")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsFilePath(gofn.New(`C:\\abc`)).Exec()))

	assert.Equal(t, 0, len(StrIsWinFilePath(gofn.New(`C:\abc`)).Exec()))
	assert.Equal(t, "is_win_file_path", StrIsWinFilePath(gofn.New("/root")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUnixFilePath(gofn.New(`/root/path`)).Exec()))

	assert.Equal(t, 0, len(StrIsDataURI(gofn.New("data:image/png;base64,iVBORw0OHwAAAABJRU5ErkJggg==")).Exec()))
	assert.Equal(t, "is_data_uri", StrIsDataURI(gofn.New("example.com")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMagnetURI(gofn.New("magnet:?xt=urn:xxxxxx:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&dn=bbb&tr=ccc")).Exec()))
	assert.Equal(t, "is_magnet_uri", StrIsMagnetURI(gofn.New("example.com")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISO3166Alpha2(gofn.New("VN")).Exec()))
	assert.Equal(t, "is_iso3166_alpha2", StrIsISO3166Alpha2(gofn.New("VND")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISO3166Alpha3(gofn.New("USA")).Exec()))
	assert.Equal(t, "is_iso3166_alpha3", StrIsISO3166Alpha3(gofn.New("USD")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISO639Alpha2(gofn.New("vi")).Exec()))
	assert.Equal(t, "is_iso639_alpha2", StrIsISO639Alpha2(gofn.New("vie")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISO639Alpha3b(gofn.New("vie")).Exec()))
	assert.Equal(t, "is_iso639_alpha3b", StrIsISO639Alpha3b(gofn.New("vi")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsDNSName(gofn.New("MX")).Exec()))
	assert.Equal(t, "is_dns_name", StrIsDNSName(gofn.New("")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3224(gofn.New("d9cf77bff9d00e47ad2d4841539bf72b0cfeff5e106819625e4f99f4")).Exec()))
	assert.Equal(t, "is_sha3_224", StrIsSHA3224(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3256(gofn.New("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Exec()))
	assert.Equal(t, "is_sha3_256", StrIsSHA3256(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3384(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()))
	assert.Equal(t, "is_sha3_384", StrIsSHA3384(gofn.New("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3512(gofn.New("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b37ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")).Exec()))
	assert.Equal(t, "is_sha3_512", StrIsSHA3512(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA512(gofn.New("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b37ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")).Exec()))
	assert.Equal(t, "is_sha512", StrIsSHA512(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA384(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()))
	assert.Equal(t, "is_sha384", StrIsSHA384(gofn.New("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA256(gofn.New("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Exec()))
	assert.Equal(t, "is_sha256", StrIsSHA256(gofn.New("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSHA1(gofn.New("4b2b79b6f371ca18f1216461cffeaddf6848a50e")).Exec()))
	assert.Equal(t, "is_sha1", StrIsSHA1(gofn.New("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsTiger192(gofn.New("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")).Exec()))
	assert.Equal(t, "is_tiger192", StrIsTiger192(gofn.New("8d0484881b88804442e390b0784003a1981db0b3")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsTiger160(gofn.New("8d0484881b88804442e390b0784003a1981db0b3")).Exec()))
	assert.Equal(t, "is_tiger160", StrIsTiger160(gofn.New("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsTiger128(gofn.New("b8d02ffdd4d054b5327eed23b20e8716")).Exec()))
	assert.Equal(t, "is_tiger128", StrIsTiger128(gofn.New("8d0484881b88804442e390b0784003a1981db0b3")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRipeMD160(gofn.New("2330032576b1cc1df37b92abec140368c7396745")).Exec()))
	assert.Equal(t, "is_ripemd160", StrIsRipeMD160(gofn.New("afbed790a7824ade4a3ae531285f8fe8")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRipeMD128(gofn.New("afbed790a7824ade4a3ae531285f8fe8")).Exec()))
	assert.Equal(t, "is_ripemd128", StrIsRipeMD128(gofn.New("2330032576b1cc1df37b92abec140368c7396745")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsCRC32(gofn.New("cd71f992")).Exec()))
	assert.Equal(t, "is_crc32", StrIsCRC32(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsCRC32b(gofn.New("94d6f306")).Exec()))
	assert.Equal(t, "is_crc32b", StrIsCRC32b(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMD5(gofn.New("ec02c59dee6faaca3189bace969c22d3")).Exec()))
	assert.Equal(t, "is_md5", StrIsMD5(gofn.New("abc")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMD4(gofn.New("538a2a786ca1bc47694c0bc3f3ac3228")).Exec()))
	assert.Equal(t, "is_md4", StrIsMD4(gofn.New("abc")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsDialString(gofn.New("golang.org:3000")).Exec()))
	assert.Equal(t, "is_dial_string", StrIsDialString(gofn.New("abc")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsIP(gofn.New("1.1.1.1")).Exec()))
	assert.Equal(t, "is_ip", StrIsIP(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsPort(gofn.New("1234")).Exec()))
	assert.Equal(t, "is_port", StrIsPort(gofn.New("70000")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsIPv4(gofn.New("1.1.1.1")).Exec()))
	assert.Equal(t, "is_ipv4", StrIsIPv4(gofn.New("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsIPv6(gofn.New("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")).Exec()))
	assert.Equal(t, "is_ipv6", StrIsIPv6(gofn.New("1.1.1.1")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsCIDR(gofn.New("255.255.255.0/32")).Exec()))
	assert.Equal(t, "is_cidr", StrIsCIDR(gofn.New("255.255.255.255")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMAC(gofn.New("01-80-C2-FF-FF-FF")).Exec()))
	assert.Equal(t, "is_mac", StrIsMAC(gofn.New("1.1.1.1")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsHost(gofn.New("example.com")).Exec()))
	assert.Equal(t, "is_host", StrIsHost(gofn.New("abc/123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsMongoID(gofn.New("507f1f77bcf86cd799439011")).Exec()))
	assert.Equal(t, "is_mongo_id", StrIsMongoID(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsLatitude(gofn.New("38.8951")).Exec()))
	assert.Equal(t, "is_latitude", StrIsLatitude(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsLongitude(gofn.New("-77.0364")).Exec()))
	assert.Equal(t, "is_longitude", StrIsLongitude(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsIMEI(gofn.New("11222222333333")).Exec()))
	assert.Equal(t, "is_imei", StrIsIMEI(gofn.New("abc123")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsIMSI(gofn.New("313460000000001")).Exec()))
	assert.Equal(t, "is_imsi", StrIsIMSI(gofn.New("11222222333333")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRsaPublicKey(gofn.New("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsFCmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ=="), 512).Exec()))
	assert.Equal(t, "is_rsa_public_key", StrIsRsaPublicKey(gofn.New("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsFCmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ=="), 1024).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRegex(gofn.New("[0-9a-z]+")).Exec()))

	assert.Equal(t, 0, len(StrIsSSN(gofn.New("111-22-3333")).Exec()))
	assert.Equal(t, "is_ssn", StrIsSSN(gofn.New("111.22.3333")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsSemver(gofn.New("1.2.3")).Exec()))
	assert.Equal(t, "is_semver", StrIsSemver(gofn.New("1.2")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsTime(gofn.New("2023-10-01"), "2006-02-01").Exec()))
	assert.Equal(t, "is_time", StrIsTime(gofn.New("2023/10/01"), "2006-02-01").Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsUnixTime(gofn.New("123456789")).Exec()))
	assert.Equal(t, "is_unix_time", StrIsUnixTime(gofn.New("12345678912345678912345")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRFC3339(gofn.New("2020-12-09T16:09:53-00:00")).Exec()))
	assert.Equal(t, "is_rfc3339", StrIsRFC3339(gofn.New("2020-12-09 16:09:53+00:00")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsRFC3339WithoutZone(gofn.New("2020-12-09T16:09:53")).Exec()))
	assert.Equal(t, "is_rfc3339_without_zone", StrIsRFC3339WithoutZone(gofn.New("2020-12-09T16:09:53-00:00")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsISO4217(gofn.New("USD")).Exec()))
	assert.Equal(t, "is_iso4217", StrIsISO4217(gofn.New("usd")).Exec()[0].Type())

	assert.Equal(t, 0, len(StrIsE164(gofn.New("+442071838750")).Exec()))
	assert.Equal(t, "is_e164", StrIsE164(gofn.New("abc123")).Exec()[0].Type())
}
