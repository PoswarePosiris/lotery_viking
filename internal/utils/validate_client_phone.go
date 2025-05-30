package utils

import "regexp"

// ValidatePhoneNumber checks if the given phone number is valid:
// - French mobile numbers in local format (06, 07)
// - French mobile numbers in international format (+336, +337)
// - Foreign mobile numbers with international prefix (+ or 00)
func ValidatePhoneNumber(clientPhone string) bool {
	// French mobile numbers (local format): 06xxxxxxxx or 07xxxxxxxx
	// French mobile numbers (international): +336xxxxxxxx or +337xxxxxxxx
	// International mobile numbers: +XX... or 00XX... (9 to 15 digits total after the prefix)
	pattern := `^(?:0[67]\d{8})$|^(?:\+33[67]\d{8})$|^(?:(?:\+|00)[1-9]\d{6,14})$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(clientPhone)
}

// Test for every country
// type CountryRule struct {
// 	CountryCode   string         // e.g., "FR"
// 	DialCode      string         // e.g., "+33"
// 	MobilePattern *regexp.Regexp // regex for mobile numbers
// 	Allowed       bool           // control allow/block
// }

// var countryRules = []CountryRule{
// 	{
// 		CountryCode:   "FR", // France
// 		DialCode:      "+33",
// 		MobilePattern: regexp.MustCompile(`^(?:0[67]\d{8})$|^(?:\+33[67]\d{8})$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CH", // Switzerland
// 		DialCode:      "+41",
// 		MobilePattern: regexp.MustCompile(`^(\+41|0041|0)7[5-9]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BE", // Belgium
// 		DialCode:      "+32",
// 		MobilePattern: regexp.MustCompile(`^(\+32|0032|0)4\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "DE", // Germany
// 		DialCode:      "+49",
// 		MobilePattern: regexp.MustCompile(`^(\+49|0049|0)1[56789]\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NL", // Netherlands
// 		DialCode:      "+31",
// 		MobilePattern: regexp.MustCompile(`^(\+31|0031|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "ES", // Spain
// 		DialCode:      "+34",
// 		MobilePattern: regexp.MustCompile(`^(\+34|0034|0)6[6-9]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "IT", // Italy
// 		DialCode:      "+39",
// 		MobilePattern: regexp.MustCompile(`^(\+39|0039|0)3\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PT", // Portugal
// 		DialCode:      "+351",
// 		MobilePattern: regexp.MustCompile(`^(\+351|00351|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PL", // Poland
// 		DialCode:      "+48",
// 		MobilePattern: regexp.MustCompile(`^(\+48|0048|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "US", // United States
// 		DialCode:      "+1",
// 		MobilePattern: regexp.MustCompile(`^\+1[2-9]\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "GB", // United Kingdom
// 		DialCode:      "+44",
// 		MobilePattern: regexp.MustCompile(`^\+447\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "IN", // India
// 		DialCode:      "+91",
// 		MobilePattern: regexp.MustCompile(`^\+91[6-9]\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AU", // Australia
// 		DialCode:      "+61",
// 		MobilePattern: regexp.MustCompile(`^(\+61|0061|0)4\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CA", // Canada
// 		DialCode:      "+1",
// 		MobilePattern: regexp.MustCompile(`^\+1[2-9]\d{9}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "CN", // China
// 		DialCode:      "+86",
// 		MobilePattern: regexp.MustCompile(`^(\+86|0086|0)1\d{10}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "JP", // Japan
// 		DialCode:      "+81",
// 		MobilePattern: regexp.MustCompile(`^(\+81|0081|0)8\d{8}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "KR", // South Korea
// 		DialCode:      "+82",
// 		MobilePattern: regexp.MustCompile(`^(\+82|0082|0)1\d{8}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "TW", // Taiwan
// 		DialCode:      "+886",
// 		MobilePattern: regexp.MustCompile(`^(\+886|00886|0)9\d{8}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "TH", // Thailand
// 		DialCode:      "+66",
// 		MobilePattern: regexp.MustCompile(`^(\+66|0066|0)8\d{8}$`),
// 		Allowed:       true,
// 	}, {
// 		CountryCode:   "VN", // Vietnam
// 		DialCode:      "+84",
// 		MobilePattern: regexp.MustCompile(`^(\+84|0084|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BR", // Brazil
// 		DialCode:      "+55",
// 		MobilePattern: regexp.MustCompile(`^(\+55|0055|0)[1-9]{2}9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "RU", // Russia
// 		DialCode:      "+7",
// 		MobilePattern: regexp.MustCompile(`^(\+7|007|8)9\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "ZA", // South Africa
// 		DialCode:      "+27",
// 		MobilePattern: regexp.MustCompile(`^(\+27|0027|0)[6-8]\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SG", // Singapore
// 		DialCode:      "+65",
// 		MobilePattern: regexp.MustCompile(`^(\+65|0065|0)8\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MY", // Malaysia
// 		DialCode:      "+60",
// 		MobilePattern: regexp.MustCompile(`^(\+60|0060|0)1[0-9]\d{7,8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "ID", // Indonesia
// 		DialCode:      "+62",
// 		MobilePattern: regexp.MustCompile(`^(\+62|0062|0)8\d{8,11}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SA", // Saudi Arabia
// 		DialCode:      "+966",
// 		MobilePattern: regexp.MustCompile(`^(\+966|00966|0)5\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AE", // United Arab Emirates
// 		DialCode:      "+971",
// 		MobilePattern: regexp.MustCompile(`^(\+971|00971|0)5[0256]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NG", // Nigeria
// 		DialCode:      "+234",
// 		MobilePattern: regexp.MustCompile(`^(\+234|00234|0)[7-9]{1}[0-1]{1}\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "KE", // Kenya
// 		DialCode:      "+254",
// 		MobilePattern: regexp.MustCompile(`^(\+254|00254|0)[17]\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "EG", // Egypt
// 		DialCode:      "+20",
// 		MobilePattern: regexp.MustCompile(`^(\+20|0020|0)1[0-2]\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AR", // Argentina
// 		DialCode:      "+54",
// 		MobilePattern: regexp.MustCompile(`^(\+54|0054|0)911\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MX", // Mexico
// 		DialCode:      "+52",
// 		MobilePattern: regexp.MustCompile(`^(\+52|0052|0)1\d{10}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CO", // Colombia
// 		DialCode:      "+57",
// 		MobilePattern: regexp.MustCompile(`^(\+57|0057|0)3\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PH", // Philippines
// 		DialCode:      "+63",
// 		MobilePattern: regexp.MustCompile(`^(\+63|0063|0)9\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "TR", // Turkey
// 		DialCode:      "+90",
// 		MobilePattern: regexp.MustCompile(`^(\+90|0090|0)5\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "UA", // Ukraine
// 		DialCode:      "+380",
// 		MobilePattern: regexp.MustCompile(`^(\+380|00380|0)[3-9]\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "GR", // Greece
// 		DialCode:      "+30",
// 		MobilePattern: regexp.MustCompile(`^(\+30|0030|0)69\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SE", // Sweden
// 		DialCode:      "+46",
// 		MobilePattern: regexp.MustCompile(`^(\+46|0046|0)7\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NO", // Norway
// 		DialCode:      "+47",
// 		MobilePattern: regexp.MustCompile(`^(\+47|0047|0)4\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "FI", // Finland
// 		DialCode:      "+358",
// 		MobilePattern: regexp.MustCompile(`^(\+358|00358|0)4\d{7,8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "DK", // Denmark
// 		DialCode:      "+45",
// 		MobilePattern: regexp.MustCompile(`^(\+45|0045|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "IE", // Ireland
// 		DialCode:      "+353",
// 		MobilePattern: regexp.MustCompile(`^(\+353|00353|0)8[35679]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AT", // Austria
// 		DialCode:      "+43",
// 		MobilePattern: regexp.MustCompile(`^(\+43|0043|0)6\d{8,10}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "HU", // Hungary
// 		DialCode:      "+36",
// 		MobilePattern: regexp.MustCompile(`^(\+36|0036|0)[237]0\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CZ", // Czech Republic
// 		DialCode:      "+420",
// 		MobilePattern: regexp.MustCompile(`^(\+420|00420|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SK", // Slovakia
// 		DialCode:      "+421",
// 		MobilePattern: regexp.MustCompile(`^(\+421|00421|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "IL", // Israel
// 		DialCode:      "+972",
// 		MobilePattern: regexp.MustCompile(`^(\+972|00972|0)5\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NZ", // New Zealand
// 		DialCode:      "+64",
// 		MobilePattern: regexp.MustCompile(`^(\+64|0064|0)2\d{8,9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "HK", // Hong Kong
// 		DialCode:      "+852",
// 		MobilePattern: regexp.MustCompile(`^(\+852|00852|0)[569]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MO", // Macau
// 		DialCode:      "+853",
// 		MobilePattern: regexp.MustCompile(`^(\+853|00853|0)6\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SG", // Singapore
// 		DialCode:      "+65",
// 		MobilePattern: regexp.MustCompile(`^(\+65|0065|0)8\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MY", // Malaysia
// 		DialCode:      "+60",
// 		MobilePattern: regexp.MustCompile(`^(\+60|0060|0)1[0-9]\d{7,8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PK", // Pakistan
// 		DialCode:      "+92",
// 		MobilePattern: regexp.MustCompile(`^(\+92|0092|0)3\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BD", // Bangladesh
// 		DialCode:      "+880",
// 		MobilePattern: regexp.MustCompile(`^(\+880|00880|0)1\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "LK", // Sri Lanka
// 		DialCode:      "+94",
// 		MobilePattern: regexp.MustCompile(`^(\+94|0094|0)7\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NP", // Nepal
// 		DialCode:      "+977",
// 		MobilePattern: regexp.MustCompile(`^(\+977|00977|0)9\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CL", // Chile
// 		DialCode:      "+56",
// 		MobilePattern: regexp.MustCompile(`^(\+56|0056|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PE", // Peru
// 		DialCode:      "+51",
// 		MobilePattern: regexp.MustCompile(`^(\+51|0051|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "VE", // Venezuela
// 		DialCode:      "+58",
// 		MobilePattern: regexp.MustCompile(`^(\+58|0058|0)4\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "EC", // Ecuador
// 		DialCode:      "+593",
// 		MobilePattern: regexp.MustCompile(`^(\+593|00593|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BO", // Bolivia
// 		DialCode:      "+591",
// 		MobilePattern: regexp.MustCompile(`^(\+591|00591|0)[67]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PY", // Paraguay
// 		DialCode:      "+595",
// 		MobilePattern: regexp.MustCompile(`^(\+595|00595|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "UY", // Uruguay
// 		DialCode:      "+598",
// 		MobilePattern: regexp.MustCompile(`^(\+598|00598|0)9\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CR", // Costa Rica
// 		DialCode:      "+506",
// 		MobilePattern: regexp.MustCompile(`^(\+506|00506|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "PA", // Panama
// 		DialCode:      "+507",
// 		MobilePattern: regexp.MustCompile(`^(\+507|00507|0)6\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "DO", // Dominican Republic
// 		DialCode:      "+1",
// 		MobilePattern: regexp.MustCompile(`^(\+1|001|1)8[024]9\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "GT", // Guatemala
// 		DialCode:      "+502",
// 		MobilePattern: regexp.MustCompile(`^(\+502|00502|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SV", // El Salvador
// 		DialCode:      "+503",
// 		MobilePattern: regexp.MustCompile(`^(\+503|00503|0)[67]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "HN", // Honduras
// 		DialCode:      "+504",
// 		MobilePattern: regexp.MustCompile(`^(\+504|00504|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "NI", // Nicaragua
// 		DialCode:      "+505",
// 		MobilePattern: regexp.MustCompile(`^(\+505|00505|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "RO", // Romania
// 		DialCode:      "+40",
// 		MobilePattern: regexp.MustCompile(`^(\+40|0040|0)7\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BG", // Bulgaria
// 		DialCode:      "+359",
// 		MobilePattern: regexp.MustCompile(`^(\+359|00359|0)8[789]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "HR", // Croatia
// 		DialCode:      "+385",
// 		MobilePattern: regexp.MustCompile(`^(\+385|00385|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "RS", // Serbia
// 		DialCode:      "+381",
// 		MobilePattern: regexp.MustCompile(`^(\+381|00381|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "SI", // Slovenia
// 		DialCode:      "+386",
// 		MobilePattern: regexp.MustCompile(`^(\+386|00386|0)[3-7]\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BA", // Bosnia and Herzegovina
// 		DialCode:      "+387",
// 		MobilePattern: regexp.MustCompile(`^(\+387|00387|0)6\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MK", // North Macedonia
// 		DialCode:      "+389",
// 		MobilePattern: regexp.MustCompile(`^(\+389|00389|0)7\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AL", // Albania
// 		DialCode:      "+355",
// 		MobilePattern: regexp.MustCompile(`^(\+355|00355|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MT", // Malta
// 		DialCode:      "+356",
// 		MobilePattern: regexp.MustCompile(`^(\+356|00356|0)9\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "CY", // Cyprus
// 		DialCode:      "+357",
// 		MobilePattern: regexp.MustCompile(`^(\+357|00357|0)9\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "LU", // Luxembourg
// 		DialCode:      "+352",
// 		MobilePattern: regexp.MustCompile(`^(\+352|00352|0)6\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "IS", // Iceland
// 		DialCode:      "+354",
// 		MobilePattern: regexp.MustCompile(`^(\+354|00354|0)\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "KZ", // Kazakhstan
// 		DialCode:      "+7",
// 		MobilePattern: regexp.MustCompile(`^(\+7|007|8)7\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "UZ", // Uzbekistan
// 		DialCode:      "+998",
// 		MobilePattern: regexp.MustCompile(`^(\+998|00998|0)\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "GE", // Georgia
// 		DialCode:      "+995",
// 		MobilePattern: regexp.MustCompile(`^(\+995|00995|0)5\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AM", // Armenia
// 		DialCode:      "+374",
// 		MobilePattern: regexp.MustCompile(`^(\+374|00374|0)\d{8}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "AZ", // Azerbaijan
// 		DialCode:      "+994",
// 		MobilePattern: regexp.MustCompile(`^(\+994|00994|0)\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "BY", // Belarus
// 		DialCode:      "+375",
// 		MobilePattern: regexp.MustCompile(`^(\+375|00375|0)\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "MD", // Moldova
// 		DialCode:      "+373",
// 		MobilePattern: regexp.MustCompile(`^(\+373|00373|0)6\d{7}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "KG", // Kyrgyzstan
// 		DialCode:      "+996",
// 		MobilePattern: regexp.MustCompile(`^(\+996|00996|0)\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "TJ", // Tajikistan
// 		DialCode:      "+992",
// 		MobilePattern: regexp.MustCompile(`^(\+992|00992|0)\d{9}$`),
// 		Allowed:       true,
// 	},
// 	{
// 		CountryCode:   "TM", // Turkmenistan
// 		DialCode:      "+993",
// 		MobilePattern: regexp.MustCompile(`^(\+993|00993|0)\d{8}$`),
// 		Allowed:       true,
// 	},

// 	// Add more countries as needed
// }

// func ValidatePhoneNumber(phone string) bool {
// 	for _, rule := range countryRules {
// 		if rule.MobilePattern.MatchString(phone) {
// 			return rule.Allowed
// 		}
// 	}
// 	return false
// }
