"""When dealing with a Task, which really is a Future, then you have the ability to execute a `callback` function once the Future has a value set on it."""

import asyncio


async def foo():
    await asyncio.sleep(10)
    return "Foo!"


def got_result(future):
    print(f"got the result! {future.result()}")


async def hello_world():
    task = asyncio.create_task(foo())
    task.add_done_callback(got_result)
    print(task)
    await asyncio.sleep(5)
    print("Hello World!")
    await asyncio.sleep(10)
    print(task)


asyncio.run(hello_world())

"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-create-task-with-callback.py 
<Task pending coro=<foo() running at asyncio-create-task-with-callback.py:6> cb=[got_result() at asyncio-create-task-with-callback.py:11]>
Hello World!
got the result! Foo!
<Task finished coro=<foo() done, defined at asyncio-create-task-with-callback.py:6> result='Foo!'>
"""