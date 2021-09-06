# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
#     specific language governing permissions and limitations
# under the License.

#!/usr/bin/env python3

import fileinput
import glob
import sys

version = sys.argv[1]

filename = glob.glob('./scripts/run-integration-tests.sh', recursive=True)[0]

oldline = 'readonly PULSAR_DEFAULT_VERSION='
newline = '{}"{}"'.format(oldline, version)
process = lambda l : newline + '\n' if l.count(oldline) else l

lines = fileinput.input(files=(filename), inplace=True)

for line in lines:
    print(process(line), end='')