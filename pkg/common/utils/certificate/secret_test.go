package certificate

import (
	"encoding/base64"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

func Test_BuildCAFromSecret(t *testing.T) {
	tss := []*corev1.Secret{
		&corev1.Secret{
			Data: nil,
		},
		&corev1.Secret{
			Data: map[string][]byte{"tls.crt": []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVIekNDQXdlZ0F3SUJBZ0lSQU0vRUlVMmRBSFRXdllxb1hGR0k5cWt3RFFZSktvWklodmNOQVFFTEJRQXcKWVRFc01Db0dBMVVFQ2hNalpHOXlhWE10YjNCbGNtRjBiM0l0ZDJWaWFHOXZheTF6WldOeVpYUXRkMkYwWTJneApNVEF2QmdOVkJBTVRLR1J2Y21sekxXOXdaWEpoZEc5eUxYZGxZbWh2YjJzdGMyVmpjbVYwTFhkaGRHTm9MVWhVClZGQXdIaGNOTWpRd09ERTVNRE13TVRFeFdoY05NalV3T0RFNU1ETXhNVEV4V2pCaE1Td3dLZ1lEVlFRS0V5TmsKYjNKcGN5MXZjR1Z5WVhSdmNpMTNaV0pvYjI5ckxYTmxZM0psZEMxM1lYUmphREV4TUM4R0ExVUVBeE1vWkc5eQphWE10YjNCbGNtRjBiM0l0ZDJWaWFHOXZheTF6WldOeVpYUXRkMkYwWTJndFNGUlVVRENDQVNJd0RRWUpLb1pJCmh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTWg0UHJ3VnJaSnByQ0hGQUZvWXhRdGs3dS9QRVlSVWlPaFYKbXJUT3ZoKzA0elliK0VtVWMzTktIWUk1dFoyTlp3bUtXOTE1Tk1UZlBBcUsxM1h2M3ZjWkdBbWtSMEg2ZzRIMgpwcTlYckwyTGlSbkpuTHlXRnhQdE1WNm1MR3U0alcwV0xZUnRZeVVUOHhabTVaamRYWkpnUVN0NU43VEVOQit0Cm13ZzBOWDBoRkhVTWtGUlZGY1Z2YzViZWdBTXVJaW1wRTZoMTExT2tRNERpVDFFSnNtczRxbzY2bEV5dzJrVEcKblZOVXk5YlpSSTJ5ZndTTFk2VkM5MVd6c0dWQU9VajFQcDkwb096aXh5T28wUWhEdldYTzZHdDMzTGpsdmFrOQpNQTM0MkovMkFwY0JYZy9rZ1dPNUphaW5HNFlNMUlZdUhlN1paWFZab1A3NWpGR0g3a3NDQXdFQUFhT0IwVENCCnpqQU9CZ05WSFE4QkFmOEVCQU1DQjRBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3SUdDQ3NHQVFVRkJ3TUIKTUF3R0ExVWRFd0VCL3dRQ01BQXdnWTRHQTFVZEVRU0JoakNCZzRJY1pHOXlhWE10YjNCbGNtRjBiM0l0YzJWeQpkbWxqWlM1a2IzSnBjNElnWkc5eWFYTXRiM0JsY21GMGIzSXRjMlZ5ZG1salpTNWtiM0pwY3k1emRtT0NMbVJ2CmNtbHpMVzl3WlhKaGRHOXlMWE5sY25acFkyVXVaRzl5YVhNdWMzWmpMbU5zZFhOMFpYSXViRzlqWVd5Q0VXUnYKY21sekxtVjRZVzF3YkdVdVkyOXRNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUFYYnJxZ2xhU2c5YkNIUEJZSQpCTkRRTmx2QkU4UXNiU1J3YTI4YUpTY0xYS1NEWTFNNVY5aWVwN3JWWmV2Wi91WnpYNWFBR0ZETDJ1TEMvOUJUCmVKRkRwZHFhaGNQc09zalRqMzdVeDRDR1hWRHVVNXR5Y3hlZTlwWnJnL0VDUm1hNUMwOG9sdGExanZKQkxPOXkKL0RNSjdPZHQ0M29oSm5yUTM4TFJwRlk4N3NweUlKZjFRbDgxOWI1UWk3eC9oOEFMOUF6Vk1BdkgyVmFoN1FYVQpPTitMbXdHS290N21RbDA0a09leVhNTFRJMmtqYWY2ZmpDUVpRZWg1TXN2REpEaG12dFB6ZVZwSmN5d0pUTUYxCitIQzhqK1Q0eERxK1pJZnF6RkhEalZjU2FreGpmOGQrUGMvbC9IS1N5VEhhbHZ1TkhobUVBcXBpMUllR1VnZDkKaGxKTgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="),
				"tls.key": []byte("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeUhnK3ZCV3RrbW1zSWNVQVdoakZDMlR1Nzg4UmhGU0k2RldhdE02K0g3VGpOaHY0ClNaUnpjMG9kZ2ptMW5ZMW5DWXBiM1hrMHhOODhDb3JYZGUvZTl4a1lDYVJIUWZxRGdmYW1yMWVzdll1SkdjbWMKdkpZWEUrMHhYcVlzYTdpTmJSWXRoRzFqSlJQekZtYmxtTjFka21CQkszazN0TVEwSDYyYkNEUTFmU0VVZFF5UQpWRlVWeFc5emx0NkFBeTRpS2FrVHFIWFhVNlJEZ09KUFVRbXlhemlxanJxVVRMRGFSTWFkVTFUTDF0bEVqYkovCkJJdGpwVUwzVmJPd1pVQTVTUFUrbjNTZzdPTEhJNmpSQ0VPOVpjN29hM2ZjdU9XOXFUMHdEZmpZbi9ZQ2x3RmUKRCtTQlk3a2xxS2NiaGd6VWhpNGQ3dGxsZFZtZy92bU1VWWZ1U3dJREFRQUJBb0lCQUFUdkdMUXVZR2hERzQ1MQpOODMvNEJHb3VVekRydTJ2MnZMRThKclVuODlOOTIya2dKYm5CWFJuOVpQMzhhVmdEcElvaXB4SDJHbWtDT2xGCmc0b09qbEZNbDgvYzhEZ0tmMlFhZXYzRHphZVlvSlkyNVJtckpBUkVpVHVJSVgzQlBCSGNRdUNJaExFcjdZYTIKbjU0RWpNSUdoblFZZHF1SFRMVWU0N3J1OWdZUFdEQ3Fjd3lMOXRic1ByUXQ0eWRQZ3FEaUpYU2I5OVlDQ2FtQwpJRm5LMkRvcldOckF3cXNTcGVKMC9nczhQL0Z2TktGeUJRVHBlaCs2ZGFTNElrOVFYQWZnTWl1UGJuODZ0ZndSCjQ4Z0hPYklETE1UeFdIS3BMTUo2bWdnaENoME5BZWxUNUYrZlVKWHM1djNtblBacGIzVm1CRlg2MUV4OGtWeE4KeFU4eVpFRUNnWUVBMHhVcVIvcEJTNm5uQitvL3dkK3JNVVFKaVpVZ2s5a0ZOeCswbUpUMk1BYkRCaS8xWU1jWApuQ0dvSGVYQTU0NFJEbDh1NVROcDgvaUk3eDYyK2M2dWV0UWp1UFBLeG9qLzV1YnJYeGNCNGpyMjdtNVd1OUowCjMxSWNFL1VocytLOEpYOHVqQ05uSWFnS2xKcm44cnhzRmhUTFVLRFhIcXY5cnJjVGQrZzZ6T3NDZ1lFQTh5RHUKM2JMei9HL0xkS2RKOEVLRGJVTjA4MytCb2NHK1hkN0FpeG9wRjBLNHhkZWNMK1hpdTRCMTNpUmRsWjB4UklWaQpXMHU4TEZsUG94WjhvUmRXQTUzamN2czltbnZSaXVMRlozc0VVZ0QyOWE5WDBsWmtRMzV1bzF4d2svc2Ftb3BJCjhySDVuSHNyQWo3WjJDaWg0QjJyaUpWbkEzRjBhWDRvS2lqcGpDRUNnWUJnRGU2YXFJQVVMWEhMd1VaWU5DOVUKRVBFQ0lkR2NWaEt3ZmdZUnRSKyt4U1QwYU5pUnZLZTZ6Zm9SK041cXdOUjBKTTVUVUswemIzTG8xYUpRVlVSQwo4c3g2dXI0WTBIa2JHeFFheG41OTgzVXhGUmIxQzVWUmxxemVUQnVWSDJXYzdwNUErVTlTR29VT0VOdmlNdXBBCnRoKzdSaVgrZUNmTjNOUElLZTZ3RndLQmdCN2ZkYldOUGRJMlYvMk9LS05ycnNpM0lpaXhia0tlR1pCTjh1RTMKQnpTKzhqdWRMMllkcVBicVR3MVFUTm5zeGc3UGFUZnluQWg5cHRFc3o5S2M1ZjA3bFdCS2F2NHM3dVRWK3Y4eQo3YytEODlncTFkb053bG1YdW5EZ0VpT1laVDV1aE9qY2xMZThKQy82enlyVnJnaUplL2l5RUF5dDRYeHIycW5QCjNDaEJBb0dCQUlyWmNmc29oSWxyeHZqL0V0SUdtVlRLOVNkOWJOVWlQRzZjSnJNekdTdjA1VGcvUDVmQjA0NjgKR1NEN3FiT0lOTmo2Q1Y2dFZKQitlOWhGR2hZRmRlTHJ4L3JCNHNVTk91aFkwcUxpakw1NlhGc2s2MGV4VG14SApBeVQvdjZIZ0hrUWVXSW1XTzFvMzhJSWxydGhuYnpDN2ZvRlBqVTZOM2ZnOEpRY2prQkpCCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==")},
		},
		&corev1.Secret{
			Data: map[string][]byte{"tls.crt": []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUVIekNDQXdlZ0F3SUJBZ0lSQU0vRUlVMmRBSFRXdllxb1hGR0k5cWt3RFFZSktvWklodmNOQVFFTEJRQXcKWVRFc01Db0dBMVVFQ2hNalpHOXlhWE10YjNCbGNtRjBiM0l0ZDJWaWFHOXZheTF6WldOeVpYUXRkMkYwWTJneApNVEF2QmdOVkJBTVRLR1J2Y21sekxXOXdaWEpoZEc5eUxYZGxZbWh2YjJzdGMyVmpjbVYwTFhkaGRHTm9MVWhVClZGQXdIaGNOTWpRd09ERTVNRE13TVRFeFdoY05NalV3T0RFNU1ETXhNVEV4V2pCaE1Td3dLZ1lEVlFRS0V5TmsKYjNKcGN5MXZjR1Z5WVhSdmNpMTNaV0pvYjI5ckxYTmxZM0psZEMxM1lYUmphREV4TUM4R0ExVUVBeE1vWkc5eQphWE10YjNCbGNtRjBiM0l0ZDJWaWFHOXZheTF6WldOeVpYUXRkMkYwWTJndFNGUlVVRENDQVNJd0RRWUpLb1pJCmh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTWg0UHJ3VnJaSnByQ0hGQUZvWXhRdGs3dS9QRVlSVWlPaFYKbXJUT3ZoKzA0elliK0VtVWMzTktIWUk1dFoyTlp3bUtXOTE1Tk1UZlBBcUsxM1h2M3ZjWkdBbWtSMEg2ZzRIMgpwcTlYckwyTGlSbkpuTHlXRnhQdE1WNm1MR3U0alcwV0xZUnRZeVVUOHhabTVaamRYWkpnUVN0NU43VEVOQit0Cm13ZzBOWDBoRkhVTWtGUlZGY1Z2YzViZWdBTXVJaW1wRTZoMTExT2tRNERpVDFFSnNtczRxbzY2bEV5dzJrVEcKblZOVXk5YlpSSTJ5ZndTTFk2VkM5MVd6c0dWQU9VajFQcDkwb096aXh5T28wUWhEdldYTzZHdDMzTGpsdmFrOQpNQTM0MkovMkFwY0JYZy9rZ1dPNUphaW5HNFlNMUlZdUhlN1paWFZab1A3NWpGR0g3a3NDQXdFQUFhT0IwVENCCnpqQU9CZ05WSFE4QkFmOEVCQU1DQjRBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3SUdDQ3NHQVFVRkJ3TUIKTUF3R0ExVWRFd0VCL3dRQ01BQXdnWTRHQTFVZEVRU0JoakNCZzRJY1pHOXlhWE10YjNCbGNtRjBiM0l0YzJWeQpkbWxqWlM1a2IzSnBjNElnWkc5eWFYTXRiM0JsY21GMGIzSXRjMlZ5ZG1salpTNWtiM0pwY3k1emRtT0NMbVJ2CmNtbHpMVzl3WlhKaGRHOXlMWE5sY25acFkyVXVaRzl5YVhNdWMzWmpMbU5zZFhOMFpYSXViRzlqWVd5Q0VXUnYKY21sekxtVjRZVzF3YkdVdVkyOXRNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUUFYYnJxZ2xhU2c5YkNIUEJZSQpCTkRRTmx2QkU4UXNiU1J3YTI4YUpTY0xYS1NEWTFNNVY5aWVwN3JWWmV2Wi91WnpYNWFBR0ZETDJ1TEMvOUJUCmVKRkRwZHFhaGNQc09zalRqMzdVeDRDR1hWRHVVNXR5Y3hlZTlwWnJnL0VDUm1hNUMwOG9sdGExanZKQkxPOXkKL0RNSjdPZHQ0M29oSm5yUTM4TFJwRlk4N3NweUlKZjFRbDgxOWI1UWk3eC9oOEFMOUF6Vk1BdkgyVmFoN1FYVQpPTitMbXdHS290N21RbDA0a09leVhNTFRJMmtqYWY2ZmpDUVpRZWg1TXN2REpEaG12dFB6ZVZwSmN5d0pUTUYxCitIQzhqK1Q0eERxK1pJZnF6RkhEalZjU2FreGpmOGQrUGMvbC9IS1N5VEhhbHZ1TkhobUVBcXBpMUllR1VnZDkKaGxKTgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==")},
		},
		&corev1.Secret{
			Data: map[string][]byte{"tls.key": []byte("LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFb3dJQkFBS0NBUUVBeUhnK3ZCV3RrbW1zSWNVQVdoakZDMlR1Nzg4UmhGU0k2RldhdE02K0g3VGpOaHY0ClNaUnpjMG9kZ2ptMW5ZMW5DWXBiM1hrMHhOODhDb3JYZGUvZTl4a1lDYVJIUWZxRGdmYW1yMWVzdll1SkdjbWMKdkpZWEUrMHhYcVlzYTdpTmJSWXRoRzFqSlJQekZtYmxtTjFka21CQkszazN0TVEwSDYyYkNEUTFmU0VVZFF5UQpWRlVWeFc5emx0NkFBeTRpS2FrVHFIWFhVNlJEZ09KUFVRbXlhemlxanJxVVRMRGFSTWFkVTFUTDF0bEVqYkovCkJJdGpwVUwzVmJPd1pVQTVTUFUrbjNTZzdPTEhJNmpSQ0VPOVpjN29hM2ZjdU9XOXFUMHdEZmpZbi9ZQ2x3RmUKRCtTQlk3a2xxS2NiaGd6VWhpNGQ3dGxsZFZtZy92bU1VWWZ1U3dJREFRQUJBb0lCQUFUdkdMUXVZR2hERzQ1MQpOODMvNEJHb3VVekRydTJ2MnZMRThKclVuODlOOTIya2dKYm5CWFJuOVpQMzhhVmdEcElvaXB4SDJHbWtDT2xGCmc0b09qbEZNbDgvYzhEZ0tmMlFhZXYzRHphZVlvSlkyNVJtckpBUkVpVHVJSVgzQlBCSGNRdUNJaExFcjdZYTIKbjU0RWpNSUdoblFZZHF1SFRMVWU0N3J1OWdZUFdEQ3Fjd3lMOXRic1ByUXQ0eWRQZ3FEaUpYU2I5OVlDQ2FtQwpJRm5LMkRvcldOckF3cXNTcGVKMC9nczhQL0Z2TktGeUJRVHBlaCs2ZGFTNElrOVFYQWZnTWl1UGJuODZ0ZndSCjQ4Z0hPYklETE1UeFdIS3BMTUo2bWdnaENoME5BZWxUNUYrZlVKWHM1djNtblBacGIzVm1CRlg2MUV4OGtWeE4KeFU4eVpFRUNnWUVBMHhVcVIvcEJTNm5uQitvL3dkK3JNVVFKaVpVZ2s5a0ZOeCswbUpUMk1BYkRCaS8xWU1jWApuQ0dvSGVYQTU0NFJEbDh1NVROcDgvaUk3eDYyK2M2dWV0UWp1UFBLeG9qLzV1YnJYeGNCNGpyMjdtNVd1OUowCjMxSWNFL1VocytLOEpYOHVqQ05uSWFnS2xKcm44cnhzRmhUTFVLRFhIcXY5cnJjVGQrZzZ6T3NDZ1lFQTh5RHUKM2JMei9HL0xkS2RKOEVLRGJVTjA4MytCb2NHK1hkN0FpeG9wRjBLNHhkZWNMK1hpdTRCMTNpUmRsWjB4UklWaQpXMHU4TEZsUG94WjhvUmRXQTUzamN2czltbnZSaXVMRlozc0VVZ0QyOWE5WDBsWmtRMzV1bzF4d2svc2Ftb3BJCjhySDVuSHNyQWo3WjJDaWg0QjJyaUpWbkEzRjBhWDRvS2lqcGpDRUNnWUJnRGU2YXFJQVVMWEhMd1VaWU5DOVUKRVBFQ0lkR2NWaEt3ZmdZUnRSKyt4U1QwYU5pUnZLZTZ6Zm9SK041cXdOUjBKTTVUVUswemIzTG8xYUpRVlVSQwo4c3g2dXI0WTBIa2JHeFFheG41OTgzVXhGUmIxQzVWUmxxemVUQnVWSDJXYzdwNUErVTlTR29VT0VOdmlNdXBBCnRoKzdSaVgrZUNmTjNOUElLZTZ3RndLQmdCN2ZkYldOUGRJMlYvMk9LS05ycnNpM0lpaXhia0tlR1pCTjh1RTMKQnpTKzhqdWRMMllkcVBicVR3MVFUTm5zeGc3UGFUZnluQWg5cHRFc3o5S2M1ZjA3bFdCS2F2NHM3dVRWK3Y4eQo3YytEODlncTFkb053bG1YdW5EZ0VpT1laVDV1aE9qY2xMZThKQy82enlyVnJnaUplL2l5RUF5dDRYeHIycW5QCjNDaEJBb0dCQUlyWmNmc29oSWxyeHZqL0V0SUdtVlRLOVNkOWJOVWlQRzZjSnJNekdTdjA1VGcvUDVmQjA0NjgKR1NEN3FiT0lOTmo2Q1Y2dFZKQitlOWhGR2hZRmRlTHJ4L3JCNHNVTk91aFkwcUxpakw1NlhGc2s2MGV4VG14SApBeVQvdjZIZ0hrUWVXSW1XTzFvMzhJSWxydGhuYnpDN2ZvRlBqVTZOM2ZnOEpRY2prQkpCCi0tLS0tRU5EIFJTQSBQUklWQVRFIEtFWS0tLS0tCg==")},
		},
	}

	for _, ts := range tss {
		if ts.Data != nil {
			if crt, ok := ts.Data[TLsCertName]; ok {
				dbyte, err := base64.StdEncoding.DecodeString(string(crt))
				if err != nil {
					t.Errorf("decode tls.crt failed.")
				}
				ts.Data[TLsCertName] = dbyte
			}
			if key, ok := ts.Data[TlsKeyName]; ok {
				kbyte, err := base64.StdEncoding.DecodeString(string(key))
				if err != nil {
					t.Errorf("decode tls.key failed")
				}
				ts.Data[TlsKeyName] = kbyte
			}
		}
		ca := BuildCAFromSecret(ts)
		if ts.Data == nil {
			if ca != nil {
				t.Errorf("secret data is nil, ca is not nil.")
			}
		} else if _, ok := ts.Data[TlsKeyName]; !ok {
			if ca != nil {
				t.Errorf("secret data tls.key is nil, ca is not nil.")
			}
		} else if _, ok = ts.Data[TLsCertName]; !ok {
			if ca != nil {
				t.Errorf("secret data tls.cert is nil, ca is not nil.")
			}
		} else {
			if ca == nil {
				t.Errorf("secret is normal ca is nil.")
			}
		}
	}
}