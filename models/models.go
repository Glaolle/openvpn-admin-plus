package models

import (
	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/bnhf/go-openvpn/server/config"
	passlib "gopkg.in/hlandau/passlib.v1"
)

var GlobalCfg Settings

func init() {
	initDB()
	createDefaultUsers()
	createDefaultSettings()
	createDefaultOVConfig()
}

func initDB() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	dbSource, err := web.AppConfig.String("dbPath")
	if err != nil {
		logs.Error(err)
	}
	dbSource = "file:" + dbSource

	err = orm.RegisterDataBase("default", "sqlite3", dbSource)
	if err != nil {
		panic(err)
	}
	orm.Debug = true
	orm.RegisterModel(
		new(User),
		new(Settings),
		new(OVConfig),
	)

	// Database alias.
	name := "default"
	// Drop table and re-create.
	force := false
	// Print log.
	verbose := true

	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		logs.Error(err)
		return
	}
}

func createDefaultUsers() {
	hash, err := passlib.Hash("b3secure")
	if err != nil {
		logs.Error("Unable to hash password", err)
	}
	user := User{
		Id:       1,
		Login:    "admin",
		Name:     "Administrator",
		Email:    "root@localhost",
		Password: hash,
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&user, "Name"); err == nil {
		if created {
			logs.Info("Default admin account created")
		} else {
			logs.Debug(user)
		}
	}

}

func createDefaultSettings() {
	s := Settings{
		Profile:       "default",
		MIAddress:     "openvpn:2080",
		MINetwork:     "tcp",
		ServerAddress: "myopenvpnserver.duckdns.org",
		OVConfigPath:  "/etc/openvpn/server/",
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&s, "Profile"); err == nil {
		GlobalCfg = s

		if created {
			logs.Info("New settings profile created")
		} else {
			logs.Debug(s)
		}
	} else {
		logs.Error(err)
	}
}

func createDefaultOVConfig() {
	c := OVConfig{
		Profile: "default",
		Config: config.Config{
			Dev:                 "tun",
			Port:                1194,
			Proto:               "udp",
			DNSServerOne:        "# push \"dhcp-option DNS 8.8.8.8\"",
			DNSServerTwo:        "# push \"dhcp-option DNS 8.8.4.4\"",
			Cipher:              "AES-256-GCM",
			Auth:                "SHA256",
			Dh:                  "none",
			Keepalive:           "10 120",
			IfconfigPoolPersist: "ipp.txt",
			Management:          "172.17.0.1 2080",
			CCEncryption:        "easy-rsa/pki/ta.key",
			Server:              "server 192.168.233.0 255.255.255.0",
			Ca:                  "easy-rsa/pki/ca.crt",
			Cert:                "easy-rsa/pki/issued/" + os.Getenv("PIVPN_SERVER") + ".crt",
			Key:                 "easy-rsa/pki/private/" + os.Getenv("PIVPN_SERVER") + ".key",
			ExtraServerOptions:  "push \"route 0.0.0.0 255.255.255.255 net_gateway\"\nclient-to-client\n# push block-outside-dns\n# push \"redirect-gateway def1\"\n# client-config-dir /etc/openvpn/ccd\n# duplicate-cn\nmax-clients 100\n# compress lz4-v2\n",
			ExtraClientOptions:  "dev tap\n# dev tun\n# lport 0\n# compress lz4-v2\n",
			PiVPNServer:         os.Getenv("PIVPN_SERVER"),
		},
	}
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(&c, "Profile"); err == nil {
		if created {
			logs.Info("New settings profile created")
		} else {
			logs.Debug(c)
		}
		path := GlobalCfg.OVConfigPath + "/" + os.Getenv("PIVPN_CONF")
		if _, err = os.Stat(path); os.IsNotExist(err) {
			destPath := GlobalCfg.OVConfigPath + "/" + os.Getenv("PIVPN_CONF")
			if err = config.SaveToFile("conf/openvpn-server-config.tpl",
				c.Config, destPath); err != nil {
				logs.Error(err)
			}
		}
	} else {
		logs.Error(err)
	}
}
