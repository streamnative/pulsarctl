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

package token

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint
var testTokenData = []struct {
	signatureAlgorithm string
	subject            string
	tokenString        string
}{
	{
		"ES256",
		"test-token",
		"eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.fq6RDkvz4T_QTJcRrc5cUgGkQHKrft10RDKB2TDqfXvfTzbzEkfs0vrTwbz4sM4RfRFLR5LBfAKuYBrqtzy8fg",
	},
	{
		"ES384",
		"test-token",
		"eyJhbGciOiJFUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.4jLGbnMfvdBbU9HpzVVXEVhC39F3Jy1Ph21b84V3LMPabB1WmFai5EgtzdrydyhpRN9WqbJB0D_1kHCHpNul0rPupyYfFzN0qAv9kJ273S2vY5sUL5_aI_AQlYtZ1eLV",
	},
	{
		"ES512",
		"test-token",
		"eyJhbGciOiJFUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.ATkRXQtvGmWPmG9h_ELNfWQUE2EzDy24UX1HExv1S4bDxntbbKCYUr-S8S5_HC9wJHXZVU4r3MHEM12Dk3I3UmHnAYSKkVu1Dxn3pyG0Dec3Sg82qpZU5sGuTt245VovyxYwNvec2pl1ZP26F3dZBi6-ptlUOmwa2m-cIUIRaNg1bNew",
	},
	{
		"RS256",
		"test-token",
		"eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.AguFkJIOMjhTBAHwgNLwyrVgZcFviJpDhARCWJI5UxOji6Uij8S66zRNrjrbYiAP_lf5LMNjSbdPQpP3U9OtblCuQYYCHbVFuFqxxB9nyGNXOY4eg6qBx0WquFrS5V8mbgvL2NLxjMCcXhfpfWDr_ZubntIo_tYzkRa_dVL6dnPkFKnPS3cfjIdSDTaOsCfuTO0uA3ewyCZE8jGpWRt7Wx9He1NlW6Amf5fJVUI4At5GTgAtX5ChroWW5G2HvciJdLsmjJIiIkGx2gyrfev1G_QTMEMTHuMxVgDgFR8ZxkvJaGrggFzrb1BS-fEpg2_OhKTndH2p7hB3kWj7A_zdcQ",
	},
	{
		"RS384",
		"test-token",
		"eyJhbGciOiJSUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.bA0dvguAJGh1siK8U4ud79IWs9Xy-BelKYax9rvkFmWZSd7XkeqRiQL3dE3Mq1MGRhN9URYz0BhJJmsQoReOhTIvm1fSUCHPefU_fqaZ9SRR6zcTESk3sZb7S9Wd_W-Ylhm-j7bJoFHNwYsBRrHUpkQwgLNIk9DrmUMLGG8SI4ZwOCo1HGNFIs4I7Ywl4NalfeuqO0yzA07cO0ZVJhW9WCjUbaEUP0KH5RKR7B2Q-GRMS2LATTFDdF2XhC8WX5vRPbL7M3Y6dBL_qeI2mbXDb93bIylkw85xGB9VkFKZ9PA2KxEIljy7qg7t-aNHj6sYIpgRbaeK4CvYsU2wa6dttUjg18aXgJ9y_GqZbrvN6C6r1LaFvKqmPTDfaGmZQedXvCy30fvIFDV--Y2cQv2IbmVWUY2N6O0n_UT54ix13aipzwRnZV3R4-mNKM0oQNlezrdKIX6R_VpTiDNXGEDdqlcgP1yDppPDyquuhANCd4LROkX6yuOx0SxKbjCH6M6b",
	},
	{
		"RS512",
		"test-token",
		"eyJhbGciOiJSUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.LGfufrYd6X0OolDxc2FbDJ_GGlD568NVBxkTxKZbC9EQCs3XFLzQYNLoZIcF0gQ1TkghVeeh2weHPMIlaueRq3MgfLJasGjTs1D20hW2cI2VZMiF4TPZLwFBHyjz5_1BjB3pSClBAHiTi98iEhEXkockSkbmNNvJtADyi05DP7DH_OznYd9PzmOJjUEF050NmHDcTUirGAWswiMv56pEiuyRXPmeOZexTneX9dmM2UmsACvMWEayvcwP1o0VWhvb6gu9lzTROrnb310smOVj2hpmVc_mXBRGcc_Tcm-y82uGti53bzG5sGERPZJLjlsTHv5rrTaL7CI9ZPalSagZorq8TtHS2WAvwfLxpfja2pcf2trzJEkl-4xJRVB0TAW6OE-6KiQSfVVNi4XGN1thF2gRcD1FSrUogeKbAOCNvkXlC_CinG9cVRoL2fEzkqLv51Sdvw7RR2KiQ6jlTdQw2Cy3rAvJ33Rg3fOMh7_YB-eY-pQz-wfyQFq8IF3ADM-tr1sffb10OP-KM20f7rY_S9Zit-9kIwCY2ZmV4UwmEk-ctW5FUMXqnjmPPK7KotRGMhmMXQUhhPgZZlkGvPFTN5VveP9C0e6-hYqpVf-5nEsJm8pPBgPwjGGtfdm3IexVaE2h_51zJcaIHCUPWAw9i_XKnfT5uxRnbXyxO_xX4W0",
	},
	{
		"HS256",
		"test-token",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.7ybKqieH_XZofazFBrnHh9gkU8H_qQVXOP2-dzF-7eY",
	},
	{
		"HS384",
		"test-token",
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.o_AJgmT4Tu59K9sF-EkWPE4i9c9nPDHyFZ9JRcyS0y6bZsQleXQMTH01lmXiP3IF",
	},
	{
		"HS512",
		"test-token",
		"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.RSOHl1bdfkLAxTxjF7YnygiI9P0xbAxk90SQitseNz5JgKKF9g68h2UsWhOHVOD2hbDQj_uFXJaViXAeL5Mkjg",
	},
	{
		"ES256",
		"test-token",
		"eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.LT7rO1rHL4_EWPKwMMbCncmi1LALtmX7vltkTHS7vApNBbf3x75tf52u-4l5V_NCX6jyhD_-pmAacLJXXHy-fw",
	},
	{
		"ES384",
		"test-token",
		"eyJhbGciOiJFUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.Z9njMcryZN4mYFPrMO2eJVvnl7vYrweU0P-4LWN6Qd_l2CNHVhPI16z6iiEO_uOxPGNq-E96Se_fc2VMQ8ht_JIhEh6l_rLtV-sxBYs9OdQKZ6DF5CVAAl9K33adSOmE",
	},
	{
		"ES512",
		"test-token",
		"eyJhbGciOiJFUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.AIPl4s6ln6mU7laZn9DBdPX2mx4t8oOMtMhjD8sjXPOpeCyu1v5MRSWsSYQA20APyRxV5Nfz-yK7rom9F_jpNGVbAfuB9Re6XaEY41LAE1tInn46VpB_gZ6vuX-y-89cssOAxQmju8YcVrXUfyWqLV8ju6V-36jggdxMk9KixOCO8i8k",
	},
	{
		"RS256",
		"test-token",
		"eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.Xta9ROaY-72bIGxQbHQlJgH7IXgfXMkk8ItyWmn-aG747U7Z72TJFGHuUd_d1kU73_iGD8Ovp89MYH0tyawSUO8-GXL2Zv3xmsNgBa7yFyFZBmxinaDx1dVZ2v_zv06GJOxZkd6hVJsV9JVWpCZkZcubNrEpf5FLvTWxD655cI5jU6WmQfAMe6HAeO5fItD5HA1FdlSWAONxFcXBgOyHWjuMDLp7Zgv08nb9BXLM5G7MuXrXmCIotgzT-qZ_gO_k1WEzWSOkn9WwwJnKglGu1LCNrFmQJ9bGSuekz6yXjCvPEa_0fDWnJkoVjmYl1gTMcrvzkT5qKehbGkraOT8u-Q",
	},
	{
		"RS384",
		"test-token",
		"eyJhbGciOiJSUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.s_fC9IAYTVYwolP5p3vbMpgoK64AcpXdpQZ_DPJRKYon29Ku8kFIOy8sctnZQy3p7ZnbKeqZgot23Auhq2SRWRPF0NgJnNIc_pSyXUWbBe0xaMKXEQGuxGZShmG1mbOt2Nuw8was4Lz9vtXNJobFeyC_fAzHV5_BcVOINSc14jc8J_FX0qhnf6HfcSHnnZtQSLXzGM7nl8z5K14oqgvFxTowIFgF90n9M2iH6S1E790uH9Wig7amrZEY_GGIKKSeNNaA0Mga5DBUZWXuwmkCHq9BA1NtvXZg9hYOIHw5U2C93xlKk2SiDO_fEQsXs1utlOx7dTsR8M2zF0LmaUiQFevNcZvKTFtQD6HAikY_fsQTIpXsIvvkd2fhqvOZq7Ca_yFPWAoYKlte4wd7vp-jS1syCPYNv-f7j5Hw4OE_Buktv51pm5rkXzLXheyqkbY8qhtaC9XEYzIX1kqaGJ7snN5dK_5i7WiDWpoj0lBJFOUC_5--2wWYnyCi-LQRo5UL",
	},
	{
		"RS512",
		"test-token",
		"eyJhbGciOiJSUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.cBp0eb9URPkuWNEddWKzXHJg1DlJ2-KnKRUVz4jN_pR_hzYnW8iIevTvxsrlcpYLxas3RGMtySSDVJt38nNX6IpJoD1xUkfqus9ngUVmD_VS6e623kXFj02VjtSeQJ4xaHXTcturt8sTkj7-dA2FDDNb-bleALQqiPP0hNGTikI5X6vJWn-t01wt9RRX9IzvprbacaN74pePLAojypQnFwEwrECcEoK8x8Mzt2Xx2VzU-ePGbA36ItR4-rji3GPiI7p5MX1-sFADF2vxyGJQEx83sPf3B_46NVwVWzM75JzqM9jdCjtFnVk_PCojQJ-2My_EiDMoqocm0-moLYVO5Am1Mp3oGcRiIiDMHfot5eRXh_Zbth_eZaNfHls0AwgKe2VW9uEX-AMbPZsPpjUYquYMFT8la8eKkpRZYZPHj2whIKkzY6mk3a2AQTjVo8tktxOJ96p2wQ0fQY13Z5ZDwWrWukHei568X7DlLvhbCxubokQIJI93PCeo_l2ft5zZEvtoRkSiDf1mFqjp8s1K1RoCzMmeFGuBvhmY3__9r_fTAjDnh3X1tZtYZhKpdq0uE8Tj16NQI2kG-8IIGsekCV_1BGfLyeM85q0AQ0NCSkDSZv7ECJIlymgZnipe5Ie5YczItm_CZ7MHn9K4EWAzs9tvUl9uPbBPXE1BEIwaBjc",
	},
	{
		"HS256",
		"test-token",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.w7rq_Zz8vFdVSW4RHg_vhkjTSxnlfiEkzEtvthW9PeE",
	},
	{
		"HS384",
		"test-token",
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.s-W_Z2dchUWC5YsbB9XogeL7Xw5bx9ZX56lCqkPfTAjBy_jZ6EXXwE4OwLYImsmv",
	},
	{
		"HS512",
		"test-token",
		"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.QS1XsGm1qkeWFLjZUNWNZsTCm1lD5N05KqWACuxaXFSsanFIbnK8cdOez-C664aCefBau5aSJr11iTrcAIcy8A",
	},
}

func TestShowCmd(t *testing.T) {
	for _, data := range testTokenData {
		doTestShowCmd(t, data.signatureAlgorithm, data.subject, data.tokenString)
	}
}

func doTestShowCmd(t *testing.T, signatureAlgorithm, subject, tokenString string) {
	args := []string{"show", "--token-string", tokenString}
	out, execErr, err := testTokenCommands(show, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	// nolint
	assert.Equal(t,
		fmt.Sprintf("The algorithm and subject of the token are {\"alg\":\"%s\"}, {\"sub\":\"%s\"}.\n", signatureAlgorithm, subject),
		out.String())
}

func TestShowWithEnvToken(t *testing.T) {
	for _, data := range testTokenData {
		os.Setenv("TOKEN", data.tokenString)
		defer os.Unsetenv("TOKEN")
		doTestShowWithEnvToken(t, data.signatureAlgorithm, data.subject)
	}
	os.Clearenv()
}

func doTestShowWithEnvToken(t *testing.T, signatureAlgorithm, subject string) {
	args := []string{"show"}
	out, execErr, err := testTokenCommands(show, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t,
		fmt.Sprintf("The algorithm and subject of the token are {\"alg\":\"%s\"}, {\"sub\":\"%s\"}.\n",
			signatureAlgorithm, subject),
		out.String())
}

func TestCommand(t *testing.T) {
	// nolint
	args := []string{"show", "--token-string", "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.iGrhUJtoY2exvMXTCkZmdebCfyK_WM4ARkJhDexDiBny-dm6EK-TIpLkosCOq8_m7SrFwyCKaAlDmqJTTlmhyq-EbPe8KakLvtRxXByoNFx_HVNRZKnIM44fXbDTavmQRz3fbTzu1rmIkgERZcfWxoCg_R5Jmo11K-DHZit1YzftHVLYOVpKmjq1bdaW5cKmwuVf8BNDnnevuSGzIsvCQSQTMxpiAYm2nJV0bJlEsvJ6-WY1ATlzNE8JGEzLWe_sUr1MJfqvPApQ9AKx02PB3EtiwxH784Sm3LOu26jvTeVg0FkILv-o1q2eOBxLoFi7C-0egph_80_P1G2H2pd9bg"}
	out, execErr, err := testTokenCommands(show, args)
	if err != nil {
		t.Fatal()
	}

	if execErr != nil {
		t.Fatal()
	}
	fmt.Print(out.String())
}

func TestNoTokenStringOrFileErrInShowCmd(t *testing.T) {
	args := []string{"show"}
	_, execErr, err := testTokenCommands(show, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoTokenStringOrFile.Error(), execErr.Error())
}

func TestTokenSpecifiedMoreThanOneErrInShowCmd(t *testing.T) {
	args := []string{"show", "--token-string", "token-string", "--token-file", "token-file"}
	_, execErr, err := testTokenCommands(show, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errTokenSpecifiedMoreThanOne.Error(), execErr.Error())
}

func TestBlankArgs(t *testing.T) {
	args := []string{"show", "--token-string", "    "}
	_, execErr, err := testTokenCommands(show, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoTokenStringOrFile.Error(), execErr.Error())
}
