## create a storage bucket if you don't have an existing one
gsutil mb gs://bosta-configs
gsutil cp slack-notifier.yml gs://bosta-configs/slack-notifier.yml


## cloud run
gcloud run deploy slack-notifier-srvc \
   --image=Your_docker_img_url \
   --no-allow-unauthenticated \
   --update-env-vars=CONFIG_PATH=gs://bosta-configs/slack-notifier.yml,PROJECT_ID=bosta-new-infrastructure


## add permissions and roles
gcloud projects add-iam-policy-binding  bosta-new-infrastructure \
   --member=serviceAccount:service-262530383813@gcp-sa-pubsub.iam.gserviceaccount.com \
   --role=roles/iam.serviceAccountTokenCreator

gcloud iam service-accounts create cloud-run-pubsub-invoker \
  --display-name "Cloud Run Pub/Sub Invoker"

gcloud run services add-iam-policy-binding slack-notifier-srvc \
   --member=serviceAccount:cloud-run-pubsub-invoker@bosta-new-infrastructure.iam.gserviceaccount.com \
   --role=roles/run.invoker

gcloud pubsub topics create cloud-builds


# last step to be executed
gcloud pubsub subscriptions create slack-notifier-sub \
   --topic=cloud-builds \
   --push-endpoint=https://slack-notifier-srvc-7ny5ksb27a-ey.a.run.app \
   --push-auth-service-account=cloud-run-pubsub-invoker@bosta-new-infrastructure.iam.gserviceaccount.com

