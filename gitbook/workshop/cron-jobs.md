# Cron Jobs

For anyone coming from the linux world cron job is a known term. But I won't assume that you are. 

Cron jobs are jobs that run at specific interval. Say you want to run some batch job everyday at 12. Or say you want clean up or process some data in the database every hour, cron jobs are your friend. 

Its pretty similar to a job except it has a schedule field that allows us to set a schedule just like linux cron jobs.

If you don't know how to use the schedule expression here is[ quick guide ](https://crontab.guru/#*_*_*_*_*)

Similar to our job we set up cron job that calls busybox sleep for 3 second. We do this every minute.

```text
schedule: "*/1 * * * *"
```

Run

```text
kubectl k8s/cron-jobs/cronjob.yaml
```

If you run

```text
watch kubectl get po
```

You will see every minute a new job is kicked off.

