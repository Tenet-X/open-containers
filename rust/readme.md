## Polkadot eco containers

### Build the runner

#### Polkadot

```shell
make polkadot_runner
```

#### Phala

##### Phala node

```shell
make phala_node_runner
```

### Push to container registries

#### gcr.io

1. configure auth

configure Docker to use the `gcloud` command-line tool to authenticate requests to Container Registry.

```shell
gcloud auth configure-docker
```

2. tag the container

```
docker tag rustbuilder:latest gcr.io/[GCP_PROJECTID]/rustbuilder:latest
```

3. push remote

```
docker push gcr.io/[GCP_PROJECTID]/rustbuilder:latest
```

4. pull from registry

```
docker pull gcr.io/[GCP_PROJECTID]/rustbuilder:latest
```