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

### Access private gcr.io registry in k8s/GKE

#### Create a Service Account

1. Create an IAM service account

```
gcloud iam service-accounts create [account-name]
```

2. Grant the service account access to Container Registry.

```
gcloud projects add-iam-policy-binding [project-id] \
  --member serviceAccount:[account-name]@[project-id].iam.gserviceaccount.com \
  --role roles/storage.objectViewer
```

3. Download the account's service account kay

```
gcloud iam service-accounts keys create key.json \
  --iam-account [account-name]@[project-id].iam.gserviceaccount.com
```

#### Define a k8s secret

```
kubectl create secret docker-registry gcr-secret \
    --docker-server=gcr.io \
    --docker-username=_json_key \
    --docker-email=[account-name]@[project-id].iam.gserviceaccount.com \
    --docker-password="$(cat key.json)"
```

You may want to tidy things up by running `rm key.json`.