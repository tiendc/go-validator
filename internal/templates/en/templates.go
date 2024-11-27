package en

var (
	FmtTemplates = map[string]string{
		// General
		"group":        `{{.Field}} must satisfy all specified validations`,
		"one_of":       `{{.Field}} must satisfy at least one of the specified validations`,
		"exact_one_of": `{{.Field}} must satisfy exact one of the specified validations`,
		"not_of":       `{{.Field}} must not satisfy any of the specified validations`,
		"if":           `{{.Field}} must satisfy the specified condition`, // deprecated: use `must`
		"must":         `{{.Field}} must satisfy the specified condition`,
		"nil":          `{{.Field}} must be nil`,
		"not_nil":      `{{.Field}} must be not nil`,
		"required":     `{{.Field}} is required`,

		// Number
		"number_eq":           `{{.Field}} must be equal to {{.TargetValue}}`,
		"number_gt":           `{{.Field}} must be greater than {{.Min}}`,
		"number_gte":          `{{.Field}} must be greater than or equal to {{.Min}}`,
		"number_lt":           `{{.Field}} must be less than {{.Max}}`,
		"number_lte":          `{{.Field}} must be less than or equal to {{.Max}}`,
		"number_range":        `{{.Field}} must be in range {{.Min}} to {{.Max}}`,
		"number_in":           `{{.Field}} must be one of {{.TargetValue}}`,
		"number_not_in":       `{{.Field}} must not be one of {{.TargetValue}}`,
		"number_divisible_by": `{{.Field}} must be divisible by {{.TargetValue}}`,
		"number_js_safe_int":  `{{.Field}} must be in Javascript safe int range`,

		// String
		"string_eq":         `{{.Field}} must be equal to {{.TargetValue}}`,
		"string_len":        `{{.Field}} length must be in range {{.Min}} to {{.Max}}`,
		"string_byte_len":   `{{.Field}} byte-length must be in range {{.Min}} to {{.Max}}`,
		"string_in":         `{{.Field}} must be one of {{.TargetValue}}`,
		"string_not_in":     `{{.Field}} must not be one of {{.TargetValue}}`,
		"string_match":      `{{.Field}} must match the unicode pattern {{.TargetValue}}`,
		"string_byte_match": `{{.Field}} must match the pattern {{.TargetValue}}`,

		// String "is" / "has"
		"string_is_email":                `{{.Field}} must be a valid email`,
		"string_is_existing_email":       `{{.Field}} must be a valid existing email`,
		"string_is_url":                  `{{.Field}} must be a valid URL`,
		"string_is_request_url":          `{{.Field}} must be a valid request URL`,
		"string_is_request_uri":          `{{.Field}} must be a valid request URI`,
		"string_is_alpha":                `{{.Field}} must contain English letters only (a-zA-Z)`,
		"string_is_alpha_numeric":        `{{.Field}} must contain English letters and digits only (a-zA-Z0-9)`,
		"string_is_utf_letter":           `{{.Field}} must contain unicode letters only`,
		"string_is_utf_letter_numeric":   `{{.Field}} must contain unicode letters and numbers only`,
		"string_is_numeric":              `{{.Field}} must contain English number letters only (0-9)`,
		"string_is_utf_numeric":          `{{.Field}} must contain unicode number letters only`,
		"string_is_utf_digit":            `{{.Field}} must contain unicode radix-10 decimal digits only`,
		"string_is_hexadecimal":          `{{.Field}} must be a valid hexadecimal number`,
		"string_is_hexcolor":             `{{.Field}} must be a valid Hex color code`,
		"string_is_rgbcolor":             `{{.Field}} must be a valid RGB color in form of rgb(R,G,B)`,
		"string_is_lower_case":           `{{.Field}} must contain lower case letters only`,
		"string_is_upper_case":           `{{.Field}} must contain upper case letters only`,
		"string_has_lower_case":          `{{.Field}} must contain lower case letters`,
		"string_has_upper_case":          `{{.Field}} must contain upper case letters`,
		"string_is_int":                  `{{.Field}} must be a valid integer number`,
		"string_is_float":                `{{.Field}} must be a valid floating number`,
		"string_has_whitespace_only":     `{{.Field}} must contain whitespace only`,
		"string_has_whitespace":          `{{.Field}} must contain whitespace`,
		"string_is_uuid_v3":              `{{.Field}} must be a valid UUID v3`,
		"string_is_uuid_v4":              `{{.Field}} must be a valid UUID v4`,
		"string_is_uuid_v5":              `{{.Field}} must be a valid UUID v5`,
		"string_is_uuid":                 `{{.Field}} must be a valid UUID`,
		"string_is_ulid":                 `{{.Field}} must be a valid ULID`,
		"string_is_credit_card":          `{{.Field}} must be a valid credit card number`,
		"string_is_isbn10":               `{{.Field}} must be a valid ISBN v10`,
		"string_is_isbn13":               `{{.Field}} must be a valid ISBN v13`,
		"string_is_isbn":                 `{{.Field}} must be a valid ISBN (either ISBN v10 or ISBN v13)`,
		"string_is_json":                 `{{.Field}} must be in valid JSON format`,
		"string_is_multibyte":            `{{.Field}} must contain multibyte characters`,
		"string_is_ascii":                `{{.Field}} must contain only ASCII characters`,
		"string_is_printable_ascii":      `{{.Field}} must contain only printable ASCII characters`,
		"string_is_full_width":           `{{.Field}} must contain full-width characters`,
		"string_is_half_width":           `{{.Field}} must contain half-width characters`,
		"string_is_variable_width":       `{{.Field}} must contain both full-width and half-width characters`,
		"string_is_base64":               `{{.Field}} must be a valid Base64 encoded string`,
		"string_is_file_path":            `{{.Field}} must be a valid file path`,
		"string_is_win_file_path":        `{{.Field}} must be a valid Windows file path`,
		"string_is_unix_file_path":       `{{.Field}} must be a valid Unix file path`,
		"string_is_data_uri":             `{{.Field}} must be a valid base64-encoded data URI`,
		"string_is_magnet_uri":           `{{.Field}} must be a valid magnet URI`,
		"string_is_iso3166_alpha2":       `{{.Field}} must be a valid ISO3166 Alpha2 country code`,
		"string_is_iso3166_alpha3":       `{{.Field}} must be a valid ISO3166 Alpha3 country code`,
		"string_is_iso639_alpha2":        `{{.Field}} must be a valid ISO639 Alpha2 language code`,
		"string_is_iso639_alpha3b":       `{{.Field}} must be a valid ISO639 Alpha3b language code`,
		"string_is_dns_name":             `{{.Field}} must be a valid DNS name`,
		"string_is_sha3_224":             `{{.Field}} must be a valid SHA3-224`,
		"string_is_sha3_256":             `{{.Field}} must be a valid SHA3-256`,
		"string_is_sha3_384":             `{{.Field}} must be a valid SHA3-384`,
		"string_is_sha3_512":             `{{.Field}} must be a valid SHA3-512`,
		"string_is_sha512":               `{{.Field}} must be a valid SHA512`,
		"string_is_sha384":               `{{.Field}} must be a valid SHA384`,
		"string_is_sha256":               `{{.Field}} must be a valid SHA256`,
		"string_is_sha1":                 `{{.Field}} must be a valid SHA1`,
		"string_is_tiger192":             `{{.Field}} must be a valid Tiger192`,
		"string_is_tiger160":             `{{.Field}} must be a valid Tiger160`,
		"string_is_tiger128":             `{{.Field}} must be a valid Tiger128`,
		"string_is_ripemd160":            `{{.Field}} must be a valid RipeMD160`,
		"string_is_ripemd128":            `{{.Field}} must be a valid RipeMD128`,
		"string_is_crc32":                `{{.Field}} must be a valid CRC32`,
		"string_is_crc32b":               `{{.Field}} must be a valid CRC32b`,
		"string_is_md5":                  `{{.Field}} must be a valid MD5`,
		"string_is_md4":                  `{{.Field}} must be a valid MD4`,
		"string_is_dial_string":          `{{.Field}} must be a valid dial string`,
		"string_is_ip":                   `{{.Field}} must be a valid IP address (either IPv4 or Ipv6 address)`,
		"string_is_port":                 `{{.Field}} must be a valid host Port`,
		"string_is_ipv4":                 `{{.Field}} must be a valid IPv4 address`,
		"string_is_ipv6":                 `{{.Field}} must be a valid IPv6 address`,
		"string_is_cidr":                 `{{.Field}} must be a valid CIDR`,
		"string_is_mac":                  `{{.Field}} must be a valid MAC address`,
		"string_is_host":                 `{{.Field}} must be a valid host IP or domain`,
		"string_is_mongo_id":             `{{.Field}} must be a valid Mongo ID`,
		"string_is_latitude":             `{{.Field}} must be a valid latitude`,
		"string_is_longitude":            `{{.Field}} must be a valid longitude`,
		"string_is_imei":                 `{{.Field}} must be a valid IMEI number`,
		"string_is_imsi":                 `{{.Field}} must be a valid IMSI code`,
		"string_is_rsa_public_key":       `{{.Field}} must be a valid RSA public key`,
		"string_is_regex":                `{{.Field}} must be a valid Regex`,
		"string_is_ssn":                  `{{.Field}} must be a valid Social Security number`,
		"string_is_semver":               `{{.Field}} must be a valid Semver`,
		"string_is_time":                 `{{.Field}} must be a valid time`,
		"string_is_unix_time":            `{{.Field}} must be a valid Unix time`,
		"string_is_rfc3339":              `{{.Field}} must be a valid RFC3339 date time`,
		"string_is_rfc3339_without_zone": `{{.Field}} must be a valid RFC3339 date time without zone`,
		"string_is_iso4217":              `{{.Field}} must be a valid ISO4217 currency code`,
		"string_is_e164":                 `{{.Field}} must be a valid E164 phone number`,

		// Slice
		"slice_len":         `{{.Field}}: number of array items must be in range {{.Min}} to {{.Max}}`,
		"slice_elem_in":     `{{.Field}}: array items must be one of {{.TargetValue}}`,
		"slice_elem_not_in": `{{.Field}}: array items must not be one of {{.TargetValue}}`,
		"slice_elem_range":  `{{.Field}}: array items must be in range {{.Min}} to {{.Max}}`,
		"slice_unique":      `{{.Field}}: array items must be unique`,
		"slice_sorted":      `{{.Field}}: array items must be sorted in ascending order`,
		"slice_sorted_desc": `{{.Field}}: array items must be sorted in descending order`,

		// Map
		"map_len":          `{{.Field}}: number of items must be in range {{.Min}} to {{.Max}}`,
		"map_has_key":      `{{.Field}}: map must contain keys {{.TargetValue}}`,
		"map_not_have_key": `{{.Field}}: map must not contain keys {{.TargetValue}}`,
		"map_key_in":       `{{.Field}}: map keys must be one of {{.TargetValue}}`,
		"map_key_not_in":   `{{.Field}}: map keys must not be one of {{.TargetValue}}`,
		"map_key_range":    `{{.Field}}: map keys must be in range {{.Min}} to {{.Max}}`,
		"map_value_in":     `{{.Field}}: map values must be one of {{.TargetValue}}`,
		"map_value_not_in": `{{.Field}}: map values must not be one of {{.TargetValue}}`,
		"map_value_range":  `{{.Field}}: map values must be in range {{.Min}} to {{.Max}}`,
		"map_value_unique": `{{.Field}}: map values must be unique`,

		// Time
		"time_eq":     `{{.Field}} must be equal to {{.TargetValue}}`,
		"time_gt":     `{{.Field}} must be greater than {{.Min}}`,
		"time_gte":    `{{.Field}} must be greater than or equal to {{.Min}}`,
		"time_lt":     `{{.Field}} must be less than {{.Max}}`,
		"time_lte":    `{{.Field}} must be less than or equal to {{.Max}}`,
		"time_valid":  `{{.Field}} must be a valid time`,
		"time_range":  `{{.Field}} must be in range {{.Min}} to {{.Max}}`,
		"time_in":     `{{.Field}} must be one of {{.TargetValue}}`,
		"time_not_in": `{{.Field}} must not be one of {{.TargetValue}}`,
	}
)
