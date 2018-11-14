#!/usr/bin/python
'''
    (C) Copyright 2018 Intel Corporation.

    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.

    GOVERNMENT LICENSE RIGHTS-OPEN SOURCE SOFTWARE
    The Government's rights to use, modify, reproduce, release, perform, display,
    or disclose this software are subject to the terms of the Apache License as
    provided in Contract No. B609815.
    Any reproduction of computer software, computer software documentation, or
    portions thereof marked with this legend must also reproduce the markings.
    '''

import os
import time
import traceback
import sys
import json
import uuid

from avocado       import Test

sys.path.append('./util')
sys.path.append('../util')
sys.path.append('../../../utils/py')
sys.path.append('./../../utils/py')
import ServerUtils
import WriteHostFile
import daos_api
from daos_api import DaosContext
from daos_api import DaosPool
from daos_api import DaosContainer

class OpenClose(Test):
    """
    Tests DAOS container open/close function with handle parameter.
    """
    def setUp(self):
        # get paths from the build_vars generated by build
        with open(os.path.join(os.path.dirname(os.path.realpath(__file__)),
                                       '../../../../.build_vars.json')) as f:
            build_paths = json.load(f)
        self.basepath = os.path.normpath(build_paths['PREFIX']  + "/../")
        self.tmp = build_paths['PREFIX'] + '/tmp'

        self.server_group = self.params.get("server_group",'/server/','daos_server')

        # setup the DAOS python API
        self.Context = DaosContext(build_paths['PREFIX'] + '/lib/')
        self.POOL = None

        self.hostlist = self.params.get("test_machines",'/run/hosts/*')
        self.hostfile = WriteHostFile.WriteHostFile(self.hostlist, self.tmp)
        print("Host file is: {}".format(self.hostfile))

        ServerUtils.runServer(self.hostfile, self.server_group, self.basepath)
        time.sleep(5)

    def tearDown(self):
        if self.hostfile is not None:
            os.remove(self.hostfile)
        if self.POOL is not None and self.POOL.attached:
            self.POOL.destroy(1)

        ServerUtils.stopServer(hosts=self.hostlist)

    def test_openhandle(self):
        """
        Test container open with container handle parameter.

        :avocado: tags=container,openclose,openhandle
        """
        # parameters used in pool create
        createmode = self.params.get("mode",'/run/createtests/createmode/*/')
        createuid  = os.geteuid()
        creategid  = os.getegid()
        createsetid = self.params.get("setname",'/run/createtests/createset/')
        createsize  = self.params.get("size",'/run/createtests/createsize/')
        coh = self.params.get("coh",'/run/createtests/container_handle/*/')

        expected_result = coh[1]

        try:
            # initialize a python pool object then create the underlying
            # daos storage
            self.POOL = DaosPool(self.Context)
            self.POOL.create(createmode, createuid, creategid,
                            createsize, createsetid, None)

            poh = self.POOL.handle

            self.POOL.connect(1 << 1)

            # Container initialization and creation
            self.Container1 = DaosContainer(self.Context)
            self.Container2 = DaosContainer(self.Context)
            self.Container1.create(poh)
            self.Container2.create(poh)

            str_cuuid = self.Container1.get_uuid_str()
            cuuid = uuid.UUID(str_cuuid)
            coh_bad = self.Container2.coh

            # Destroying Container2 and use it's handle
            self.Container2.destroy()

            # Defining 'good' and 'bad' container handles
            if coh[0] == 'GOOD':
                coh[0] = self.Container1.coh
            else:
                coh[0] = coh_bad

            # opening Container handle.
            # Once with right coh and once with wrong coh
            self.Container1.open(poh, cuuid, 2, None, coh[0])
            self.Container1.close(coh[0])
            if expected_result in ['FAIL']:
                self.fail("Test was expected to fail but it passed.\n")

        except ValueError as e:
            print e
            print traceback.format_exc()
            if expected_result == 'PASS':
                self.fail("Test was expected to pass but it failed.\n")
        finally:
            # cleanup the Pool and Container
            self.Container1.destroy()
            self.POOL.disconnect()
            self.POOL.destroy(1)
            self.POOL = None

    def test_closehandle(self):
        """
        Test container close function with container handle paramter.

        :avocado: tags=container,openclose,closehandle
        """
        # parameters used in pool create
        createmode = self.params.get("mode",'/run/createtests/createmode/*/')
        createuid  = os.geteuid()
        creategid  = os.getegid()
        createsetid = self.params.get("setname",'/run/createtests/createset/')
        createsize  = self.params.get("size",'/run/createtests/createsize/')
        coh = self.params.get("coh",'/run/createtests/container_handle/*/')

        expected_result = coh[1]

        try:
            # initialize a python pool object then create the underlying
            # daos storage
            self.POOL = DaosPool(self.Context)
            self.POOL.create(createmode, createuid, creategid,
                            createsize, createsetid, None)

            poh = self.POOL.handle
            self.POOL.connect(1 << 1)

            # Container initialization and creation
            self.Container1 = DaosContainer(self.Context)
            self.Container2 = DaosContainer(self.Context)
            self.Container1.create(poh)
            self.Container2.create(poh)

            str_cuuid = self.Container1.get_uuid_str()
            cuuid = uuid.UUID(str_cuuid)
            coh_bad = self.Container2.coh

            # Destroying Container2 and use it's handle
            self.Container2.destroy()

             # Defining 'good' and 'bad' container handles
            if coh[0] == 'GOOD':
                coh[0] = self.Container1.coh
            else:
                coh[0] = coh_bad

            # opening Container handle.
            # Once with right coh and once with wrong coh
            self.Container1.open(poh, cuuid, 2, None, self.Container1.coh)
            self.Container1.close(coh[0])

            self.Container1.destroy()
            self.POOL.disconnect()
            self.POOL.destroy(1)

            if expected_result in ['FAIL']:
                self.fail("Test was expected to fail but it passed.\n")

        except ValueError as e:
            print e
            print traceback.format_exc()
            if expected_result == 'PASS':
                self.fail("Test was expected to pass but it failed.\n")

        finally:
            self.POOL = None
