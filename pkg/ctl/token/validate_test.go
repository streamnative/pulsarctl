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
	"testing"

	"github.com/stretchr/testify/assert"
)

// nolint
var tokenData = []struct {
	secretFile         string
	publicFile         string
	tokenStr           string
	signatureAlgorithm string
	subject            string
}{
	{
		"",
		"../../../test/key/pulsar-admin-es256-public.key",
		"eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.fq6RDkvz4T_QTJcRrc5cUgGkQHKrft10RDKB2TDqfXvfTzbzEkfs0vrTwbz4sM4RfRFLR5LBfAKuYBrqtzy8fg",
		"ES256",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsar-admin-es384-public.key",
		"eyJhbGciOiJFUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.4jLGbnMfvdBbU9HpzVVXEVhC39F3Jy1Ph21b84V3LMPabB1WmFai5EgtzdrydyhpRN9WqbJB0D_1kHCHpNul0rPupyYfFzN0qAv9kJ273S2vY5sUL5_aI_AQlYtZ1eLV",
		"ES384",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsar-admin-es512-public.key",
		"eyJhbGciOiJFUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.ATkRXQtvGmWPmG9h_ELNfWQUE2EzDy24UX1HExv1S4bDxntbbKCYUr-S8S5_HC9wJHXZVU4r3MHEM12Dk3I3UmHnAYSKkVu1Dxn3pyG0Dec3Sg82qpZU5sGuTt245VovyxYwNvec2pl1ZP26F3dZBi6-ptlUOmwa2m-cIUIRaNg1bNew",
		"ES512",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsar-admin-rs256-public.key",
		"eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.AguFkJIOMjhTBAHwgNLwyrVgZcFviJpDhARCWJI5UxOji6Uij8S66zRNrjrbYiAP_lf5LMNjSbdPQpP3U9OtblCuQYYCHbVFuFqxxB9nyGNXOY4eg6qBx0WquFrS5V8mbgvL2NLxjMCcXhfpfWDr_ZubntIo_tYzkRa_dVL6dnPkFKnPS3cfjIdSDTaOsCfuTO0uA3ewyCZE8jGpWRt7Wx9He1NlW6Amf5fJVUI4At5GTgAtX5ChroWW5G2HvciJdLsmjJIiIkGx2gyrfev1G_QTMEMTHuMxVgDgFR8ZxkvJaGrggFzrb1BS-fEpg2_OhKTndH2p7hB3kWj7A_zdcQ",
		"RS256",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsar-admin-rs384-public.key",
		"eyJhbGciOiJSUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.bA0dvguAJGh1siK8U4ud79IWs9Xy-BelKYax9rvkFmWZSd7XkeqRiQL3dE3Mq1MGRhN9URYz0BhJJmsQoReOhTIvm1fSUCHPefU_fqaZ9SRR6zcTESk3sZb7S9Wd_W-Ylhm-j7bJoFHNwYsBRrHUpkQwgLNIk9DrmUMLGG8SI4ZwOCo1HGNFIs4I7Ywl4NalfeuqO0yzA07cO0ZVJhW9WCjUbaEUP0KH5RKR7B2Q-GRMS2LATTFDdF2XhC8WX5vRPbL7M3Y6dBL_qeI2mbXDb93bIylkw85xGB9VkFKZ9PA2KxEIljy7qg7t-aNHj6sYIpgRbaeK4CvYsU2wa6dttUjg18aXgJ9y_GqZbrvN6C6r1LaFvKqmPTDfaGmZQedXvCy30fvIFDV--Y2cQv2IbmVWUY2N6O0n_UT54ix13aipzwRnZV3R4-mNKM0oQNlezrdKIX6R_VpTiDNXGEDdqlcgP1yDppPDyquuhANCd4LROkX6yuOx0SxKbjCH6M6b",
		"RS384",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsar-admin-rs512-public.key",
		"eyJhbGciOiJSUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.LGfufrYd6X0OolDxc2FbDJ_GGlD568NVBxkTxKZbC9EQCs3XFLzQYNLoZIcF0gQ1TkghVeeh2weHPMIlaueRq3MgfLJasGjTs1D20hW2cI2VZMiF4TPZLwFBHyjz5_1BjB3pSClBAHiTi98iEhEXkockSkbmNNvJtADyi05DP7DH_OznYd9PzmOJjUEF050NmHDcTUirGAWswiMv56pEiuyRXPmeOZexTneX9dmM2UmsACvMWEayvcwP1o0VWhvb6gu9lzTROrnb310smOVj2hpmVc_mXBRGcc_Tcm-y82uGti53bzG5sGERPZJLjlsTHv5rrTaL7CI9ZPalSagZorq8TtHS2WAvwfLxpfja2pcf2trzJEkl-4xJRVB0TAW6OE-6KiQSfVVNi4XGN1thF2gRcD1FSrUogeKbAOCNvkXlC_CinG9cVRoL2fEzkqLv51Sdvw7RR2KiQ6jlTdQw2Cy3rAvJ33Rg3fOMh7_YB-eY-pQz-wfyQFq8IF3ADM-tr1sffb10OP-KM20f7rY_S9Zit-9kIwCY2ZmV4UwmEk-ctW5FUMXqnjmPPK7KotRGMhmMXQUhhPgZZlkGvPFTN5VveP9C0e6-hYqpVf-5nEsJm8pPBgPwjGGtfdm3IexVaE2h_51zJcaIHCUPWAw9i_XKnfT5uxRnbXyxO_xX4W0",
		"RS512",
		"test-token",
	},
	{
		"../../../test/key/pulsar-admin-hs256-secret.key",
		"",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.7ybKqieH_XZofazFBrnHh9gkU8H_qQVXOP2-dzF-7eY",
		"HS256",
		"test-token",
	},
	{
		"../../../test/key/pulsar-admin-hs384-secret.key",
		"",
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.o_AJgmT4Tu59K9sF-EkWPE4i9c9nPDHyFZ9JRcyS0y6bZsQleXQMTH01lmXiP3IF",
		"HS384",
		"test-token",
	},
	{
		"../../../test/key/pulsar-admin-hs512-secret.key",
		"",
		"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.RSOHl1bdfkLAxTxjF7YnygiI9P0xbAxk90SQitseNz5JgKKF9g68h2UsWhOHVOD2hbDQj_uFXJaViXAeL5Mkjg",
		"HS512",
		"test-token",
	},

	{
		"",
		"../../../test/key/pulsarctl-es256-public.key",
		"eyJhbGciOiJFUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.LT7rO1rHL4_EWPKwMMbCncmi1LALtmX7vltkTHS7vApNBbf3x75tf52u-4l5V_NCX6jyhD_-pmAacLJXXHy-fw",
		"ES256",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsarctl-es384-public.key",
		"eyJhbGciOiJFUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.Z9njMcryZN4mYFPrMO2eJVvnl7vYrweU0P-4LWN6Qd_l2CNHVhPI16z6iiEO_uOxPGNq-E96Se_fc2VMQ8ht_JIhEh6l_rLtV-sxBYs9OdQKZ6DF5CVAAl9K33adSOmE",
		"ES384",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsarctl-es512-public.key",
		"eyJhbGciOiJFUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.AIPl4s6ln6mU7laZn9DBdPX2mx4t8oOMtMhjD8sjXPOpeCyu1v5MRSWsSYQA20APyRxV5Nfz-yK7rom9F_jpNGVbAfuB9Re6XaEY41LAE1tInn46VpB_gZ6vuX-y-89cssOAxQmju8YcVrXUfyWqLV8ju6V-36jggdxMk9KixOCO8i8k",
		"ES512",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsarctl-rs256-public.key",
		"eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.Xta9ROaY-72bIGxQbHQlJgH7IXgfXMkk8ItyWmn-aG747U7Z72TJFGHuUd_d1kU73_iGD8Ovp89MYH0tyawSUO8-GXL2Zv3xmsNgBa7yFyFZBmxinaDx1dVZ2v_zv06GJOxZkd6hVJsV9JVWpCZkZcubNrEpf5FLvTWxD655cI5jU6WmQfAMe6HAeO5fItD5HA1FdlSWAONxFcXBgOyHWjuMDLp7Zgv08nb9BXLM5G7MuXrXmCIotgzT-qZ_gO_k1WEzWSOkn9WwwJnKglGu1LCNrFmQJ9bGSuekz6yXjCvPEa_0fDWnJkoVjmYl1gTMcrvzkT5qKehbGkraOT8u-Q",
		"RS256",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsarctl-rs384-public.key",
		"eyJhbGciOiJSUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.s_fC9IAYTVYwolP5p3vbMpgoK64AcpXdpQZ_DPJRKYon29Ku8kFIOy8sctnZQy3p7ZnbKeqZgot23Auhq2SRWRPF0NgJnNIc_pSyXUWbBe0xaMKXEQGuxGZShmG1mbOt2Nuw8was4Lz9vtXNJobFeyC_fAzHV5_BcVOINSc14jc8J_FX0qhnf6HfcSHnnZtQSLXzGM7nl8z5K14oqgvFxTowIFgF90n9M2iH6S1E790uH9Wig7amrZEY_GGIKKSeNNaA0Mga5DBUZWXuwmkCHq9BA1NtvXZg9hYOIHw5U2C93xlKk2SiDO_fEQsXs1utlOx7dTsR8M2zF0LmaUiQFevNcZvKTFtQD6HAikY_fsQTIpXsIvvkd2fhqvOZq7Ca_yFPWAoYKlte4wd7vp-jS1syCPYNv-f7j5Hw4OE_Buktv51pm5rkXzLXheyqkbY8qhtaC9XEYzIX1kqaGJ7snN5dK_5i7WiDWpoj0lBJFOUC_5--2wWYnyCi-LQRo5UL",
		"RS384",
		"test-token",
	},
	{
		"",
		"../../../test/key/pulsarctl-rs512-public.key",
		"eyJhbGciOiJSUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.cBp0eb9URPkuWNEddWKzXHJg1DlJ2-KnKRUVz4jN_pR_hzYnW8iIevTvxsrlcpYLxas3RGMtySSDVJt38nNX6IpJoD1xUkfqus9ngUVmD_VS6e623kXFj02VjtSeQJ4xaHXTcturt8sTkj7-dA2FDDNb-bleALQqiPP0hNGTikI5X6vJWn-t01wt9RRX9IzvprbacaN74pePLAojypQnFwEwrECcEoK8x8Mzt2Xx2VzU-ePGbA36ItR4-rji3GPiI7p5MX1-sFADF2vxyGJQEx83sPf3B_46NVwVWzM75JzqM9jdCjtFnVk_PCojQJ-2My_EiDMoqocm0-moLYVO5Am1Mp3oGcRiIiDMHfot5eRXh_Zbth_eZaNfHls0AwgKe2VW9uEX-AMbPZsPpjUYquYMFT8la8eKkpRZYZPHj2whIKkzY6mk3a2AQTjVo8tktxOJ96p2wQ0fQY13Z5ZDwWrWukHei568X7DlLvhbCxubokQIJI93PCeo_l2ft5zZEvtoRkSiDf1mFqjp8s1K1RoCzMmeFGuBvhmY3__9r_fTAjDnh3X1tZtYZhKpdq0uE8Tj16NQI2kG-8IIGsekCV_1BGfLyeM85q0AQ0NCSkDSZv7ECJIlymgZnipe5Ie5YczItm_CZ7MHn9K4EWAzs9tvUl9uPbBPXE1BEIwaBjc",
		"RS512",
		"test-token",
	},
	{
		"../../../test/key/pulsarctl-hs256-secret.key",
		"",
		"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.w7rq_Zz8vFdVSW4RHg_vhkjTSxnlfiEkzEtvthW9PeE",
		"HS256",
		"test-token",
	},
	{
		"../../../test/key/pulsarctl-hs384-secret.key",
		"",
		"eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.s-W_Z2dchUWC5YsbB9XogeL7Xw5bx9ZX56lCqkPfTAjBy_jZ6EXXwE4OwLYImsmv",
		"HS384",
		"test-token",
	},
	{
		"../../../test/key/pulsarctl-hs512-secret.key",
		"",
		"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ0ZXN0LXRva2VuIn0.QS1XsGm1qkeWFLjZUNWNZsTCm1lD5N05KqWACuxaXFSsanFIbnK8cdOez-C664aCefBau5aSJr11iTrcAIcy8A",
		"HS512",
		"test-token",
	},
}

func TestValidateCmd(t *testing.T) {
	t.FailNow()
	for _, data := range tokenData {
		doTestValidateCmd(t, data.secretFile, data.publicFile, data.tokenStr, data.signatureAlgorithm, data.subject)
	}
}

func doTestValidateCmd(t *testing.T, secretKeyFile, publicKeyFile, tokenString, signatureAlgorithm, subject string) {
	args := []string{
		"validate",
		"--secret-key-file", secretKeyFile,
		"--public-key-file", publicKeyFile,
		"--token-string", tokenString,
		"--signature-algorithm", signatureAlgorithm,
	}
	out, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.Nil(t, execErr)
	assert.Equal(t, fmt.Sprintf("The subject is %s and it will never expire.\n", subject), out.String())
}

func TestNoTokenStringOrFileErr(t *testing.T) {
	args := []string{"validate", "--public-key-file", "public.key"}
	_, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoTokenStringOrFile.Error(), execErr.Error())
}

func TestTokenSpecifiedMoreThanOneErr(t *testing.T) {
	args := []string{"validate", "--public-key-file", "public.key",
		"--token-string", "token", "--token-file", "token-file"}
	_, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errTokenSpecifiedMoreThanOne.Error(), execErr.Error())
}

func TestNoValidateKeySpecifiedErr(t *testing.T) {
	args := []string{"validate", "--token-string", "token", "--token-file", "token-file"}
	_, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errTokenSpecifiedMoreThanOne.Error(), execErr.Error())
}

func TestValidateKeySpecifiedMoreThanOneErr(t *testing.T) {
	args := []string{"validate", "--secret-key-file", "secret-file", "--public-key-file", "public-key-file",
		"--token-string", "token"}
	_, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errValidateKeySpecifiedMoreThanOne.Error(), execErr.Error())
}

func TestTrimSpaceForArgs(t *testing.T) {
	args := []string{"validate", "--secret-key-file", "    ", "--token-string", "    "}
	_, execErr, err := testTokenCommands(validate, args)
	assert.Nil(t, err)
	assert.NotNil(t, execErr)
	assert.Equal(t, errNoTokenStringOrFile.Error(), execErr.Error())
}
