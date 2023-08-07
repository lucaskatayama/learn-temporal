# Multiple Runtimes

Using other runtimes to implement Temporal Workers

- [Golang](./golang)
- [Python](./python)
- [NodeJS](./nodejs)

# Running workflows

```shell
$ temporal workflow start \
  --task-queue your-task-queue \
  --type YourWorkflow \
  --input '"Jane Doe"'
```
