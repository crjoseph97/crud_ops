package domain

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//To use in mysql queries application_type = "e-paper"
const (
	ErrInvalidAdminID            = "Invalid admin ID"
	ErrInvalidAdminName          = "Invalid admin Name"
	ErrInvalidSchedule           = "Report schedule not found"
	ErrNameAlreadyInUse          = "Name already in use"
	ReportScheduleNameReg string = `^[A-Za-z\d\s_@.&\-\(\)\[\]]+$`
)

// LoadLocation load the location
var LoadLocation = make(map[string]*time.Location)

// MetricDefaultOrder for sorting purposes
var MetricDefaultOrder = []string{
	"engagement_time",
	"revenue in USD",
	"revenue in INR",
	"revenue",
	"subscriptions",
	"views",
	"users",
}

// CountryMapKey ISO code mapping
var CountryMapKey = map[string]string{
	"jersey":                              "JEY",
	"st. kitts & nevis":                   "KNA",
	"comoros":                             "COM",
	"san marino":                          "SMR",
	"dominica":                            "DMA",
	"monaco":                              "MCO",
	"guernsey":                            "GGY",
	"liechtenstein":                       "LIE",
	"guadeloupe":                          "GLP",
	"aruba":                               "ABW",
	"georgia":                             "GEO",
	"puerto rico":                         "PRI",
	"bosnia and herzegovina":              "BIH",
	"northern cyprus":                     "CYP",
	"northern mariana islands":            "MNP",
	"ecuador":                             "ECU",
	"gibraltar":                           "GIB",
	"togo":                                "TGO",
	"macau sar china":                     "MAC",
	"france":                              "FRA",
	"sierra leone":                        "SLE",
	"curaçao":                             "CUW",
	"tajikistan":                          "TJK",
	"maldives":                            "MDV",
	"myanmar":                             "MMR",
	"nepal":                               "NPL",
	"palestinian territories":             "PSE",
	"colombia":                            "COL",
	"hungary":                             "HUN",
	"greece":                              "GRC",
	"kyrgyzstan":                          "KGZ",
	"armenia":                             "ARM",
	"ireland":                             "IRL",
	"iraq":                                "IRQ",
	"liberia":                             "LBR",
	"netherlands":                         "NLD",
	"somalia":                             "SOM",
	"serbia":                              "SRB",
	"sint maarten":                        "SXM",
	"afghanistan":                         "AFG",
	"chile":                               "CHL",
	"western sahara":                      "ESH",
	"sweden":                              "SWE",
	"vanuatu":                             "VUT",
	"fiji":                                "FJI",
	"russia":                              "RUS",
	"timor-leste":                         "TLS",
	"china":                               "CHN",
	"swaziland":                           "SWZ",
	"congo - kinshasa":                    "COD",
	"antarctica":                          "ATA",
	"french southern and antarctic lands": "ATF",
	"zimbabwe":                            "ZWE",
	"egypt":                               "EGY",
	"indonesia":                           "IDN",
	"iceland":                             "ISL",
	"israel":                              "ISR",
	"pakistan":                            "PAK",
	"albania":                             "ALB",
	"morocco":                             "MAR",
	"mauritius":                           "MUS",
	"belarus":                             "BLR",
	"algeria":                             "DZA",
	"guinea-bissau":                       "GNB",
	"guinea bissau":                       "GNB",
	"latvia":                              "LVA",
	"oman":                                "OMN",
	"el salvador":                         "SLV",
	"west bank":                           "PSE",
	"vatican city":                        "VAT",
	"vatican city state":                  "VAT",
	"holy see (vatican city state)":       "VAT",
	"ivory coast":                         "CIV",
	"cameroon":                            "CMR",
	"namibia":                             "NAM",
	"saudi arabia":                        "SAU",
	"spain":                               "ESP",
	"south korea":                         "KOR",
	"seychelles":                          "SYC",
	"cambodia":                            "KHM",
	"solomon islands":                     "SLB",
	"mongolia":                            "MNG",
	"belgium":                             "BEL",
	"honduras":                            "HND",
	"falkland islands":                    "FLK",
	"eritrea":                             "ERI",
	"estonia":                             "EST",
	"north korea":                         "PRK",
	"sudan":                               "SDN",
	"uganda":                              "UGA",
	"cyprus":                              "CYP",
	"kenya":                               "KEN",
	"kazakhstan":                          "KAZ",
	"niger":                               "NER",
	"finland":                             "FIN",
	"japan":                               "JPN",
	"french guiana":                       "GUF",
	"laos":                                "LAO",
	"republic of serbia":                  "SRB",
	"venezuela":                           "VEN",
	"canada":                              "CAN",
	"germany":                             "DEU",
	"mauritania":                          "MRT",
	"uzbekistan":                          "UZB",
	"bosnia & herzegovina":                "BIH",
	"belize":                              "BLZ",
	"mexico":                              "MEX",
	"gambia":                              "GMB",
	"myanmar (burma)":                     "MMR",
	"azerbaijan":                          "AZE",
	"ghana":                               "GHA",
	"bermuda":                             "BMU",
	"united kingdom":                      "GBR",
	"papua new guinea":                    "PNG",
	"central african republic":            "CAF",
	"switzerland":                         "CHE",
	"guyana":                              "GUY",
	"italy":                               "ITA",
	"united states":                       "USA",
	"united arab emirates":                "ARE",
	"benin":                               "BEN",
	"paraguay":                            "PRY",
	"dominican republic":                  "DOM",
	"gabon":                               "GAB",
	"botswana":                            "BWA",
	"australia":                           "AUS",
	"cuba":                                "CUB",
	"equatorial guinea":                   "GNQ",
	"brunei":                              "BRN",
	"costa rica":                          "CRI",
	"iran":                                "IRN",
	"austria":                             "AUT",
	"the bahamas":                         "BHS",
	"romania":                             "ROU",
	"turks & caicos islands":              "TCA",
	"burkina faso":                        "BFA",
	"isle of man":                         "IMN",
	"hong kong sar china":                 "HKG",
	"guatemala":                           "GTM",
	"suriname":                            "SUR",
	"malta":                               "MLT",
	"croatia":                             "HRV",
	"uruguay":                             "URY",
	"côte d’ivoire":                       "CIV",
	"bahamas":                             "BHS",
	"bhutan":                              "BTN",
	"mozambique":                          "MOZ",
	"slovenia":                            "SVN",
	"singapore":                           "SGP",
	"libya":                               "LBY",
	"lithuania":                           "LTU",
	"moldova":                             "MDA",
	"nigeria":                             "NGA",
	"bolivia":                             "BOL",
	"guinea":                              "GIN",
	"ethiopia":                            "ETH",
	"kuwait":                              "KWT",
	"new caledonia":                       "NCL",
	"east timor":                          "TLS",
	"vietnam":                             "VNM",
	"republic of the congo":               "COG",
	"czech republic":                      "CZE",
	"trinidad & tobago":                   "TTO",
	"micronesia":                          "FSM",
	"jamaica":                             "JAM",
	"tunisia":                             "TUN",
	"madagascar":                          "MDG",
	"qatar":                               "QAT",
	"chad":                                "TCD",
	"tanzania":                            "TZA",
	"cayman islands":                      "CYM",
	"angola":                              "AGO",
	"kosovo":                              "RKS",
	"yemen":                               "YEM",
	"guam":                                "GUM",
	"rwanda":                              "RWA",
	"turkmenistan":                        "TKM",
	"new zealand":                         "NZL",
	"poland":                              "POL",
	"portugal":                            "PRT",
	"luxembourg":                          "LUX",
	"montenegro":                          "MNE",
	"malaysia":                            "MYS",
	"lebanon":                             "LBN",
	"senegal":                             "SEN",
	"mali":                                "MLI",
	"philippines":                         "PHL",
	"turkey":                              "TUR",
	"denmark":                             "DNK",
	"macedonia":                           "MKD",
	"brazil":                              "BRA",
	"barbados":                            "BRB",
	"greenland":                           "GRL",
	"malawi":                              "MWI",
	"syria":                               "SYR",
	"bangladesh":                          "BGD",
	"djibouti":                            "DJI",
	"palau":                               "PLW",
	"ukraine":                             "UKR",
	"panama":                              "PAN",
	"somaliland":                          "SOM",
	"jordan":                              "JOR",
	"haiti":                               "HTI",
	"india":                               "IND",
	"thailand":                            "THA",
	"lesotho":                             "LSO",
	"norway":                              "NOR",
	"peru":                                "PER",
	"south africa":                        "ZAF",
	"argentina":                           "ARG",
	"sri lanka":                           "LKA",
	"bulgaria":                            "BGR",
	"zambia":                              "ZMB",
	"slovakia":                            "SVK",
	"taiwan":                              "TWN",
	"nicaragua":                           "NIC",
	"south sudan":                         "SSD",
	"martinique":                          "MTQ",
	"trinidad and tobago":                 "TTO",
	"bahrain":                             "BHR",
	"burundi":                             "BDI",
	"democratic republic of the congo":    "COD",
}

// Response ...
type Response struct {
	Labels interface{} `json:"labels,omitempty"`
	Values interface{} `json:"values,omitempty"`
}

// GenericItem ...
type GenericItem struct {
	Key   interface{} `json:"key,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// ReplaceNamedParamsInQuery replaces the named parameters in the query
// with the corresponding values for logging purpose
// example:
//	query = select * from table where id = :id and category = :category
//	params = {"id": 1, "category": "c"}
//	returns select * from table where id = 1 and category = 'c'
func ReplaceNamedParamsInQuery(query string, params map[string]interface{}) string {
	// TODO: This is a naive method for replacement. A better implementation can be added as required. (-_-) zzz
	replacerGroupList := make([]string, 0, 2*len(params))
	for pName, pValue := range params {
		replacerGroupList = append(replacerGroupList, fmt.Sprintf(":%s", pName), getValueAsString(pValue))
	}
	return strings.NewReplacer(replacerGroupList...).Replace(query)
}

// ReplacePositionalParamsInQuery replaces the positional parameters in the query
// with the corresponding values for logging purpose
// example:
//	query = select * from table where id = ? and category = ?
//	params = [1, "c"]
//	returns select * from table where id = 1 and category = 'c'
func ReplacePositionalParamsInQuery(query string, params ...interface{}) string {
	// TODO: This is a naive method for replacement. A better implementation can be added as required. (-_-) zzz
	for _, param := range params {
		query = strings.Replace(query, "?", getValueAsString(param), 1)
	}
	return query
}

var spaceRemover = regexp.MustCompile("\\s\\s+")

// RemoveUnwantedSpaces replaces multiple adjacent space with single
func RemoveUnwantedSpaces(str string) string {
	return spaceRemover.ReplaceAllString(str, " ")
}

func getValueAsString(value interface{}) string {
	if value == nil {
		return "NULL"
	}
	switch typedValue := value.(type) {
	case
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", typedValue)
	case float32, float64:
		return fmt.Sprintf("%f", typedValue)
	case *string:
		return fmt.Sprintf("'%s'", *typedValue)
	default:
		return fmt.Sprintf("'%v'", typedValue)
	}
}

// FormatX converts to their actual format
func FormatX(x interface{}) interface{} {
	switch t := x.(type) {
	case time.Time:
		return t.UnixNano() / 1e6
	case int64:
		return t
	case float64:
		return t
	case string:
		return t
	case []byte:
		return string(t)
	default:
		return t
	}
}

// FormatUnixTime converts to Unix time
func FormatUnixTime(x interface{}) (int64, error) {
	var (
		t   time.Time
		err error
	)
	switch a := x.(type) {
	case time.Time:
		t = a
	case string:
		t, err = time.Parse(DATEFORMAT, a)
		if err != nil {
			t, err = time.Parse(TIMEFORMAT, a)
			if err != nil {
				t, err = time.Parse(TIMEZONEFORMAT, a)
				if err != nil {
					t, err = time.Parse(time.RFC3339, a)
				}
			}
		}
	case []byte:
		t, err = time.Parse(DATEFORMAT, string(a))
		if err != nil {
			t, err = time.Parse(TIMEFORMAT, string(a))
			if err != nil {
				t, err = time.Parse(TIMEZONEFORMAT, string(a))
				if err != nil {
					t, err = time.Parse(time.RFC3339, string(a))
				}
			}
		}
	default:
		return -1, errors.New("Invalid time")
	}

	return t.UnixNano() / 1e6, err
}

// FormatNumber returns to numeric format
func FormatNumber(y interface{}) interface{} {
	switch t := y.(type) {
	case time.Time:
		return t.UnixNano() / 1e6
	case int64:
		return t
	case float64:
		return t
	case float32:
		return t
	case string:
		return 0
	case []byte:
		byteToInt, _ := strconv.Atoi(string(t))
		return byteToInt
	default:
		return 0
	}
}

// ParseFloat converts to float64
func ParseFloat(n interface{}) float64 {
	switch t := n.(type) {
	case int64:
		return float64(t)
	case float64:
		return t
	default:
		return 0
	}
}

// ParseInt converts to int64
func ParseInt(n interface{}) int64 {
	switch t := n.(type) {
	case int64:
		return t
	case float64:
		return int64(t)
	default:
		return 0
	}
}

// FormatString converts to string
func FormatString(d interface{}) string {
	switch t := d.(type) {
	case time.Time:
		return t.String()
	case int64:
		return fmt.Sprint(t)
	case float64:
		return fmt.Sprint(t)
	case float32:
		return fmt.Sprint(t)
	case string:
		return t
	case []byte:
		return string(t)
	case bool:
		if t {
			return "true"
		}
		return "false"
	default:
		return ""
	}
}

// ACCOUNTID for the genie-epaper service.
const ACCOUNTID int64 = 4

// Request Contants to be set as keys the echo.Context
const (
	SUPERADMIN       string = "super_admin"
	ADMINID          string = "admin_id"
	ADMINNAME        string = "admin_name"
	QUERY            string = "query"
	ESINDEX          string = "elasticsearch_index"
	REQUESTID        string = "request_id"
	ADMINPERMISSIONS string = "token_object"
	DATEFORMAT       string = "2006-01-02"
	TIMEFORMAT       string = "2006-01-02 15:04:05"
	TIMEZONEFORMAT   string = "2006-01-02 15:04:05 -07:00"
)

// DebugVar Constants are to provide debuging information in echo.Context
type DebugVar string

// Available Debug Variables
const (
	// ADDQUERY adds the query in the response
	ADDQUERY         = DebugVar("add_query")
	DEBUGSTATS       = DebugVar("live_debug_stats")
	AUTOREFRESHCACHE = DebugVar("auto_refresh_cache")
	BURSTCACHE       = DebugVar("burst_cache")
	USELIVEES        = DebugVar("use_live_es")
)

// RemoveDuplicates removes the duplicate value
func RemoveDuplicates(input []string) []string {
	check := make(map[string]bool)
	list := make([]string, 0)
	for _, a := range input {
		if _, value := check[a]; !value {
			check[a] = true
			list = append(list, a)
		}
	}

	return list
}
