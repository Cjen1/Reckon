import logging
import time
import threading
import reckon.reckon_types as t
from typing import Union
from typing_extensions import Literal


class Shared:
    def __init__(self, partitioned=[], fault_state=None):
        self.partitioned = partitioned
        self.fault_state = fault_state


class IntermittentPartialPartition(t.AbstractFault):
    def __init__(
        self,
        kind: Union[Literal["create"], Literal["remove"]],
        shared: Shared,
        sleep_duration,
        cluster,
        system,
    ):
        self._kind = kind
        self._shared = shared
        self._sleep_duration = sleep_duration
        self._cluster = cluster
        self._system = system

    def add_partition(self, host, remote):
        cmd = "iptables -I OUTPUT -d {0} -j DROP".format(remote.IP())
        logging.debug("cmd on {0} = {1}".format(host.name, cmd))
        host.cmd(cmd, shell=True)
        self._shared.partitioned.append(host)

    def heal_partition(self):
        cmd = "iptables -D OUTPUT 1"
        for host in self._shared.partitioned:
            host.cmd(cmd, shell=True)
            logging.debug("cmd on {0} = {1}".format(host.name, cmd))
        self._shared.partitioned = []

    def initiate_failure(self):
        leader = self._system.get_leader(self._cluster)
        non_leader = [h for h in self._cluster if not h == leader][0]
        self.failure_fixed = threading.Event()

        def thread_fn(self=self, leader=leader, non_leader=non_leader):
            while not self.failure_fixed.isSet():
                print("Partitioning {0} {1}".format(leader.name, non_leader.name))
                self.add_partition(leader, non_leader)
                self.add_partition(non_leader, leader)
                time.sleep(self.sleep_duration)
                self.heal_partition()
                time.sleep(self.sleep_duration)

        thread = threading.Thread(target=thread_fn, args=[])
        thread.start()

    def remove_failure(self):
        if self._shared.fault_state:
            self._shared.fault_state.set()

    def apply_fault(self):
        if self._kind == "create":
            self.initiate_failure()
        elif self._kind == "remove":
            self.remove_failure


class IntermittentPPartitionFailure(t.AbstractFailureGenerator):
    def __init__(self, sleep_duration):
        self.sleep_duration = sleep_duration

    def get_failures(self, cluster, system, restarters, stoppers):
        shared = Shared()
        return [
            t.NullFault(),
            IntermittentPartialPartition(
                "create", shared, self.sleep_duration, cluster, system
            ),
            IntermittentPartialPartition(
                "remove", shared, self.sleep_duration, cluster, system
            ),
            t.NullFault(),
        ]
