// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netutils

import "testing"

func TestNewIPV6Addr(t *testing.T) {
	cases := []struct {
		in            string
		wantAddr      IPV6Addr
		want          string
		stepUp        string
		stepDown      string
		preflen       uint8
		netAddr       string
		broadcastAddr string
	}{
		{
			in: "::1",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			want:     "::1",
			stepUp:   "::2",
			stepDown: "::",

			preflen:       126,
			netAddr:       "::",
			broadcastAddr: "::3",
		},
		{
			in: "3ffe::1",
			wantAddr: IPV6Addr{
				0x3ffe, 0, 0, 0, 0, 0, 0, 1,
			},
			want:     "3ffe::1",
			stepUp:   "3ffe::2",
			stepDown: "3ffe::",

			preflen:       126,
			netAddr:       "3ffe::",
			broadcastAddr: "3ffe::3",
		},
		{
			in: "3ffe:0:0:0:0:0:0:1",
			wantAddr: IPV6Addr{
				0x3ffe, 0, 0, 0, 0, 0, 0, 1,
			},
			want:     "3ffe::1",
			stepUp:   "3ffe::2",
			stepDown: "3ffe::",

			preflen:       126,
			netAddr:       "3ffe::",
			broadcastAddr: "3ffe::3",
		},
		{
			in: "3FFe:0:0:0:0:0:0:1",
			wantAddr: IPV6Addr{
				0x3ffe, 0, 0, 0, 0, 0, 0, 1,
			},
			want:     "3ffe::1",
			stepUp:   "3ffe::2",
			stepDown: "3ffe::",

			preflen:       124,
			netAddr:       "3ffe::",
			broadcastAddr: "3ffe::f",
		},
		{
			in: "::",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			want:     "::",
			stepUp:   "::1",
			stepDown: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",

			preflen:       64,
			netAddr:       "::",
			broadcastAddr: "::ffff:ffff:ffff:ffff",
		},
		{
			in: "::",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			want:     "::",
			stepUp:   "::1",
			stepDown: "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",

			preflen:       68,
			netAddr:       "::",
			broadcastAddr: "::fff:ffff:ffff:ffff",
		},
		{
			in: "2001 : db8: 3333 : 4444 : 5555 : 6666 : 7777 : 8888",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0x3333, 0x4444, 0x5555, 0x6666, 0x7777, 0x8888,
			},
			want:     "2001:db8:3333:4444:5555:6666:7777:8888",
			stepUp:   "2001:db8:3333:4444:5555:6666:7777:8889",
			stepDown: "2001:db8:3333:4444:5555:6666:7777:8887",

			preflen:       125,
			netAddr:       "2001:db8:3333:4444:5555:6666:7777:8888",
			broadcastAddr: "2001:db8:3333:4444:5555:6666:7777:888f",
		},
		{
			in: "2001 : db8 : 3333 : 4444 : CCCC : DDDD : EEEE : FFFF",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0x3333, 0x4444, 0xcccc, 0xdddd, 0xeeee, 0xffff,
			},
			want:     "2001:db8:3333:4444:cccc:dddd:eeee:ffff",
			stepUp:   "2001:db8:3333:4444:cccc:dddd:eeef::",
			stepDown: "2001:db8:3333:4444:cccc:dddd:eeee:fffe",

			preflen:       96,
			netAddr:       "2001:db8:3333:4444:cccc:dddd::",
			broadcastAddr: "2001:db8:3333:4444:cccc:dddd:ffff:ffff",
		},
		{
			in: "2001: db8: :",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0, 0, 0, 0, 0, 0,
			},
			want:     "2001:db8::",
			stepUp:   "2001:db8::1",
			stepDown: "2001:db7:ffff:ffff:ffff:ffff:ffff:ffff",

			preflen:       16,
			netAddr:       "2001::",
			broadcastAddr: "2001:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
		},
		{
			in: ": : 1234 : 5678",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0, 0, 0x1234, 0x5678,
			},
			want:     "::1234:5678",
			stepUp:   "::1234:5679",
			stepDown: "::1234:5677",

			preflen:       112,
			netAddr:       "::1234:0",
			broadcastAddr: "::1234:ffff",
		},
		{
			in: "2001 : db8: : 1234 : 5678",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0, 0, 0, 0, 0x1234, 0x5678,
			},
			want:     "2001:db8::1234:5678",
			stepUp:   "2001:db8::1234:5679",
			stepDown: "2001:db8::1234:5677",

			preflen:       128,
			netAddr:       "2001:db8::1234:5678",
			broadcastAddr: "2001:db8::1234:5678",
		},
		{
			in: "2001:0db8:0001:0000:0000:0ab9:C0A8:0102",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0x1, 0, 0, 0xab9, 0xc0a8, 0x0102,
			},
			want:     "2001:db8:1::ab9:c0a8:102",
			stepUp:   "2001:db8:1::ab9:c0a8:103",
			stepDown: "2001:db8:1::ab9:c0a8:101",

			preflen:       124,
			netAddr:       "2001:db8:1::ab9:c0a8:100",
			broadcastAddr: "2001:db8:1::ab9:c0a8:10f",
		},
		{
			in: "2001 : db8: 3333 : 4444 : 5555 : 6666 : 1 . 2 . 3 . 4",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0x3333, 0x4444, 0x5555, 0x6666, 0x102, 0x304,
			},
			want:     "2001:db8:3333:4444:5555:6666:102:304",
			stepUp:   "2001:db8:3333:4444:5555:6666:102:305",
			stepDown: "2001:db8:3333:4444:5555:6666:102:303",

			preflen:       120,
			netAddr:       "2001:db8:3333:4444:5555:6666:102:300",
			broadcastAddr: "2001:db8:3333:4444:5555:6666:102:3ff",
		},
		{
			in: ": : 11 . 22 . 33 . 44",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0, 0, 0xb16, 0x212c,
			},
			want:     "::b16:212c",
			stepUp:   "::b16:212d",
			stepDown: "::b16:212b",

			preflen:       124,
			netAddr:       "::b16:2120",
			broadcastAddr: "::b16:212f",
		},
		{
			in: "2001 : db8: : 123 . 123 . 123 . 123",
			wantAddr: IPV6Addr{
				0x2001, 0xdb8, 0, 0, 0, 0, 0x7b7b, 0x7b7b,
			},
			want:     "2001:db8::7b7b:7b7b",
			stepUp:   "2001:db8::7b7b:7b7c",
			stepDown: "2001:db8::7b7b:7b7a",

			preflen:       96,
			netAddr:       "2001:db8::",
			broadcastAddr: "2001:db8::ffff:ffff",
		},
		{
			in: ": : 1234 : 5678 : 91 . 123 . 4 . 56",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0x1234, 0x5678, 0x5b7b, 0x0438,
			},
			want:     "::1234:5678:5b7b:438",
			stepUp:   "::1234:5678:5b7b:439",
			stepDown: "::1234:5678:5b7b:437",

			preflen:       64,
			netAddr:       "::",
			broadcastAddr: "::ffff:ffff:ffff:ffff",
		},
		{
			in: ": : 1234 : 5678 : 1 . 2 . 3 . 4",
			wantAddr: IPV6Addr{
				0, 0, 0, 0, 0x1234, 0x5678, 0x102, 0x304,
			},
			want:     "::1234:5678:102:304",
			stepUp:   "::1234:5678:102:305",
			stepDown: "::1234:5678:102:303",

			preflen:       72,
			netAddr:       "::1200:0:0:0",
			broadcastAddr: "::12ff:ffff:ffff:ffff",
		},
	}
	for _, c := range cases {
		addr6, err := NewIPV6Addr(c.in)
		if err != nil {
			t.Errorf("NewIPV6Addr %s fail %s", c.in, err)
		} else if !addr6.Equals(c.wantAddr) {
			t.Errorf("in %s want %s got %s", c.in, c.wantAddr.String(), addr6.String())
		} else if addr6.String() != c.want {
			t.Errorf("in %s got %s want %s", c.in, addr6.String(), c.want)
		} else {
			up := addr6.StepUp()
			if up.String() != c.stepUp {
				t.Errorf("in %s stepup got %s want %s", addr6.String(), up.String(), c.stepUp)
			} else if addr6.String() != up.StepDown().String() {
				t.Errorf("%s != %s", addr6.String(), up.StepDown().String())
			} else {
				down := addr6.StepDown()
				if down.String() != c.stepDown {
					t.Errorf("in %s stepDown got %s want %s", addr6.String(), down.String(), c.stepDown)
				} else if addr6.String() != down.StepUp().String() {
					t.Errorf("%s != %s", addr6.String(), down.StepUp().String())
				} else {
					netAddr := addr6.NetAddr(c.preflen)
					if netAddr.String() != c.netAddr {
						t.Errorf("%s preflen %d netaddr %s want %s", addr6.String(), c.preflen, netAddr.String(), c.netAddr)
					} else {
						baddr := addr6.BroadcastAddr(c.preflen)
						if baddr.String() != c.broadcastAddr {
							t.Errorf("%s preflen %d broadcastAddr %s want %s", addr6.String(), c.preflen, baddr.String(), c.broadcastAddr)
						} else {
							haddr := addr6.HostAddr(baddr, c.preflen)
							if haddr.String() != baddr.String() {
								t.Errorf("%s preflen host %s get %s want %s", addr6.String(), baddr.String(), haddr.String(), baddr.String())
							}
						}
					}
				}
			}
		}
	}
}

func TestCompareIPV6(t *testing.T) {
	cases := []struct {
		addr1      string
		addr2      string
		wantLe     bool
		wantLt     bool
		wantGe     bool
		wantGt     bool
		wantEquals bool
	}{
		{
			addr1:      "::1",
			addr2:      "::2",
			wantLe:     true,
			wantLt:     true,
			wantGe:     false,
			wantGt:     false,
			wantEquals: false,
		},
		{
			addr1:      "3ffe::1",
			addr2:      "::2",
			wantLe:     false,
			wantLt:     false,
			wantGe:     true,
			wantGt:     true,
			wantEquals: false,
		},
		{
			addr1:      "3ffe::1",
			addr2:      "3ffe::1",
			wantLe:     true,
			wantLt:     false,
			wantGe:     true,
			wantGt:     false,
			wantEquals: true,
		},
	}

	for _, c := range cases {
		addr1, err := NewIPV6Addr(c.addr1)
		if err != nil {
			t.Fatalf("NewIPV6Addr %s fail %s", c.addr1, err)
		}
		addr2, err := NewIPV6Addr(c.addr2)
		if err != nil {
			t.Fatalf("NewIPV6Addr %s fail %s", c.addr2, err)
		}
		gotLe := addr1.Le(addr2)
		gotLt := addr1.Lt(addr2)
		gotGe := addr1.Ge(addr2)
		gotGt := addr1.Gt(addr2)
		gotEquals := addr1.Equals(addr2)
		if gotLe != c.wantLe {
			t.Errorf("%s %s <= %s %s got %v want %v", c.addr1, addr1.String(), c.addr2, addr2.String(), gotLe, c.wantLe)
		}
		if gotLt != c.wantLt {
			t.Errorf("%s %s  < %s %s got %v want %v", c.addr1, addr1.String(), c.addr2, addr2.String(), gotLt, c.wantLt)
		}
		if gotGe != c.wantGe {
			t.Errorf("%s %s >= %s %s got %v want %v", c.addr1, addr1.String(), c.addr2, addr2.String(), gotGe, c.wantGe)
		}
		if gotGt != c.wantGt {
			t.Errorf("%s %s > %s %s got %v want %v", c.addr1, addr1.String(), c.addr2, addr2.String(), gotGt, c.wantGt)
		}
		if gotEquals != c.wantEquals {
			t.Errorf("%s == %s got %v want %v", c.addr1, c.addr2, gotEquals, c.wantEquals)
		}
	}
}

func TestRandomAddress(t *testing.T) {
	cases := []struct {
		addr1 string
		addr2 string
	}{
		{
			addr1: "3ffe::1",
			addr2: "3ffe::1",
		},
		{
			addr1: "3ffe::1",
			addr2: "3ffe::ffff",
		},
		{
			addr1: "::1",
			addr2: "::ffff",
		},
		{
			addr1: "::1234:1",
			addr2: "::1234:ffff",
		},
		{
			addr1: "::1234:1",
			addr2: "::1238:ffff",
		},
	}
	for _, c := range cases {
		addr1, err := NewIPV6Addr(c.addr1)
		if err != nil {
			t.Fatalf("NewIPV6Addr %s fail %s", c.addr1, err)
		}
		addr2, err := NewIPV6Addr(c.addr2)
		if err != nil {
			t.Fatalf("NewIPV6Addr %s fail %s", c.addr2, err)
		}
		ipRange := NewIPV6AddrRange(addr1, addr2)
		for i := 0; i < 100; i++ {
			randomAddr := ipRange.Random()
			if randomAddr.Lt(ipRange.StartIp()) || randomAddr.Gt(ipRange.EndIp()) {
				t.Errorf("random %s out of range %s", randomAddr, ipRange.String())
			}
		}
	}
}

func TestDeriveIPv6Addr(t *testing.T) {
	cases := []struct {
		ipAddr   string
		macAddr  string
		startIp6 string
		endIp6   string
		maskLen6 uint8
		want     string
	}{
		{
			ipAddr:   "192.168.222.171",
			macAddr:  "00:24:b4:6d:a8:56",
			startIp6: "fd:3ffe:3200:1222::2",
			endIp6:   "fd:3ffe:3200:1222::fffe",
			maskLen6: 64,
			want:     "fd:3ffe:3200:1222::a856",
		},
		{
			ipAddr:   "192.168.222.171",
			macAddr:  "00:24:b4:6d:a8:56",
			startIp6: "fd:3ffe:3200:1222::2",
			endIp6:   "fd:3ffe:3200:1222:ffff:ffff:ffff:fffe",
			maskLen6: 64,
			want:     "fd:3ffe:3200:1222:c0a8:deab:b46d:a856",
		},
	}

	for _, c := range cases {
		got := DeriveIPv6AddrFromIPv4AddrMac(c.ipAddr, c.macAddr, c.startIp6, c.endIp6, c.maskLen6)
		if got != c.want {
			t.Errorf("DeriveIPv6AddrFromIPv4AddrMac %s %s %s %s %d want %s got %s", c.ipAddr, c.macAddr, c.startIp6, c.endIp6, c.maskLen6, c.want, got)
		}
	}
}
