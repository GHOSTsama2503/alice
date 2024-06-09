#!/bin/python3
import logging
import os

from . import clients, env
from .modules import language, logger


log = logging.getLogger(__name__)


def create_default_paths():
    os.makedirs(env.SESSIONS_PATH, exist_ok=True)
    os.makedirs(env.STORAGE_PATH, exist_ok=True)
    os.makedirs(env.TEMP_PATH, exist_ok=True)


async def main():
    log.info("Creating default paths...")
    create_default_paths()

    log.info("Loading language files...")
    language.load_locales()

    await clients.start_main_client()

    log.info("Setting bot commands...")
    await clients.set_commands()

    log.info("Started successfully.")


if __name__ == "__main__":
    log.info("Initializing....")

    try:
        clients.bot.loop.create_task(main())
        clients.bot.loop.run_forever()

    except KeyboardInterrupt:
        pass

    except:
        log.error("Unexpected error:", exc_info=True)
