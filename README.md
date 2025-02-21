# kasongoDDoS
### **Kasongo - Advanced Slowloris DDoS Tool 🚀🔥**  

**⚠️ Disclaimer:** This tool is for **educational and security testing purposes only**. Unauthorized use is **illegal**. Use responsibly.  

### **🛠 Overview**  
Kasongo is a powerful **Go-based DDoS tool** inspired by Slowloris but with advanced evasion techniques. It **keeps connections open** to exhaust server resources, making it difficult to mitigate.  

### **🔹 Features**  
✅ **Advanced Socket Handling** – Efficiently manages high-volume connections.  
✅ **Tor Integration** – Routes traffic through Tor to bypass IP blocking.  
✅ **User-Agent & Header Randomization** – Evades fingerprinting defenses.  
✅ **Multi-Threading** – Sends thousands of concurrent requests.  
✅ **Target Flexibility** – Specify the target via command-line, no manual code changes.  

### **🚀 Usage**  
**Run against a target:**  
```sh
go run kasongo.go example.com 80 100 true
```  
- **example.com** → Target domain  
- **80** → Target port  
- **100** → Number of connections  
- **true** → Use Tor (set to **false** if not using Tor)  

**Start Tor before launching an attack:**  
```sh
sudo systemctl start tor
```

### **⚠️ Legal Warning**  
This tool **must only be used on systems you own or have permission to test**. Unauthorized use is **illegal and unethical**. **Use at your own risk.**  

👉 **For cybersecurity professionals and penetration testers only!** 🚀
