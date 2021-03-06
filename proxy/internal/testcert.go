package internal

// LocalhostCert is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at Jan 29 16:00:00 2084 GMT.
// generated from src/crypto/tls:
// go run generate_cert.go  --rsa-bits 1024 --host 127.0.0.1,::1,example.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var LocalhostCert = []byte(`-----BEGIN CERTIFICATE-----
MIICEzCCAXygAwIBAgIQS3cofn+2H4NxFntgaMRAPTANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCB
iQKBgQDp+sQVBNYwZ4YSskddAtTYq2NPdWYawNw9YQDBU9ft3fIm1r9UoyL/57bo
gCgFAkglXo06sAfuk+W6OXRPplEwxCU/mAiAjMLKES1V3oZnI42sTeiskdvb8j6E
47EpbWSA2OU4Nqulbh6vkGrzYzUdlmwwz+rGvfmHp1EOjMVzvQIDAQABo2gwZjAO
BgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUw
AwEB/zAuBgNVHREEJzAlggtleGFtcGxlLmNvbYcEfwAAAYcQAAAAAAAAAAAAAAAA
AAAAATANBgkqhkiG9w0BAQsFAAOBgQBChdgkaHaw83GFx8aDWoE3K4+h9YqXuvEP
b2OWAYlzY/U99BA9P0lE4vGpaIAeCFxalJ2AK3yHjt+eezy3sw0bMeG8ZNYcOyIV
exS95UdAKFt93a5zIWrkYQvhuzln1IOxPJQZ4rkq4nikLj2WuyGR7QnuVBdgPqP7
RN4BPb5Sog==
-----END CERTIFICATE-----`)

// LocalhostKey is the private key for LocalhostCert.
var LocalhostKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDp+sQVBNYwZ4YSskddAtTYq2NPdWYawNw9YQDBU9ft3fIm1r9U
oyL/57bogCgFAkglXo06sAfuk+W6OXRPplEwxCU/mAiAjMLKES1V3oZnI42sTeis
kdvb8j6E47EpbWSA2OU4Nqulbh6vkGrzYzUdlmwwz+rGvfmHp1EOjMVzvQIDAQAB
AoGAGVoXduOQRaxh5ZK1kslkwJlJaGmjB5EQDAJ/r3LjOZ3LyBOKpaQLfcjgk66X
J3vIz2vAR7SdF2elA5mIFb1CnJ4HW4cWHzgFQdUnUtoUNuMPy/9QREFfeag9GMPx
dZNiypiKqHDSY5ovUL92gtv5W0/w00lYpFiBaYLl+WHvQ6ECQQDvZpULZCEmZHwL
hZun4ObzLwFNZ9sNPgwJybnxVYaolXACeh4Ewur0kZlY9DJMqo7Rz82JWuFarkgU
GQK/L231AkEA+jP0+q7jfI8NJqwpWFDjwKiI7fadClcdUgXvW2c5wc2pEe4KiAqs
ZOWPGsH7SxigGRLzw01SCoInX5yw689JqQJARIOTPENXyWkQpyuBtLYE4qwdL039
vvh28YYuFQdpFm5ONCdG2A4AuCXDQVYB3zcg0KMsK5c6z3z5W+cchiLI0QJBAIDS
ZYz4pNoKEVxbAgKdy1XzsGTNN/gN+GO1+JJYKK23RRidNkDrNe3RIAhH3inBKRUf
4/AnjFkqwDkDRTh0htkCQQDfrRZr+gazwzDTSp23+l6MEbqBbc+TTC3c40zpNj4a
egxjd5+SkMj6zXEJxAOgo+LmQDGWsu1YQ+XXL87VPwIP
-----END RSA PRIVATE KEY-----`)

// LocalhostCert2 is a PEM-encoded TLS cert with SAN IPs
// "127.0.0.1" and "[::1]", expiring at Jan 29 16:00:00 2084 GMT.
// generated from src/crypto/tls:
// go run generate_cert.go  --rsa-bits 1024 --host 127.0.0.1,::1,example2.com --ca --start-date "Jan 1 00:00:00 1970" --duration=1000000h
var LocalhostCert2 = []byte(`-----BEGIN CERTIFICATE-----
MIICNTCCAZ6gAwIBAgIQeD/ltjjdLHO9L2c5eMG5rDANBgkqhkiG9w0BAQsFADAS
MRAwDgYDVQQKEwdBY21lIENvMCAXDTcwMDEwMTAwMDAwMFoYDzIwODQwMTI5MTYw
MDAwWjASMRAwDgYDVQQKEwdBY21lIENvMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCB
iQKBgQDzwrR53c4RyXpitbeRD9CY6PRhqYgnrCOBy0GUuGs5hJgMqSMXuIH4Vs4h
lOH19hb9o733O+qJM6s4D8GNfz2LC/SC/DOHqXv0DeB6lGJ47I2Wv8569uFjNh3K
pi5yYlAqNdjQ1TYjUZmDytiQxp8eCLCKGpbWvjWzop50GTefqwIDAQABo4GJMIGG
MA4GA1UdDwEB/wQEAwICpDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8E
BTADAQH/MB0GA1UdDgQWBBSLtvT+Rtv9N+Tw1ZbXU+jSxYqYxTAvBgNVHREEKDAm
ggxleGFtcGxlMi5jb22HBH8AAAGHEAAAAAAAAAAAAAAAAAAAAAEwDQYJKoZIhvcN
AQELBQADgYEAkO5LWq8SHx8dUgMDryCyrJamsFT3Z/Lt1zMfJfNdTRSvsg7Fy3XR
IOtxPqh2gT7OsSeU6fjjbDUTuGmH/BckwZTFMkRho/WaEgbP6XWWjkl+6euJBvtG
lBFElB/HVPa5puggihR9H1pE3s+SdtslwfOf8XsUA4xlcrhpU5kuMa4=
-----END CERTIFICATE-----`)

// LocalhostKey2 is the private key for LocalhostCert2.
var LocalhostKey2 = []byte(`-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAPPCtHndzhHJemK1
t5EP0Jjo9GGpiCesI4HLQZS4azmEmAypIxe4gfhWziGU4fX2Fv2jvfc76okzqzgP
wY1/PYsL9IL8M4epe/QN4HqUYnjsjZa/znr24WM2HcqmLnJiUCo12NDVNiNRmYPK
2JDGnx4IsIoalta+NbOinnQZN5+rAgMBAAECgYEAlRnYuN5SiRC7WpuacBHDX3TG
3uILFXE2utKwB58Sfzk6pCvk+kJyxYubRHFEEeX4RCcfMJYmrMu9BGqm0r0sz6nb
CSZk2Crn7eKgLK01+t2K+2s1R2oNB/fxkmxVUTbxiZ+Bt7xsAvFnnQRl06r9NNYo
XqAadfRFGlmMkSEAbbkCQQD9feneVqIdPPGQOadUq8VYY/M8onMSG72O0qr6Dx8X
j3hXf5D91pvXos+h3TwSICQ354BcakI4VK/EXxfmNUSXAkEA9iwkdXYQXkjgqpwH
3jMxG3DLAl9aHnykFSm8G2vQj2427ePnLppHclXKTPfu3E0qUNT2WgOK64N3qC1F
NkZ8DQJBAPSOlKFfnVlt4XOOW8QRT/wduZ4G79NJlhCDaFaFXi7ByI1J0h1C/ekE
9yInKXwnLCoPG0SNc0ObWFOwloMPYxMCQEIH5yOmro9Lxw+cWLPuUU7F+35Aa2Dg
F/chQbatPb0rWAqJZhpnAaEWh/QLUQPAowgZh5bvelTf57mxou4DDAUCQHynSBDh
GPDiBeTKX+VF50yHR8P6YrbeLIwasw19HA2BVhKKP05ZbQnWtN/ekhhBbI5fTwn0
njeVQ3HbrfnY1Ik=
-----END PRIVATE KEY-----`)
