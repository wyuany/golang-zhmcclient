/*
 * =============================================================================
 * IBM Confidential
 * © Copyright IBM Corp. 2020, 2021
 *
 * The source code for this program is not published or otherwise divested of
 * its trade secrets, irrespective of what has been deposited with the
 * U.S. Copyright Office.
 * =============================================================================
 */

package zhmcclient

import (
	"bufio"
	"encoding/json"
	"net/url"
	"os"
)

type HmcErrorCode int

const (
	ERR_CODE_HMC_INVALID_URL HmcErrorCode = iota + 1000
	ERR_CODE_HMC_BAD_REQUEST
	ERR_CODE_HMC_EMPTY_RESPONSE
	ERR_CODE_HMC_READ_RESPONSE_FAIL
	ERR_CODE_HMC_TRACE_REQUEST_FAIL
	ERR_CODE_HMC_EXECUTE_FAIL
	ERR_CODE_HMC_MARSHAL_FAIL
	ERR_CODE_HMC_UNMARSHAL_FAIL
	ERR_CODE_EMPTY_JOB_URI
)

const (
	ERR_MSG_INSECURE_URL   = "https is used for the client for secure reason."
	ERR_MSG_EMPTY_RESPONSE = "http response is empty."
	ERR_MSG_EMPTY_JOB_URI  = "empty job-uri."
)

type HmcError struct {
	Reason  int    `json:"reason"`
	Message string `json:"message"`
}

func getHmcErrorFromErr(reason HmcErrorCode, err error) *HmcError {
	return &HmcError{
		Reason:  int(reason),
		Message: err.Error(),
	}
}

func getHmcErrorFromMsg(reason HmcErrorCode, err string) *HmcError {
	return &HmcError{
		Reason:  (int)(reason),
		Message: err,
	}
}

func BuildUrlFromQuery(url *url.URL, query map[string]string) *url.URL {
	if query != nil {
		q := url.Query()
		for key, value := range query {
			q.Add(key, value)
		}
		url.RawQuery = q.Encode()
	}
	return url
}

func RetrieveBytes(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)
	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)
	return bytes, err
}

func GenerateErrorFromResponse(responseBodyStream []byte) *HmcError {
	errBody := &HmcError{}
	err := json.Unmarshal(responseBodyStream, errBody)

	if err != nil {
		return &HmcError{
			Reason:  -1,
			Message: "Unknown error.",
		}

	}

	return errBody
}
