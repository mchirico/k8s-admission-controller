![CI](https://github.com/mchirico/k8s-admission-controller/workflows/CI/badge.svg)
# Reference: [TGIK](https://youtu.be/RVDK0m2XQeg?list=PL7bmigfV0EqQzxcNpmcdTJ9eFRPBe-iZa&t=1182)

# Steps to Test

```

make k8s119
make cluster
make
make kind
make load

```

# Steps to troubleshoot

```
make unload
kubectl run -it --rm --restart=Never --image=infoblox/dnstools:latest dnstools

# in another terminal, copy certs 
kubectl cp certs dnstools:/certs
make load

# Now, back on dnstools
curl --cacert /certs/clientCom.pem -v https://warden.validation.svc:443/

# Look for "SSL certificate verify ok."

```

# Steps to generate new certs

```
git clone https://github.com/michaelklishin/tls-gen tls-gen
cd tls-gen/basic
make regen CN=warden.validation.svc

```

Next, you need to combine files

```
cd result
openssl pkcs12 -in client_key.p12 -out clientCom.pem -nodes \
 -passin pass:
cat ca_certificate.pem >> clientCom.pem
```

Now, create base64 for webhook

```
cat clientCom.pem|base64 -w 0 > base64.txt
# If you're on a mac, instead
cat clientCom.pem|base64 > base64.txt
```

The contents of base64.txt goes in [webhook](https://github.com/mchirico/k8s-admission-controller/blob/6d46cc0d52c8f06cd36427933b7b0d843a59d341/webhook.yaml#L25)

<img src='https://user-images.githubusercontent.com/755710/93027093-bf554900-f5d8-11ea-97f8-d86b9468977a.png' width=560 />



# theITHollow Warden

This project is for setting up a basic Kubernetes validating Admission
Controller using python.

Steps to build your own admission controller.

1. Create your custom logic. There is an example admission controller shown in
   the /app directory that looks for a "billing" label to be applied to pods and
   deployments.

2. Update certgen.sh to match your admission controller. You may need to update
   the service and namespace where the controller lives.

3. Run the certgen.sh script to create the self-signed certificates for the
   admission controller.

4. Get the base64 value of the ca.crt file created by the certgen.sh script. 
`cat certs/ca.crt | base64`

5. Paste the base64 value into the caBundle location in the webhook.yaml file.

6. Build the container using the Dockerfile within the directory. Push the image
   to your image repository

7. Update the warden-k8s.yaml file to point to your new image.

8. apply the warden-k8s.yaml file to deploy your admission controller within the
   cluster.

9. Apply the webhook.yaml file to deploy the validation configuration to the
   Kubernetes API server.

10. Test your app. If using the default warden.py included with this repository,
    there are three test manifests in the [test-pods](/test-pods) folder.
