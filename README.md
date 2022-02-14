# Terrajet CloudFlare Provider

`provider-jet-cloudflare` is a [Crossplane](https://crossplane.io/) provider that
is built using [Terrajet](https://github.com/crossplane/terrajet) code
generation tools and exposes XRM-conformant managed resources for the 
CloudFlare API.

## Developing

Run code-generation pipeline:

```console
make generate
```

Install CRDS

```console
kubectl apply -f package/crds
```

Add credentials and provider config

```console
export CLOUDFLARE_API_TOKEN=<cloudflare-api-token>
cat examples/providerconfig/secret.yaml | envsubst | kubectl apply -f -
kubectl apply -f examples/providerconfig/providerconfig.yaml
```

Add localhost config, this should create a localhost entry in the provided zone

```console
export ZONE_ID=<zoneid>
cat examples/record/record.yaml | envsubst | kubectl apply -f -
```

Run against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build image:

```console
make image
```

Push image:

```console
make push
```

Build binary:

```console
make build
```

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/LucaLanziani/provider-jet-cloudflare/issues).

## Governance and Owners

provider-jet-cloudflare is run according to the same
[Governance](https://github.com/crossplane/crossplane/blob/master/GOVERNANCE.md)
and [Ownership](https://github.com/crossplane/crossplane/blob/master/OWNERS.md)
structure as the core Crossplane project.

## Code of Conduct

provider-jet-cloudflare adheres to the same [Code of
Conduct](https://github.com/crossplane/crossplane/blob/master/CODE_OF_CONDUCT.md)
as the core Crossplane project.

## Licensing

provider-jet-cloudflare is under the Apache 2.0 license.
