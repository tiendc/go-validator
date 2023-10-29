package validation

import (
	"regexp"

	"github.com/tiendc/go-validator/base"
	strFunc "github.com/tiendc/go-validator/base/string"
)

const (
	strType = "string"
)

func StrLen[T base.String](s *T, min, max int) SingleValidator {
	return ptrCall3[T]("len", strType, "Min", "Max", strFunc.RuneLen[T])(s, min, max)
}
func StrByteLen[T base.String](s *T, min, max int) SingleValidator {
	return ptrCall3[T]("byte_len", strType, "Min", "Max", strFunc.ByteLen[T])(s, min, max)
}

func StrEQ[T base.String](v *T, s T) SingleValidator {
	return ptrCall2[T]("eq", strType, "TargetValue", strFunc.EQ[T])(v, s)
}
func StrIn[T base.String](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("in", strType, "TargetValue", strFunc.In[T])(v, s...)
}
func StrNotIn[T base.String](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("not_in", strType, "TargetValue", strFunc.NotIn[T])(v, s...)
}

func StrMatch[T base.String](v *T, re *regexp.Regexp) SingleValidator {
	return ptrCall2[T]("match", strType, "TargetValue", strFunc.RuneMatch[T])(v, re)
}
func StrByteMatch[T base.String](v *T, re *regexp.Regexp) SingleValidator {
	return ptrCall2[T]("byte_match", strType, "TargetValue", strFunc.ByteMatch[T])(v, re)
}

func StrIsEmail[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_email", strType, strFunc.IsEmail[T])(s)
}
func StrIsExistingEmail[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_existing_email", strType, strFunc.IsExistingEmail[T])(s)
}
func StrIsURL[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_url", strType, strFunc.IsURL[T])(s)
}
func StrIsRequestURL[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_request_url", strType, strFunc.IsRequestURL[T])(s)
}
func StrIsRequestURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_request_uri", strType, strFunc.IsRequestURI[T])(s)
}
func StrIsAlpha[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_alpha", strType, strFunc.IsAlpha[T])(s)
}
func StrIsUTFLetter[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_letter", strType, strFunc.IsUTFLetter[T])(s)
}
func StrIsAlphanumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_alpha_numeric", strType, strFunc.IsAlphanumeric[T])(s)
}
func StrIsUTFLetterNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_letter_numeric", strType, strFunc.IsUTFLetterNumeric[T])(s)
}
func StrIsNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_numeric", strType, strFunc.IsNumeric[T])(s)
}
func StrIsUTFNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_numeric", strType, strFunc.IsUTFNumeric[T])(s)
}
func StrIsUTFDigit[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_digit", strType, strFunc.IsUTFDigit[T])(s)
}
func StrIsHexadecimal[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_hexadecimal", strType, strFunc.IsHexadecimal[T])(s)
}
func StrIsHexcolor[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_hexcolor", strType, strFunc.IsHexcolor[T])(s)
}
func StrIsRGBcolor[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rgbcolor", strType, strFunc.IsRGBcolor[T])(s)
}
func StrIsLowerCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_lower_case", strType, strFunc.IsLowerCase[T])(s)
}
func StrIsUpperCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_upper_case", strType, strFunc.IsUpperCase[T])(s)
}
func StrHasLowerCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_lower_case", strType, strFunc.HasLowerCase[T])(s)
}
func StrHasUpperCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_upper_case", strType, strFunc.HasUpperCase[T])(s)
}
func StrIsInt[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_int", strType, strFunc.IsInt[T])(s)
}
func StrIsFloat[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_float", strType, strFunc.IsFloat[T])(s)
}
func StrHasWhitespaceOnly[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_whitespace_only", strType, strFunc.HasWhitespaceOnly[T])(s)
}
func StrHasWhitespace[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_whitespace", strType, strFunc.HasWhitespace[T])(s)
}
func StrIsUUIDv3[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v3", strType, strFunc.IsUUIDv3[T])(s)
}
func StrIsUUIDv4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v4", strType, strFunc.IsUUIDv4[T])(s)
}
func StrIsUUIDv5[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v5", strType, strFunc.IsUUIDv5[T])(s)
}
func StrIsUUID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid", strType, strFunc.IsUUID[T])(s)
}
func StrIsULID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ulid", strType, strFunc.IsULID[T])(s)
}
func StrIsCreditCard[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_credit_card", strType, strFunc.IsCreditCard[T])(s)
}
func StrIsISBN10[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn10", strType, strFunc.IsISBN10[T])(s)
}
func StrIsISBN13[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn13", strType, strFunc.IsISBN13[T])(s)
}
func StrIsISBN[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn", strType, strFunc.IsISBN[T])(s)
}
func StrIsJSON[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_json", strType, strFunc.IsJSON[T])(s)
}
func StrIsMultibyte[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_multibyte", strType, strFunc.IsMultibyte[T])(s)
}
func StrIsASCII[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ascii", strType, strFunc.IsASCII[T])(s)
}
func StrIsPrintableASCII[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_printable_ascii", strType, strFunc.IsPrintableASCII[T])(s)
}
func StrIsFullWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_full_width", strType, strFunc.IsFullWidth[T])(s)
}
func StrIsHalfWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_half_width", strType, strFunc.IsHalfWidth[T])(s)
}
func StrIsVariableWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_variable_width", strType, strFunc.IsVariableWidth[T])(s)
}
func StrIsBase64[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_base64", strType, strFunc.IsBase64[T])(s)
}
func StrIsFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_file_path", strType, strFunc.IsFilePath[T])(s)
}
func StrIsWinFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_win_file_path", strType, strFunc.IsWinFilePath[T])(s)
}
func StrIsUnixFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_unix_file_path", strType, strFunc.IsUnixFilePath[T])(s)
}
func StrIsDataURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_data_uri", strType, strFunc.IsDataURI[T])(s)
}
func StrIsMagnetURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_magnet_uri", strType, strFunc.IsMagnetURI[T])(s)
}
func StrIsISO3166Alpha2[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso3166_alpha2", strType, strFunc.IsISO3166Alpha2[T])(s)
}
func StrIsISO3166Alpha3[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso3166_alpha3", strType, strFunc.IsISO3166Alpha3[T])(s)
}
func StrIsISO639Alpha2[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso639_alpha2", strType, strFunc.IsISO639Alpha2[T])(s)
}
func StrIsISO639Alpha3b[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso639_alpha3b", strType, strFunc.IsISO639Alpha3b[T])(s)
}
func StrIsDNSName[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_dns_name", strType, strFunc.IsDNSName[T])(s)
}
func StrIsSHA3224[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_224", strType, strFunc.IsSHA3224[T])(s)
}
func StrIsSHA3256[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_256", strType, strFunc.IsSHA3256[T])(s)
}
func StrIsSHA3384[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_384", strType, strFunc.IsSHA3384[T])(s)
}
func StrIsSHA3512[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_512", strType, strFunc.IsSHA3512[T])(s)
}
func StrIsSHA512[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha512", strType, strFunc.IsSHA512[T])(s)
}
func StrIsSHA384[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha384", strType, strFunc.IsSHA384[T])(s)
}
func StrIsSHA256[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha256", strType, strFunc.IsSHA256[T])(s)
}
func StrIsTiger192[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger192", strType, strFunc.IsTiger192[T])(s)
}
func StrIsTiger160[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger160", strType, strFunc.IsTiger160[T])(s)
}
func StrIsRipeMD160[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ripemd160", strType, strFunc.IsRipeMD160[T])(s)
}
func StrIsSHA1[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha1", strType, strFunc.IsSHA1[T])(s)
}
func StrIsTiger128[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger128", strType, strFunc.IsTiger128[T])(s)
}
func StrIsRipeMD128[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ripemd128", strType, strFunc.IsRipeMD128[T])(s)
}
func StrIsCRC32[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_crc32", strType, strFunc.IsCRC32[T])(s)
}
func StrIsCRC32b[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_crc32b", strType, strFunc.IsCRC32b[T])(s)
}
func StrIsMD5[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_md5", strType, strFunc.IsMD5[T])(s)
}
func StrIsMD4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_md4", strType, strFunc.IsMD4[T])(s)
}
func StrIsDialString[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_dial_string", strType, strFunc.IsDialString[T])(s)
}
func StrIsIP[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ip", strType, strFunc.IsIP[T])(s)
}
func StrIsPort[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_port", strType, strFunc.IsPort[T])(s)
}
func StrIsIPv4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ipv4", strType, strFunc.IsIPv4[T])(s)
}
func StrIsIPv6[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ipv6", strType, strFunc.IsIPv6[T])(s)
}
func StrIsCIDR[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_cidr", strType, strFunc.IsCIDR[T])(s)
}
func StrIsMAC[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_mac", strType, strFunc.IsMAC[T])(s)
}
func StrIsHost[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_host", strType, strFunc.IsHost[T])(s)
}
func StrIsMongoID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_mongo_id", strType, strFunc.IsMongoID[T])(s)
}
func StrIsLatitude[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_latitude", strType, strFunc.IsLatitude[T])(s)
}
func StrIsLongitude[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_longitude", strType, strFunc.IsLongitude[T])(s)
}
func StrIsIMEI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_imei", strType, strFunc.IsIMEI[T])(s)
}
func StrIsIMSI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_imsi", strType, strFunc.IsIMSI[T])(s)
}
func StrIsRsaPublicKey[T base.String](s *T, keyLen int) SingleValidator {
	return ptrCall2[T]("is_rsa_public_key", "key_len", strType, strFunc.IsRsaPublicKey[T])(s, keyLen)
}
func StrIsRegex[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_regex", strType, strFunc.IsRegex[T])(s)
}
func StrIsSSN[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ssn", strType, strFunc.IsSSN[T])(s)
}
func StrIsSemver[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_semver", strType, strFunc.IsSemver[T])(s)
}
func StrIsTime[T base.String](s *T, layout string) SingleValidator {
	return ptrCall2[T]("is_time", "layout", strType, strFunc.IsTime[T])(s, layout)
}
func StrIsUnixTime[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_unix_time", strType, strFunc.IsUnixTime[T])(s)
}
func StrIsRFC3339[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rfc3339", strType, strFunc.IsRFC3339[T])(s)
}
func StrIsRFC3339WithoutZone[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rfc3339_without_zone", strType, strFunc.IsRFC3339WithoutZone[T])(s)
}
func StrIsISO4217[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso4217", strType, strFunc.IsISO4217[T])(s)
}
func StrIsE164[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_e164", strType, strFunc.IsE164[T])(s)
}
