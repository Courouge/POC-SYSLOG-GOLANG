L'application est codé en Golang

Les middlewares utilisés sont: Messages syslog -> kafka-> App Golang -> Files

Possibilité de construire une Table de reference (KSQL) ou dans Elasticsearch

Exemple de messages syslog:
Mar 22 10:25:14 10.22.100.10 systemd-logind: New session 834 of user root.
Mar 22 10:25:14 10.22.100.10 systemd: Started Session 189 of user foo.
Mar 22 10:25:14 10.22.100.11 systemd: Starting Session 598 of user foo1.
Mar 22 10:25:29 10.22.100.11 systemd-logind: Removed session foo2.
Mar 22 10:26:44 10.22.100.11 systemd: Started Session 673 of user root.
Mar 22 10:26:44 10.22.100.12 systemd-logind: New session 492 of user foo.
Mar 22 10:26:44 10.22.100.12 systemd: Starting Session 2067 of user foo3.

05 10:15:51+01:00 10.105.162.1 Action="accept" UUid="{0x5c079746,0x3a,0x1fda8c0,0xc0000000}" rule="24" rule_uid="{0FAD6E56-5FF5-4CEF-95A2-72E479045765}" rule_name="XenDesktop XenApp - IGEL" service_id="UDP_30005" src="10.28.130.23" dst="10.105.33.112" proto="17" product="VPN-1 & FireWall-1" service="30005" s_port="52202" product_family="Network"


Messages in kafka:
{"logsource":"server1","program":"connect-distributed","host":"10.22.100.10","priority":30,"timestamp":"Mar 15 11:02:38","severity_label":"Informational","@version":"1","message":"[2019-03-15 11:02:38,569] INFO [Worker clientId=connect-1, groupId=connect-kafka] Discovered group coordinator 10.22.100.11:9092 (id: 2147483645 rack: null) (org.apache.kafka.clients.consumer.internals.AbstractCoordinator:654)","@timestamp":"2019-03-15T10:02:38.000Z","facility":3,"severity":6,"facility_label":"system"}
{"logsource":"server2","program":"connect-distributed","host":"10.22.100.11","priority":30,"timestamp":"Mar 16 11:02:38","severity_label":"Informational","@version":"1","message":"[2019-03-16 11:02:38,569] INFO [Worker clientId=connect-1, groupId=connect-kafka] Discovered group coordinator 10.22.100.11:9092 (id: 2147483645 rack: null) (org.apache.kafka.clients.consumer.internals.AbstractCoordinator:654)","@timestamp":"2019-03-15T10:02:38.000Z","facility":3,"severity":6,"facility_label":"system"}

Base KSQL:
Equipement  |  Type
10.105.162.1    FW


Messages out kafka:
Date| « FW »|Nom_techno|@IP du FW|hostname du FW|@IP source|@IP dest|Port Source|Port Dest|Action
05 10:15:51+01:00 | FW |  ? | 10.105.162.1 | ? | 10.28.130.23 | 10.105.33.112 | ? | ? | accept

Application Golang:
-> read messages in kafka 
-> mapping messages in kafka
-> Write in file or write in kafka
