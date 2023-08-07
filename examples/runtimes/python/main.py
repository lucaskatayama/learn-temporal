from datetime import timedelta
from dataclasses import dataclass
import asyncio
from temporalio import workflow, activity
from temporalio.common import RetryPolicy
from temporalio.client import Client
from temporalio.worker import Worker



@dataclass
class YourParams:
    greeting: str
    name: str


@activity.defn(name="greet")
async def greet(input: YourParams) -> str:
    # return f"{input.greeting}, {input.name}!"
    raise ValueError("klasndlas")


@workflow.defn(name="YourWorkflow")
class YourWorkflow:
    @workflow.run
    async def run(self, name: str) -> str:
        return await workflow.execute_activity(
            "greet",
            YourParams("Hello from Python", name),
            start_to_close_timeout=timedelta(seconds=10),
            retry_policy=RetryPolicy(
                initial_interval=timedelta(1),
                backoff_coefficient=2,
                maximum_interval=timedelta(milliseconds=100),
                maximum_attempts=0,
            )
        )


async def main():
    client = await Client.connect("localhost:7233", identity="python-client")
    worker = Worker(
        client,
        task_queue="your-task-queue",
        workflows=[YourWorkflow],
        activities=[greet],
    )
    await worker.run()


if __name__ == "__main__":
    asyncio.run(main())
