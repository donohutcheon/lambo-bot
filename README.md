# lambo-bot

## Commands 

### Run Docker Locally:
```shell
docker build . --tag lambo-bot --rm
docker run --env PORT=8080 lambo-bot
```

### gcloud Commands

https://github.com/google-github-actions/setup-gcloud/tree/master/example-workflows/gke

Create new project (project name must be unique)
```
gcloud projects create lambo-bot --name="lambo-bot"
```

```
gcloud iam service-accounts create lambo-bot-github-action --description="Service account for Gihub action deploys" --display-name="lambo-bot-github-action"
gcloud projects add-iam-policy-binding lambo-bot \
    --member="serviceAccount:lambo-bot-github-action@lambo-bot.iam.gserviceaccount.com" \
    --role="roles/container.developer"
gcloud projects add-iam-policy-binding lambo-bot \
    --member="serviceAccount:lambo-bot-github-action@lambo-bot.iam.gserviceaccount.com" \
    --role="roles/compute.storageAdmin"
gcloud iam service-accounts keys create key.json \
  --iam-account lambo-bot-github-action@lambo-bot.iam.gserviceaccount.com
  
gcloud projects add-iam-policy-binding lambo-bot \
    --member="serviceAccount:lambo-bot-github-action@lambo-bot.iam.gserviceaccount.com" \
    --role="roles/storage.buckets.create"
  
gcloud projects add-iam-policy-binding lambo-bot \
    --member="serviceAccount:lambo-bot-github-action@lambo-bot.iam.gserviceaccount.com" \
    --role="roles/compute.loadBalancerAdmin"
```

Create new gke cluster:
```shell
gcloud container clusters create lambo-bot --release-channel regular --zone europe-west1 --node-locations europe-west1

gcloud beta container --project "lambo-bot" clusters create "lambo-bot" \
    --region "europe-west1" --no-enable-basic-auth --cluster-version "1.17.17-gke.1101" \
    --release-channel "stable" --machine-type "g1-small" --image-type "COS" \
    --disk-type "pd-standard" --disk-size "32" --metadata disable-legacy-endpoints=true \
    --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" \
    --num-nodes "3" --no-enable-stackdriver-kubernetes --enable-ip-alias \
    --network "projects/lambo-bot/global/networks/default" --subnetwork "projects/lambo-bot/regions/europe-west1/subnetworks/default" \
    --default-max-pods-per-node "110" --no-enable-master-authorized-networks \
    --addons HorizontalPodAutoscaling,HttpLoadBalancing --enable-autoupgrade --enable-autorepair \
    --max-surge-upgrade 1 --max-unavailable-upgrade 0
    
gcloud beta container --project "lambo-bot" clusters create "lambo-bot" --zone "us-east1-d" --no-enable-basic-auth --cluster-version "1.18.12-gke.1210" --release-channel "regular" --machine-type "g1-small" --image-type "COS" --disk-type "pd-standard" --disk-size "32" --metadata disable-legacy-endpoints=true --scopes "https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" --num-nodes "2" --no-enable-stackdriver-kubernetes --enable-ip-alias --network "projects/lambo-bot/global/networks/default" --subnetwork "projects/lambo-bot/regions/us-east1/subnetworks/default" --default-max-pods-per-node "110" --no-enable-master-authorized-networks --addons HorizontalPodAutoscaling,HttpLoadBalancing,GcePersistentDiskCsiDriver --enable-autoupgrade --enable-autorepair --max-surge-upgrade 1 --max-unavailable-upgrade 0 --enable-shielded-nodes --node-locations "us-east1-d"
```


```shell
gcloud compute addresses create NAME --project=lambo-bot --network-tier=STANDARD --region=us-east1-d
gcloud compute addresses create lambo-bot-address --global
```