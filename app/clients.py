import asyncio
import logging
import os
import sys

from pyrogram import Client
from pyrogram.enums import ParseMode
from pyrogram.types import BotCommand

from . import env
from .modules import language

log = logging.getLogger(__name__)
plugins = {"root": env.PLUGINS_PATH}

multi_clients = {}
work_loads = {}


if env.BOT_SESSION:
    bot = Client(env.CLIENT_NAME, session_string=env.BOT_SESSION, plugins=plugins)

elif env.BOT_TOKEN and env.API_ID and env.API_HASH:
    bot = Client(
        env.CLIENT_NAME,
        api_id=env.API_ID,
        api_hash=env.API_HASH,
        bot_token=env.BOT_TOKEN,
        plugins=plugins,
        workdir=env.SESSIONS_PATH,
    )

else:
    log.critical("Undefined required environment variables.")
    sys.exit()


async def start_main_client(disable_parse_mode: bool = True):
    """
    Initialize main Telegram client.
    `disable_parse_mode` (bool): Completely disable parsing for peace of mind.
    """

    if disable_parse_mode:
        bot.set_parse_mode(ParseMode.DISABLED)

    try:
        log.info("Starting main client...")
        await bot.start()

        log.info(f"Bot -> {bot.me.first_name} (https://t.me/{bot.me.username})")

        if env.START_MESSAGE:
            await bot.send_message(env.ADMIN_ID, text="Started!")

    except Exception:
        log.error("Error starting main client:", exc_info=True)


async def set_commands():
    for language_code in language.locales:
        bot_commands: list[BotCommand] = []

        commands: dict[str, str] = language.bot_commands(language_code)
        for command, description in commands.items():
            bot_commands.append(BotCommand(command, description))

        await bot.set_bot_commands(bot_commands, language_code=language_code)
