package stringvalidation

import (
	extVld "github.com/asaskevich/govalidator"

	"github.com/tiendc/go-validator/base"
)

func IsEmail[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsEmail(string(s)), nil
}
func IsExistingEmail[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsExistingEmail(string(s)), nil
}
func IsURL[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsURL(string(s)), nil
}
func IsRequestURL[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRequestURL(string(s)), nil
}
func IsRequestURI[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRequestURI(string(s)), nil
}
func IsAlpha[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsAlpha(string(s)), nil
}
func IsUTFLetter[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUTFLetter(string(s)), nil
}
func IsAlphanumeric[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsAlphanumeric(string(s)), nil
}
func IsUTFLetterNumeric[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUTFLetterNumeric(string(s)), nil
}
func IsNumeric[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsNumeric(string(s)), nil
}
func IsUTFNumeric[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUTFNumeric(string(s)), nil
}
func IsUTFDigit[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUTFDigit(string(s)), nil
}
func IsHexadecimal[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsHexadecimal(string(s)), nil
}
func IsHexcolor[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsHexcolor(string(s)), nil
}
func IsRGBcolor[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRGBcolor(string(s)), nil
}
func IsLowerCase[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsLowerCase(string(s)), nil
}
func IsUpperCase[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUpperCase(string(s)), nil
}
func HasLowerCase[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.HasLowerCase(string(s)), nil
}
func HasUpperCase[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.HasUpperCase(string(s)), nil
}
func IsInt[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsInt(string(s)), nil
}
func IsFloat[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsFloat(string(s)), nil
}
func HasWhitespaceOnly[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.HasWhitespaceOnly(string(s)), nil
}
func HasWhitespace[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.HasWhitespace(string(s)), nil
}
func IsUUIDv3[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUUIDv3(string(s)), nil
}
func IsUUIDv4[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUUIDv4(string(s)), nil
}
func IsUUIDv5[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUUIDv5(string(s)), nil
}
func IsUUID[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUUID(string(s)), nil
}
func IsULID[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsULID(string(s)), nil
}
func IsCreditCard[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsCreditCard(string(s)), nil
}
func IsISBN13[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISBN13(string(s)), nil
}
func IsISBN10[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISBN10(string(s)), nil
}
func IsISBN[T base.String](s T) (bool, []base.ErrorParam) {
	success, _ := IsISBN10(s)
	if success {
		return true, nil
	}
	return IsISBN13(s)
}
func IsJSON[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsJSON(string(s)), nil
}
func IsMultibyte[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMultibyte(string(s)), nil
}
func IsASCII[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsASCII(string(s)), nil
}
func IsPrintableASCII[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsPrintableASCII(string(s)), nil
}
func IsFullWidth[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsFullWidth(string(s)), nil
}
func IsHalfWidth[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsHalfWidth(string(s)), nil
}
func IsVariableWidth[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsVariableWidth(string(s)), nil
}
func IsBase64[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsBase64(string(s)), nil
}
func IsFilePath[T base.String](s T) (bool, []base.ErrorParam) {
	success, _ := IsWinFilePath(s)
	if success {
		return true, nil
	}
	return IsUnixFilePath(s)
}
func IsWinFilePath[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsWinFilePath(string(s)), nil
}
func IsUnixFilePath[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUnixFilePath(string(s)), nil
}
func IsDataURI[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsDataURI(string(s)), nil
}
func IsMagnetURI[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMagnetURI(string(s)), nil
}
func IsISO3166Alpha2[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISO3166Alpha2(string(s)), nil
}
func IsISO3166Alpha3[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISO3166Alpha3(string(s)), nil
}
func IsISO639Alpha2[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISO693Alpha2(string(s)), nil // FIXME: IsISO693 is a wrong name?
}
func IsISO639Alpha3b[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISO693Alpha3b(string(s)), nil // FIXME: IsISO693 is a wrong name?
}
func IsDNSName[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsDNSName(string(s)), nil
}
func IsSHA3224[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA3224(string(s)), nil
}
func IsSHA3256[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA3256(string(s)), nil
}
func IsSHA3384[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA3384(string(s)), nil
}
func IsSHA3512[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA3512(string(s)), nil
}
func IsSHA512[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA512(string(s)), nil
}
func IsSHA384[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA384(string(s)), nil
}
func IsSHA256[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA256(string(s)), nil
}
func IsSHA1[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSHA1(string(s)), nil
}
func IsTiger192[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsTiger192(string(s)), nil
}
func IsTiger160[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsTiger160(string(s)), nil
}
func IsTiger128[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsTiger128(string(s)), nil
}
func IsRipeMD160[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRipeMD160(string(s)), nil
}
func IsRipeMD128[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRipeMD128(string(s)), nil
}
func IsCRC32[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsCRC32(string(s)), nil
}
func IsCRC32b[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsCRC32b(string(s)), nil
}
func IsMD5[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMD5(string(s)), nil
}
func IsMD4[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMD4(string(s)), nil
}
func IsDialString[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsDialString(string(s)), nil
}
func IsIP[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsIP(string(s)), nil
}
func IsPort[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsPort(string(s)), nil
}
func IsIPv4[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsIPv4(string(s)), nil
}
func IsIPv6[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsIPv6(string(s)), nil
}
func IsCIDR[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsCIDR(string(s)), nil
}
func IsMAC[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMAC(string(s)), nil
}
func IsHost[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsHost(string(s)), nil
}
func IsMongoID[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsMongoID(string(s)), nil
}
func IsLatitude[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsLatitude(string(s)), nil
}
func IsLongitude[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsLongitude(string(s)), nil
}
func IsIMEI[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsIMEI(string(s)), nil
}
func IsIMSI[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsIMSI(string(s)), nil
}
func IsRsaPublicKey[T base.String](s T, keyLen int) (bool, []base.ErrorParam) {
	return extVld.IsRsaPublicKey(string(s), keyLen), nil
}
func IsRegex[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRegex(string(s)), nil
}
func IsSSN[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSSN(string(s)), nil
}
func IsSemver[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsSemver(string(s)), nil
}
func IsTime[T base.String](s T, layout string) (bool, []base.ErrorParam) {
	return extVld.IsTime(string(s), layout), nil
}
func IsUnixTime[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsUnixTime(string(s)), nil
}
func IsRFC3339[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRFC3339(string(s)), nil
}
func IsRFC3339WithoutZone[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsRFC3339WithoutZone(string(s)), nil
}
func IsISO4217[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsISO4217(string(s)), nil
}
func IsE164[T base.String](s T) (bool, []base.ErrorParam) {
	return extVld.IsE164(string(s)), nil
}
