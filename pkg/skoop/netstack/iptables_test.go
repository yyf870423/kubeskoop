package netstack

import (
	"fmt"
	"testing"
)

func TestParseIPTables(t *testing.T) {
	ipt := ParseIPTables(iptablesDump)
	fmt.Println(ipt)
}

var iptablesDump = `<iptables-rules version="1.0">
<!-- # Generated by iptables*-save v1.4.21 on Tue Sep 20 10:54:54 2022 -->
  <table name="raw" >
    <chain name="PREROUTING" policy="ACCEPT" packet-count="137944872" byte-count="66096647455" >
      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >udp</p>
        </match>
        <udp >
          <dport >53</dport>
        </udp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >tcp</p>
        </match>
        <tcp >
          <dport >53</dport>
        </tcp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

    </chain>
    <chain name="OUTPUT" policy="ACCEPT" packet-count="254128030" byte-count="19878216315" >
      <rule >
       <conditions>
        <match >
          <s >169.254.20.10/32</s>
          <p >tcp</p>
        </match>
        <tcp >
          <sport >8080</sport>
        </tcp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >tcp</p>
        </match>
        <tcp >
          <dport >8080</dport>
        </tcp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >udp</p>
        </match>
        <udp >
          <dport >53</dport>
        </udp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >tcp</p>
        </match>
        <tcp >
          <dport >53</dport>
        </tcp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <s >169.254.20.10/32</s>
          <p >udp</p>
        </match>
        <udp >
          <sport >53</sport>
        </udp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <s >169.254.20.10/32</s>
          <p >tcp</p>
        </match>
        <tcp >
          <sport >53</sport>
        </tcp>
       </conditions>
       <actions>
        <NOTRACK  />
       </actions>

      </rule>

    </chain>
  </table>
<!-- # Completed on Tue Sep 20 10:54:54 2022 -->
<!-- # Generated by iptables*-save v1.4.21 on Tue Sep 20 10:54:54 2022 -->
  <table name="mangle" >
    <chain name="PREROUTING" policy="ACCEPT" packet-count="138117562" byte-count="68007964179" />
    <chain name="INPUT" policy="ACCEPT" packet-count="136335555" byte-count="62666290673" />
    <chain name="FORWARD" policy="ACCEPT" packet-count="1781476" byte-count="5341388171" />
    <chain name="OUTPUT" policy="ACCEPT" packet-count="254247535" byte-count="19884679027" />
    <chain name="POSTROUTING" policy="ACCEPT" packet-count="133009687" byte-count="16428648145" />
    <chain name="KUBE-IPTABLES-HINT" packet-count="0" byte-count="0" />
    <chain name="KUBE-KUBELET-CANARY" packet-count="0" byte-count="0" />
  </table>
<!-- # Completed on Tue Sep 20 10:54:54 2022 -->
<!-- # Generated by iptables*-save v1.4.21 on Tue Sep 20 10:54:54 2022 -->
  <table name="filter" >
    <chain name="INPUT" policy="ACCEPT" packet-count="82" byte-count="6139" >
      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >udp</p>
        </match>
        <udp >
          <dport >53</dport>
        </udp>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <d >169.254.20.10/32</d>
          <p >tcp</p>
        </match>
        <tcp >
          <dport >53</dport>
        </tcp>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes health check rules&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <call >
          <KUBE-NODE-PORT />
        </call>
       </actions>

      </rule>

      <rule >
       <actions>
        <call >
          <KUBE-FIREWALL />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="FORWARD" policy="ACCEPT" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes forwarding rules&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <call >
          <KUBE-FORWARD />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="OUTPUT" policy="ACCEPT" packet-count="76" byte-count="6002" >
      <rule >
       <conditions>
        <match >
          <s >169.254.20.10/32</s>
          <p >udp</p>
        </match>
        <udp >
          <sport >53</sport>
        </udp>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <s >169.254.20.10/32</s>
          <p >tcp</p>
        </match>
        <tcp >
          <sport >53</sport>
        </tcp>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <actions>
        <call >
          <KUBE-FIREWALL />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-FIREWALL" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes firewall for dropping marked packets&quot;</comment>
        </comment>
        <mark >
          <mark >0x8000/0x8000</mark>
        </mark>
       </conditions>
       <actions>
        <DROP  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <s  invert="1">127.0.0.0/8</s>
          <d >127.0.0.0/8</d>
        </match>
        <comment >
          <comment >&quot;block incoming localnet connections&quot;</comment>
        </comment>
        <conntrack >
          <ctstate  invert="1">RELATED,ESTABLISHED,DNAT</ctstate>
        </conntrack>
       </conditions>
       <actions>
        <DROP  />
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-FORWARD" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes forwarding rules&quot;</comment>
        </comment>
        <mark >
          <mark >0x4000/0x4000</mark>
        </mark>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes forwarding conntrack rule&quot;</comment>
        </comment>
        <conntrack >
          <ctstate >RELATED,ESTABLISHED</ctstate>
        </conntrack>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-NODE-PORT" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;Kubernetes health check node port&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-HEALTH-CHECK-NODE-PORT dst</match-set>
        </set>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-KUBELET-CANARY" packet-count="0" byte-count="0" />
  </table>
<!-- # Completed on Tue Sep 20 10:54:54 2022 -->
<!-- # Generated by iptables*-save v1.4.21 on Tue Sep 20 10:54:54 2022 -->
  <table name="nat" >
    <chain name="PREROUTING" policy="ACCEPT" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes service portals&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <call >
          <KUBE-SERVICES />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="OUTPUT" policy="ACCEPT" packet-count="6" byte-count="360" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes service portals&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <call >
          <KUBE-SERVICES />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="POSTROUTING" policy="ACCEPT" packet-count="6" byte-count="360" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes postrouting rules&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <call >
          <KUBE-POSTROUTING />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-FIREWALL" packet-count="0" byte-count="0" >
      <rule >
       <actions>
        <call >
          <KUBE-MARK-DROP />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-LOAD-BALANCER" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;Kubernetes service load balancer ip + port with externalTrafficPolicy=local&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-LOAD-BALANCER-LOCAL dst,dst</match-set>
        </set>
       </conditions>
       <actions>
        <RETURN  />
       </actions>

      </rule>

      <rule >
       <actions>
        <call >
          <KUBE-MARK-MASQ />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-MARK-DROP" packet-count="0" byte-count="0" >
      <rule >
       <actions>
        <MARK >
          <set-xmark >0x8000/0x8000</set-xmark>
        </MARK>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-MARK-MASQ" packet-count="0" byte-count="0" >
      <rule >
       <actions>
        <MARK >
          <set-xmark >0x4000/0x4000</set-xmark>
        </MARK>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-NODE-PORT" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <match >
          <p >tcp</p>
        </match>
        <comment >
          <comment >&quot;Kubernetes nodeport TCP port with externalTrafficPolicy=local&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-NODE-PORT-LOCAL-TCP dst</match-set>
        </set>
       </conditions>
       <actions>
        <RETURN  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <p >tcp</p>
        </match>
        <comment >
          <comment >&quot;Kubernetes nodeport TCP port for masquerade purpose&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-NODE-PORT-TCP dst</match-set>
        </set>
       </conditions>
       <actions>
        <call >
          <KUBE-MARK-MASQ />
        </call>
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-POSTROUTING" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;Kubernetes endpoints dst ip:port, source ip for solving hairpin purpose&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-LOOP-BACK dst,dst,src</match-set>
        </set>
       </conditions>
       <actions>
        <MASQUERADE  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <mark >
          <mark  invert="1">0x4000/0x4000</mark>
        </mark>
       </conditions>
       <actions>
        <RETURN  />
       </actions>

      </rule>

      <rule >
       <actions>
        <MARK >
          <set-xmark >0x4000/0x0</set-xmark>
        </MARK>
       </actions>

      </rule>

      <rule >
       <conditions>
        <comment >
          <comment >&quot;kubernetes service traffic requiring SNAT&quot;</comment>
        </comment>
       </conditions>
       <actions>
        <MASQUERADE  />
       </actions>

      </rule>

    </chain>
    <chain name="KUBE-SERVICES" packet-count="0" byte-count="0" >
      <rule >
       <conditions>
        <comment >
          <comment >&quot;Kubernetes service lb portal&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-LOAD-BALANCER dst,dst</match-set>
        </set>
       </conditions>
       <actions>
        <call >
          <KUBE-LOAD-BALANCER />
        </call>
       </actions>

      </rule>

      <rule >
       <conditions>
        <match >
          <s  invert="1">172.16.0.0/12</s>
        </match>
        <comment >
          <comment >&quot;Kubernetes service cluster ip + port for masquerade purpose&quot;</comment>
        </comment>
        <set >
          <match-set >KUBE-CLUSTER-IP dst,dst</match-set>
        </set>
       </conditions>
       <actions>
        <call >
          <KUBE-MARK-MASQ />
        </call>
       </actions>

      </rule>

      <rule >
       <conditions>
        <addrtype >
          <dst-type >LOCAL</dst-type>
        </addrtype>
       </conditions>
       <actions>
        <call >
          <KUBE-NODE-PORT />
        </call>
       </actions>

      </rule>

      <rule >
       <conditions>
        <set >
          <match-set >KUBE-CLUSTER-IP dst,dst</match-set>
        </set>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

      <rule >
       <conditions>
        <set >
          <match-set >KUBE-LOAD-BALANCER dst,dst</match-set>
        </set>
       </conditions>
       <actions>
        <ACCEPT  />
       </actions>

      </rule>

    </chain>
    <chain name="INPUT" policy="ACCEPT" packet-count="0" byte-count="0" />
    <chain name="KUBE-KUBELET-CANARY" packet-count="0" byte-count="0" />
  </table>
<!-- # Completed on Tue Sep 20 10:54:54 2022 -->
</iptables-rules>
`