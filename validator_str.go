package validation

import (
	"regexp"

	"github.com/tiendc/go-validator/base"
	strFunc "github.com/tiendc/go-validator/base/string"
)

const (
	strType = "string"
)

// StrLen validates the input string must have length of runes in the specified range
func StrLen[T base.String](s *T, min, max int) SingleValidator {
	return ptrCall3[T]("len", strType, "Min", "Max", strFunc.RuneLen[T])(s, min, max)
}

// StrByteLen validates the input string must have length of bytes in the specified range
func StrByteLen[T base.String](s *T, min, max int) SingleValidator {
	return ptrCall3[T]("byte_len", strType, "Min", "Max", strFunc.ByteLen[T])(s, min, max)
}

// StrEQ validates the input string must equal to the specified value
func StrEQ[T base.String](v *T, s T) SingleValidator {
	return ptrCall2[T]("eq", strType, "TargetValue", strFunc.EQ[T])(v, s)
}

// StrIn validates the input string must be in the specified values
func StrIn[T base.String](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("in", strType, "TargetValue", strFunc.In[T])(v, s...)
}

// StrNotIn validates the input string must be not in the specified values
func StrNotIn[T base.String](v *T, s ...T) SingleValidator {
	return ptrCall2N[T]("not_in", strType, "TargetValue", strFunc.NotIn[T])(v, s...)
}

// StrMatch validates the input string must have runes matching the specified regex
func StrMatch[T base.String](v *T, re *regexp.Regexp) SingleValidator {
	return ptrCall2[T]("match", strType, "TargetValue", strFunc.RuneMatch[T])(v, re)
}

// StrByteMatch validates the input string must have bytes matching the specified regex
func StrByteMatch[T base.String](v *T, re *regexp.Regexp) SingleValidator {
	return ptrCall2[T]("byte_match", strType, "TargetValue", strFunc.ByteMatch[T])(v, re)
}

// StrIsEmail validates the input string must be a valid email
func StrIsEmail[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_email", strType, strFunc.IsEmail[T])(s)
}

// StrIsExistingEmail validates the input string must be a valid email of existing domain
func StrIsExistingEmail[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_existing_email", strType, strFunc.IsExistingEmail[T])(s)
}

// StrIsURL validates the input string must be a valid URL
func StrIsURL[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_url", strType, strFunc.IsURL[T])(s)
}

// StrIsRequestURL validates the input string must be a valid request URL (URL confirm to RFC 3986)
func StrIsRequestURL[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_request_url", strType, strFunc.IsRequestURL[T])(s)
}

// StrIsRequestURI validates the input string must be a valid request URI
func StrIsRequestURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_request_uri", strType, strFunc.IsRequestURI[T])(s)
}

// StrIsAlpha validates the input string must contain only characters in range a-zA-Z
func StrIsAlpha[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_alpha", strType, strFunc.IsAlpha[T])(s)
}

// StrIsUTFLetter validates the input string must contain only UTF letters
func StrIsUTFLetter[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_letter", strType, strFunc.IsUTFLetter[T])(s)
}

// StrIsAlphanumeric validates the input string must contain only characters in range a-zA-Z0-9
func StrIsAlphanumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_alpha_numeric", strType, strFunc.IsAlphanumeric[T])(s)
}

// StrIsUTFLetterNumeric validates the input string must contain only UTF letters and numerics
func StrIsUTFLetterNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_letter_numeric", strType, strFunc.IsUTFLetterNumeric[T])(s)
}

// StrIsNumeric validates the input string must contain only characters in range 0-9
func StrIsNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_numeric", strType, strFunc.IsNumeric[T])(s)
}

// StrIsUTFNumeric validates the input string must contain only UTF numeric characters
func StrIsUTFNumeric[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_numeric", strType, strFunc.IsUTFNumeric[T])(s)
}

// StrIsUTFDigit validates the input string must contain only UTF digits
func StrIsUTFDigit[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_utf_digit", strType, strFunc.IsUTFDigit[T])(s)
}

// StrIsHexadecimal validates the input string must be in hex format
func StrIsHexadecimal[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_hexadecimal", strType, strFunc.IsHexadecimal[T])(s)
}

// StrIsHexcolor validates the input string must be a hex color such as #fab or #aabbcc
func StrIsHexcolor[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_hexcolor", strType, strFunc.IsHexcolor[T])(s)
}

// StrIsRGBcolor validates the input string must be a rgb color such as rgb(10,20,30)
func StrIsRGBcolor[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rgbcolor", strType, strFunc.IsRGBcolor[T])(s)
}

// StrIsLowerCase validates the input string must be in lower case
func StrIsLowerCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_lower_case", strType, strFunc.IsLowerCase[T])(s)
}

// StrIsUpperCase validates the input string must be in upper case
func StrIsUpperCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_upper_case", strType, strFunc.IsUpperCase[T])(s)
}

// StrHasLowerCase validates the input string must contain at least 1 lower case character
func StrHasLowerCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_lower_case", strType, strFunc.HasLowerCase[T])(s)
}

// StrHasUpperCase validates the input string must contain at least 1 upper case character
func StrHasUpperCase[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_upper_case", strType, strFunc.HasUpperCase[T])(s)
}

// StrIsInt validates the input string must be an integral number
func StrIsInt[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_int", strType, strFunc.IsInt[T])(s)
}

// StrIsFloat validates the input string must be a floating number
func StrIsFloat[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_float", strType, strFunc.IsFloat[T])(s)
}

// StrHasWhitespaceOnly validates the input string must contain whitespaces only
func StrHasWhitespaceOnly[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_whitespace_only", strType, strFunc.HasWhitespaceOnly[T])(s)
}

// StrHasWhitespace validates the input string must contain at least 1 whitespace character
func StrHasWhitespace[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("has_whitespace", strType, strFunc.HasWhitespace[T])(s)
}

// StrIsUUIDv3 validates the input string must be a UUID v3
func StrIsUUIDv3[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v3", strType, strFunc.IsUUIDv3[T])(s)
}

// StrIsUUIDv4 validates the input string must be a UUID v4
func StrIsUUIDv4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v4", strType, strFunc.IsUUIDv4[T])(s)
}

// StrIsUUIDv5 validates the input string must be a UUID v5
func StrIsUUIDv5[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid_v5", strType, strFunc.IsUUIDv5[T])(s)
}

// StrIsUUID validates the input string must be a UUID v3 or UUID v4 or UUID v5
func StrIsUUID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_uuid", strType, strFunc.IsUUID[T])(s)
}

// StrIsULID validates the input string must be a ULID
func StrIsULID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ulid", strType, strFunc.IsULID[T])(s)
}

// StrIsCreditCard validates the input string must be a credit card number
func StrIsCreditCard[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_credit_card", strType, strFunc.IsCreditCard[T])(s)
}

// StrIsISBN10 validates the input string must be a ISBN v10
func StrIsISBN10[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn10", strType, strFunc.IsISBN10[T])(s)
}

// StrIsISBN13 validates the input string must be a ISBN v13
func StrIsISBN13[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn13", strType, strFunc.IsISBN13[T])(s)
}

// StrIsISBN validates the input string must be a ISBN v10 or ISBN v13
func StrIsISBN[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_isbn", strType, strFunc.IsISBN[T])(s)
}

// StrIsJSON validates the input string must be in JSON format
func StrIsJSON[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_json", strType, strFunc.IsJSON[T])(s)
}

// StrIsMultibyte validates the input string must be a multiple-byte string
func StrIsMultibyte[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_multibyte", strType, strFunc.IsMultibyte[T])(s)
}

// StrIsASCII validates the input string must contain only ASCII characters
func StrIsASCII[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ascii", strType, strFunc.IsASCII[T])(s)
}

// StrIsPrintableASCII validates the input string must contain only printable ASCII characters
func StrIsPrintableASCII[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_printable_ascii", strType, strFunc.IsPrintableASCII[T])(s)
}

// StrIsFullWidth validates the input string must contain only full-width characters
func StrIsFullWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_full_width", strType, strFunc.IsFullWidth[T])(s)
}

// StrIsHalfWidth validates the input string must contain only half-width characters
func StrIsHalfWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_half_width", strType, strFunc.IsHalfWidth[T])(s)
}

// StrIsVariableWidth validates the input string must contain variable-width characters
func StrIsVariableWidth[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_variable_width", strType, strFunc.IsVariableWidth[T])(s)
}

// StrIsBase64 validates the input string must be in Base64 format
func StrIsBase64[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_base64", strType, strFunc.IsBase64[T])(s)
}

// StrIsFilePath validates the input string must be a file path in both Windows or Unix
func StrIsFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_file_path", strType, strFunc.IsFilePath[T])(s)
}

// StrIsWinFilePath validates the input string must be a file path in Windows
func StrIsWinFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_win_file_path", strType, strFunc.IsWinFilePath[T])(s)
}

// StrIsUnixFilePath validates the input string must be a file path in Unix
func StrIsUnixFilePath[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_unix_file_path", strType, strFunc.IsUnixFilePath[T])(s)
}

// StrIsDataURI validates the input string must be a data URI
func StrIsDataURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_data_uri", strType, strFunc.IsDataURI[T])(s)
}

// StrIsMagnetURI validates the input string must be a magnet URI
func StrIsMagnetURI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_magnet_uri", strType, strFunc.IsMagnetURI[T])(s)
}

// StrIsISO3166Alpha2 validates the input string must be one of ISO3166 Alpha2 country codes
func StrIsISO3166Alpha2[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso3166_alpha2", strType, strFunc.IsISO3166Alpha2[T])(s)
}

// StrIsISO3166Alpha3 validates the input string must be one of ISO3166 Alpha3 country codes
func StrIsISO3166Alpha3[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso3166_alpha3", strType, strFunc.IsISO3166Alpha3[T])(s)
}

// StrIsISO639Alpha2 validates the input string must be one of ISO639 Alpha2 language codes
func StrIsISO639Alpha2[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso639_alpha2", strType, strFunc.IsISO639Alpha2[T])(s)
}

// StrIsISO639Alpha3b validates the input string must be one of ISO639 Alpha3b language codes
func StrIsISO639Alpha3b[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso639_alpha3b", strType, strFunc.IsISO639Alpha3b[T])(s)
}

// StrIsDNSName validates the input string must be a domain name
func StrIsDNSName[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_dns_name", strType, strFunc.IsDNSName[T])(s)
}

// StrIsSHA3224 validates the input string must be in SHA3-224 format
func StrIsSHA3224[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_224", strType, strFunc.IsSHA3224[T])(s)
}

// StrIsSHA3256 validates the input string must be in SHA3-256 format
func StrIsSHA3256[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_256", strType, strFunc.IsSHA3256[T])(s)
}

// StrIsSHA3384 validates the input string must be in SHA3-384 format
func StrIsSHA3384[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_384", strType, strFunc.IsSHA3384[T])(s)
}

// StrIsSHA3512 validates the input string must be in SHA3-512 format
func StrIsSHA3512[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha3_512", strType, strFunc.IsSHA3512[T])(s)
}

// StrIsSHA512 validates the input string must be in SHA512 format
func StrIsSHA512[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha512", strType, strFunc.IsSHA512[T])(s)
}

// StrIsSHA384 validates the input string must be in SHA384 format
func StrIsSHA384[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha384", strType, strFunc.IsSHA384[T])(s)
}

// StrIsSHA256 validates the input string must be in SHA256 format
func StrIsSHA256[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha256", strType, strFunc.IsSHA256[T])(s)
}

// StrIsTiger192 validates the input string must be in Tiger192 format
func StrIsTiger192[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger192", strType, strFunc.IsTiger192[T])(s)
}

// StrIsTiger160 validates the input string must be in Tiger160 format
func StrIsTiger160[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger160", strType, strFunc.IsTiger160[T])(s)
}

// StrIsRipeMD160 validates the input string must be in RipeMD160 format
func StrIsRipeMD160[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ripemd160", strType, strFunc.IsRipeMD160[T])(s)
}

// StrIsSHA1 validates the input string must be in SHA1 format
func StrIsSHA1[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_sha1", strType, strFunc.IsSHA1[T])(s)
}

// StrIsTiger128 validates the input string must be in Tiger128 format
func StrIsTiger128[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_tiger128", strType, strFunc.IsTiger128[T])(s)
}

// StrIsRipeMD128 validates the input string must be in RipeMD128 format
func StrIsRipeMD128[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ripemd128", strType, strFunc.IsRipeMD128[T])(s)
}

// StrIsCRC32 validates the input string must be in CRC32 format
func StrIsCRC32[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_crc32", strType, strFunc.IsCRC32[T])(s)
}

// StrIsCRC32b validates the input string must be in CRC32b format
func StrIsCRC32b[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_crc32b", strType, strFunc.IsCRC32b[T])(s)
}

// StrIsMD5 validates the input string must be in MD5 format
func StrIsMD5[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_md5", strType, strFunc.IsMD5[T])(s)
}

// StrIsMD4 validates the input string must be in MD4 format
func StrIsMD4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_md4", strType, strFunc.IsMD4[T])(s)
}

// StrIsDialString validates the input string must be a dial string such as a hostname, IP, or a port
func StrIsDialString[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_dial_string", strType, strFunc.IsDialString[T])(s)
}

// StrIsIP validates the input string must be in IP v4 or IP v6 format
func StrIsIP[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ip", strType, strFunc.IsIP[T])(s)
}

// StrIsPort validates the input string must be a valid port number (range 1-65535)
func StrIsPort[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_port", strType, strFunc.IsPort[T])(s)
}

// StrIsIPv4 validates the input string must be in IP v4 format
func StrIsIPv4[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ipv4", strType, strFunc.IsIPv4[T])(s)
}

// StrIsIPv6 validates the input string must be in IP v6 format
func StrIsIPv6[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ipv6", strType, strFunc.IsIPv6[T])(s)
}

// StrIsCIDR validates the input string must be in CIDR format
func StrIsCIDR[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_cidr", strType, strFunc.IsCIDR[T])(s)
}

// StrIsMAC validates the input string must be in MAC address format
func StrIsMAC[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_mac", strType, strFunc.IsMAC[T])(s)
}

// StrIsHost validates the input string must be a hostname or an IP
func StrIsHost[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_host", strType, strFunc.IsHost[T])(s)
}

// StrIsMongoID validates the input string must be a Mongo ID
func StrIsMongoID[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_mongo_id", strType, strFunc.IsMongoID[T])(s)
}

// StrIsLatitude validates the input string must be a latitude number
func StrIsLatitude[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_latitude", strType, strFunc.IsLatitude[T])(s)
}

// StrIsLongitude validates the input string must be a longitude number
func StrIsLongitude[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_longitude", strType, strFunc.IsLongitude[T])(s)
}

// StrIsIMEI validates the input string must be an IMEI
func StrIsIMEI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_imei", strType, strFunc.IsIMEI[T])(s)
}

// StrIsIMSI validates the input string must be an IMSI
func StrIsIMSI[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_imsi", strType, strFunc.IsIMSI[T])(s)
}

// StrIsRsaPublicKey validates the input string must be an RSA public key
func StrIsRsaPublicKey[T base.String](s *T, keyLen int) SingleValidator {
	return ptrCall2[T]("is_rsa_public_key", "key_len", strType, strFunc.IsRsaPublicKey[T])(s, keyLen)
}

// StrIsRegex validates the input string must be a regex
func StrIsRegex[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_regex", strType, strFunc.IsRegex[T])(s)
}

// StrIsSSN validates the input string must be a SSN
func StrIsSSN[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_ssn", strType, strFunc.IsSSN[T])(s)
}

// StrIsSemver validates the input string must be a Semver
func StrIsSemver[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_semver", strType, strFunc.IsSemver[T])(s)
}

// StrIsTime validates the input string must be a date time of the specified layout
func StrIsTime[T base.String](s *T, layout string) SingleValidator {
	return ptrCall2[T]("is_time", "layout", strType, strFunc.IsTime[T])(s, layout)
}

// StrIsUnixTime validates the input string must be a unix time value
func StrIsUnixTime[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_unix_time", strType, strFunc.IsUnixTime[T])(s)
}

// StrIsRFC3339 validates the input string must be a date time of RFC3339 layout
func StrIsRFC3339[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rfc3339", strType, strFunc.IsRFC3339[T])(s)
}

// StrIsRFC3339WithoutZone validates the input string must be a date time of RFC3339 layout without time zone
func StrIsRFC3339WithoutZone[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_rfc3339_without_zone", strType, strFunc.IsRFC3339WithoutZone[T])(s)
}

// StrIsISO4217 validates the input string must be one of ISO4217 currency codes
func StrIsISO4217[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_iso4217", strType, strFunc.IsISO4217[T])(s)
}

// StrIsE164 validates the input string must be in E164 format
func StrIsE164[T base.String](s *T) SingleValidator {
	return ptrCall1[T]("is_e164", strType, strFunc.IsE164[T])(s)
}
