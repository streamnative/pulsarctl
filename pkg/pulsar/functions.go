// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package pulsar

import (
    `bytes`
    `encoding/json`
    `fmt`
    `io`
    `mime/multipart`
    `net/http`
    `os`
    `path/filepath`
    `strings`
)

type Functions interface {
    // Create a new function.
    CreateFunc(data *FunctionConfig, fileName string) error

    // Create a new function by providing url from which fun-pkg can be downloaded. supported url: http/file
    // eg:
    //  File: file:/dir/fileName.jar
    //  Http: http://www.repo.com/fileName.jar
    //
    // @param functionConfig
    //      the function configuration object
    // @param pkgUrl
    //      url from which pkg can be downloaded
    CreateFuncWithUrl(data *FunctionConfig, pkgUrl string) error
}

type functions struct {
    client 		*client
    basePath 	string
}

func (c *client) Functions() Functions {
    return &functions{
        client: c,
        basePath:"/functions",
    }
}

func (f *functions) CreateFunc(funcConf *FunctionConfig, fileName string) error {
    endpoint := f.client.endpoint(f.basePath, funcConf.Tenant, funcConf.Namespace, funcConf.Name)

    // buffer to store our request as bytes
    var requestBody bytes.Buffer

    multiPartWriter := multipart.NewWriter(&requestBody)

    res, err := json.Marshal(funcConf)
    if err != nil {
        return err
    }

    filedWriter, err := multiPartWriter.CreateFormField("functionConfig")
    if err != nil {
        return err
    }

    _, err = filedWriter.Write(res)
    if err != nil {

        return err
    }

    if fileName != "" && !strings.HasPrefix(fileName, "builtin://") {
        // If the function code is built in, we don't need to submit here
        file, err := os.Open(fileName)
        if err != nil {
            return err
        }
        defer file.Close()

        part, err := multiPartWriter.CreateFormFile("data", filepath.Base(file.Name()))

        if err != nil {
            return err
        }

        // copy the actual file content to the filed's writer
        _, err = io.Copy(part, file)
        if err != nil {
            return err
        }
    }

    // In here, we completed adding the file and the fields, let's close the multipart writer
    // So it writes the ending boundary
    if err = multiPartWriter.Close(); err != nil {
        return err
    }

    req, err := http.NewRequest("POST", endpoint, &requestBody)
    if err != nil {
        return err
    }

    // we need to set the content type from the writer, it includes necessary boundary as well.
    req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())

    // Do the request
    client := new(http.Client)

    // TODO: need add auth headers
    response, err := client.Do(req)
    if err != nil {
        return err
    }

    if response.StatusCode < 200 || response.StatusCode >= 300 {
        return fmt.Errorf("response status:%s, response status code:%d", response.Status, response.StatusCode)
    }

    return nil
}

func (f *functions) CreateFuncWithUrl(data *FunctionConfig, pkgUrl string) error {
    return nil
}



