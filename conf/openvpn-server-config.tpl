management {{ .Management }}
dev {{ .Dev }}
proto {{ .Proto }}
port {{ .Port }}

ca {{ .Ca }}
cert {{ .Cert }}
key {{ .Key }}
dh {{ .Dh }}
#ecdh-curve prime256v1
tls-groups prime256v1

topology subnet
{{ .Server }}
ifconfig-pool-persist {{ .IfconfigPoolPersist }}
{{ .DNSServerOne }}
{{ .DNSServerTwo }}

keepalive {{ .Keepalive }}
remote-cert-tls client
tls-version-min 1.2
tls-crypt {{ .CCEncryption }}
cipher {{ .Cipher }}
auth {{ .Auth }}

persist-key
persist-tun
crl-verify crl.pem
user nobody
group nogroup
explicit-exit-notify 1

status openvpn-status.log 20
status-version 3
# syslog
log openvpn.log
verb 3
mute 10

{{ .ExtraServerOptions }}