from mininet.net import Mininet
from mininet.node import Controller, Host
from mininet.log import setLogLevel
from mininet.link import TCLink

import reckon.reckon_types as t

import math

setLogLevel("info")


class SimpleTopologyProvider(t.AbstractTopologyGenerator):
    def __init__(self, number_nodes, number_clients, link_latency=None, link_loss=None, link_jitter=None):
        self.number_nodes = number_nodes
        self.number_clients = number_clients

        # Since we have a star topology we use link_latency = link_latency / 2
        self.per_link_latency = (
                None if not link_latency else f"{link_latency / 2}ms"
        )

        
        self.per_link_jitter = (
                None if not link_jitter else f"{math.sqrt(((link_jitter * link_latency) ** 2) / 2)}ms"
        ) if (link_jitter is not None and link_jitter > 0) else None

        # since we have 2 links, when we want the abstraction of one direct link
        # we use link_loss = 1 - sqrt(1 - L)
        self.per_link_loss = (
            None if not link_loss else (1 - math.sqrt(1 - link_loss / 100)) * 100
        )
        if self.per_link_loss == 0:
            self.per_link_loss = None

        self.switch_num = 0
        self.host_num = 0
        self.client_num = 0

    def add_switch(self):
        name = "s%s" % str(self.switch_num + 1)
        self.switch_num += 1
        return self.net.addSwitch(name)

    def add_host(self) -> Host:
        name = "h%s" % str(self.host_num + 1)
        self.host_num += 1
        return self.net.addHost(name)

    def add_client(self) -> Host:
        name = "mc%s" % str(self.client_num + 1)
        self.client_num += 1
        return self.net.addHost(name)

    def setup(self):
        self.net = Mininet(controller=Controller, link=TCLink)
        self.net.addController("c0")
        sw = self.add_switch()

        hosts = [self.add_host() for _ in range(self.number_nodes)]
        clients = [self.add_client() for _ in range(self.number_clients)]

        for host in hosts + clients:
            self.net.addLink(host, sw, delay=self.per_link_latency, loss=self.per_link_loss, jitter=self.per_link_jitter)

        self.net.start()

        return (self.net, hosts, clients)
