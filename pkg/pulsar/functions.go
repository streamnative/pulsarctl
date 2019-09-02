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
    `net/textproto`
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
    client   *client
    basePath string
}

func (c *client) Functions() Functions {
    return &functions{
        client:   c,
        basePath: "/functions",
    }
}

func (f *functions) createStringFromField(w *multipart.Writer, value string) (io.Writer, error) {
    h := make(textproto.MIMEHeader)
    h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s" `, "functionConfig"))
    h.Set("Content-Type", "application/json")
    return w.CreatePart(h)
}

func (f *functions) createTextFromFiled(w *multipart.Writer, value string) (io.Writer, error) {
    h := make(textproto.MIMEHeader)
    h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s" `, "url"))
    h.Set("Content-Type", "text/plain")
    return w.CreatePart(h)
}

func (f *functions) CreateFunc(funcConf *FunctionConfig, fileName string) error {
    endpoint := f.client.endpoint(f.basePath, funcConf.Tenant, funcConf.Namespace, funcConf.Name)

    // buffer to store our request as bytes
    bodyBuf := bytes.NewBufferString("")

    multiPartWriter := multipart.NewWriter(bodyBuf)

    jsonData, err := json.Marshal(funcConf)
    if err != nil {
        return err
    }

    stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
    if err != nil {
        return err
    }

    _, err = stringWriter.Write(jsonData)
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

    contentType := multiPartWriter.FormDataContentType()
    err = f.client.post(endpoint, nil, nil, bodyBuf, contentType)
    if err != nil {
        return err
    }

    return nil
}

func (f *functions) CreateFuncWithUrl(funcConf *FunctionConfig, pkgUrl string) error {
    endpoint := f.client.endpoint(f.basePath, funcConf.Tenant, funcConf.Namespace, funcConf.Name)
    // buffer to store our request as bytes
    bodyBuf := bytes.NewBufferString("")

    multiPartWriter := multipart.NewWriter(bodyBuf)

    textWriter, err := f.createTextFromFiled(multiPartWriter, "url")
    if err != nil {
        return err
    }

    _, err = textWriter.Write([]byte(pkgUrl))
    if err != nil {
        return err
    }

    jsonData, err := json.Marshal(funcConf)
    if err != nil {
        return err
    }

    stringWriter, err := f.createStringFromField(multiPartWriter, "functionConfig")
    if err != nil {
        return err
    }

    _, err = stringWriter.Write(jsonData)
    if err != nil {
        return err
    }

    if err = multiPartWriter.Close(); err != nil {
        return err
    }

    url := fmt.Sprintf("http://localhost:8080%s", endpoint)
    req, err := http.NewRequest(http.MethodPost, url, bodyBuf)
    if err != nil {
        return err
    }
    req.Header.Set("Content-Type", multiPartWriter.FormDataContentType())
    client := new(http.Client)
    response, err := client.Do(req)
    if err != nil {
        return err
    }
    if response.StatusCode < 200 || response.StatusCode >= 300 {
        return fmt.Errorf("response status:%s, response status code:%d", response.Status, response.StatusCode)
    }

    return nil
}
