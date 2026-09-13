Price Scraper

[![Version](https://img.shields.io/badge/version-1.0.0-blue)](version)
[![License](https://img.shields.io/badge/license-MIT-green)](license)

Price Scraper is a software that scrapes sites for prices and notice you when changes bigger tha 5%. you can with this app, extract price from any site to simply. 

---

## ✨ Features

- 🚀 **simply** - the software writen in simply. can read easily and edit easily
- ⚡ **quick** - it software dont use of channel, but is quick too


---

## 📦 Installation

### Prerequisites
- a go ompiler

### Step-by-step Guide

for use of this software for your needs, you need to edit 3 files
main.go
codes/myhtml/myhtml.go
codes/sites/site.go

this software send email for you if changes bigger than 5%, for send email uncomment the following part in codes/sites/site.go
```
import(
	//"Price_Scraper/codes/sendmail"
	)
...
...
//if ok == true {
//sendmail.SendMail(price, url, product)
//}
```



---
### 📝 Changelog

v1.0.0 (2026-07-06)

🎉 Initial release

Basic functionality implemented




### 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

### 📧 Contact
📧 Email: alishadan84@gmail.com

💼 LinkedIn: ali-shadan-48359152/

=======
