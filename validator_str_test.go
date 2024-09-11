package validation

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiendc/gofn"
)

func Test_StrLen(t *testing.T) {
	errs := StrLen(gofn.ToPtr("chào"), 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrLen(gofn.ToPtr("chào"), 1, 3).Validate(ctxBg)
	assert.Equal(t, "len", errs[0].Type())
}

func Test_StrByteLen(t *testing.T) {
	errs := StrByteLen(gofn.ToPtr("ab "), 2, 10).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrByteLen(gofn.ToPtr("chào"), 1, 4).Validate(ctxBg)
	assert.Equal(t, "byte_len", errs[0].Type())
}

func Test_StrEQ(t *testing.T) {
	errs := StrEQ(gofn.ToPtr("abc"), "abc").Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrEQ(gofn.ToPtr("abc"), "aBc").Validate(ctxBg)
	assert.Equal(t, "eq", errs[0].Type())
}

func Test_StrIn(t *testing.T) {
	errs := StrIn(gofn.ToPtr("chào"), "", "a", "chào").Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrIn(gofn.ToPtr("a"), "", "b").Validate(ctxBg)
	assert.Equal(t, "in", errs[0].Type())
}

func Test_StrNotIn(t *testing.T) {
	errs := StrNotIn(gofn.ToPtr("chào"), "", "a", "b").Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrNotIn(gofn.ToPtr("a"), "", "a", "b").Validate(ctxBg)
	assert.Equal(t, "not_in", errs[0].Type())
}

func Test_StrMatch(t *testing.T) {
	re := regexp.MustCompile("[0-9]+")
	errs := StrMatch(gofn.ToPtr("123"), re).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrMatch(gofn.ToPtr("abc"), re).Validate(ctxBg)
	assert.Equal(t, "match", errs[0].Type())
}

func Test_StrByteMatch(t *testing.T) {
	re := regexp.MustCompile("[0-9]+")
	errs := StrByteMatch(gofn.ToPtr("123"), re).Validate(ctxBg)
	assert.Equal(t, 0, len(errs))

	errs = StrByteMatch(gofn.ToPtr("abc"), re).Validate(ctxBg)
	assert.Equal(t, "byte_match", errs[0].Type())
}

// nolint: lll
func Test_StrIs_Functions(t *testing.T) {
	assert.Equal(t, 0, len(StrIsEmail(gofn.ToPtr("test@example.com")).Validate(ctxBg)))
	assert.Equal(t, "is_email", StrIsEmail(gofn.ToPtr("Test@eXample")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsExistingEmail(gofn.ToPtr("test@example.com")).Validate(ctxBg)))
	assert.Equal(t, "is_existing_email", StrIsExistingEmail(gofn.ToPtr("Test@eXample")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsURL(gofn.ToPtr("http://host.com")).Validate(ctxBg)))
	assert.Equal(t, "is_url", StrIsURL(gofn.ToPtr("host")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRequestURL(gofn.ToPtr("http://host.com:3000")).Validate(ctxBg)))
	assert.Equal(t, "is_request_url", StrIsRequestURL(gofn.ToPtr("abc3000")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRequestURI(gofn.ToPtr("http://host.com:3000")).Validate(ctxBg)))
	assert.Equal(t, "is_request_uri", StrIsRequestURI(gofn.ToPtr("abc3000")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsAlpha(gofn.ToPtr("abc")).Validate(ctxBg)))
	assert.Equal(t, "is_alpha", StrIsAlpha(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUTFLetter(gofn.ToPtr("Tiến")).Validate(ctxBg)))
	assert.Equal(t, "is_utf_letter", StrIsUTFLetter(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsAlphanumeric(gofn.ToPtr("abc012")).Validate(ctxBg)))
	assert.Equal(t, "is_alpha_numeric", StrIsAlphanumeric(gofn.ToPtr("abc-123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUTFLetterNumeric(gofn.ToPtr("abc012")).Validate(ctxBg)))
	assert.Equal(t, "is_utf_letter_numeric", StrIsUTFLetterNumeric(gofn.ToPtr("abc+123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsNumeric(gofn.ToPtr("012")).Validate(ctxBg)))
	assert.Equal(t, "is_numeric", StrIsNumeric(gofn.ToPtr("-123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUTFNumeric(gofn.ToPtr("012")).Validate(ctxBg)))
	assert.Equal(t, "is_utf_numeric", StrIsUTFNumeric(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUTFDigit(gofn.ToPtr("123456")).Validate(ctxBg)))
	assert.Equal(t, "is_utf_digit", StrIsUTFDigit(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsHexadecimal(gofn.ToPtr("abdcef")).Validate(ctxBg)))
	assert.Equal(t, "is_hexadecimal", StrIsHexadecimal(gofn.ToPtr("abdcefg")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsHexcolor(gofn.ToPtr("#ffffff")).Validate(ctxBg)))
	assert.Equal(t, "is_hexcolor", StrIsHexcolor(gofn.ToPtr("#fffffg")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRGBcolor(gofn.ToPtr("rgb(1,1,1)")).Validate(ctxBg)))
	assert.Equal(t, "is_rgbcolor", StrIsRGBcolor(gofn.ToPtr("rgb(1,1,a)")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsLowerCase(gofn.ToPtr("abc123")).Validate(ctxBg)))
	assert.Equal(t, "is_lower_case", StrIsLowerCase(gofn.ToPtr("aBc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUpperCase(gofn.ToPtr("ABC123")).Validate(ctxBg)))
	assert.Equal(t, "is_upper_case", StrIsUpperCase(gofn.ToPtr("aBc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrHasLowerCase(gofn.ToPtr("aBC123")).Validate(ctxBg)))
	assert.Equal(t, "has_lower_case", StrHasLowerCase(gofn.ToPtr("ABC123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrHasUpperCase(gofn.ToPtr("aBc123")).Validate(ctxBg)))
	assert.Equal(t, "has_upper_case", StrHasUpperCase(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsInt(gofn.ToPtr("-123")).Validate(ctxBg)))
	assert.Equal(t, "is_int", StrIsInt(gofn.ToPtr("123a")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsFloat(gofn.ToPtr("-123.123")).Validate(ctxBg)))
	assert.Equal(t, "is_float", StrIsFloat(gofn.ToPtr("123.123f")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrHasWhitespaceOnly(gofn.ToPtr("   ")).Validate(ctxBg)))
	assert.Equal(t, "has_whitespace_only", StrHasWhitespaceOnly(gofn.ToPtr("ab c")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrHasWhitespace(gofn.ToPtr(" b  c")).Validate(ctxBg)))
	assert.Equal(t, "has_whitespace", StrHasWhitespace(gofn.ToPtr("abc")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv3(gofn.ToPtr("2037da5d-a759-3ba8-bfeb-84519bb669c6")).Validate(ctxBg)))
	assert.Equal(t, "is_uuid_v3", StrIsUUIDv3(gofn.ToPtr("2037da5d-a759-3ba8-bfeb-84519bb669g6")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv4(gofn.ToPtr("fb3e2e7c-e478-4d76-aa84-9880d6eb67f4")).Validate(ctxBg)))
	assert.Equal(t, "is_uuid_v4", StrIsUUIDv4(gofn.ToPtr("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUUIDv5(gofn.ToPtr("aa43954a-ced3-5b84-9931-d3516b2e1867")).Validate(ctxBg)))
	assert.Equal(t, "is_uuid_v5", StrIsUUIDv5(gofn.ToPtr("aa43954a-ced3-5b84-9931-d3516b2e186g")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUUID(gofn.ToPtr("2037da5d-a759-3ba8-bfeb-84519bb669c6")).Validate(ctxBg)))
	assert.Equal(t, "is_uuid", StrIsUUID(gofn.ToPtr("fb3e2e7c-e478-4d76-aa84-9880d6eb67g4")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsULID(gofn.ToPtr("01G65Z755AFWAKHE12NY0CQ9FH")).Validate(ctxBg)))
	assert.Equal(t, "is_ulid", StrIsULID(gofn.ToPtr("01G65Z755AFWAKHE12NY0CQ9FHx")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsCreditCard(gofn.ToPtr("4242-8888-8888-8888")).Validate(ctxBg)))
	assert.Equal(t, "is_credit_card", StrIsCreditCard(gofn.ToPtr("5242-8888-8888-8888")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISBN13(gofn.ToPtr("978-3-16-148410-0")).Validate(ctxBg)))
	assert.Equal(t, "is_isbn13", StrIsISBN13(gofn.ToPtr("0-306-40615-2")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISBN10(gofn.ToPtr("0-306-40615-2")).Validate(ctxBg)))
	assert.Equal(t, "is_isbn10", StrIsISBN10(gofn.ToPtr("978-3-16-148410-0")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISBN(gofn.ToPtr("978-3-16-148410-0")).Validate(ctxBg)))

	assert.Equal(t, 0, len(StrIsJSON(gofn.ToPtr("123")).Validate(ctxBg)))
	assert.Equal(t, "is_json", StrIsJSON(gofn.ToPtr(`{"k":v}`)).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMultibyte(gofn.ToPtr("Chào buổi sáng")).Validate(ctxBg)))
	assert.Equal(t, "is_multibyte", StrIsMultibyte(gofn.ToPtr("abc 123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsASCII(gofn.ToPtr("hello")).Validate(ctxBg)))
	assert.Equal(t, "is_ascii", StrIsASCII(gofn.ToPtr("hello Tiến")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsPrintableASCII(gofn.ToPtr("hello")).Validate(ctxBg)))
	assert.Equal(t, "is_printable_ascii", StrIsPrintableASCII(gofn.ToPtr("hello\t")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsFullWidth(gofn.ToPtr("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")).Validate(ctxBg)))
	assert.Equal(t, "is_full_width", StrIsFullWidth(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsHalfWidth(gofn.ToPtr("abc123")).Validate(ctxBg)))
	assert.Equal(t, "is_half_width", StrIsHalfWidth(gofn.ToPtr("Ｔｙｐｅ　ｏｒ　Ｐａｓｔｅ")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsVariableWidth(gofn.ToPtr("Ｔｙｐｅ　ｏｒ　Paste")).Validate(ctxBg)))
	assert.Equal(t, "is_variable_width", StrIsVariableWidth(gofn.ToPtr("Ｔｙｐｅ")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsBase64(gofn.ToPtr("aGVsbG8gd29ybGQ=")).Validate(ctxBg)))
	assert.Equal(t, "is_base64", StrIsBase64(gofn.ToPtr("hello")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsFilePath(gofn.ToPtr(`C:\\abc`)).Validate(ctxBg)))

	assert.Equal(t, 0, len(StrIsWinFilePath(gofn.ToPtr(`C:\abc`)).Validate(ctxBg)))
	assert.Equal(t, "is_win_file_path", StrIsWinFilePath(gofn.ToPtr("/root")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUnixFilePath(gofn.ToPtr(`/root/path`)).Validate(ctxBg)))

	assert.Equal(t, 0, len(StrIsDataURI(gofn.ToPtr("data:image/png;base64,iVBORw0OHwAAAABJRU5ErkJggg==")).Validate(ctxBg)))
	assert.Equal(t, "is_data_uri", StrIsDataURI(gofn.ToPtr("example.com")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMagnetURI(gofn.ToPtr("magnet:?xt=urn:xxxxxx:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa&dn=bbb&tr=ccc")).Validate(ctxBg)))
	assert.Equal(t, "is_magnet_uri", StrIsMagnetURI(gofn.ToPtr("example.com")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISO3166Alpha2(gofn.ToPtr("VN")).Validate(ctxBg)))
	assert.Equal(t, "is_iso3166_alpha2", StrIsISO3166Alpha2(gofn.ToPtr("VND")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISO3166Alpha3(gofn.ToPtr("USA")).Validate(ctxBg)))
	assert.Equal(t, "is_iso3166_alpha3", StrIsISO3166Alpha3(gofn.ToPtr("USD")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISO639Alpha2(gofn.ToPtr("vi")).Validate(ctxBg)))
	assert.Equal(t, "is_iso639_alpha2", StrIsISO639Alpha2(gofn.ToPtr("vie")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISO639Alpha3b(gofn.ToPtr("vie")).Validate(ctxBg)))
	assert.Equal(t, "is_iso639_alpha3b", StrIsISO639Alpha3b(gofn.ToPtr("vi")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsDNSName(gofn.ToPtr("MX")).Validate(ctxBg)))
	assert.Equal(t, "is_dns_name", StrIsDNSName(gofn.ToPtr("")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3224(gofn.ToPtr("d9cf77bff9d00e47ad2d4841539bf72b0cfeff5e106819625e4f99f4")).Validate(ctxBg)))
	assert.Equal(t, "is_sha3_224", StrIsSHA3224(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3256(gofn.ToPtr("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Validate(ctxBg)))
	assert.Equal(t, "is_sha3_256", StrIsSHA3256(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3384(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)))
	assert.Equal(t, "is_sha3_384", StrIsSHA3384(gofn.ToPtr("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA3512(gofn.ToPtr("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b37ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")).Validate(ctxBg)))
	assert.Equal(t, "is_sha3_512", StrIsSHA3512(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA512(gofn.ToPtr("44a94f581e500c84b59bb38b990f499b4619c6774fd51e0e8534f89ee18dc3c6ea6ed096b37ff5a38517cec3ceb46bba2bd6989aef708da76fa6345810f9b1c4")).Validate(ctxBg)))
	assert.Equal(t, "is_sha512", StrIsSHA512(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA384(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)))
	assert.Equal(t, "is_sha384", StrIsSHA384(gofn.ToPtr("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA256(gofn.ToPtr("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Validate(ctxBg)))
	assert.Equal(t, "is_sha256", StrIsSHA256(gofn.ToPtr("3d479bdfb2d870868fef4f8dd56941e741c1d9c306f5ab0e6918e5f26ee1a0237c97606efb61c8cb4f173e3a526761dd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSHA1(gofn.ToPtr("4b2b79b6f371ca18f1216461cffeaddf6848a50e")).Validate(ctxBg)))
	assert.Equal(t, "is_sha1", StrIsSHA1(gofn.ToPtr("19697671a75511d50bbb5382ab6f53e5799481bb27968f0af7f42ca69474aac8")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsTiger192(gofn.ToPtr("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")).Validate(ctxBg)))
	assert.Equal(t, "is_tiger192", StrIsTiger192(gofn.ToPtr("8d0484881b88804442e390b0784003a1981db0b3")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsTiger160(gofn.ToPtr("8d0484881b88804442e390b0784003a1981db0b3")).Validate(ctxBg)))
	assert.Equal(t, "is_tiger160", StrIsTiger160(gofn.ToPtr("8d0484881b88804442e390b0784003a1981db0b31b5bf7af")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsTiger128(gofn.ToPtr("b8d02ffdd4d054b5327eed23b20e8716")).Validate(ctxBg)))
	assert.Equal(t, "is_tiger128", StrIsTiger128(gofn.ToPtr("8d0484881b88804442e390b0784003a1981db0b3")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRipeMD160(gofn.ToPtr("2330032576b1cc1df37b92abec140368c7396745")).Validate(ctxBg)))
	assert.Equal(t, "is_ripemd160", StrIsRipeMD160(gofn.ToPtr("afbed790a7824ade4a3ae531285f8fe8")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRipeMD128(gofn.ToPtr("afbed790a7824ade4a3ae531285f8fe8")).Validate(ctxBg)))
	assert.Equal(t, "is_ripemd128", StrIsRipeMD128(gofn.ToPtr("2330032576b1cc1df37b92abec140368c7396745")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsCRC32(gofn.ToPtr("cd71f992")).Validate(ctxBg)))
	assert.Equal(t, "is_crc32", StrIsCRC32(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsCRC32b(gofn.ToPtr("94d6f306")).Validate(ctxBg)))
	assert.Equal(t, "is_crc32b", StrIsCRC32b(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMD5(gofn.ToPtr("ec02c59dee6faaca3189bace969c22d3")).Validate(ctxBg)))
	assert.Equal(t, "is_md5", StrIsMD5(gofn.ToPtr("abc")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMD4(gofn.ToPtr("538a2a786ca1bc47694c0bc3f3ac3228")).Validate(ctxBg)))
	assert.Equal(t, "is_md4", StrIsMD4(gofn.ToPtr("abc")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsDialString(gofn.ToPtr("golang.org:3000")).Validate(ctxBg)))
	assert.Equal(t, "is_dial_string", StrIsDialString(gofn.ToPtr("abc")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsIP(gofn.ToPtr("1.1.1.1")).Validate(ctxBg)))
	assert.Equal(t, "is_ip", StrIsIP(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsPort(gofn.ToPtr("1234")).Validate(ctxBg)))
	assert.Equal(t, "is_port", StrIsPort(gofn.ToPtr("70000")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsIPv4(gofn.ToPtr("1.1.1.1")).Validate(ctxBg)))
	assert.Equal(t, "is_ipv4", StrIsIPv4(gofn.ToPtr("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsIPv6(gofn.ToPtr("FEDC:BA98:768A:0C98:FEBA:CB87:7678:1111")).Validate(ctxBg)))
	assert.Equal(t, "is_ipv6", StrIsIPv6(gofn.ToPtr("1.1.1.1")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsCIDR(gofn.ToPtr("255.255.255.0/32")).Validate(ctxBg)))
	assert.Equal(t, "is_cidr", StrIsCIDR(gofn.ToPtr("255.255.255.255")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMAC(gofn.ToPtr("01-80-C2-FF-FF-FF")).Validate(ctxBg)))
	assert.Equal(t, "is_mac", StrIsMAC(gofn.ToPtr("1.1.1.1")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsHost(gofn.ToPtr("example.com")).Validate(ctxBg)))
	assert.Equal(t, "is_host", StrIsHost(gofn.ToPtr("abc/123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsMongoID(gofn.ToPtr("507f1f77bcf86cd799439011")).Validate(ctxBg)))
	assert.Equal(t, "is_mongo_id", StrIsMongoID(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsLatitude(gofn.ToPtr("38.8951")).Validate(ctxBg)))
	assert.Equal(t, "is_latitude", StrIsLatitude(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsLongitude(gofn.ToPtr("-77.0364")).Validate(ctxBg)))
	assert.Equal(t, "is_longitude", StrIsLongitude(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsIMEI(gofn.ToPtr("11222222333333")).Validate(ctxBg)))
	assert.Equal(t, "is_imei", StrIsIMEI(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsIMSI(gofn.ToPtr("313460000000001")).Validate(ctxBg)))
	assert.Equal(t, "is_imsi", StrIsIMSI(gofn.ToPtr("11222222333333")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRsaPublicKey(gofn.ToPtr("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsFCmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ=="), 512).Validate(ctxBg)))
	assert.Equal(t, "is_rsa_public_key", StrIsRsaPublicKey(gofn.ToPtr("MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ1oCVw+tgIf52KApencj1hHW/KtvqwfnmQsFCmb4IhywfFAPbJ5qx1jX1HPDb+v/yMXzGbvlcE2kFzjYFy/LUsCAwEAAQ=="), 1024).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRegex(gofn.ToPtr("[0-9a-z]+")).Validate(ctxBg)))

	assert.Equal(t, 0, len(StrIsSSN(gofn.ToPtr("111-22-3333")).Validate(ctxBg)))
	assert.Equal(t, "is_ssn", StrIsSSN(gofn.ToPtr("111.22.3333")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsSemver(gofn.ToPtr("1.2.3")).Validate(ctxBg)))
	assert.Equal(t, "is_semver", StrIsSemver(gofn.ToPtr("1.2")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsTime(gofn.ToPtr("2023-10-01"), "2006-02-01").Validate(ctxBg)))
	assert.Equal(t, "is_time", StrIsTime(gofn.ToPtr("2023/10/01"), "2006-02-01").Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsUnixTime(gofn.ToPtr("123456789")).Validate(ctxBg)))
	assert.Equal(t, "is_unix_time", StrIsUnixTime(gofn.ToPtr("12345678912345678912345")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRFC3339(gofn.ToPtr("2020-12-09T16:09:53-00:00")).Validate(ctxBg)))
	assert.Equal(t, "is_rfc3339", StrIsRFC3339(gofn.ToPtr("2020-12-09 16:09:53+00:00")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsRFC3339WithoutZone(gofn.ToPtr("2020-12-09T16:09:53")).Validate(ctxBg)))
	assert.Equal(t, "is_rfc3339_without_zone", StrIsRFC3339WithoutZone(gofn.ToPtr("2020-12-09T16:09:53-00:00")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsISO4217(gofn.ToPtr("USD")).Validate(ctxBg)))
	assert.Equal(t, "is_iso4217", StrIsISO4217(gofn.ToPtr("usd")).Validate(ctxBg)[0].Type())

	assert.Equal(t, 0, len(StrIsE164(gofn.ToPtr("+442071838750")).Validate(ctxBg)))
	assert.Equal(t, "is_e164", StrIsE164(gofn.ToPtr("abc123")).Validate(ctxBg)[0].Type())
}
