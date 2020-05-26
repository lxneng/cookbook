"""uses the `FIRST_COMPLETED` option, whichever task finishes first is what will be returned."""

import asyncio
from random import randrange


async def foo(n):
    s = randrange(5)
    print(f"{n} will sleep for: {s} seconds")
    await asyncio.sleep(s)
    print(f"n: {n}!")


async def main():
    tasks = [foo(1), foo(2), foo(3)]
    result = await asyncio.wait(tasks, return_when=asyncio.FIRST_COMPLETED)
    print(result)


asyncio.run(main())

"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-wait-with-return-when.py 
1 will sleep for: 4 seconds
2 will sleep for: 0 seconds
3 will sleep for: 4 seconds
n: 2!
({<Task finished coro=<foo() done, defined at asyncio-wait-with-return-when.py:7> result=None>}, {<Task pending coro=<foo() running at asyncio-wait-with-return-when.py:10> wait_for=<Future pending cb=[<TaskWakeupMethWrapper object at 0x7fc6071c8e10>()]>>, <Task pending coro=<foo() running at asyncio-wait-with-return-when.py:10> wait_for=<Future pending cb=[<TaskWakeupMethWrapper object at 0x7fc6071c8dd0>()]>>})
"""