package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/tebeka/selenium"
)

var formURL = "https://docs.google.com/forms/d/e/1FAIpQLSfxBSibny-UusYDsnQqHuST6u2KseMjwY8YuFP6djJG3xv_9A/viewform?usp=pp_url"

var xpaths = []string{
	"/html/body/div/div[2]/form/div[2]/div/div[3]/div[1]/div[1]/div/span/span",
	"/html/body/div/div[2]/form/div[2]/div/div[3]/div[1]/div[1]/div[2]/span/span",
	"/html/body/div/div[2]/form/div[2]/div/div[3]/div[1]/div[1]/div[2]/span/span",
	"/html/body/div/div[2]/form/div[2]/div/div[3]/div[1]/div[1]/div[2]/span/span",
	"/html/body/div/div[2]/form/div[2]/div/div[3]/div[1]/div[1]/div[2]/span/span",
}

const (
	seleniumPath = "./chromedriver"
	port         = 4444
)

func runPrefill(wd selenium.WebDriver, prefillURL string) {
	if err := wd.Get(prefillURL); err != nil {
		log.Printf("Error %s: %v", prefillURL, err)
		return
	}
	time.Sleep(2 * time.Second)

	for _, xp := range xpaths {
		elem, err := wd.FindElement(selenium.ByXPATH, xp)
		if err != nil {
			log.Printf("Error %s: %v", xp, err)
			continue
		}
		if err := elem.Click(); err != nil {
			log.Printf("Error %s: %v", xp, err)
		} else {
			log.Printf("Success %s", xp)
		}
		time.Sleep(1 * time.Second)
	}
}

func getRandomValue(options []string) string {
	return options[rand.Intn(len(options))]
}

func getRandom1to5() string {
	return strconv.Itoa(rand.Intn(5) + 1)
}

func getRandomMultipleSelections(options []string) []string {
	var selections []string
	for _, option := range options {
		if rand.Intn(2) == 0 {
			selections = append(selections, option)
		}
	}
	if len(selections) == 0 && len(options) > 0 {
		selections = append(selections, getRandomValue(options))
	}
	return selections
}

func getJobByAgeRange(ageRange string) string {
	switch ageRange {
	case "Dưới 18":
		values := []string{"Học sinh/Sinh viên", "Khác"}
		return values[rand.Intn(len(values))]
	case "18-24":
		values := []string{"Học sinh/Sinh viên", "Nhân viên văn phòng", "Khác"}
		return values[rand.Intn(len(values))]
	case "25-30":
		values := []string{"Nhân viên văn phòng", "Công chức/Viên chức", "Khác"}
		return values[rand.Intn(len(values))]
	case "Trên 30":
		values := []string{"Nhân viên văn phòng", "Công chức/Viên chức", "Khác"}
		return values[rand.Intn(len(values))]
	default:
		return "Khác"
	}
}

func getEducationByAgeRangeAndJob(ageRange string, job string) string {
	switch ageRange {
	case "Dưới 18":
		return "Trung học phổ thông"
	case "18-24":
		if job == "Học sinh/Sinh viên" {
			return "Trung học phổ thông"
		}
		return "Cao đẳng/Đại học"
	case "25-30":
		if job == "Nhân viên văn phòng" {
			return "Cao đẳng/Đại học"
		}
		return "Sau đại học"
	case "Trên 30":
		if job == "Nhân viên văn phòng" {
			return "Cao đẳng/Đại học"
		}
		return "Sau đại học"
	default:
		return "Khác"
	}
}

func getIncomeByAgeRangeAndJob(ageRange string, job string) string {
	switch ageRange {
	case "Dưới 18":
		return "Dưới 5 triệu VNĐ"
	case "18-24":
		if job == "Học sinh/Sinh viên" {
			return "Dưới 5 triệu VNĐ"
		}
		return "5-10 triệu VNĐ"
	case "25-30":
		if job == "Nhân viên văn phòng" {
			return "10-20 triệu VNĐ"
		}
		return "Trên 20 triệu VNĐ"
	case "Trên 30":
		if job == "Nhân viên văn phòng" {
			return "10-20 triệu VNĐ"
		}
		return "Trên 20 triệu VNĐ"
	default:
		return "Dưới 5 triệu VNĐ"
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	service, err := selenium.NewChromeDriverService(seleniumPath, port)
	fmt.Println("Service:", service)
	if err != nil {
		log.Fatalf("Error when run ChromeDriver: %v", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatalf("Error WebDriver connection: %v", err)
	}
	defer wd.Quit()

	ageRange := []string{"Dưới 18", "18-24", "25-30", "Trên 30"}
	gender := []string{"Nam", "Nữ", "Không muốn tiết lộ"}
	// job := []string{"Học sinh/Sinh viên", "Công chức/Viên chức", "Cao đẳng/Đại học", "Khác"}
	// education := []string{"Trung học phổ thông", "Cao đẳng/Đại học", "Sau đại học", "Khác"}
	// income := []string{"Dưới 5 triệu VNĐ", "5 triệu - 10 triệu VNĐ", "10-20 triệu VNĐ", "Trên 20 triệu VNĐ"}
	spending := []string{"0% - 10%", "10% - 25", "25% - 50%", "Trên 50%"}
	frequency := []string{"Không thường xuyên", "Thường xuyên"}
	method := []string{"Tại cửa hàng", "Trực tuyến (ví dụ như sàn thương mại điện tử,website...)", "Omni-channel (xem sản phẩm trước trên web rồi đến cửa hàng mua sản phẩm)"}
	answers := []string{
		"An toàn hơn so với mật khẩu truyền thống",
		"Nhanh chóng và tiện lợi khi thực hiện giao dịch",
		"Không cần nhớ mật khẩu, giảm nguy cơ bị lộ thông tin",
		"Dễ sử dụng, chỉ cần vân tay hoặc khuôn mặt",
		"QR code giúp thanh toán nhanh, không cần nhập số tài khoản",
		"Bảo mật cao hơn vì khó giả mạo",
		"Giảm rủi ro bị hack tài khoản",
		"Tiện lợi khi sử dụng trên thiết bị di động",
		"Không cần mang theo nhiều thẻ hoặc giấy tờ",
		"Phù hợp với xu hướng thanh toán không dùng tiền mặt",
		"Hạn chế sai sót khi nhập thông tin giao dịch",
		"Giúp giao dịch nhanh chóng, đặc biệt trong tình huống khẩn cấp",
		"Bảo vệ thông tin cá nhân tốt hơn so với nhập mã PIN",
		"Dễ dàng tích hợp với các hệ thống ngân hàng hiện đại",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Mang lại trải nghiệm người dùng tốt hơn",
		"Có thể sử dụng trên nhiều nền tảng khác nhau",
		"Tích hợp với nhiều ứng dụng ngân hàng và ví điện tử",
		"Hỗ trợ xác thực hai lớp an toàn",
		"Giúp tránh bị đánh cắp thông tin đăng nhập",
		"Thân thiện với người cao tuổi và trẻ em",
		"Công nghệ tiên tiến giúp nâng cao bảo mật",
		"QR code giảm rủi ro lộ thông tin thẻ tín dụng",
		"Tích hợp dễ dàng với các thiết bị di động",
		"Tự động xác thực nhanh mà không cần nhập tay",
		"Không cần mang theo thẻ ATM hay CMND",
		"Dễ dàng quản lý và theo dõi giao dịch",
		"Không bị ảnh hưởng bởi lừa đảo qua điện thoại",
		"Có thể dùng khi mất mạng hoặc kết nối kém",
		"Mở khóa giao dịch chỉ bằng dấu vân tay",
		"Giảm nguy cơ bị giả mạo danh tính",
		"Tăng tính bảo mật khi giao dịch trực tuyến",
		"Không cần nhập thông tin cá nhân mỗi lần giao dịch",
		"Không bị ảnh hưởng bởi tấn công phishing",
		"Tích hợp với hệ thống ngân hàng số",
		"Bảo mật cao hơn so với SMS OTP",
		"Không bị gián đoạn do mất sóng điện thoại",
		"Tiện lợi khi thanh toán tại quầy",
		"Dễ dàng sử dụng cho mọi đối tượng",
		"Tích hợp với nhiều ứng dụng ngân hàng và ví điện tử",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Hỗ trợ xác thực hai lớp an toàn",
		"Nhanh chóng và tiện lợi khi thực hiện giao dịch",
		"Giúp giao dịch nhanh chóng, đặc biệt trong tình huống khẩn cấp",
		"Có thể dùng khi mất mạng hoặc kết nối kém",
		"Không cần nhớ mật khẩu, giảm nguy cơ bị lộ thông tin",
		"Dễ dàng quản lý và theo dõi giao dịch",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Công nghệ tiên tiến giúp nâng cao bảo mật",
		"Dễ dàng sử dụng cho mọi đối tượng",
		"Có thể dùng khi mất mạng hoặc kết nối kém",
		"An toàn hơn so với mật khẩu truyền thống",
		"Tiện lợi khi sử dụng trên thiết bị di động",
		"Tích hợp với nhiều ứng dụng ngân hàng và ví điện tử",
		"Không cần nhớ mật khẩu, giảm nguy cơ bị lộ thông tin",
		"Bảo vệ thông tin cá nhân tốt hơn so với nhập mã PIN",
		"Tích hợp dễ dàng với các thiết bị di động",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Tiện lợi khi thanh toán tại quầy",
		"Không bị gián đoạn do mất sóng điện thoại",
		"Có thể dùng khi mất mạng hoặc kết nối kém",
		"Có thể sử dụng trên nhiều nền tảng khác nhau",
		"Không cần mang theo nhiều thẻ hoặc giấy tờ",
		"Giảm nguy cơ bị giả mạo danh tính",
		"Dễ dàng sử dụng cho mọi đối tượng",
		"QR code giảm rủi ro lộ thông tin thẻ tín dụng",
		"Công nghệ tiên tiến giúp nâng cao bảo mật",
		"Tự động xác thực nhanh mà không cần nhập tay",
		"Dễ dàng quản lý và theo dõi giao dịch",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Giúp tránh bị đánh cắp thông tin đăng nhập",
		"Không cần nhập thông tin cá nhân mỗi lần giao dịch",
		"Bảo vệ thông tin cá nhân tốt hơn so với nhập mã PIN",
		"Tiện lợi khi thanh toán tại quầy",
		"Giúp giao dịch nhanh chóng, đặc biệt trong tình huống khẩn cấp",
		"Có thể sử dụng trên nhiều nền tảng khác nhau",
		"Tích hợp với hệ thống ngân hàng số",
		"Không bị ảnh hưởng bởi tấn công phishing",
		"Tích hợp với hệ thống ngân hàng số",
		"Hạn chế sai sót khi nhập thông tin giao dịch",
		"Dễ dàng tích hợp với các hệ thống ngân hàng hiện đại",
		"Tự động xác thực nhanh mà không cần nhập tay",
		"Giảm rủi ro bị hack tài khoản",
		"Giảm nguy cơ bị giả mạo danh tính",
		"Tích hợp với nhiều ứng dụng ngân hàng và ví điện tử",
		"Không lo bị mất thẻ hoặc quên mã PIN",
		"Giúp tránh bị đánh cắp thông tin đăng nhập",
		"Tích hợp với nhiều ứng dụng ngân hàng và ví điện tử",
		"Không cần nhớ mật khẩu, giảm nguy cơ bị lộ thông tin",
		"QR code giúp thanh toán nhanh, không cần nhập số tài khoản",
		"Tăng tính bảo mật khi giao dịch trực tuyến",
		"Dễ dàng quản lý và theo dõi giao dịch",
		"Tích hợp với hệ thống ngân hàng số",
		"Không bị ảnh hưởng bởi lừa đảo qua điện thoại",
		"Không bị ảnh hưởng bởi lừa đảo qua điện thoại",
		"Tăng tính bảo mật khi giao dịch trực tuyến",
		"Không cần nhập thông tin cá nhân mỗi lần giao dịch",
		"Không bị gián đoạn do mất sóng điện thoại",
		"Bảo mật cao hơn vì khó giả mạo",
		"Giảm rủi ro bị hack tài khoản",
		"Giảm nguy cơ bị giả mạo danh tính",
		"Giúp tránh bị đánh cắp thông tin đăng nhập",
		"Không cần nhớ mật khẩu, giảm nguy cơ bị lộ thông tin",
		"Không cần mang theo nhiều thẻ hoặc giấy tờ",
		"Tích hợp dễ dàng với các thiết bị di động",
	}

	for {
		randomAgeRange := getRandomValue(ageRange)
		randomJob := getJobByAgeRange(randomAgeRange)
		randomEducation := getEducationByAgeRangeAndJob(randomAgeRange, randomJob)
		randomIncome := getIncomeByAgeRangeAndJob(randomAgeRange, randomJob)
		formData := url.Values{
			"entry.1389336972": {randomAgeRange},
			"entry.448103239":  {randomEducation},
			"entry.1733332143": {randomJob},
			"entry.255775007":  {randomIncome},
			"entry.263343525":  {getRandomValue(frequency)},
			"entry.490324943":  {getRandomValue(spending)},
			"entry.342160537":  getRandomMultipleSelections(method),
			"entry.1114995741": {getRandomValue(answers)},
			"entry.1505334718": {getRandomValue(gender)},
			"entry.1794013419": {"Có"},
			"entry.549884811":  {getRandom1to5()},
			"entry.1495300220": {getRandom1to5()},
			"entry.2081391900": {getRandom1to5()},
			"entry.373757695":  {getRandom1to5()},
			"entry.1191883904": {getRandom1to5()},
			"entry.680697604":  {getRandom1to5()},
			"entry.116330850":  {getRandom1to5()},
			"entry.2080991417": {getRandom1to5()},
			"entry.1007662968": {getRandom1to5()},
			"entry.1817113763": {getRandom1to5()},
			"entry.1982148623": {getRandom1to5()},
			"entry.539759207":  {getRandom1to5()},
			"entry.1225023986": {getRandom1to5()},
			"entry.1503706948": {getRandom1to5()},
			"entry.1883672035": {getRandom1to5()},
			"entry.390240386":  {getRandom1to5()},
			"entry.789970352":  {getRandom1to5()},
			"entry.1312460030": {getRandom1to5()},
			"entry.216317031":  {getRandom1to5()},
			"entry.1187451875": {getRandom1to5()},
			"entry.1508926085": {getRandom1to5()},
			"entry.9022419":    {getRandom1to5()},
			"entry.229595952":  {getRandom1to5()},
			"entry.539927733":  {getRandom1to5()},
			"entry.360968025":  {getRandom1to5()},
			"entry.1743344943": {getRandom1to5()},
			"entry.122503798":  {getRandom1to5()},
			"entry.1098420754": {getRandom1to5()},
			"entry.1499009540": {getRandom1to5()},
			"entry.62998863":   {getRandom1to5()},
			"entry.470937265":  {getRandom1to5()},
			"entry.1984752829": {getRandom1to5()},
			"entry.281381110":  {getRandom1to5()},
			"entry.1945761227": {getRandom1to5()},
			"entry.409109335":  {getRandom1to5()},
			"entry.1984981826": {getRandom1to5()},
			"entry.1248068872": {getRandom1to5()},
			"entry.1496695166": {getRandom1to5()},
			"entry.772224077":  {getRandom1to5()},
			"entry.1890720584": {getRandom1to5()},
			"entry.497021995":  {getRandom1to5()},
			"entry.1882498434": {getRandom1to5()},
			"entry.112491107":  {"Anh Thư"},
			// "fvv":                 {"1"},
			// "dlut":                {"1"},
			// "submissionTimestamp": {strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)},
		}
		prefillURL := formURL + "&" + formData.Encode()
		fmt.Println("Prefill URL:", prefillURL)
		runPrefill(wd, prefillURL)
		sleepMinutes := rand.Intn(4) + 2 // generates a random number in [2, 5]
		fmt.Printf("Sleeping for %d minutes...\n", sleepMinutes)
		time.Sleep(time.Duration(sleepMinutes) * time.Minute)
	}

}
