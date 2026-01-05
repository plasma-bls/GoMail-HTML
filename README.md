# 📝 GoMail-HTML

## 📋 What is it?

A lightweight SMTP client in Go for sending dynamic HTML5 emails with external template support.

## 💡 Why should you care?

Facilitates the process of sending emails simultaneously with HTML5-Integrated content. Can be used for (Spear)Phishing

## ✨ Features

- - Sends emails simultaneously
- Highly customizable
- High performance
- Cross-Platform

## 📋 Prerequisites

- Golang Up-to-date
- Basic System Requirements

## ⚙️ Installation

Follow these simple steps to get started:

```bash
# Clone the repository
git clone https://github.com/plasma-bls/GoMail-HTML

# Navigate to the project directory
cd GoMail-HTML/
```

## 🛠️ Configuration
 
Replace the values in the .JSON config file with your SMTP host, your username and password;
after that you insert the receiver/s, -- if there are more than 1 receivers, they must be separated by a ","
otherwise, it will not work

### NOTE: The password field doesn't contain the actual account password. Instead, it uses an App-Specific Password (ASP) or a Static Token, you may create one on your SMTP Host's account. 




## 💡 Usage

```bash
# Usage Example
go run main.go
```

## 🗺️ Roadmap

- Command line integration with flags (eg. "--config /tmp/html.html")

## 🤝 Contributing

We love contributions! Here's how you can help:

1. 🍴 Fork the repository
2. 🌟 Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. 💾 Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. 📤 Push to the branch (`git push origin feature/AmazingFeature`)
5. 🔃 Open a Pull Request

Please feel free to submit issues and enhancement requests!

## 📄 License

This project is licensed under the Unlicense License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Plasma**
- GitHub: [https://github.com/plasma-bls/GoMail-HTML](https://github.com/plasma-bls/GoMail-HTML)

---

⭐ **If you found this project helpful, please give it a star!** ⭐
<p>If GoMail-HTML is useful to you, leave a star! <img src="https://i.postimg.cc/pTH5R15s/star-2b50.webp" width="30" height="30";"></p>
