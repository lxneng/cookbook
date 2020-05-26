"""Utilize a timeout to prevent waiting endlessly for an asynchronous task to finish."""

import asyncio


async def foo(n):
    await asyncio.sleep(10)
    print(f"n: {n}!")


async def main():
    try:
        await asyncio.wait_for(foo(1), timeout=5)
    except asyncio.TimeoutError:
        print("timeout!")


asyncio.run(main())

"""
OUTPUT

(base) eric@pop-os:~/Sync/cookbook/Python$ python asyncio-wait-for-with-timeout.py 
timeout!
"""