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

### **How to Install and Run Kasongo on Kali Linux**  

#### **Step 1: Install Dependencies**  
Before cloning and running Kasongo, install the necessary dependencies:  
```sh
sudo apt update && sudo apt install -y golang tor proxychains4
```

#### **Step 2: Clone the Repository**  
```sh
git clone https://github.com/STURUNGI777/kasongoDDoS.git
cd kasongoDDoS
```

#### **Step 3: Build the Executable**  
```sh
go build -o kasongo kasongo.go
```

#### **Step 4: Enable Tor for Anonymity**  
Start the Tor service to route traffic through the Tor network:  
```sh
sudo service tor start
```
Confirm Tor is running:  
```sh
curl --socks5-hostname 127.0.0.1:9050 https://check.torproject.org
```

#### **Step 5: Run Kasongo with Tor (DDoS Mode)**  
Use **Proxychains** to route the attack through Tor:  
```sh
proxychains4 ./kasongo -host <target.com> -port 443 -threads 500
```
**Example:**  
```sh
proxychains4 ./kasongo -host example.com -port 443 -threads 1000
```

#### **Step 6: Running Without Tor (Direct Mode)**  
If you don’t want to use Tor, simply run:  
```sh
./kasongo -host <target.com> -port 443 -threads 500
```

#### **⚠️ Warning: Use Responsibly**  
This tool is for educational and security testing **ONLY**. Unauthorized use is illegal. 🚨
