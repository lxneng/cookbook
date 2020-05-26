"""Convert a coroutine into a Task and schedule it onto the event loop."""

import asyncio


async def foo():
    await asyncio.sleep(10)
    print("Foo!")


async def hello_world():
    task = asyncio.create_task(foo())
    print(task)
    await asyncio.sleep(5)
    print("Hello World!")
    await asyncio.sleep(10)
    print(task)


asyncio.run(hello_world())

"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-create-task.py 
<Task pending coro=<foo() running at asyncio-create-task.py:6>>
Hello World!
Foo!
<Task finished coro=<foo() done, defined at asyncio-create-task.py:6> result=None>
"""