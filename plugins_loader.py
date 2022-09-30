from functools import partial
from typing import Callable, Coroutine
from pyrogram import Client
from pyrogram.handlers.handler import Handler
from modules.logger import log
from inspect import getmembers, getfullargspec, isfunction, ismodule

def add_plugins(client: Client, module: object, function_prefix: str = ""):
    async def worker(coro: Coroutine, *args):
        client.loop.create_task(coro(*args))

    submodule_list = getmembers(module, ismodule)
    for submodule_name, submodule in submodule_list:
        function_list = getmembers(submodule, isfunction)
        for name, func in function_list:
            if not name.startswith(function_prefix):
                continue
            defaults = getfullargspec(func).defaults
            if defaults == None or len(defaults) != 2:
                continue
            if not issubclass(defaults[0], Handler):
                continue
            # if not isinstance(defaults[1], (Filter, tuple)):
            #     continue

            function: Callable = partial(worker, func)
            handler: type = defaults[0]
            filter = defaults[1]
            try:
                client.add_handler(handler(function, filter))
                log.info(f"Loaded: type={handler.__name__}, func={submodule_name}.{name}")
            except:
                log.info(f"Failed to load: func={submodule_name}.{name}")
