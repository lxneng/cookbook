"""Wait for multiple asynchronous tasks to complete."""

import asyncio


async def foo(n):
    await asyncio.sleep(5)  # wait 5s before continuing
    print(f"n: {n}!")


async def main():
    tasks = [foo(1), foo(2), foo(3)]
    await asyncio.gather(*tasks)


asyncio.run(main())

"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-gather.py 
n: 1!
n: 2!
n: 3!
"""