# Extract public key
# TestIssuer is based on existing scripts. It is the common name of the certificate.
security find-certificate -c TestIssuer -p > cert.pem
openssl x509 -noout -pubkey -in cert.pem > pubkey.pem