import { proxyActivities, RetryPolicy } from '@temporalio/workflow';

const { greet } = proxyActivities({
  startToCloseTimeout: '1 minute',
  retry: {
    InitialInterval:    1,
    BackoffCoefficient: 2.0,
    MaximumInterval:    100, // 100 * InitialInterval
    MaximumAttempts:    0, // Unlimited
  }
});

/** A workflow that simply calls an activity */
export async function YourWorkflow(name) {
  return await greet({greeting: "Hello from NodeJS", name});
}
