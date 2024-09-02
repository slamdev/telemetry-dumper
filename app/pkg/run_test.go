package pkg

import (
	"testing"

	"github.com/klauspost/compress/snappy"
	"github.com/xhhuango/json"
)

func TestNewApp(t *testing.T) {
	raw := `{"body":"/wYAAHNOYVBwWQBMBQCBAcrEtyjwTAq0KAq+AQolCglob3N0Lm5hbWUSGAoWaG90cm9kLWJmZDdmODRkYy16MjRnNAoSCgdvcy50eXBlEgcKBWxpbnV4ChcKDHNlcnZpY2UuBT6wBwoFcm91dGUKHgoWdGVsZW1ldHJ5LnNkay5sYW5ndWFnZRIECgJnbwolChJ0MiAABT8YDwoNb3BlbhU6DAohChUVDQVHaHZlcnNpb24SCAoGMS4yOC4wEscmCkcKPWdvLjI6AKAuaW8vY29udHJpYi9pbnN0cnVtZW50YXRpb24vbmV0L2h0dHAvb3RlbAEJ0BIGMC41My4wEpkEChAUgbMS/iuy8f587xfaYx2iEghKrciAoIHUqCIIfRuVKlkMz8YqBi9yAec8MAI59bCuECGQ8RdBNMnNEwEJDEoUCgsBVTwubWV0aG9kEgUKA0dFVEoVDRYkc2NoZW1lEgYKBAElHEoaCg1uZXQuOYAQCQoHMC4FAgRKFB0cMHBvcnQSAxiTP0ohChIBMlBzb2NrLnBlZXIuYWRkchILCgkxMjcFOQgxSho+IwAFPogEGLrVAkorChN1c2VyX2FnZW50Lm9yaWdpbmFsEhQKEkdvLQGTMC1jbGllbnQvMS4xShcNtRR0YXJnZXQhdQn4DEodChQBhRxwcm90b2NvbDWXCAUKAwE4CBYKCgFLAC5FDB03CCIKHAkYHGVzcG9uc2VfIaNEZW50X2xlbmd0aBICGDpKFwoQBSRQc3RhdHVzX2NvZGUSAxjIAVp1CWjqKWxoEhVIVFRQIHJlcXVlc3QgcmVjZWl2ZWQaDwoGMnUBKBoxCgN1cmwSKgooCbioP2Ryb3BvZmY9MTE1JTJDMjc3JnBpY2t1cD02OCUyQzQ5OBoPCgVsZXZlbCGiNGluZm96AIUBAAMAABKaUhwCRPMiJqaoL/K7Igh9ywEDsDNv2V0cBC4ACbAMQesC/P4cAv4cAn4cAgC4/hwC/hwChhwCADtmHAIMdgnKOSlsohwCADJJHAgrCil+HAIEMzhBHQQ2OGodAgCYUh0CRFICjAKCFdo0Igg1cYc5lhkPM10dCIVQQkUUEEHIEskU/jkE/jkEejkEALz+HQL+HQKGHQIAOWYdAgx0CcnRKWyiHQIAMEkdCCkKJ34dAgAygUgIOTUaujgERP2G1fZzWvMBIghkkYJCFo8Qal0bCOJNAUUSEEEnCdoX/hsC/hsCehsC/jgE/jgE/jgEBLKMKWyiGwKmOAQAN4U4CDUwOWYdAlZxCEQOX275M7yxbCIIwIsq4055Da5dHQiSRdKJOAxzI7gY/h0C/h0Ceh0C/nEI/nEI/nEIBIGCKWyiHQKmcQgcODk5JTJDMzBmHAJWVAZEVv8ZfqfhsFUiCM8f2Smqa2h0XRwI6XHNiTkIuE/O/hwC/hwCfhwC/lQG/lQG/lQGBMivKWyiHAKqVAYUMCUyQzExZhsCVjcERACJJniG69McIgiKppq3M9msTF0bCEl30UUSEEEQWYEa/jcE/jcEejcE/hsC/hsCihsCcqgMBAa/KWyiGwKuNwQQJTJDNjVqiwpWqAxEYWmOt+CgAs8iCOzD/2fAp9OSXRwMZgvfF0ETDEGxJ7P+HAL+HAJ+HAL+cAj+cAj+cAgEWkkpbKIcAqZwCCQ5NjQlMkM4ODIauo0KRDYAfCb0UH6cIgir7+H0lg7Xk10dCEB+vok5DAOOEhtBJv7hEP7hEP7hEP7hEPbhEHLFDgQR4Sls/h0CSh0CBDc0DsUOBDIwZjoEBBonDngRDHM6Ly9CShM0c2NoZW1hcy8xLjI2LjA=","headers":{"Accept-Encoding":["gzip"],"Content-Encoding":["snappy"],"Content-Length":["1370"],"Content-Type":["application/x-protobuf"],"User-Agent":["Grafana Alloy/v1.3.0 (linux/amd64)"]},"req":"/tempo/v1/traces"}`
	var req map[string]interface{}
	if err := json.Unmarshal([]byte(raw), &req); err != nil {
		t.Fatal(err)
	}
	body := req["body"].(string)
	decompressedBytes, err := snappy.Decode(nil, []byte(body))
	if err != nil {
		t.Fatal(err)
	}

	println(decompressedBytes)
}
