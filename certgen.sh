
keydir="certs"
rm -rf ${keydir}
mkdir -p ${keydir}
cd "$keydir"


openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 100000 -out ca.crt -subj "/CN=admission_ca"
cat >server.conf <<EOF
[req]
req_extensions = v3_req
distinguished_name     = req_distinguished_name

[req_distinguished_name]
countryName            = US
stateOrProvinceName    = Nevada
organizationName       = Sample
organizationalUnitName = kDepName
commonName             = mike
emailAddress           = mc@cwxstat.com

[ v3_req ]

# Extensions to add to a certificate request

basicConstraints = CA:FALSE
keyUsage = nonRepudiation, digitalSignature, keyEncipherment
extendedKeyUsage = clientAuth, serverAuth
subjectAltName = @alt_names

[alt_names]
DNS.1   = validation.svc
DNS.2   = warden.validation.svc
DNS.3   = aipiggybot.com
EOF

openssl genrsa -out server.key 2048
#openssl req -new -key server.key -out server.csr -subj "/CN=warden.validation.svc" -config server.conf
openssl req -new -out server.csr -key server.key -config server.conf


openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 100000 -extensions v3_req -extfile server.conf

cp server.crt wardencrt.pem
cp server.key wardenkey.pem


# CA cert and private key
## openssl req -x509 -new -nodes -key ca.key -days 100000 -out ca.crt -subj "/CN=admission_ca"
# private key for the webhook server
## openssl genrsa -out warden.key 2048
# Generate and sign the key
## openssl req -new -key warden.key -subj "/CN=warden.validation.svc" \
##    | openssl x509 -req -CA ca.crt -CAkey ca.key -CAcreateserial -out warden.crt 
# Create .pem versions
## cp warden.crt wardencrt.pem \
##    | cp warden.key wardenkey.pem
