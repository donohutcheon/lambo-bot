# lambo-bot

## Commands 

### Run Docker Locally:
```shell
docker build . --tag lambo-bot --rm
docker run --env PORT=8080 lambo-bot
```

### gcloud Commands

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
```

Create new gke cluster:
```shell
gcloud container clusters create lambo-bot --release-channel regular --zone europe-west1 --node-locations europe-west1
```