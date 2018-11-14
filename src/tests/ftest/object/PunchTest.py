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
import avocado
from avocado import Test, main

sys.path.append('./util')
sys.path.append('../util')
sys.path.append('../../../utils/py')
sys.path.append('./../../utils/py')

import ServerUtils
import WriteHostFile
from daos_api import DaosContext, DaosPool, DaosContainer

class PunchTest(Test):
    """
    Simple test to verify the 3 different punch calls.
    """
    def setUp(self):

        try:
            # get paths from the build_vars generated by build
            with open('../../../.build_vars.json') as f:
                build_paths = json.load(f)
                self.basepath = os.path.normpath(build_paths['PREFIX'] + "/../")
                self.tmp = build_paths['PREFIX'] + '/tmp'

                self.server_group = self.params.get("server_group",'/server/',
                                                    'daos_server')

                # setup the DAOS python API
                self.Context = DaosContext(build_paths['PREFIX'] + '/lib/')

                self.hostlist = self.params.get("test_machines",'/run/hosts/*')
                self.hostfile = WriteHostFile.WriteHostFile(self.hostlist, self.tmp)

                ServerUtils.runServer(self.hostfile, self.server_group,
                                      self.basepath)

                # parameters used in pool create
                createmode = self.params.get("mode",'/run/pool/createmode/')
                createsetid = self.params.get("setname",'/run/pool/createset/')
                createsize  = self.params.get("size",'/run/pool/createsize/')

                createuid  = os.geteuid()
                creategid  = os.getegid()

                # initialize a python pool object then create the underlying
                # daos storage
                self.pool = DaosPool(self.Context)
                self.pool.create(createmode, createuid, creategid,
                                 createsize, createsetid, None)
                self.pool.connect(1 << 1)

                # create a container
                self.container = DaosContainer(self.Context)
                self.container.create(self.pool.handle)

                # now open it
                self.container.open()

        except ValueError as e:
            print(e)
            print(traceback.format_exc())
            self.fail("Test failed during setup.\n")

    def tearDown(self):

        try:
            self.container.close()

            # wait a few seconds and then destroy
            time.sleep(5)
            self.container.destroy()

            # cleanup the pool
            self.pool.disconnect()
            self.pool.destroy(1)

            ServerUtils.stopServer(hosts=self.hostlist)
            if self.hostfile is not None:
                os.remove(self.hostfile)

        except ValueError as e:
            print(e)
            print(traceback.format_exc())
            self.fail("Test failed during teardown.\n")

    def test_dkey_punch(self):
        """
        The most basic test of the dkey punch function.

        :avocado: tags=object,punch,dkeypunch,regression,vm,small
        """
        try:

            # create an object and write some data into it
            thedata = "a string that I want to stuff into an object"
            dkey = "this is the dkey"
            akey = "this is the akey"

            obj, epoch = self.container.write_an_obj(thedata, len(thedata)+1,
                                                     dkey, akey)

            # read the data back and make sure its correct
            thedata2 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                                  obj, epoch)
            if thedata != thedata2.value:
                print("data I wrote:" + thedata)
                print("data I read back" + thedata2.value)
                self.fail("Wrote data, read it back, didn't match\n")

            # repeat above, but know that the write_an_obj call is advancing
            # the epoch so the original copy remains and the new copy is in
            # a new epoch.
            thedata3 = "a different string"
            # note using the same keys so writing to the same spot
            obj, epoch2 = self.container.write_an_obj(thedata3, len(thedata3)+1,
                                                      dkey, akey, obj)

            # read the data back and make sure its correct
            thedata4 = self.container.read_an_obj(len(thedata3)+1, dkey, akey,
                                             obj, epoch2)
            if thedata3 != thedata4.value:
                print("data I wrote:" + thedata3)
                print("data I read back" + thedata4.value)
                self.fail("wrote in new epoch, read it back, didn't match\n")

            # the original data should still be there too
            thedata5 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                             obj, epoch)
            if thedata != thedata5.value:
                self.fail("original data isn't there any more\n")

            # repeat, so there will be 3 epochs
            thedata6 = "a really different string"

            # note using the same keys so writing to the same spot
            obj, epoch3 = self.container.write_an_obj(thedata6, len(thedata6)+1,
                                                      dkey, akey, obj)

            # read the data back and make sure its correct
            thedata7 = self.container.read_an_obj(len(thedata6)+1, dkey, akey,
                                                  obj, epoch3)
            if thedata6 != thedata7.value:
                print("data I wrote:" + thedata6)
                print("data I read back" + thedata7.value)
                self.fail("wrote in new epoch, read it back, didn't match\n")

            # now punch the data from the middle epoch
            obj.punch_dkeys(epoch2, [dkey])

        except ValueError as e:
            print(e)
            print(traceback.format_exc())
            self.fail("Test failed.\n")

        try:
            # read the data from the middle epoch
            thedata8 = self.container.read_an_obj(len(thedata3)+1, dkey, akey,
                                                  obj, epoch2)

            if len(thedata8.value) is not 0:
                print("data8: {} {}", thedata8.value, len(thedata8.value))
                self.fail("punch from middle epoch didn't work")

            # read the data from the last epoch
            thedata9 = self.container.read_an_obj(len(thedata6)+1, dkey, akey,
                                                  obj, epoch3)

            if len(thedata9.value) is not 0:
                print("data9: {} {}", thedata9.value, len(thedata9.value))
                self.fail("after punch data in the last epoch should be gone")

            # lastly check the first epoch
            thedata10 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                                  obj, epoch)

            if thedata != thedata10.value:
                self.fail("Epoch preceeding the punch should still have data\n")

        except ValueError as e:
            print(e)
            self.fail("Test failed.\n")

    def test_akey_punch(self):
        """
        The most basic test of the akey punch function.

        :avocado: tags=object,punch,akeypunch,regression,vm,small
        """
        try:
            # create an object and write some data into it
            dkey = "this is the dkey"
            data1 = [("this is akey 1", "this is data value 1"),
                    ("this is akey 2", "this is data value 2"),
                    ("this is akey 3", "this is data value 3")]
            obj, epoch1 = self.container.write_multi_akeys(dkey, data1)

            # do this again, note that the epoch has been advanced by
            # the write_multi_akeys function
            data2 = [("this is akey 1", "this is data value 4"),
                     ("this is akey 2", "this is data value 5"),
                     ("this is akey 3", "this is data value 6")]
            obj, epoch2 = self.container.write_multi_akeys(dkey, data2, obj)

            # do this again, note that the epoch has been advanced by
            # the write_multi_akeys function
            data3 = [("this is akey 1", "this is data value 7"),
                     ("this is akey 2", "this is data value 8"),
                     ("this is akey 3", "this is data value 9")]
            obj, epoch3 = self.container.write_multi_akeys(dkey, data3, obj)

            # read back the 1st epoch's data and check 1 value just to make sure
            # everything is on the up and up
            readbuf = [(data1[0][0], len(data1[0][1]) + 1),
                       (data1[1][0], len(data1[1][1]) + 1),
                       (data1[2][0], len(data1[2][1]) + 1)]
            retrieved_data = self.container.read_multi_akeys(dkey, readbuf, obj,
                                                             epoch1)
            if retrieved_data[data1[1][0]] != data1[1][1]:
                print("middle akey, 1st epoch {}".format(
                        retrieved_data[data1[1][0]]))
                self.fail("data retrieval failure")

            # now punch one akey from the middle epoch
            print("punching: {}".format([data2[1][0]]))
            obj.punch_akeys(epoch2, dkey, [data2[1][0]])

            # verify its gone from the epoch where it was punched
            readbuf = [(data2[1][0], len(data2[1][1]) + 1)]
            retrieved_data = self.container.read_multi_akeys(dkey, readbuf, obj,
                                                             epoch2)

            if len(retrieved_data[data2[1][0]]) != 0:
                print("retrieved: {}".format(retrieved_data))
                print("retrieved punched data but it was still there")
                self.fail("punched data still present")

        except ValueError as e:
            print(e)
            print(traceback.format_exc())
            self.fail("Test failed.\n")

    @avocado.skip('Currently this test fails')
    def test_obj_punch(self):
        """
        The most basic test of the object punch function.  Really similar
        to above except the whole object is deleted.

        :avocado: tags=object,punch,objpunch,regression,vm,small
        """
        try:

            # create an object and write some data into it
            thedata = "a string that I want to stuff into an object"
            dkey = "this is the dkey"
            akey = "this is the akey"

            obj, epoch = self.container.write_an_obj(thedata, len(thedata)+1,
                                                     dkey, akey)

            # read the data back and make sure its correct
            thedata2 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                                  obj, epoch)
            if thedata != thedata2.value:
                print("data I wrote:" + thedata)
                print("data I read back" + thedata2.value)
                self.fail("Wrote data, read it back, didn't match\n")

            # repeat above, but know that the write_an_obj call is advancing
            # the epoch so the original copy remains and the new copy is in
            # a new epoch.
            thedata3 = "a different string"
            # note using the same keys so writing to the same spot
            obj, epoch2 = self.container.write_an_obj(thedata3, len(thedata3)+1,
                                                      dkey, akey, obj)

            # read the data back and make sure its correct
            thedata4 = self.container.read_an_obj(len(thedata3)+1, dkey, akey,
                                             obj, epoch2)
            if thedata3 != thedata4.value:
                print("data I wrote:" + thedata3)
                print("data I read back" + thedata4.value)
                self.fail("wrote in new epoch, read it back, didn't match\n")

            # the original data should still be there too
            thedata5 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                             obj, epoch)
            if thedata != thedata5.value:
                self.fail("original data isn't there any more\n")

            # repeat, so there will be 3 epochs
            thedata6 = "a really different string"

            # note using the same keys so writing to the same spot
            obj, epoch3 = self.container.write_an_obj(thedata6, len(thedata6)+1,
                                                      dkey, akey, obj)

            # read the data back and make sure its correct
            thedata7 = self.container.read_an_obj(len(thedata6)+1, dkey, akey,
                                                  obj, epoch3)
            if thedata6 != thedata7.value:
                print("data I wrote:" + thedata6)
                print("data I read back" + thedata7.value)
                self.fail("wrote in new epoch, read it back, didn't match\n")

            # now punch the object from the middle epoch
            obj.punch(epoch2)

        except ValueError as e:
            print (e)
            print (traceback.format_exc())
            self.fail("Test failed.\n")

        try:
            # read the data from the middle epoch, should be gone
            thedata8 = self.container.read_an_obj(len(thedata3)+1, dkey, akey,
                                                  obj, epoch2)
            if len(thedata8.value) is not 0:
                print("data8: {} {}", thedata8.value, len(thedata8.value))
                self.fail("punch from middle epoch didn't work")

        except ValueError as e:
            print(e)
            self.fail("READ FROM DELETED OBJECT FAILED.\n")

        try:
            # read the data from the last epoch
            thedata9 = self.container.read_an_obj(len(thedata6)+1, dkey, akey,
                                                  obj, epoch3)

            if len(thedata9.value) is not 0:
                print("data9: {} {}", thedata8.value, len(thedata8.value))
                self.fail("after punch data in the last epoch should be gone")

        except ValueError as e:
            print(e)
            self.fail("READ FROM DELETED OBJECT FAILED.\n")

        try:
            # lastly check the first epoch, this one should still be there
            thedata10 = self.container.read_an_obj(len(thedata)+1, dkey, akey,
                                                  obj, epoch)

            if thedata != thedata10.value:
                self.fail("Epoch preceeding the punch should still have data\n")

        except ValueError as e:
            print(e)
            self.fail("Test failed.\n")

if __name__ == "__main__":
    main()
