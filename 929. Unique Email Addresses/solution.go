//https://leetcode.com/problems/unique-email-addresses/
//字符串遍历
//Runtime: 24 ms, faster than 14.60% of Go online submissions for Unique Email Addresses.
//Memory Usage: 6.2 MB, less than 9.76% of Go online submissions for Unique Email Addresses.
//
//func numUniqueEmails(emails []string) int {
	emailExist := make(map[string]int)

	for _, email := range emails {
		var processedEmail string
		var skip bool
		for index, v  := range []rune(email) {
			str := string(v)
			if str == "@" {
				processedEmail = processedEmail + email[index:]
				break
			}
			if str == "+" {
				skip = true
				continue
			}

			if str == "." || skip  {
				continue
			}
			processedEmail = processedEmail + str


		}
		emailExist[processedEmail] = 1

	}
	return len(emailExist)
}


// 优化版
// Runtime: 8 ms, faster than 100.00% of Go online submissions for Unique Email Addresses.
// Memory Usage: 5.9 MB, less than 78.05% of Go online submissions for Unique Email Addresses.

func numUniqueEmails(emails []string) int {
	emailExist := make(map[string]int)

	for _, email := range emails {
        
        e := strings.Split(email, "@")
        local, domain := e[0], e[1]
        skipIndex := strings.Index(email, "+")
        if skipIndex > 0 {
          local = local[:skipIndex]  
        }
        
        local = strings.Replace(local, ".", "", -1)
        ret := local + "@"+ domain
                
		emailExist[ret] = 1

	}
	return len(emailExist)
}
