package govultr

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestBareMetalServerServiceHandler_AppInfo(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/get_app_info", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"app_info": "Minecraft Server Details\n\nYour Minecraft server is accessible at: \n\n45.74.196.185:25565\n\nYou can access the console of your Minecraft server by using the \"screen\" utility from the following login:\nUser: minecraft\nPass: NXwdsdZjwJasdZbsc\n\nRead more about this app on Vultr Docs: \n\nhttps://www.vultr.com/docs/one-click-minecraft\n"
		}
		`
		fmt.Fprint(writer, response)
	})

	appInfo, err := client.BareMetalServer.AppInfo(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.AppInfo returned error: %v", err)
	}

	expected := &AppInfo{
		AppInfo: "Minecraft Server Details\n\nYour Minecraft server is accessible at: \n\n45.74.196.185:25565\n\nYou can access the console of your Minecraft server by using the \"screen\" utility from the following login:\nUser: minecraft\nPass: NXwdsdZjwJasdZbsc\n\nRead more about this app on Vultr Docs: \n\nhttps://www.vultr.com/docs/one-click-minecraft\n",
	}

	if !reflect.DeepEqual(appInfo, expected) {
		t.Errorf("BareMetalServer.AppInfo returned %+v, expected %+v", appInfo, expected)
	}
}

func TestBareMetalServerServiceHandler_Bandwidth(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/bandwidth", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"incoming_bytes": [
				[
					"2017-04-01",
					91571055
				],
				[
					"2017-04-02",
					78355758
				],
				[
					"2017-04-03",
					85827590
				]
			],
			"outgoing_bytes": [
				[
					"2017-04-01",
					3084731
				],
				[
					"2017-04-02",
					1810478
				],
				[
					"2017-04-03",
					2729604
				]
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	bandwidth, err := client.BareMetalServer.Bandwidth(ctx, "123")

	if err != nil {
		t.Errorf("BareMetalServer.Bandwidth returned %+v", err)
	}

	expected := []map[string]string{
		{"date": "2017-04-01", "incoming": "91571055", "outgoing": "3084731"},
		{"date": "2017-04-02", "incoming": "78355758", "outgoing": "1810478"},
		{"date": "2017-04-03", "incoming": "85827590", "outgoing": "2729604"},
	}

	if !reflect.DeepEqual(bandwidth, expected) {
		t.Errorf("BareMetalServer.Bandwidth returned %+v, expected %+v", bandwidth, expected)
	}
}

func TestBareMetalServerServiceHandler_ChangeApp(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/app_change", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.ChangeApp(ctx, "900000", "15")

	if err != nil {
		t.Errorf("BareMetalServer.ChangeApp returned %+v, ", err)
	}
}

func TestBareMetalServerServiceHandler_ChangeOS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/os_change", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.ChangeOS(ctx, "900000", "302")

	if err != nil {
		t.Errorf("BareMetalServer.ChangeOS return %+v ", err)
	}
}

func TestBareMetalServerServiceHandler_Create(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/create", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"SUBID": "900000"
		}
		`
		fmt.Fprint(writer, response)
	})

	options := &BareMetalServerOptions{
		StartupScriptID: "1",
		SnapshotID:      "1",
		EnableIPV6:      "yes",
		Label:           "go-bm-test",
		SSHKeyIDs:       []string{"6b80207b1821f"},
		AppID:           "1",
		UserData:        "ZWNobyBIZWxsbyBXb3JsZAo=",
		NotifyActivate:  "yes",
		Hostname:        "test",
		Tag:             "go-test",
		ReservedIPV4:    "111.111.111.111",
	}

	bm, err := client.BareMetalServer.Create(ctx, "1", "1", "1", options)

	if err != nil {
		t.Errorf("BareMetalServer.Create returned error: %v", err)
	}

	expected := &BareMetalServer{BareMetalServerID: "900000"}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.Create returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_Delete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/destroy", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Delete(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Delete returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_EnableIPV6(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/ipv6_enable", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.EnableIPV6(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.EnableIPV6 returned %+v", err)
	}
}

func TestBareMetalServerServiceHandler_List(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"900000": {
					"SUBID": "900000",
					"os": "CentOS 6 x64",
					"ram": "65536 MB",
					"disk": "2x 240 GB SSD",
					"main_ip": "203.0.113.10",
					"cpu_count": 1,
					"location": "New Jersey",
					"DCID": "1",
					"default_password": "ab81u!ryranq",
					"date_created": "2017-04-12 18:45:41",
					"status": "active",
					"netmask_v4": "255.255.255.0",
					"gateway_v4": "203.0.113.1",
					"METALPLANID": 28,
					"v6_networks": [
						{
							"v6_network": "2001:DB8:9000::",
							"v6_main_ip": "2001:DB8:9000::100",
							"v6_network_size": 64
						}
					],
					"label": "my label",
					"tag": "my tag",
					"OSID": "127",
					"APPID": "0"
				}
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.List(ctx)

	if err != nil {
		t.Errorf("BareMetalServer.List returned error: %v", err)
	}

	expected := []BareMetalServer{
		{
			BareMetalServerID: "900000",
			Os:                "CentOS 6 x64",
			RAM:               "65536 MB",
			Disk:              "2x 240 GB SSD",
			MainIP:            "203.0.113.10",
			CPUs:              1,
			Location:          "New Jersey",
			RegionID:          1,
			DefaultPassword:   "ab81u!ryranq",
			DateCreated:       "2017-04-12 18:45:41",
			Status:            "active",
			NetmaskV4:         "255.255.255.0",
			GatewayV4:         "203.0.113.1",
			BareMetalPlanID:   28,
			V6Networks: []V6Network{
				{
					Network:     "2001:DB8:9000::",
					MainIP:      "2001:DB8:9000::100",
					NetworkSize: "64",
				},
			},
			Label: "my label",
			Tag:   "my tag",
			OsID:  "127",
			AppID: "0",
		},
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.List returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_ListByLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"900000": {
					"SUBID": "900000",
					"os": "CentOS 6 x64",
					"ram": "65536 MB",
					"disk": "2x 240 GB SSD",
					"main_ip": "203.0.113.10",
					"cpu_count": 1,
					"location": "New Jersey",
					"DCID": "1",
					"default_password": "ab81u!ryranq",
					"date_created": "2017-04-12 18:45:41",
					"status": "active",
					"netmask_v4": "255.255.255.0",
					"gateway_v4": "203.0.113.1",
					"METALPLANID": 28,
					"v6_networks": [
						{
							"v6_network": "2001:DB8:9000::",
							"v6_main_ip": "2001:DB8:9000::100",
							"v6_network_size": 64
						}
					],
					"label": "my label",
					"tag": "my tag",
					"OSID": "127",
					"APPID": "0"
				}
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.ListByLabel(ctx, "my label")

	if err != nil {
		t.Errorf("BareMetalServer.ListByLabel returned error: %v", err)
	}

	expected := []BareMetalServer{
		{
			BareMetalServerID: "900000",
			Os:                "CentOS 6 x64",
			RAM:               "65536 MB",
			Disk:              "2x 240 GB SSD",
			MainIP:            "203.0.113.10",
			CPUs:              1,
			Location:          "New Jersey",
			RegionID:          1,
			DefaultPassword:   "ab81u!ryranq",
			DateCreated:       "2017-04-12 18:45:41",
			Status:            "active",
			NetmaskV4:         "255.255.255.0",
			GatewayV4:         "203.0.113.1",
			BareMetalPlanID:   28,
			V6Networks: []V6Network{
				{
					Network:     "2001:DB8:9000::",
					MainIP:      "2001:DB8:9000::100",
					NetworkSize: "64",
				},
			},
			Label: "my label",
			Tag:   "my tag",
			OsID:  "127",
			AppID: "0",
		},
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.ListByLabel returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_ListByMainIP(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"900000": {
					"SUBID": "900000",
					"os": "CentOS 6 x64",
					"ram": "65536 MB",
					"disk": "2x 240 GB SSD",
					"main_ip": "203.0.113.10",
					"cpu_count": 1,
					"location": "New Jersey",
					"DCID": "1",
					"default_password": "ab81u!ryranq",
					"date_created": "2017-04-12 18:45:41",
					"status": "active",
					"netmask_v4": "255.255.255.0",
					"gateway_v4": "203.0.113.1",
					"METALPLANID": 28,
					"v6_networks": [
						{
							"v6_network": "2001:DB8:9000::",
							"v6_main_ip": "2001:DB8:9000::100",
							"v6_network_size": 64
						}
					],
					"label": "my label",
					"tag": "my tag",
					"OSID": "127",
					"APPID": "0"
				}
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.ListByMainIP(ctx, "203.0.113.10")

	if err != nil {
		t.Errorf("BareMetalServer.ListByMainIP returned error: %v", err)
	}

	expected := []BareMetalServer{
		{
			BareMetalServerID: "900000",
			Os:                "CentOS 6 x64",
			RAM:               "65536 MB",
			Disk:              "2x 240 GB SSD",
			MainIP:            "203.0.113.10",
			CPUs:              1,
			Location:          "New Jersey",
			RegionID:          1,
			DefaultPassword:   "ab81u!ryranq",
			DateCreated:       "2017-04-12 18:45:41",
			Status:            "active",
			NetmaskV4:         "255.255.255.0",
			GatewayV4:         "203.0.113.1",
			BareMetalPlanID:   28,
			V6Networks: []V6Network{
				{
					Network:     "2001:DB8:9000::",
					MainIP:      "2001:DB8:9000::100",
					NetworkSize: "64",
				},
			},
			Label: "my label",
			Tag:   "my tag",
			OsID:  "127",
			AppID: "0",
		},
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.ListByMainIP returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_ListByTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"900000": {
					"SUBID": "900000",
					"os": "CentOS 6 x64",
					"ram": "65536 MB",
					"disk": "2x 240 GB SSD",
					"main_ip": "203.0.113.10",
					"cpu_count": 1,
					"location": "New Jersey",
					"DCID": "1",
					"default_password": "ab81u!ryranq",
					"date_created": "2017-04-12 18:45:41",
					"status": "active",
					"netmask_v4": "255.255.255.0",
					"gateway_v4": "203.0.113.1",
					"METALPLANID": 28,
					"v6_networks": [
						{
							"v6_network": "2001:DB8:9000::",
							"v6_main_ip": "2001:DB8:9000::100",
							"v6_network_size": 64
						}
					],
					"label": "my label",
					"tag": "my tag",
					"OSID": "127",
					"APPID": "0"
				}
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.ListByTag(ctx, "my tag")

	if err != nil {
		t.Errorf("BareMetalServer.ListByTag returned error: %v", err)
	}

	expected := []BareMetalServer{
		{
			BareMetalServerID: "900000",
			Os:                "CentOS 6 x64",
			RAM:               "65536 MB",
			Disk:              "2x 240 GB SSD",
			MainIP:            "203.0.113.10",
			CPUs:              1,
			Location:          "New Jersey",
			RegionID:          1,
			DefaultPassword:   "ab81u!ryranq",
			DateCreated:       "2017-04-12 18:45:41",
			Status:            "active",
			NetmaskV4:         "255.255.255.0",
			GatewayV4:         "203.0.113.1",
			BareMetalPlanID:   28,
			V6Networks: []V6Network{
				{
					Network:     "2001:DB8:9000::",
					MainIP:      "2001:DB8:9000::100",
					NetworkSize: "64",
				},
			},
			Label: "my label",
			Tag:   "my tag",
			OsID:  "127",
			AppID: "0",
		},
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.ListByTag returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_GetServer(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
			{
				"SUBID": "900000",
				"os": "CentOS 6 x64",
				"ram": "65536 MB",
				"disk": "2x 240 GB SSD",
				"main_ip": "203.0.113.10",
				"cpu_count": 1,
				"location": "New Jersey",
				"DCID": "1",
				"default_password": "ab81u!ryranq",
				"date_created": "2017-04-12 18:45:41",
				"status": "active",
				"netmask_v4": "255.255.255.0",
				"gateway_v4": "203.0.113.1",
				"METALPLANID": 28,
				"v6_networks": [
					{
						"v6_network": "2001:DB8:9000::",
						"v6_main_ip": "2001:DB8:9000::100",
						"v6_network_size": 64
					}
				],
				"label": "my label",
				"tag": "my tag",
				"OSID": "127",
				"APPID": "0"
			}
		`
		fmt.Fprint(writer, response)
	})

	bm, err := client.BareMetalServer.GetServer(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.GetServer returned error: %v", err)
	}

	expected := &BareMetalServer{
		BareMetalServerID: "900000",
		Os:                "CentOS 6 x64",
		RAM:               "65536 MB",
		Disk:              "2x 240 GB SSD",
		MainIP:            "203.0.113.10",
		CPUs:              1,
		Location:          "New Jersey",
		RegionID:          1,
		DefaultPassword:   "ab81u!ryranq",
		DateCreated:       "2017-04-12 18:45:41",
		Status:            "active",
		NetmaskV4:         "255.255.255.0",
		GatewayV4:         "203.0.113.1",
		BareMetalPlanID:   28,
		V6Networks: []V6Network{
			{
				Network:     "2001:DB8:9000::",
				MainIP:      "2001:DB8:9000::100",
				NetworkSize: "64",
			},
		},
		Label: "my label",
		Tag:   "my tag",
		OsID:  "127",
		AppID: "0",
	}

	if !reflect.DeepEqual(bm, expected) {
		t.Errorf("BareMetalServer.GetServer returned %+v, expected %+v", bm, expected)
	}
}

func TestBareMetalServerServiceHandler_GetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/get_user_data", func(writer http.ResponseWriter, request *http.Request) {
		response := `{"userdata": "ZWNobyBIZWxsbyBXb3JsZA=="}`
		fmt.Fprint(writer, response)
	})

	userData, err := client.BareMetalServer.GetUserData(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.GetUserData return %+v ", err)
	}

	expected := &UserData{UserData: "ZWNobyBIZWxsbyBXb3JsZA=="}

	if !reflect.DeepEqual(userData, expected) {
		t.Errorf("BareMetalServer.GetUserData returned %+v, expected %+v", userData, expected)
	}
}

func TestBareMetalServerServiceHandler_Halt(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/halt", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Halt(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Halt returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_IPV4Info(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list_ipv4", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"900000": [
				{
					"ip": "203.0.113.10",
					"netmask": "255.255.255.0",
					"gateway": "203.0.113.1",
					"type": "main_ip"
				}
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	ipv4, err := client.BareMetalServer.IPV4Info(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.IPV4Info returned %+v", err)
	}

	expected := []BareMetalServerIPV4{
		{
			IP:      "203.0.113.10",
			Netmask: "255.255.255.0",
			Gateway: "203.0.113.1",
			Type:    "main_ip",
		},
	}

	if !reflect.DeepEqual(ipv4, expected) {
		t.Errorf("BareMetalServer.IPV4Info returned %+v, expected %+v", ipv4, expected)
	}
}

func TestBareMetalServerServiceHandler_IPV6Info(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/list_ipv6", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"900000": [
				{
					"ip": "2001:DB8:9000::100",
					"network": "2001:DB8:9000::",
					"network_size": 64,
					"type": "main_ip"
				}
			]
		}
		`
		fmt.Fprint(writer, response)
	})

	ipv4, err := client.BareMetalServer.IPV6Info(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.IPV6Info returned %+v", err)
	}

	expected := []BareMetalServerIPV6{
		{
			IP:          "2001:DB8:9000::100",
			Network:     "2001:DB8:9000::",
			NetworkSize: 64,
			Type:        "main_ip",
		},
	}

	if !reflect.DeepEqual(ipv4, expected) {
		t.Errorf("BareMetalServer.IPV6Info returned %+v, expected %+v", ipv4, expected)
	}
}

func TestBareMetalServerServiceHandler_ListApps(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/app_change_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"1": {
				"APPID": "1",
				"name": "LEMP",
				"short_name": "lemp",
				"deploy_name": "LEMP on CentOS 6 x64",
				"surcharge": 0
			}
		}
		`
		fmt.Fprint(writer, response)
	})

	application, err := client.BareMetalServer.ListApps(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.ListApps returned %+v, ", err)
	}

	expected := []Application{
		{
			AppID:      "1",
			Name:       "LEMP",
			ShortName:  "lemp",
			DeployName: "LEMP on CentOS 6 x64",
			Surcharge:  0,
		},
	}

	if !reflect.DeepEqual(application, expected) {
		t.Errorf("BareMetalServer.ListApps returned %+v, expected %+v", application, expected)
	}
}

func TestBareMetalServerServiceHandler_ListOS(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/os_change_list", func(writer http.ResponseWriter, request *http.Request) {
		response := `
		{
			"127": {
				"OSID": "127",
				"name": "CentOS 6 x64",
				"arch": "x64",
				"family": "centos",
				"windows": false,
				"surcharge": 0
			}
		}
		`
		fmt.Fprint(writer, response)
	})

	os, err := client.BareMetalServer.ListOS(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.ListOS return %+v ", err)
	}

	expected := []OS{
		{
			OsID:    127,
			Name:    "CentOS 6 x64",
			Arch:    "x64",
			Family:  "centos",
			Windows: false,
		},
	}

	if !reflect.DeepEqual(os, expected) {
		t.Errorf("BareMetalServer.ListOS returned %+v, expected %+v", os, expected)
	}
}

func TestBareMetalServerServiceHandler_Reboot(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/reboot", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Reboot(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Reboot returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_Reinstall(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/reinstall", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.Reinstall(ctx, "900000")

	if err != nil {
		t.Errorf("BareMetalServer.Reinstall returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_SetLabel(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/label_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.SetLabel(ctx, "900000", "new label")

	if err != nil {
		t.Errorf("BareMetalServer.SetLabel returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_SetTag(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/tag_set", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.SetTag(ctx, "900000", "new tag")

	if err != nil {
		t.Errorf("BareMetalServer.SetTag returned %+v, expected %+v", err, nil)
	}
}

func TestBareMetalServerServiceHandler_SetUserData(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/baremetal/set_user_data", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer)
	})

	err := client.BareMetalServer.SetUserData(ctx, "900000", "ZWNobyBIZWxsbyBXb3JsZA==")

	if err != nil {
		t.Errorf("Server.SetUserData return %+v ", err)
	}
}
